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
	case *constraint.ListConstraint:
		var acts []TestAction
		acts, err = addConstraintActions(field, fieldSet, c.Constraint, variableName)
		if err != nil {
			return
		}
		actions = append(actions, acts...)
		var entryActs []TestAction
		entryActs, err = addConstraintActions(field, fieldSet, c.EntryConstraint, variableName)
		if err != nil {
			return
		}
		if len(entryActs) > 0 {
			actions = append(actions, &CheckListEntries{
				constraintAction: constraintAction{
					Field:    field,
					FieldSet: fieldSet,
					Variable: variableName,
				},
				Validations: entryActs,
			})
		}
	case *constraint.ExactConstraint:
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
