package constraint

import (
	"github.com/hasty/alchemy/matter/types"
)

type EmptyLimit struct {
}

func (c *EmptyLimit) AsciiDocString(dataType *types.DataType) string {
	return "empty"
}

func (c *EmptyLimit) DataModelString(dataType *types.DataType) string {
	return "empty"
}

func (c *EmptyLimit) Equal(o ConstraintLimit) bool {
	_, ok := o.(*EmptyLimit)
	return ok
}

func (c *EmptyLimit) Min(cc Context) (min types.DataTypeExtreme) {
	return types.DataTypeExtreme{Type: types.DataTypeExtremeTypeEmpty, Format: types.NumberFormatHex}
}

func (c *EmptyLimit) Max(cc Context) (max types.DataTypeExtreme) {
	return c.Min(cc)
}

func (c *EmptyLimit) Default(cc Context) (max types.DataTypeExtreme) {
	return c.Min(cc)
}

func (c *EmptyLimit) Clone() ConstraintLimit {
	return &EmptyLimit{}
}
