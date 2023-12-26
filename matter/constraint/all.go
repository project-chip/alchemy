package constraint

import "github.com/hasty/alchemy/matter"

type AllConstraint struct {
	Value string
}

func NewAllConstraint(value string) *AllConstraint {
	return &AllConstraint{Value: value}
}

func (c *AllConstraint) Type() matter.ConstraintType {
	return matter.ConstraintTypeAll
}

func (c *AllConstraint) AsciiDocString(dataType *matter.DataType) string {
	return c.Value
}

func (c *AllConstraint) Equal(o matter.Constraint) bool {
	_, ok := o.(*AllConstraint)
	return ok
}

func (c *AllConstraint) Min(cc *matter.ConstraintContext) (min matter.DataTypeExtreme) {
	return
}

func (c *AllConstraint) Max(cc *matter.ConstraintContext) (max matter.DataTypeExtreme) {
	return
}

func (c *AllConstraint) Default(cc *matter.ConstraintContext) (max matter.DataTypeExtreme) {
	return
}

func (c *AllConstraint) Clone() matter.Constraint {
	return &AllConstraint{}
}
