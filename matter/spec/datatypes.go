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
	"github.com/project-chip/alchemy/matter/constraint"
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
			sp.resolveDataTypes(spec, cluster, a, a.Type)
		}
		for _, s := range cluster.Structs {
			for _, f := range s.Fields {
				sp.resolveDataTypes(spec, cluster, f, f.Type)
			}
		}
		for _, s := range cluster.Events {
			for _, f := range s.Fields {
				sp.resolveDataTypes(spec, cluster, f, f.Type)
			}
		}
		for _, s := range cluster.Commands {
			for _, f := range s.Fields {
				sp.resolveDataTypes(spec, cluster, f, f.Type)
			}
		}
	}
	for _, s := range spec.structIndex {
		for _, f := range s.Fields {
			sp.resolveDataTypes(spec, nil, f, f.Type)
		}
	}
}

func (sp *Builder) resolveDataTypes(spec *Specification, cluster *matter.Cluster, field *matter.Field, dataType *types.DataType) {
	sp.resolveFieldDataType(spec, cluster, field, dataType)
	sp.resolveFieldConstraintReferences(spec, cluster, field, field.Constraint)
	sp.resolveFieldConstraintLimit(spec, cluster, field, field.Fallback)
}

func (sp *Builder) resolveFieldDataType(spec *Specification, cluster *matter.Cluster, field *matter.Field, dataType *types.DataType) {
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
		sp.resolveDataTypes(spec, cluster, field, dataType.EntryType)
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
			sp.resolveDataTypes(spec, cluster, f, f.Type)
		}
	}
}

func (sp *Builder) resolveFieldConstraintReferences(spec *Specification, cluster *matter.Cluster, field *matter.Field, con constraint.Constraint) {
	switch con := con.(type) {
	case *constraint.ExactConstraint:
		sp.resolveFieldConstraintLimit(spec, cluster, field, con.Value)
	case *constraint.ListConstraint:
		sp.resolveFieldConstraintReferences(spec, cluster, field, con.Constraint)
		sp.resolveFieldConstraintReferences(spec, cluster, field, con.EntryConstraint)
	case *constraint.MaxConstraint:
		sp.resolveFieldConstraintLimit(spec, cluster, field, con.Maximum)
	case *constraint.MinConstraint:
		sp.resolveFieldConstraintLimit(spec, cluster, field, con.Minimum)
	case *constraint.RangeConstraint:
		sp.resolveFieldConstraintLimit(spec, cluster, field, con.Minimum)
		sp.resolveFieldConstraintLimit(spec, cluster, field, con.Maximum)
	case constraint.Set:
		for _, c := range con {
			sp.resolveFieldConstraintReferences(spec, cluster, field, c)
		}
	}
}

func (sp *Builder) resolveFieldConstraintLimit(spec *Specification, cluster *matter.Cluster, field *matter.Field, l constraint.Limit) {
	switch l := l.(type) {
	case *constraint.CharacterLimit:
		sp.resolveFieldConstraintLimit(spec, cluster, field, l.ByteCount)
		sp.resolveFieldConstraintLimit(spec, cluster, field, l.CodepointCount)
	case *constraint.LengthLimit:
		sp.resolveFieldConstraintLimit(spec, cluster, field, l.Reference)
	case *constraint.IdentifierLimit:
		sp.findEntityForIdentifierLimit(spec, cluster, field, l)
		sp.resolveFieldConstraintLimit(spec, cluster, field, l.Field)
	case *constraint.MathExpressionLimit:
		sp.resolveFieldConstraintLimit(spec, cluster, field, l.Left)
		sp.resolveFieldConstraintLimit(spec, cluster, field, l.Right)
	case *constraint.ReferenceLimit:
		if l.Entity == nil {
			slog.Info("resolving reference", "ref", l.Reference, log.Path("path", field))
			l.Entity = getCustomDataTypeFromReference(spec, cluster, l.Reference)
		}
		sp.resolveFieldConstraintLimit(spec, cluster, field, l.Field)
	}
}

func (*Builder) findEntityForIdentifierLimit(spec *Specification, cluster *matter.Cluster, field *matter.Field, l *constraint.IdentifierLimit) {
	if l.Entity != nil {
		return
	}
	l.Entity = getCustomDataTypeFromIdentifier(spec, cluster, field, l.ID)
	if l.Entity != nil {
		return
	}
	if cluster != nil {
		for _, a := range cluster.Attributes {
			if strings.EqualFold(a.Name, l.ID) {
				l.Entity = a
				return
			}
		}
	}
	if field.Type != nil {
		var fieldSet matter.FieldSet
		switch e := field.Type.Entity.(type) {
		case *matter.Struct:
			fieldSet = e.Fields
		case *matter.Command:
			fieldSet = e.Fields
		case *matter.Event:
			fieldSet = e.Fields
		case *matter.Enum:
			for _, v := range e.Values {
				if strings.EqualFold(v.Name, l.ID) {
					l.Entity = v
					return
				}
			}
			slog.Warn("Unable to find matching enum for identifier limit", log.Path("source", field), slog.String("identifier", l.ID), slog.String("enum", e.Name), log.Path("enumSource", e))

		case *matter.Bitmap:
			for _, v := range e.Bits {
				if strings.EqualFold(v.Name(), l.ID) {
					l.Entity = v
					return
				}
			}
			slog.Warn("Unable to find matching bit for identifier limit", log.Path("source", field), slog.String("identifier", l.ID), log.Type("bitmap", e.Name))

		case nil:
		default:
			slog.Warn("referenced constraint field has a type without fields", log.Path("source", field), slog.String("identifier", l.ID), log.Type("type", e))
		}
		if len(fieldSet) > 0 {
			childField := fieldSet.GetField(l.ID)
			if childField != nil {
				l.Entity = childField
				return
			}
		}
	}

}

func getCustomDataType(spec *Specification, dataTypeName string, cluster *matter.Cluster, field *matter.Field) (e types.Entity) {
	e = getCustomDataTypeFromFieldReference(spec, cluster, field)
	if e != nil {
		// We have a reference to a data type; use that
		return
	}
	e = getCustomDataTypeFromIdentifier(spec, cluster, field, dataTypeName)
	return
}

func getCustomDataTypeFromIdentifier(spec *Specification, cluster *matter.Cluster, field *matter.Field, identifier string) types.Entity {
	entities := spec.entities[identifier]
	if len(entities) == 0 {
		canonicalName := CanonicalName(identifier)
		if canonicalName != identifier {
			return getCustomDataTypeFromIdentifier(spec, cluster, field, canonicalName)
		}
	} else if len(entities) == 1 {
		for m := range entities {
			return m
		}
	} else {
		return disambiguateDataType(entities, cluster, field)
	}
	return nil
}

func getCustomDataTypeFromFieldReference(spec *Specification, cluster *matter.Cluster, field *matter.Field) (e types.Entity) {
	switch source := field.Type.Source.(type) {
	case *asciidoc.CrossReference:
		return getCustomDataTypeFromReference(spec, cluster, source.ID)
	default:
		return
	}
}

func getCustomDataTypeFromReference(spec *Specification, cluster *matter.Cluster, reference string) (e types.Entity) {
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
			return
		}
	}
	return nil
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
