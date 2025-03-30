package testscript

import (
	"fmt"

	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/constraint"
)

func addConstraintActions(field *matter.Field, fieldSet matter.FieldSet, c constraint.Constraint, variableName string) (actions []TestAction, err error) {
	if constraint.IsBlank(c) {
		return
	}
	switch c := c.(type) {
	case *constraint.MinConstraint:
		actions = append(actions, &CheckMinConstraint{
			constraintAction: constraintAction{
				Field:    field,
				FieldSet: fieldSet,
				Variable: variableName,
			},
			Constraint: c})
	case *constraint.MaxConstraint:
		actions = append(actions, &CheckMaxConstraint{
			constraintAction: constraintAction{
				Field:    field,
				FieldSet: fieldSet,
				Variable: variableName,
			},
			Constraint: c})
	case *constraint.RangeConstraint:
		actions = append(actions, &CheckRangeConstraint{
			constraintAction: constraintAction{
				Field:    field,
				FieldSet: fieldSet,
				Variable: variableName,
			},
			Constraint: c})
	case *constraint.AllConstraint, *constraint.DescribedConstraint, *constraint.GenericConstraint:
		return
	case constraint.Set:
		for _, c := range c {
			var acts []TestAction
			acts, err = addConstraintActions(field, fieldSet, c, variableName)
			if err != nil {
				return
			}
			actions = append(actions, acts...)
		}
	default:
		err = fmt.Errorf("unexpected constraint type setting test step response checks: %T", c)
		return
	}
	return
}

/*func setConstraintChecks(step *TestStep, field *matter.Field, c constraint.Constraint) (err error) {
	if constraint.IsBlank(c) {
		return
	}
	switch c := c.(type) {
	case *constraint.MinConstraint:
		if field.Type.IsArray() {
			ensureConstraints(step).MinLength = setLimitString(field, c.Minimum)
		} else {
			ensureConstraints(step).MinValue = setLimitString(field, c.Minimum)
		}
	case *constraint.MaxConstraint:
		if field.Type.IsArray() {
			ensureConstraints(step).MaxLength = setLimitString(field, c.Maximum)
		} else {
			ensureConstraints(step).MaxValue = setLimitString(field, c.Maximum)
		}
	case *constraint.RangeConstraint:
		if field.Type.IsArray() {
			ensureConstraints(step).MinLength = setLimitString(field, c.Minimum)
			ensureConstraints(step).MaxLength = setLimitString(field, c.Maximum)
		} else {
			ensureConstraints(step).MinValue = setLimitString(field, c.Minimum)
			ensureConstraints(step).MaxValue = setLimitString(field, c.Maximum)
		}
	case *constraint.AllConstraint, *constraint.DescribedConstraint, *constraint.GenericConstraint:
		return
	case constraint.Set:
		for _, c := range c {
			setConstraintChecks(step, field, c)
		}
	default:
		err = fmt.Errorf("unexpected constraint type setting test step response checks: %T", c)
		return
	}
	return
}

func ensureConstraints(step *testplan.Step) *parse.StepResponseConstraints {
	if step.Response.Constraints == nil {
		step.Response.Constraints = &parse.StepResponseConstraints{}
	}
	return step.Response.Constraints
}

func setLimitString(field *matter.Field, limit constraint.Limit) string {
	switch limit := limit.(type) {
	case *constraint.BooleanLimit:
		if limit.Value {
			return "True"
		}
		return "False"
	case *constraint.IdentifierLimit:
		if limit.Entity != nil {
			switch entity := limit.Entity.(type) {
			case *matter.Field:
				switch entity.EntityType() {
				case types.EntityTypeAttribute:
					return "self." + entity.Name
				case types.EntityTypeStructField:
					return "struct." + entity.Name
				default:
					slog.Warn("Unexpected field type in constraint", slog.String("type", entity.EntityType().String()))
				}
			default:
				slog.Warn("Unexpected entity type in constraint", log.Type("type", entity))
			}
		}
	case *constraint.ReferenceLimit:
		if limit.Entity != nil {
			switch entity := limit.Entity.(type) {
			case *matter.Field:
				switch entity.EntityType() {
				case types.EntityTypeAttribute:
					return "self." + entity.Name
				case types.EntityTypeStructField:
					return "struct." + entity.Name
				default:
					slog.Warn("Unexpected field type in constraint", slog.String("type", entity.EntityType().String()))
				}
			default:
				slog.Warn("Unexpected entity type in constraint", log.Type("type", entity))
			}
		}
	default:
		slog.Warn("Unexpected limit type in constraint", log.Type("type", limit))
	}
	return limit.DataModelString(field.Type)

}
*/
