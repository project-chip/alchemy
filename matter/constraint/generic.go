package constraint

import (
	"github.com/hasty/alchemy/matter/types"
)

type GenericConstraint struct {
	Value string
}

func (c *GenericConstraint) Type() ConstraintType {
	return ConstraintTypeGeneric
}

func (c *GenericConstraint) AsciiDocString(dataType *types.DataType) string {
	return c.Value
}

func (c *GenericConstraint) Equal(o Constraint) bool {
	if oc, ok := o.(*GenericConstraint); ok {
		return oc.Value == c.Value
	}
	return false
}

func (c *GenericConstraint) Min(cc Context) (min types.DataTypeExtreme) {
	return
}

func (c *GenericConstraint) Max(cc Context) (max types.DataTypeExtreme) {
	return
}

func (c *GenericConstraint) Default(cc Context) (max types.DataTypeExtreme) {
	return
}

func (c *GenericConstraint) Clone() Constraint {
	return &GenericConstraint{Value: c.Value}
}
