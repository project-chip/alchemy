package spec

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

type Builder struct {
	specRoot string

	Spec *Specification

	ignoreHierarchy bool
}

func NewBuilder(specRoot string, options ...BuilderOption) Builder {
	b := Builder{specRoot: specRoot}
	for _, o := range options {
		o(&b)
	}
	return b
}

func (sp Builder) Name() string {
	return "Building spec"
}

func (sp *Builder) Process(cxt context.Context, inputs []*pipeline.Data[*Doc]) (outputs []*pipeline.Data[*Doc], err error) {
	docs := make([]*Doc, 0, len(inputs))
	for _, i := range inputs {
		docs = append(docs, i.Content)
	}
	var referencedDocs []*Doc
	sp.Spec, referencedDocs, err = sp.buildSpec(docs)
	for _, d := range referencedDocs {
		outputs = append(outputs, pipeline.NewData(d.Path.Absolute, d))
	}
	return
}

func (sp *Builder) buildSpec(docs []*Doc) (spec *Specification, referencedDocs []*Doc, err error) {

	buildTree(docs)

	spec = newSpec(sp.specRoot)

	docGroups := buildDocumentGroups(docs, spec)

	for _, dg := range docGroups {
		referencedDocs = append(referencedDocs, dg.Docs...)
	}

	indexCrossReferences(referencedDocs)

	err = indexAnchors(referencedDocs)
	if err != nil {
		return
	}

	var basicInformationCluster, bridgedBasicInformationCluster *matter.Cluster

	for _, d := range referencedDocs {
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
				if m.ID.Valid() {
					if existing, ok := spec.DeviceTypesByID[m.ID.Value()]; ok {
						slog.Error("duplicate device type ID", slog.String("deviceTypeId", m.ID.HexString()), log.Path("previousSource", existing), log.Path("newSource", m))
					} else {
						spec.DeviceTypesByID[m.ID.Value()] = m
					}

				}
				existing, ok := spec.DeviceTypesByName[m.Name]
				if ok {
					slog.Warn("Duplicate Device Type Name", slog.String("deviceTypeId", m.ID.HexString()), slog.String("deviceTypeName", m.Name), slog.String("existingDeviceTypeId", existing.ID.HexString()))
				}
				spec.DeviceTypesByName[m.Name] = m
			case *matter.Namespace:
				spec.Namespaces = append(spec.Namespaces, m)
			case *matter.Bitmap:
				slog.Debug("Found global bitmap", "name", m.Name, "path", d.Path)
				existing, ok := spec.bitmapIndex[m.Name]
				if ok {
					slog.Error("multiple bitmaps with same name", "name", m.Name, log.Path("previousSource", existing), log.Path("newSource", m))
				} else {
					spec.bitmapIndex[m.Name] = m
				}
				spec.addEntityByName(m.Name, m, nil)
				spec.GlobalObjects[m] = struct{}{}
			case *matter.Enum:
				slog.Debug("Found global enum", "name", m.Name, "path", d.Path)
				existing, ok := spec.enumIndex[m.Name]
				if ok {
					slog.Error("multiple enums with same name", "name", m.Name, log.Path("previousSource", existing), log.Path("newSource", m))
				} else {
					spec.enumIndex[m.Name] = m
				}
				spec.addEntityByName(m.Name, m, nil)
				spec.GlobalObjects[m] = struct{}{}
			case *matter.Struct:
				slog.Debug("Found global struct", "name", m.Name, "path", d.Path)
				existing, ok := spec.structIndex[m.Name]
				if ok {
					slog.Error("multiple structs with same name", "name", m.Name, log.Path("previousSource", existing), log.Path("newSource", m))
				} else {
					spec.structIndex[m.Name] = m
				}
				spec.addEntityByName(m.Name, m, nil)
				spec.GlobalObjects[m] = struct{}{}
			case *matter.TypeDef:
				slog.Debug("Found global typedef", "name", m.Name, "path", d.Path)
				existing, ok := spec.typeDefIndex[m.Name]
				if ok {
					slog.Warn("multiple global typedefs with same name", "name", m.Name, log.Path("previousSource", existing), log.Path("newSource", m))
				} else {
					spec.typeDefIndex[m.Name] = m
				}
				spec.addEntityByName(m.Name, m, nil)
				spec.GlobalObjects[m] = struct{}{}
			case *matter.Command:
				existing, ok := spec.commandIndex[m.Name]
				if ok {
					slog.Error("multiple commands with same name", "name", m.Name, log.Path("previousSource", existing), log.Path("newSource", m))
				} else {
					spec.commandIndex[m.Name] = m
				}
				spec.addEntityByName(m.Name, m, nil)
				spec.GlobalObjects[m] = struct{}{}
			case *matter.Event:
				existing, ok := spec.eventIndex[m.Name]
				if ok {
					slog.Error("multiple events with same name", "name", m.Name, log.Path("previousSource", existing), log.Path("newSource", m))
				} else {
					spec.eventIndex[m.Name] = m
				}
				spec.addEntityByName(m.Name, m, nil)
				spec.GlobalObjects[m] = struct{}{}
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

	}

	if !sp.ignoreHierarchy {
		resolveHierarchy(spec)
	}
	associateDeviceTypeRequirementWithClusters(spec)

	sp.resolveDataTypeReferences(spec)
	sp.resolveConformances(spec)
	sp.resolveConstraints(spec)
	err = updateBridgedBasicInformationCluster(basicInformationCluster, bridgedBasicInformationCluster)
	if err != nil {
		return
	}

	buildClusterReferences(spec)

	return
}

