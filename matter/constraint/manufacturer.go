package constraint

import (
	"github.com/hasty/alchemy/matter/types"
)

type ManufacturerLimit struct {
	Value string
}

func (c *ManufacturerLimit) AsciiDocString(dataType *types.DataType) string {
	return c.Value
}

func (c *ManufacturerLimit) Equal(o ConstraintLimit) bool {
	if oc, ok := o.(*ManufacturerLimit); ok {
		return oc.Value == c.Value
	}
	return false
}

func (c *ManufacturerLimit) Min(cc Context) (min types.DataTypeExtreme) {
	return
}

func (c *ManufacturerLimit) Max(cc Context) (max types.DataTypeExtreme) {
	return
}

func (c *ManufacturerLimit) Default(cc Context) (max types.DataTypeExtreme) {
	return
}

func (c *ManufacturerLimit) Clone() ConstraintLimit {
	return &ManufacturerLimit{Value: c.Value}
}
