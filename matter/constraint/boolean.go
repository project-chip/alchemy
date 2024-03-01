package constraint

import (
	"strconv"

	"github.com/hasty/alchemy/matter/types"
)

type BooleanLimit struct {
	Value bool
}

func (c *BooleanLimit) AsciiDocString(dataType *types.DataType) string {
	return strconv.FormatBool(c.Value)
}

func (c *BooleanLimit) DataModelString(dataType *types.DataType) string {
	return strconv.FormatBool(c.Value)
}

func (c *BooleanLimit) Equal(o ConstraintLimit) bool {
	if oc, ok := o.(*BooleanLimit); ok {
		return oc.Value == c.Value
	}
	return false
}

func (c *BooleanLimit) Min(cc Context) (min types.DataTypeExtreme) {
	var val uint64
	if c.Value {
		val = 1
	}
	return types.DataTypeExtreme{
		Type:   types.DataTypeExtremeTypeUInt64,
		Format: types.NumberFormatInt,
		UInt64: val,
	}
}

func (c *BooleanLimit) Max(cc Context) (max types.DataTypeExtreme) {
	return c.Min(cc)
}

func (c *BooleanLimit) Default(cc Context) (max types.DataTypeExtreme) {
	return c.Min(cc)
}

func (c *BooleanLimit) Clone() ConstraintLimit {
	return &BooleanLimit{Value: c.Value}
}