func buildDocumentGroups(docs []*Doc, spec *Specification) (docGroups []*DocGroup) {
	for _, d := range docs {
		if len(d.parents) > 0 {
			continue
		}

		var isDocRoot bool
		path := d.Path.Relative
		for _, docRoot := range errata.DocRoots {
			if strings.EqualFold(path, docRoot) {
				isDocRoot = true
				break
			}
		}

		if !isDocRoot {
			continue
		}

		dg := NewDocGroup(d.Path.Relative)
		docGroups = append(docGroups, dg)
		setSpec(d, spec, dg)
	}
	return
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
			if cr.Cluster != nil {
				continue
			}
			if c, ok := spec.ClustersByID[cr.ClusterID.Value()]; ok {
				cr.Cluster = c
			} else {
				if c, ok := spec.ClustersByName[cr.ClusterName]; ok {
					cr.Cluster = c
					slog.Warn("linking cluster requirement by name on device type since cluster ID was not recognized",
						slog.String("clusterId", cr.ClusterID.HexString()),
						slog.String("clusterName", cr.ClusterName),
						slog.String("deviceType", dt.Name),
						log.Path("source", cr))
				} else {
					slog.Error("unknown cluster ID for cluster requirement on device type",
						slog.String("clusterId", cr.ClusterID.HexString()),
						slog.String("clusterName", cr.ClusterName),
						slog.String("deviceType", dt.Name),
						log.Path("source", cr))
				}
			}
		}
		for _, er := range dt.ElementRequirements {
			if er.Cluster != nil {
				continue
			}
			if c, ok := spec.ClustersByID[er.ClusterID.Value()]; ok {
				er.Cluster = c
			} else {
				if c, ok := spec.ClustersByName[er.ClusterName]; ok {
					er.Cluster = c
					slog.Warn("linking element requirement by cluster name on device type since cluster ID was not recognized",
						slog.String("clusterId", er.ClusterID.HexString()),
						slog.String("clusterName", er.ClusterName),
						slog.String("deviceType", dt.Name),
						log.Path("source", er))
				} else {
					slog.Error("unknown cluster ID for element requirement on device type",
						slog.String("clusterId", er.ClusterID.HexString()),
						slog.String("clusterName", er.ClusterName),
						slog.String("deviceType", dt.Name),
						log.Path("source", er))
				}
			}
		}
	}
}

