package constraint

import "github.com/hasty/alchemy/matter"

type ManufacturerLimit struct {
	Value string
}

func (c *ManufacturerLimit) AsciiDocString(dataType *matter.DataType) string {
	return c.Value
}

func (c *ManufacturerLimit) Equal(o matter.ConstraintLimit) bool {
	if oc, ok := o.(*ManufacturerLimit); ok {
		return oc.Value == c.Value
	}
	return false
}

func (c *ManufacturerLimit) Min(cc *matter.ConstraintContext) (min matter.DataTypeExtreme) {
	return
}

func (c *ManufacturerLimit) Max(cc *matter.ConstraintContext) (max matter.DataTypeExtreme) {
	return
}

func (c *ManufacturerLimit) Default(cc *matter.ConstraintContext) (max matter.DataTypeExtreme) {
	return
}

func (c *ManufacturerLimit) Clone() matter.ConstraintLimit {
	return &ManufacturerLimit{Value: c.Value}
}
