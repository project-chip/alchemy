package constraint

import (
	"fmt"

	"github.com/hasty/alchemy/matter/types"
)

type ListConstraint struct {
	Constraint      Constraint
	EntryConstraint Constraint
}

func (c *ListConstraint) Type() ConstraintType {
	return ConstraintTypeList
}

func (c *ListConstraint) AsciiDocString(dataType *types.DataType) string {
	return fmt.Sprintf("%s[%s]", c.Constraint.AsciiDocString(dataType), c.EntryConstraint.AsciiDocString(dataType))
}

func (c *ListConstraint) Equal(o Constraint) bool {
	if oc, ok := o.(*ListConstraint); ok {
		return oc.Constraint.Equal(c.Constraint) && oc.EntryConstraint.Equal(c.EntryConstraint)
	}
	return false
}

func (c *ListConstraint) Min(cc Context) (min types.DataTypeExtreme) {
	return c.Constraint.Min(cc)
}

func (c *ListConstraint) Max(cc Context) (max types.DataTypeExtreme) {
	return c.Constraint.Max(cc)
}

func (c *ListConstraint) Default(cc Context) (max types.DataTypeExtreme) {
	return
}

func (c *ListConstraint) Clone() Constraint {
	return &ListConstraint{Constraint: c.Constraint.Clone(), EntryConstraint: c.EntryConstraint.Clone()}
}
