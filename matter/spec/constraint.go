package spec

import (
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/types"
)

func (sp *Builder) resolveConstraints(spec *Specification) {
	for cluster := range spec.Clusters {
		for _, a := range cluster.Attributes {
			sp.resolveFieldConstraints(spec, cluster, cluster.Attributes, a, a.Type)
		}
		for _, s := range cluster.Structs {
			for _, f := range s.Fields {
				sp.resolveFieldConstraints(spec, cluster, s.Fields, f, f.Type)
			}
		}
		for _, command := range cluster.Commands {
			for _, f := range command.Fields {
				sp.resolveFieldConstraints(spec, cluster, command.Fields, f, f.Type)
			}
		}
		for _, event := range cluster.Events {
			for _, f := range event.Fields {
				sp.resolveFieldConstraints(spec, cluster, event.Fields, f, f.Type)
			}
		}
	}
	for _, s := range spec.structIndex {
		if s.Parent() == nil {
			for _, f := range s.Fields {
				sp.resolveFieldConstraints(spec, nil, s.Fields, f, f.Type)
			}
		}
	}
	for _, deviceType := range spec.DeviceTypes {
		conditionFinder := func(identifier string) types.Entity {
			for _, con := range deviceType.Conditions {
				if strings.EqualFold(con.Feature, identifier) {
					return con
				}
			}
			return nil
		}
		for _, er := range deviceType.ElementRequirements {
			resolveFieldConstraintReferences(spec, er.Cluster, conditionFinder, er, nil, er.Constraint)
		}
		for _, dtr := range deviceType.DeviceTypeRequirements {
			resolveFieldConstraintReferences(spec, nil, conditionFinder, dtr, nil, dtr.Constraint)
		}
	}
}

func (sp *Builder) resolveFieldConstraints(spec *Specification, cluster *matter.Cluster, fieldSet matter.FieldSet, field *matter.Field, dataType *types.DataType) {
	fieldFinder := func(identifier string) types.Entity {
		for _, of := range fieldSet {
			if strings.EqualFold(of.Name, identifier) {
				return of
			}
		}
		return nil
	}
	var fieldDataTypeEntity types.Entity
	if dataType != nil {
		fieldDataTypeEntity = dataType.Entity
	}
	resolveFieldConstraintReferences(spec, cluster, fieldFinder, field, fieldDataTypeEntity, field.Constraint)
	resolveFieldConstraintLimit(spec, cluster, fieldFinder, field, fieldDataTypeEntity, field.Fallback)
}

func resolveFieldConstraintReferences(spec *Specification, cluster *matter.Cluster, finder findNamedEntity, source log.Source, entity types.Entity, con constraint.Constraint) {
	switch con := con.(type) {
	case *constraint.ExactConstraint:
		resolveFieldConstraintLimit(spec, cluster, finder, source, entity, con.Value)
	case *constraint.ListConstraint:
		resolveFieldConstraintReferences(spec, cluster, finder, source, entity, con.Constraint)
		resolveFieldConstraintReferences(spec, cluster, finder, source, entity, con.EntryConstraint)
	case *constraint.MaxConstraint:
		resolveFieldConstraintLimit(spec, cluster, finder, source, entity, con.Maximum)
	case *constraint.MinConstraint:
		resolveFieldConstraintLimit(spec, cluster, finder, source, entity, con.Minimum)
	case *constraint.RangeConstraint:
		resolveFieldConstraintLimit(spec, cluster, finder, source, entity, con.Minimum)
		resolveFieldConstraintLimit(spec, cluster, finder, source, entity, con.Maximum)
	case constraint.Set:
		for _, c := range con {
			resolveFieldConstraintReferences(spec, cluster, finder, source, entity, c)
		}
	}
}

func resolveFieldConstraintLimit(spec *Specification, cluster *matter.Cluster, finder findNamedEntity, source log.Source, entity types.Entity, l constraint.Limit) {
	switch l := l.(type) {
	case *constraint.CharacterLimit:
		resolveFieldConstraintLimit(spec, cluster, finder, source, entity, l.ByteCount)
		resolveFieldConstraintLimit(spec, cluster, finder, source, entity, l.CodepointCount)
	case *constraint.LengthLimit:
		resolveFieldConstraintLimit(spec, cluster, finder, source, entity, l.Reference)
	case *constraint.IdentifierLimit:
		if l.Entity == nil {
			l.Entity = findEntityForFieldIdentifier(spec, cluster, finder, source, entity, l.ID)
			if l.Entity == nil {
				slog.Error("failed to resolve constraint identifier", "ref", l.ID, log.Path("path", source))
			}
		}
		if l.Entity != nil && l.Field != nil {
			resolveFieldConstraintLimit(spec, cluster, makeFinder(l.Entity), source, l.Entity, l.Field)
		}
	case *constraint.MathExpressionLimit:
		resolveFieldConstraintLimit(spec, cluster, finder, source, entity, l.Left)
		resolveFieldConstraintLimit(spec, cluster, finder, source, entity, l.Right)
	case *constraint.ReferenceLimit:
		if l.Entity == nil {
			l.Entity = getCustomDataTypeFromReference(spec, cluster, l.Reference, l.Label)
			if l.Entity == nil {
				slog.Error("failed to resolve constraint reference", "ref", l.Reference, log.Path("path", source))
			}
		}
		if l.Entity != nil && l.Field != nil {
			resolveFieldConstraintLimit(spec, cluster, makeFinder(l.Entity), source, l.Entity, l.Field)
		}
	case nil, *constraint.ManufacturerLimit, *constraint.IntLimit, *constraint.NullLimit, *constraint.StatusCodeLimit, *constraint.EmptyLimit, *constraint.StringLimit, *constraint.GenericLimit, *constraint.BooleanLimit, *constraint.TemperatureLimit, *constraint.UnspecifiedLimit, *constraint.ExpLimit, *constraint.HexLimit, *constraint.PercentLimit:
	default:
		slog.Warn("Unexpected field constraint limit type", log.Type("type", l))
	}
}

