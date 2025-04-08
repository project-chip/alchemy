package testscript

import (
	"fmt"
	"strings"

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

func checkBitmapRange(field *matter.Field, fieldSet matter.FieldSet, variableName string) (actions []TestAction) {
	if field.Type == nil {
		return
	}
	switch e := field.Type.Entity.(type) {
	case *matter.Bitmap:
		if len(e.Bits) == 0 {
			return
		}
		var t uint64
		var names []string
		for _, b := range e.Bits {
			mask, err := b.Mask()
			if err != nil {
				return
			}
			t |= mask
			names = append(names, b.Name())
		}
		if t == 0 {
			return
		}
		actions = append(actions, &CheckMaxConstraint{
			constraintAction: constraintAction{
				action: action{
					Comments: []string{fmt.Sprintf("Check bitmap value less than or equal to (%s)", strings.Join(names, " | "))},
				},
				Field:    field,
				FieldSet: fieldSet,
				Variable: variableName,
			},
			Constraint: &constraint.MaxConstraint{Maximum: &constraint.HexLimit{Value: t}}})
	}

	return
}