func addClusterToSpec(spec *Specification, doc *Doc, cluster *matter.Cluster) {
	spec.Clusters[cluster] = struct{}{}
	if cluster.ID.Valid() {
		existing, ok := spec.ClustersByID[cluster.ID.Value()]
		if ok {
			slog.Warn("Duplicate cluster ID", slog.String("clusterId", cluster.ID.HexString()), slog.String("clusterName", cluster.Name), slog.String("existingClusterName", existing.Name))
		}
		spec.ClustersByID[cluster.ID.Value()] = cluster
	} else {
		idText := cluster.ID.Text()
		if !strings.EqualFold(idText, "n/a") {
			if strings.EqualFold(idText, "ID-TBD") {
				slog.Warn("Cluster has not yet been assigned an ID; this may cause issues with generated code", slog.String("clusterName", cluster.Name))
			} else {
				slog.Warn("Cluster has invalid ID; this may cause issues with generated code", slog.String("clusterId", idText), slog.String("clusterName", cluster.Name))

			}
		}
	}
	existing, ok := spec.ClustersByName[cluster.Name]
	if ok {
		slog.Warn("Duplicate cluster Name", slog.String("clusterId", cluster.ID.HexString()), slog.String("clusterName", cluster.Name), slog.String("existingClusterId", existing.ID.HexString()))
	}
	spec.ClustersByName[cluster.Name] = cluster

	for _, en := range cluster.Bitmaps {
		_, ok := spec.bitmapIndex[en.Name]
		if ok {
			slog.Debug("multiple bitmaps with same name", "name", en.Name)
		} else {
			spec.bitmapIndex[en.Name] = en
		}
		spec.addEntityByName(en.Name, en, cluster)
	}
	for _, en := range cluster.Enums {
		_, ok := spec.enumIndex[en.Name]
		if ok {
			slog.Debug("multiple enums with same name", "name", en.Name)
		} else {
			spec.enumIndex[en.Name] = en
		}
		spec.addEntityByName(en.Name, en, cluster)
	}
	for _, en := range cluster.Structs {
		_, ok := spec.structIndex[en.Name]
		if ok {
			slog.Debug("multiple structs with same name", "name", en.Name)
		} else {
			spec.structIndex[en.Name] = en
		}
		spec.addEntityByName(en.Name, en, cluster)
	}
	for _, en := range cluster.TypeDefs {
		_, ok := spec.typeDefIndex[en.Name]
		if ok {
			slog.Debug("multiple structs with same name", "name", en.Name)
		} else {
			spec.typeDefIndex[en.Name] = en
		}
		spec.addEntityByName(en.Name, en, cluster)
	}
	noteDocRefs(spec, doc, cluster)
}

func noteDocRefs(spec *Specification, doc *Doc, cluster *matter.Cluster) {
	for _, bm := range cluster.Bitmaps {
		spec.DocRefs[bm] = doc
	}
	for _, e := range cluster.Enums {
		spec.DocRefs[e] = doc
	}
	for _, s := range cluster.Structs {
		spec.DocRefs[s] = doc
	}
	for _, td := range cluster.TypeDefs {
		spec.DocRefs[td] = doc
	}
	for _, a := range cluster.Attributes {
		spec.DocRefs[a] = doc
	}
	for _, e := range cluster.Events {
		spec.DocRefs[e] = doc
	}
	for _, cmd := range cluster.Commands {
		spec.DocRefs[cmd] = doc
	}
}

func getTagNamespace(spec *Specification, field *matter.Field) {
	for _, ns := range spec.Namespaces {
		if strings.EqualFold(ns.Name, field.Type.Name) {
			field.Type.Entity = ns
			return
		}
	}
	if field.Type.Name != "tag" {
		// Warn on unknown tag namespace, except for the example namespace "tag"
		slog.Warn("failed to match tag name space", slog.String("name", field.Name), log.Path("field", field), slog.String("namespace", field.Type.Name))
	}
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
		doc, ok := spec.DocRefs[c]
		if ok {
			// We may have created some new entities during the inherit, so make sure their doc refs are set
			noteDocRefs(spec, doc, c)
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
		ba.Fallback = a.Fallback
		ba.Access = a.Access
	}
	return nil
}