type findNamedEntity func(name string) types.Entity

func makeFinder(entity types.Entity) findNamedEntity {
	switch entity := entity.(type) {

	case *matter.Bitmap:
		return func(identifier string) types.Entity {
			for _, bmv := range entity.Bits {
				if strings.EqualFold(bmv.Name(), identifier) {
					return bmv
				}
			}
			return nil
		}
	case *matter.Enum:
		return func(identifier string) types.Entity {
			for _, ev := range entity.Values {
				if strings.EqualFold(ev.Name, identifier) {
					return ev
				}
			}
			return nil
		}
	case *matter.Struct:
		return func(identifier string) types.Entity {
			for _, ev := range entity.Fields {
				if strings.EqualFold(ev.Name, identifier) {
					return ev
				}
			}
			return nil
		}
	case *matter.Command:
		return func(identifier string) types.Entity {
			for _, ev := range entity.Fields {
				if strings.EqualFold(ev.Name, identifier) {
					return ev
				}
			}
			return nil
		}
	case *matter.Event:
		return func(identifier string) types.Entity {
			for _, ev := range entity.Fields {
				if strings.EqualFold(ev.Name, identifier) {
					return ev
				}
			}
			return nil
		}
	case *matter.Field:
		return makeFinder(entity.Type.Entity)

	default:
		slog.Warn("Unable to make named entry finder for type", log.Type("type", entity))
		return nil
	}
}

func findEntityForFieldIdentifier(spec *Specification, cluster *matter.Cluster, finder findNamedEntity, source log.Source, entity types.Entity, identifier string) (e types.Entity) {

	if spec.BaseDeviceType != nil {
		for _, c := range spec.BaseDeviceType.Conditions {
			if strings.EqualFold(c.Feature, identifier) {
				e = c
				return
			}
		}
	}
	if finder != nil {
		e = finder(identifier)
		if e != nil {
			return
		}
	}

	e = getCustomDataTypeFromIdentifier(spec, cluster, source, identifier)
	if e != nil {
		return
	}

	field, isField := entity.(*matter.Field)
	if isField {
		entity = field.Type.Entity
	}
	if entity != nil {
		var fieldSet matter.FieldSet
		switch entity := entity.(type) {
		case *matter.Struct:
			fieldSet = entity.Fields
		case *matter.Command:
			fieldSet = entity.Fields
		case *matter.Event:
			fieldSet = entity.Fields
		case *matter.Enum:
			for _, v := range entity.Values {
				if strings.EqualFold(v.Name, identifier) {
					e = v
					return
				}
			}
			slog.Warn("Unable to find matching enum for identifier limit", log.Path("source", source), slog.String("identifier", identifier), slog.String("enum", entity.Name), log.Path("enumSource", entity))

		case *matter.Bitmap:
			for _, v := range entity.Bits {
				if strings.EqualFold(v.Name(), identifier) {
					e = v
					return
				}
			}
			slog.Warn("Unable to find matching bit for identifier limit", log.Path("source", source), slog.String("identifier", identifier), log.Type("bitmap", entity.Name))

		default:
			slog.Warn("referenced constraint field has a type without fields", log.Path("source", source), slog.String("identifier", identifier), log.Type("type", entity))
		}
		if len(fieldSet) > 0 {
			childField := fieldSet.Get(identifier)
			if childField != nil {
				e = childField
				return
			}
		}
	}
	if cluster != nil {
		if cluster.Features != nil {
			for fb := range cluster.Features.FeatureBits() {
				if strings.EqualFold(fb.Code, identifier) {
					e = fb
					return
				}
			}
		}
		for _, a := range cluster.Attributes {
			if strings.EqualFold(a.Name, identifier) {
				e = a
				return
			}
		}
	}
	if spec.BaseDeviceType != nil {
		for _, con := range spec.BaseDeviceType.Conditions {
			if strings.EqualFold(identifier, con.Feature) {
				e = con
				return
			}
		}
	}
	return
}
