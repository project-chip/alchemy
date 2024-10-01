package spec

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

type Builder struct {
	Spec *Specification

	IgnoreHierarchy bool
}

func NewBuilder() Builder {
	return Builder{}
}

func (sp Builder) Name() string {
	return "Building spec"
}

func (sp Builder) Type() pipeline.ProcessorType {
	return pipeline.ProcessorTypeCollective
}

func (sp *Builder) Process(cxt context.Context, inputs []*pipeline.Data[*Doc]) (outputs []*pipeline.Data[*Doc], err error) {
	docs := make([]*Doc, 0, len(inputs))
	for _, i := range inputs {
		docs = append(docs, i.Content)
	}
	sp.Spec, err = sp.buildSpec(docs)
	outputs = inputs
	return
}

func (sp *Builder) buildSpec(docs []*Doc) (spec *Specification, err error) {

	buildTree(docs)

	spec = newSpec()

	buildDocumentGroups(docs, spec)

	indexCrossReferences(docs)

	err = indexAnchors(docs)
	if err != nil {
		return
	}

	var basicInformationCluster, bridgedBasicInformationCluster *matter.Cluster

	for _, d := range docs {
		slog.Debug("building spec", "path", d.Path)

		dt, dterr := d.DocType()
		if dterr == nil {
			switch dt {
			case matter.DocTypeBaseDeviceType:
				spec.BaseDeviceType, err = d.toBaseDeviceType()
				if err != nil {
					return
				}
			}
		}

		var entities []types.Entity
		entities, err = d.Entities()
		if err != nil {
			slog.Warn("error building entities", "doc", d.Path, "error", err)
			continue
		}
		for _, m := range entities {
			switch m := m.(type) {
			case *matter.ClusterGroup:
				for _, c := range m.Clusters {
					addClusterToSpec(spec, d, c)
				}
			case *matter.Cluster:
				switch m.Name {
				case "Basic Information":
					basicInformationCluster = m
				case "Bridged Device Basic Information":
					bridgedBasicInformationCluster = m
				}
				addClusterToSpec(spec, d, m)
			case *matter.DeviceType:
				spec.DeviceTypes = append(spec.DeviceTypes, m)
			case *matter.Namespace:
				spec.Namespaces = append(spec.Namespaces, m)
			default:
				slog.Warn("unknown entity type", "path", d.Path, "type", fmt.Sprintf("%T", m))
			}
			switch m := m.(type) {
			case *matter.ClusterGroup:
				for _, c := range m.Clusters {
					spec.DocRefs[c] = d
				}
			default:
				spec.DocRefs[m] = d
			}
		}

		err = addGlobalEntities(spec, d)

		if err != nil {
			slog.Warn("error building global objects", "doc", d.Path, "error", err)
			continue
		}

	}

	if !sp.IgnoreHierarchy {
		resolveHierarchy(spec)
	}
	sp.resolveDataTypeReferences(spec)
	err = updateBridgedBasicInformationCluster(basicInformationCluster, bridgedBasicInformationCluster)
	if err != nil {
		return
	}

	buildClusterReferences(spec)
	associateDeviceTypeRequirementWithClusters(spec)

	return
}

func buildDocumentGroups(docs []*Doc, spec *Specification) {
	for _, d := range docs {
		if len(d.parents) > 0 {
			continue
		}

		dg := NewDocGroup(d.Path.Relative)
		setSpec(d, spec, dg)
	}
}

func buildClusterReferences(spec *Specification) {
	for _, c := range spec.ClustersByName {
		if c.Features != nil {
			spec.ClusterRefs.Add(c, c.Features)
		}
		for _, en := range c.Bitmaps {
			spec.ClusterRefs.Add(c, en)
		}
		for _, en := range c.Enums {
			spec.ClusterRefs.Add(c, en)
		}
		for _, en := range c.Structs {
			spec.ClusterRefs.Add(c, en)
		}
	}
}

func indexAnchors(docs []*Doc) (err error) {
	for _, d := range docs {
		var anchors map[string][]*Anchor
		anchors, err = d.Anchors()
		if err != nil {
			return
		}
		for id, anchor := range anchors {
			d.group.anchors[id] = append(d.group.anchors[id], anchor...)
		}
		for id, anchor := range d.anchorsByLabel {
			d.group.anchorsByLabel[id] = append(d.group.anchorsByLabel[id], anchor...)
		}
	}
	return
}

func indexCrossReferences(docs []*Doc) {
	for _, d := range docs {
		crossReferences := d.CrossReferences()
		for id, xrefs := range crossReferences {
			d.group.crossReferences[id] = append(d.group.crossReferences[id], xrefs...)
		}
	}
}

