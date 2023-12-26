package constraint

import (
	"github.com/hasty/alchemy/matter"
)

type NullLimit struct {
}

func (c *NullLimit) AsciiDocString(dataType *matter.DataType) string {
	return "null"
}

func (c *NullLimit) Equal(o matter.ConstraintLimit) bool {
	_, ok := o.(*NullLimit)
	return ok
}

func (c *NullLimit) Min(cc *matter.ConstraintContext) (min matter.DataTypeExtreme) {
	return matter.DataTypeExtreme{Type: matter.DataTypeExtremeTypeNull, Format: matter.NumberFormatInt}
}

func (c *NullLimit) Max(cc *matter.ConstraintContext) (max matter.DataTypeExtreme) {
	return c.Min(cc)
}

func (c *NullLimit) Default(cc *matter.ConstraintContext) (max matter.DataTypeExtreme) {
	return c.Min(cc)
}

func (c *NullLimit) Clone() matter.ConstraintLimit {
	return &NullLimit{}
}
