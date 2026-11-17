package spec

import (
	"log/slog"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/shopspring/decimal"
)

func validateConstraint(spec *Specification, source types.Entity, c constraint.Constraint) {
	for c := range constraint.Traverse(c) {
		switch c := c.(type) {
		case *constraint.GenericConstraint:
			if c.Value != "" {
				spec.addError(&InvalidConstraintError{Constraint: c.Value, Source: source})
			}
		default:
			for _, l := range constraint.TraverseConstraintLimits(c) {
				validateLimit(spec, source, l)
			}
		}
	}
}

func validateFieldConstraint(spec *Specification, field *matter.Field) {
	if field.Constraint == nil {
		return
	}
	validateConstraint(spec, field, field.Constraint)
}

func isFieldLimitValid(field *matter.Field, limit constraint.Limit) bool {
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

func validateFieldLimit(spec *Specification, field *matter.Field, limit constraint.Limit) {
	if !isFieldLimitValid(field, limit) {
		spec.addError(&InvalidConstraintLimitError{Limit: limit, Field: field, Source: field})
	}
	validateLimit(spec, field, limit)
}

func validateLimit(spec *Specification, source types.Entity, limit constraint.Limit) {
	switch limit := limit.(type) {
	case *constraint.IdentifierLimit:
		if limit.Entity == nil {
			slog.Error("Failed to resolve constraint identifier", "ref", limit.ID, log.Path("source", source))
			spec.addError(&UnknownConstraintIdentifierError{Identifier: limit, Source: source})
		}
	case *constraint.ReferenceLimit:
		if limit.Entity == nil {
			slog.Error("Failed to resolve constraint reference", "ref", limit.Reference, log.Path("source", source))
			spec.addError(&UnknownConstraintReferenceError{Reference: limit, Source: source})
		}
	}
}

func validateFallback(spec *Specification, field *matter.Field) {
	if field.Fallback == nil {
		return
	}
	for l := range constraint.TraverseLimit(field.Fallback) {
		validateFieldLimit(spec, field, l)
	}
}
