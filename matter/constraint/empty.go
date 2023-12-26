package constraint

import (
	"github.com/hasty/alchemy/matter"
)

type EmptyLimit struct {
}

func (c *EmptyLimit) AsciiDocString(dataType *matter.DataType) string {
	return "empty"
}

func (c *EmptyLimit) Equal(o matter.ConstraintLimit) bool {
	_, ok := o.(*EmptyLimit)
	return ok
}

func (c *EmptyLimit) Min(cc *matter.ConstraintContext) (min matter.DataTypeExtreme) {
	return matter.DataTypeExtreme{Type: matter.DataTypeExtremeTypeEmpty, Format: matter.NumberFormatHex}
}

func (c *EmptyLimit) Max(cc *matter.ConstraintContext) (max matter.DataTypeExtreme) {
	return c.Min(cc)
}

func (c *EmptyLimit) Default(cc *matter.ConstraintContext) (max matter.DataTypeExtreme) {
	return c.Min(cc)
}

func (c *EmptyLimit) Clone() matter.ConstraintLimit {
	return &EmptyLimit{}
}