func associateDeviceTypeRequirementWithClusters(spec *Specification) {
	for _, dt := range spec.DeviceTypes {
		for _, cr := range dt.ClusterRequirements {
			if c, ok := spec.ClustersByID[cr.ClusterID.Value()]; ok {
				cr.Cluster = c
			} else {
				slog.Warn("unknown cluster ID for cluster requirement on device type", "clusterId", cr.ClusterID.HexString(), "clusterName", cr.ClusterName, "deviceType", dt.Name)
			}
		}
		for _, er := range dt.ElementRequirements {
			if c, ok := spec.ClustersByID[er.ClusterID.Value()]; ok {
				er.Cluster = c
			} else {
				slog.Warn("unknown cluster ID for element requirement on device type", "clusterId", er.ClusterID.HexString(), "clusterName", er.ClusterName, "deviceType", dt.Name)

			}
		}
	}
}

func addClusterToSpec(spec *Specification, d *Doc, m *matter.Cluster) {
	spec.Clusters[m] = struct{}{}
	if m.ID.Valid() {
		existing, ok := spec.ClustersByID[m.ID.Value()]
		if ok {
			slog.Warn("Duplicate cluster ID", slog.String("clusterId", m.ID.HexString()), slog.String("clusterName", m.Name), slog.String("existingClusterName", existing.Name))
		}
		spec.ClustersByID[m.ID.Value()] = m
	}
	existing, ok := spec.ClustersByName[m.Name]
	if ok {
		slog.Warn("Duplicate cluster Name", slog.String("clusterId", m.ID.HexString()), slog.String("clusterName", m.Name), slog.String("existingClusterId", existing.ID.HexString()))
	}
	spec.ClustersByName[m.Name] = m

	for _, en := range m.Bitmaps {
		_, ok := spec.bitmapIndex[en.Name]
		if ok {
			slog.Debug("multiple bitmaps with same name", "name", en.Name)
		} else {
			spec.bitmapIndex[en.Name] = en
		}
		spec.DocRefs[en] = d
		spec.addEntityByName(en.Name, en, m)
	}
	for _, en := range m.Enums {
		_, ok := spec.enumIndex[en.Name]
		if ok {
			slog.Debug("multiple enums with same name", "name", en.Name)
		} else {
			spec.enumIndex[en.Name] = en
		}
		spec.DocRefs[en] = d
		spec.addEntityByName(en.Name, en, m)
	}
	for _, en := range m.Structs {
		_, ok := spec.structIndex[en.Name]
		if ok {
			slog.Debug("multiple structs with same name", "name", en.Name)
		} else {
			spec.structIndex[en.Name] = en
		}
		spec.DocRefs[en] = d
		spec.addEntityByName(en.Name, en, m)
	}
	for _, en := range m.TypeDefs {
		_, ok := spec.typeDefIndex[en.Name]
		if ok {
			slog.Debug("multiple structs with same name", "name", en.Name)
		} else {
			spec.typeDefIndex[en.Name] = en
		}
		spec.DocRefs[en] = d
		spec.addEntityByName(en.Name, en, m)
	}
}

func getTagNamespace(spec *Specification, field *matter.Field) {
	for _, ns := range spec.Namespaces {
		if strings.EqualFold(ns.Name, field.Type.Name) {
			field.Type.Entity = ns
			return
		}
	}
	slog.Warn("failed to match tag name space", slog.String("name", field.Name), log.Path("field", field), slog.String("namespace", field.Type.Name))
}

func clusterName(cluster *matter.Cluster) string {
	if cluster != nil {
		return cluster.Name
	}
	return "none"
}

func resolveHierarchy(spec *Specification) {
	for _, c := range spec.ClustersByID {
		if c.Hierarchy == "Base" {
			continue
		}
		base, ok := spec.ClustersByName[c.Hierarchy]
		if !ok {
			slog.Warn("Failed to find base cluster", "cluster", c.Name, "baseCluster", c.Hierarchy)
			continue
		}
		linkedEntities, err := c.Inherit(base)
		if err != nil {
			slog.Warn("Failed to inherit from base cluster", "cluster", c.Name, "baseCluster", c.Hierarchy, "error", err)
		}
		// These entities were inherited from a base cluster, but not modified
		for _, linkedEntity := range linkedEntities {
			spec.ClusterRefs.Add(c, linkedEntity)
			spec.addEntity(linkedEntity, c)
		}
		assignCustomDataTypes(c)
	}
}

func updateBridgedBasicInformationCluster(basicInformationCluster *matter.Cluster, bridgedBasicInformationCluster *matter.Cluster) error {
	if basicInformationCluster == nil {
		return fmt.Errorf("missing Basic Information Cluster in spec")
	}
	if bridgedBasicInformationCluster == nil {
		return fmt.Errorf("missing Basic Information Cluster in spec")
	}
	am := make(map[uint64]*matter.Field, len(basicInformationCluster.Attributes))
	for _, a := range basicInformationCluster.Attributes {
		am[a.ID.Value()] = a
	}
	for _, ba := range bridgedBasicInformationCluster.Attributes {
		id := ba.ID.Value()
		a, ok := am[id]
		if !ok {
			continue
		}
		ba.Type = a.Type.Clone()
		ba.Constraint = a.Constraint.Clone()
		ba.Quality = a.Quality
		ba.Default = a.Default
		ba.Access = a.Access
	}
	return nil
}
