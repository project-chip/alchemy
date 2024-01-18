package constraint

import (
	"github.com/hasty/alchemy/matter/types"
)

type NullLimit struct {
}

func (c *NullLimit) AsciiDocString(dataType *types.DataType) string {
	return "null"
}

func (c *NullLimit) Equal(o ConstraintLimit) bool {
	_, ok := o.(*NullLimit)
	return ok
}

func (c *NullLimit) Min(cc Context) (min types.DataTypeExtreme) {
	return types.DataTypeExtreme{Type: types.DataTypeExtremeTypeNull, Format: types.NumberFormatAuto}
}

func (c *NullLimit) Max(cc Context) (max types.DataTypeExtreme) {
	return c.Min(cc)
}

func (c *NullLimit) Default(cc Context) (max types.DataTypeExtreme) {
	return c.Min(cc)
}

func (c *NullLimit) Clone() ConstraintLimit {
	return &NullLimit{}
}
