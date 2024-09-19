package spec

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
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

	for _, d := range docs {
		if len(d.parents) > 0 {
			continue
		}

		dg := NewDocGroup(d.Path.Relative)
		setSpec(d, spec, dg)
	}

	for _, d := range docs {
		crossReferences := d.CrossReferences()
		for id, xrefs := range crossReferences {
			d.group.crossReferences[id] = append(d.group.crossReferences[id], xrefs...)
		}
	}

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
			case matter.DocTypeDataModel:

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
					addClusterToSpec(spec, d, c, d.spec)
				}
			case *matter.Cluster:
				switch m.Name {
				case "Basic Information":
					basicInformationCluster = m
				case "Bridged Device Basic Information":
					bridgedBasicInformationCluster = m
				}
				addClusterToSpec(spec, d, m, d.spec)
			case *matter.DeviceType:
				spec.DeviceTypes = append(spec.DeviceTypes, m)
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

	for _, dt := range spec.DeviceTypes {
		for _, cr := range dt.ClusterRequirements {
			if c, ok := spec.ClustersByID[cr.ID.Value()]; ok {
				cr.Cluster = c
			} else {
				slog.Warn("unknown cluster ID for cluster requirement on device type", "clusterId", cr.ID.HexString(), "clusterName", cr.ClusterName, "deviceType", dt.Name)
			}
		}
		for _, er := range dt.ElementRequirements {
			if c, ok := spec.ClustersByID[er.ID.Value()]; ok {
				er.Cluster = c
			} else {
				slog.Warn("unknown cluster ID for element requirement on device type", "clusterId", er.ID.HexString(), "clusterName", er.ClusterName, "deviceType", dt.Name)

			}
		}
	}
	return
}

func addClusterToSpec(spec *Specification, d *Doc, m *matter.Cluster, specIndex *Specification) {
	if m.ID.Valid() {
		spec.ClustersByID[m.ID.Value()] = m
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
		specIndex.addEntity(en.Name, en, m)
	}
	for _, en := range m.Enums {
		_, ok := spec.enumIndex[en.Name]
		if ok {
			slog.Debug("multiple enums with same name", "name", en.Name)
		} else {
			spec.enumIndex[en.Name] = en
		}
		spec.DocRefs[en] = d
		specIndex.addEntity(en.Name, en, m)
	}
	for _, en := range m.Structs {
		_, ok := spec.structIndex[en.Name]
		if ok {
			slog.Debug("multiple structs with same name", "name", en.Name)
		} else {
			spec.structIndex[en.Name] = en
		}
		spec.DocRefs[en] = d
		specIndex.addEntity(en.Name, en, m)
	}
}

func (sp *Builder) resolveDataTypeReferences(spec *Specification) {
	for _, s := range spec.structIndex {
		for _, f := range s.Fields {
			sp.resolveDataType(spec, nil, f, f.Type)
		}
	}
	for _, cluster := range spec.ClustersByName {
		for _, a := range cluster.Attributes {
			sp.resolveDataType(spec, cluster, a, a.Type)
		}
		for _, s := range cluster.Structs {
			for _, f := range s.Fields {
				sp.resolveDataType(spec, cluster, f, f.Type)
			}
		}
		for _, s := range cluster.Events {
			for _, f := range s.Fields {
				sp.resolveDataType(spec, cluster, f, f.Type)
			}
		}
		for _, s := range cluster.Commands {
			for _, f := range s.Fields {
				sp.resolveDataType(spec, cluster, f, f.Type)
			}
		}
	}

}

