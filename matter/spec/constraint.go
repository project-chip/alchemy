package spec

import (
	"log/slog"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/shopspring/decimal"
)

func (sp *Builder) resolveConstraints() {
	specEntityFinder := newSpecEntityFinder(sp.Spec, nil, nil)
	for cluster := range sp.Spec.Clusters {
		specEntityFinder.cluster = cluster
		for _, a := range cluster.Attributes {
			sp.resolveFieldConstraints(cluster.Attributes, a, a.Type, specEntityFinder)
		}
		attributeFinder := newFieldFinder(cluster.Attributes, specEntityFinder)
		for _, s := range cluster.Structs {
			structFinder := newFieldFinder(s.Fields, attributeFinder)
			for _, f := range s.Fields {
				sp.resolveFieldConstraints(s.Fields, f, f.Type, structFinder)
			}
		}
		commandFinder := newCommandFinder(cluster.Commands, attributeFinder)
		for _, command := range cluster.Commands {
			for _, f := range command.Fields {
				sp.resolveFieldConstraints(command.Fields, f, f.Type, commandFinder)
			}
		}
		eventFinder := newEventFinder(cluster.Events, attributeFinder)
		for _, event := range cluster.Events {
			for _, f := range event.Fields {
				sp.resolveFieldConstraints(event.Fields, f, f.Type, eventFinder)
			}
		}
	}
	specEntityFinder.cluster = nil
	for o := range sp.Spec.GlobalObjects {
		switch o := o.(type) {
		case *matter.Struct:
			for _, f := range o.Fields {
				sp.resolveFieldConstraints(o.Fields, f, f.Type, specEntityFinder)
			}
		}
	}
	for _, deviceType := range sp.Spec.DeviceTypes {
		conditionFinder := newConditionFinder(deviceType, sp.Spec.BaseDeviceType, specEntityFinder)
		for _, er := range deviceType.ElementRequirements {
			sp.resolveFieldConstraintReferences(conditionFinder, er, er.Constraint)
		}
		for _, dtr := range deviceType.DeviceTypeRequirements {
			sp.resolveFieldConstraintReferences(conditionFinder, dtr, dtr.Constraint)
		}
	}
}

func (sp *Builder) resolveFieldConstraints(fieldSet matter.FieldSet, field *matter.Field, dataType *types.DataType, finder entityFinder) {
	var fieldFinder entityFinder = newFieldFinder(fieldSet, finder)
	if dataType != nil && dataType.Entity != nil {
		dataTypeFinder := makeEntityFinder(dataType.Entity, fieldFinder)
		if dataTypeFinder != nil {
			fieldFinder = dataTypeFinder
		}
	}
	sp.resolveFieldConstraintReferences(fieldFinder, field, field.Constraint)
	sp.resolveFieldConstraintLimit(fieldFinder, field, field.Fallback)
}

func (sp *Builder) resolveFieldConstraintReferences(finder entityFinder, source types.Entity, con constraint.Constraint) {
	for _, limit := range constraint.TraverseConstraintLimits(con) {
		sp.resolveFieldConstraintLimit(finder, source, limit)
	}
}

func (sp *Builder) resolveFieldConstraintLimit(finder entityFinder, source types.Entity, limit constraint.Limit) {
	switch limit := limit.(type) {
	case *constraint.IdentifierLimit:
		if limit.Entity == nil {
			limit.Entity = finder.findEntityByIdentifier(limit.ID, source)
		}
		if limit.Entity != nil && limit.Field != nil {
			f := makeEntityFinder(limit.Entity, finder)
			for l := range constraint.TraverseLimit(limit.Field) {
				sp.resolveFieldConstraintLimit(f, source, l)
			}
		}
	case *constraint.ReferenceLimit:
		if limit.Entity == nil {
			limit.Entity = finder.findEntityByReference(limit.Reference, limit.Label, source)
		}
		if limit.Entity != nil && limit.Field != nil {
			f := makeEntityFinder(limit.Entity, finder)
			for l := range constraint.TraverseLimit(limit.Field) {
				sp.resolveFieldConstraintLimit(f, source, l)
			}
		}
	}
}

func validateConstraint(spec *Specification, field *matter.Field) {
	for _, l := range constraint.TraverseConstraintLimits(field.Constraint) {
		validateLimit(spec, field, l)
	}
}

func isLimitValid(field *matter.Field, limit constraint.Limit) bool {
	switch l := limit.(type) {
	case *constraint.NullLimit:
		if !field.Quality.Has(matter.QualityNullable) {
			return false
		}
	case *constraint.TemperatureLimit:
		switch field.Type.BaseType {
		case types.BaseDataTypeTemperature:
			if !l.Value.Shift(2).IsInteger() {
				return false
			}
			if l.Value.GreaterThan(decimal.New(32767, -2)) || l.Value.LessThan(decimal.New(-32767, -2)) {
				return false
			}
		case types.BaseDataTypeTemperatureDifference:
			if !l.Value.Shift(2).IsInteger() {
				return false
			}
			if l.Value.GreaterThan(decimal.New(32767, -2)) || l.Value.LessThan(decimal.New(-27315, -2)) {
				return false
			}
		case types.BaseDataTypeUnsignedTemperature:
			if !l.Value.Shift(1).IsInteger() {
				return false
			}
			if l.Value.GreaterThan(decimal.New(255, 0)) || l.Value.LessThan(decimal.New(0, 0)) {
				return false
			}
		case types.BaseDataTypeSignedTemperature:
			if !l.Value.Shift(1).IsInteger() {
				return false
			}
			if l.Value.GreaterThan(decimal.New(127, 0)) || l.Value.LessThan(decimal.New(-128, 0)) {
				return false
			}
		default:
			return false
		}
	case *constraint.PercentLimit:
		switch field.Type.BaseType {
		case types.BaseDataTypePercent:
			if !l.Value.IsInteger() {
				return false
			}
		case types.BaseDataTypePercentHundredths:
			if !l.Value.Shift(2).IsInteger() {
				return false
			}
		default:
			return false
		}
	}
	return true
}

func validateLimit(spec *Specification, field *matter.Field, limit constraint.Limit) {
	if !isLimitValid(field, limit) {
		spec.addError(&InvalidConstraintLimitError{Limit: limit, Field: field, Source: field})
	}
	switch limit := limit.(type) {
	case *constraint.IdentifierLimit:
		if limit.Entity == nil {
			slog.Error("Failed to resolve constraint identifier", "ref", limit.ID, log.Path("source", field))
			//suggestions := make(map[types.Entity]int)
			//failure.finder.suggestIdentifiers(limit.ID, suggestions)
			//suggest.ListPossibilities(limit.ID, suggestions)
			spec.addError(&UnknownConstraintIdentifierError{Identifier: limit, Source: field})
		}
	case *constraint.ReferenceLimit:
		if limit.Entity == nil {
			slog.Error("Failed to resolve constraint reference", "ref", limit.Reference, log.Path("source", field))
			spec.addError(&UnknownConstraintReferenceError{Reference: limit, Source: field})
		}
	}
}

func validateFallback(spec *Specification, field *matter.Field) {
	if field.Fallback != nil {
		for l := range constraint.TraverseLimit(field.Fallback) {
			validateLimit(spec, field, l)
		}
	}
}
