package constraint

import (
	"fmt"

	"github.com/hasty/alchemy/matter"
)

type ListConstraint struct {
	Constraint      matter.Constraint
	EntryConstraint matter.Constraint
}

func (c *ListConstraint) AsciiDocString() string {
	return fmt.Sprintf("%s[%s]", c.Constraint.AsciiDocString(), c.EntryConstraint.AsciiDocString())
}

func (c *ListConstraint) Equal(o matter.Constraint) bool {
	if oc, ok := o.(*ListConstraint); ok {
		return oc.Constraint.Equal(c.Constraint) && oc.EntryConstraint.Equal(c.EntryConstraint)
	}
	return false
}

func (c *ListConstraint) MinMax(cc *matter.ConstraintContext) (min matter.ConstraintExtreme, max matter.ConstraintExtreme) {
	return c.Constraint.MinMax(cc)
}