func (sp *Builder) resolveDataType(spec *Specification, cluster *matter.Cluster, field *matter.Field, dataType *types.DataType) {
	if dataType == nil {
		if !conformance.IsDeprecated(field.Conformance) && (cluster == nil || cluster.Hierarchy == "Base") {
			var clusterName string
			if cluster != nil {
				clusterName = cluster.Name
			}
			if !sp.IgnoreHierarchy {
				slog.Warn("missing type on field", log.Path("path", field), slog.String("id", field.ID.HexString()), slog.String("name", field.Name), slog.String("cluster", clusterName))
			}
		}
		return
	}
	switch dataType.BaseType {
	case types.BaseDataTypeList:
		sp.resolveDataType(spec, cluster, field, dataType.EntryType)
	case types.BaseDataTypeCustom:
		if dataType.Entity == nil {
			dataType.Entity = getCustomDataType(spec, dataType.Name, cluster, field)
			if dataType.Entity == nil {
				slog.Error("unknown custom data type", slog.String("cluster", clusterName(cluster)), slog.String("field", field.Name), slog.String("type", dataType.Name), log.Path("source", field))
			}
		}
		if cluster == nil || dataType.Entity == nil {
			return
		}
		spec.ClusterRefs.Add(cluster, dataType.Entity)
		slog.Debug("setting cluster", "name", cluster.Name, "type", dataType.Name)
		s, ok := dataType.Entity.(*matter.Struct)
		if !ok {
			return
		}
		for _, f := range s.Fields {
			sp.resolveDataType(spec, cluster, f, f.Type)
		}
	}
}

func getCustomDataType(spec *Specification, dataTypeName string, cluster *matter.Cluster, field *matter.Field) (e types.Entity) {
	entities := spec.entities[dataTypeName]
	if len(entities) == 0 {
		canonicalName := CanonicalName(dataTypeName)
		if canonicalName != dataTypeName {
			e = getCustomDataType(spec, canonicalName, cluster, field)
		} else {
			e = getCustomDataTypeFromReference(spec, dataTypeName, cluster, field)
		}
	} else if len(entities) == 1 {
		for m := range entities {
			e = m
		}
	} else {
		e = disambiguateDataType(entities, cluster, field)
	}
	return
}

func getCustomDataTypeFromReference(spec *Specification, dataTypeName string, cluster *matter.Cluster, field *matter.Field) (e types.Entity) {
	switch source := field.Type.Source.(type) {
	case *asciidoc.CrossReference:
		doc, ok := spec.DocRefs[cluster]
		if !ok {
			return
		}
		anchor := doc.FindAnchor(source.ID)
		if anchor == nil {
			return
		}
		switch el := anchor.Element.(type) {
		case *asciidoc.Section:
			entities := doc.entitiesBySection[el]
			if len(entities) == 1 {
				e = entities[0]
				return
			}
		}
		return nil
	default:
		return
	}
}

func disambiguateDataType(entities map[types.Entity]*matter.Cluster, cluster *matter.Cluster, field *matter.Field) types.Entity {
	// If there are multiple entities with the same name, prefer the one on the current cluster
	for m, c := range entities {
		if c == cluster {
			return m
		}
	}

	// OK, if the data type is defined on the direct parent of this cluster, take that one
	if cluster.Hierarchy != "Base" {
		for m, c := range entities {
			if c != nil && c.Name == cluster.Hierarchy {
				return m
			}
		}
	}

	var nakedEntities []types.Entity
	for m, c := range entities {
		if c == nil {
			nakedEntities = append(nakedEntities, m)
		}
	}
	if len(nakedEntities) == 1 {
		return nakedEntities[0]
	}

	// Can't disambiguate out this data model
	slog.Warn("ambiguous data type", "cluster", clusterName(cluster), "field", field.Name, log.Path("source", field))
	for m, c := range entities {
		var clusterName string
		if c != nil {
			clusterName = c.Name
		} else {
			clusterName = "naked"
		}
		slog.Warn("ambiguous data type", "model", m.Source(), "cluster", clusterName)
	}
	return nil
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
		base.Base = true
		linkedEntites, err := c.Inherit(base)
		if err != nil {
			slog.Warn("Failed to inherit from base cluster", "cluster", c.Name, "baseCluster", c.Hierarchy, "error", err)
		}
		for _, linkedEntity := range linkedEntites {
			spec.ClusterRefs.Add(c, linkedEntity)
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
