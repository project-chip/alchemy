package constraint

import "github.com/hasty/alchemy/matter"

type AllConstraint struct {
	Field *matter.Field
	Value string
}

func (c *AllConstraint) AsciiDocString() string {
	return c.Value
}

func (c *AllConstraint) Equal(o matter.Constraint) bool {
	_, ok := o.(*AllConstraint)
	return ok
}

func (c *AllConstraint) MinMax(cc *matter.ConstraintContext) (min matter.ConstraintExtreme, max matter.ConstraintExtreme) {
	return
}
