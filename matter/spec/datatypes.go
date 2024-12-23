package spec

import (
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

func (s *Section) toDataTypes(d *Doc, pc *parseContext, parentEntity types.Entity) (bitmaps matter.BitmapSet, enums matter.EnumSet, structs matter.StructSet, typedefs matter.TypeDefSet, err error) {

	traverseSections(d, s, errata.SpecPurposeDataTypes, func(s *Section, parent parse.HasElements, index int) parse.SearchShould {
		switch s.SecType {
		case matter.SectionDataTypeBitmap:
			var mb *matter.Bitmap
			mb, err = s.toBitmap(d, pc, parentEntity)
			if err != nil {
				if err != ErrNotEnoughRowsInTable || d.parsed {
					slog.Warn("Error converting section to bitmap", log.Element("source", d.Path, s.Base), slog.Any("error", err))
				}
				err = nil
			} else {
				bitmaps = append(bitmaps, mb)
			}
		case matter.SectionDataTypeEnum:
			var me *matter.Enum
			me, err = s.toEnum(d, pc, parentEntity)
			if err != nil {
				if err != ErrNotEnoughRowsInTable || d.parsed {
					slog.Warn("Error converting section to enum", log.Element("source", d.Path, s.Base), slog.Any("error", err))
				}
				err = nil
			} else {
				enums = append(enums, me)
			}
		case matter.SectionDataTypeStruct:
			var me *matter.Struct
			me, err = s.toStruct(d, pc, parentEntity)
			if err != nil {
				if err != ErrNotEnoughRowsInTable || d.parsed {
					slog.Warn("Error converting section to struct", log.Element("source", d.Path, s.Base), slog.Any("error", err))
				}
				err = nil
			} else {
				structs = append(structs, me)
			}
		case matter.SectionDataTypeDef:
			var me *matter.TypeDef
			me, err = s.toTypeDef(d, pc, parentEntity)
			if err != nil {
				if err != ErrNotEnoughRowsInTable || d.parsed {
					slog.Warn("Error converting section to typedef", log.Element("source", d.Path, s.Base), slog.Any("error", err))
				}
				err = nil
			} else {
				typedefs = append(typedefs, me)
			}
		default:
		}
		return parse.SearchShouldContinue
	})

	return
}

func (sp *Builder) resolveDataTypeReferences(spec *Specification) {
	for cluster := range spec.Clusters {
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
	for _, s := range spec.structIndex {
		for _, f := range s.Fields {
			sp.resolveDataType(spec, nil, f, f.Type)
		}
	}
}

func (sp *Builder) resolveDataType(spec *Specification, cluster *matter.Cluster, field *matter.Field, dataType *types.DataType) {
	if dataType == nil {
		if !conformance.IsDeprecated(field.Conformance) && !conformance.IsDisallowed(field.Conformance) && !sp.IgnoreHierarchy && (cluster == nil || cluster.Hierarchy == "Base") {
			var clusterName string
			if cluster != nil {
				clusterName = cluster.Name
			}
			slog.Warn("missing type on field", log.Path("path", field), slog.String("id", field.ID.HexString()), slog.String("name", field.Name), slog.String("cluster", clusterName))
		}
		return
	}
	if dataType.Entity != nil {
		// This has already been resolved by some other process
		return
	}
	switch dataType.BaseType {
	case types.BaseDataTypeTag:
		getTagNamespace(spec, field)
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
		s, ok := dataType.Entity.(*matter.Struct)
		if !ok {
			return
		}
		// If this data type is a struct, we need to resolve all data types on its fields
		for _, f := range s.Fields {
			sp.resolveDataType(spec, cluster, f, f.Type)
		}
	}
}

func getCustomDataType(spec *Specification, dataTypeName string, cluster *matter.Cluster, field *matter.Field) (e types.Entity) {
	e = getCustomDataTypeFromReference(spec, cluster, field)
	if e != nil {
		// We have a reference to a data type; use that
		return
	}
	entities := spec.entities[dataTypeName]
	if len(entities) == 0 {
		canonicalName := CanonicalName(dataTypeName)
		if canonicalName != dataTypeName {
			e = getCustomDataType(spec, canonicalName, cluster, field)
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

func getCustomDataTypeFromReference(spec *Specification, cluster *matter.Cluster, field *matter.Field) (e types.Entity) {
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
			entities := anchor.Document.entitiesBySection[el]
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
	if cluster != nil && cluster.ParentCluster != nil {
		for m, c := range entities {
			if c != nil && c == cluster.ParentCluster {
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
		slog.Warn("ambiguous data type", matter.LogEntity("source", m), "cluster", clusterName)
	}
	return nil
}
