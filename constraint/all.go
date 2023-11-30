package constraint

import "github.com/hasty/alchemy/matter"

type AllConstraint struct {
	Field *matter.Field
	Value string
}

func (c *AllConstraint) AsciiDocString(dataType *matter.DataType) string {
	return c.Value
}

func (c *AllConstraint) Equal(o matter.Constraint) bool {
	_, ok := o.(*AllConstraint)
	return ok
}

func (c *AllConstraint) Min(cc *matter.ConstraintContext) (min matter.ConstraintExtreme) {
	return
}

func (c *AllConstraint) Max(cc *matter.ConstraintContext) (max matter.ConstraintExtreme) {
	return
}

func (c *AllConstraint) Default(cc *matter.ConstraintContext) (max matter.ConstraintExtreme) {
	return
}
