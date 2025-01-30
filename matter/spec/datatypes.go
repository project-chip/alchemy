package spec

import (
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

func (s *Section) toDataTypes(d *Doc, pc *parseContext, parentEntity types.Entity) (dataTypes []types.Entity, err error) {

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
				dataTypes = append(dataTypes, mb)
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
				dataTypes = append(dataTypes, me)
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
				dataTypes = append(dataTypes, me)
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
				dataTypes = append(dataTypes, me)
			}
		case matter.SectionDataTypeConstant:
			id, _ := getAnchorElements(s.Base, nil)
			switch id {
			case "ref_RespMaxConstant":
				c := matter.NewConstant(s)
				c.Name = "RESP_MAX"
				c.Value = 900
				pc.orderedEntities = append(pc.orderedEntities, c)
				pc.entitiesByElement[s.Base] = append(pc.entitiesByElement[s.Base], c)
				dataTypes = append(dataTypes, c)
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
			sp.resolveFieldDataTypes(spec, cluster, cluster.Attributes, a, a.Type)
		}
		for _, s := range cluster.Structs {
			for _, f := range s.Fields {
				sp.resolveFieldDataTypes(spec, cluster, s.Fields, f, f.Type)
			}
		}
		for _, event := range cluster.Events {
			for _, f := range event.Fields {
				sp.resolveFieldDataTypes(spec, cluster, event.Fields, f, f.Type)
			}
		}
		for _, command := range cluster.Commands {
			for _, f := range command.Fields {
				sp.resolveFieldDataTypes(spec, cluster, command.Fields, f, f.Type)
			}
		}
	}

	for _, s := range spec.structIndex {
		for _, f := range s.Fields {
			sp.resolveFieldDataTypes(spec, nil, s.Fields, f, f.Type)
		}
	}
}

func (sp *Builder) resolveFieldDataTypes(spec *Specification, cluster *matter.Cluster, fieldSet matter.FieldSet, field *matter.Field, dataType *types.DataType) {
	if dataType == nil {
		if !conformance.IsDeprecated(field.Conformance) && !conformance.IsDisallowed(field.Conformance) && !sp.ignoreHierarchy && (cluster == nil || cluster.Hierarchy == "Base") {
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
		sp.resolveFieldDataTypes(spec, cluster, fieldSet, field, dataType.EntryType)
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
			sp.resolveFieldDataTypes(spec, cluster, fieldSet, f, f.Type)
		}
	}
}

func getCustomDataType(spec *Specification, dataTypeName string, cluster *matter.Cluster, field *matter.Field) (e types.Entity) {
	e = getCustomDataTypeFromFieldReference(spec, cluster, field.Type.Source)
	if e != nil {
		// We have a reference to a data type; use that
		return
	}
	e = getCustomDataTypeFromIdentifier(spec, cluster, field, dataTypeName)
	return
}

func getCustomDataTypeFromIdentifier(spec *Specification, cluster *matter.Cluster, source log.Source, identifier string) types.Entity {
	entities := spec.entities[identifier]
	if len(entities) == 0 {
		canonicalName := CanonicalName(identifier)
		if canonicalName != identifier {
			return getCustomDataTypeFromIdentifier(spec, cluster, source, canonicalName)
		}
	} else if len(entities) == 1 {
		for m := range entities {
			return m
		}
	} else {
		return disambiguateDataType(entities, cluster, identifier, source)
	}
	return nil
}

func getCustomDataTypeFromFieldReference(spec *Specification, cluster *matter.Cluster, source asciidoc.Element) (e types.Entity) {
	switch source := source.(type) {
	case *asciidoc.CrossReference:
		return getCustomDataTypeFromReference(spec, cluster, source.ID, asciidoc.ValueToString(source.Elements()))
	default:
		return
	}
}

func getCustomDataTypeFromReference(spec *Specification, cluster *matter.Cluster, reference string, label string) (e types.Entity) {
	doc, ok := spec.DocRefs[cluster]
	if !ok {
		return
	}
	anchor := doc.FindAnchor(reference)
	if anchor == nil {
		return
	}
	switch el := anchor.Element.(type) {
	case *asciidoc.Section:
		entities := anchor.Document.entitiesBySection[el]
		if len(entities) == 1 {
			e = entities[0]

		}
	}
	if e != nil && label != "" {
		switch entity := e.(type) {
		case *matter.Enum:
			if strings.EqualFold(entity.Name, label) {
				return
			}
			for _, ev := range entity.Values {
				if strings.EqualFold(label, ev.Name) {
					e = ev
					return
				}
			}
		case *matter.Struct:
			if strings.EqualFold(entity.Name, label) {
				return
			}
			for _, f := range entity.Fields {
				if strings.EqualFold(label, f.Name) {
					e = f
					return
				}
			}
		case *matter.Command:
			if strings.EqualFold(entity.Name, label) {
				return
			}
			for _, f := range entity.Fields {
				if strings.EqualFold(label, f.Name) {
					e = f
					return
				}
			}
		case *matter.Field:
			if strings.EqualFold(entity.Name, label) {
				return
			}
			slog.Warn("Unhandled reference field with label", slog.String("clusterName", cluster.Name), slog.String("field", entity.Name))
		case *matter.Constant:
			if strings.EqualFold(entity.Name, label) {
				return
			}
			slog.Warn("Unhandled reference constant with label", slog.String("clusterName", cluster.Name), slog.String("constant", entity.Name))
		default:
			slog.Warn("Unhandled reference type with label", slog.String("clusterName", cluster.Name), log.Type("entityType", e))
		}
	}
	return
}

func disambiguateDataType(entities map[types.Entity]*matter.Cluster, cluster *matter.Cluster, identifier string, source log.Source) types.Entity {
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
	slog.Warn("ambiguous data type", "cluster", clusterName(cluster), "identifier", identifier, log.Path("source", source))
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
