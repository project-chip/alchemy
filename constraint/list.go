package constraint

import (
	"fmt"

	"github.com/hasty/alchemy/matter"
)

type ListConstraint struct {
	Constraint      matter.Constraint
	EntryConstraint matter.Constraint
}

func (c *ListConstraint) AsciiDocString(dataType *matter.DataType) string {
	return fmt.Sprintf("%s[%s]", c.Constraint.AsciiDocString(dataType), c.EntryConstraint.AsciiDocString(dataType))
}

func (c *ListConstraint) Equal(o matter.Constraint) bool {
	if oc, ok := o.(*ListConstraint); ok {
		return oc.Constraint.Equal(c.Constraint) && oc.EntryConstraint.Equal(c.EntryConstraint)
	}
	return false
}

func (c *ListConstraint) Min(cc *matter.ConstraintContext) (min matter.ConstraintExtreme) {
	return c.Constraint.Min(cc)
}

func (c *ListConstraint) Max(cc *matter.ConstraintContext) (max matter.ConstraintExtreme) {
	return c.Constraint.Max(cc)
}
