package constraint

import "github.com/hasty/alchemy/matter"

type GenericConstraint struct {
	Value string
}

func (c *GenericConstraint) AsciiDocString(dataType *matter.DataType) string {
	return c.Value
}

func (c *GenericConstraint) Equal(o matter.Constraint) bool {
	if oc, ok := o.(*GenericConstraint); ok {
		return oc.Value == c.Value
	}
	return false
}

func (c *GenericConstraint) Min(cc *matter.ConstraintContext) (min matter.ConstraintExtreme) {
	return
}

func (c *GenericConstraint) Max(cc *matter.ConstraintContext) (max matter.ConstraintExtreme) {
	return
}
