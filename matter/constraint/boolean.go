package constraint

import (
	"strconv"

	"github.com/hasty/alchemy/matter"
)

type BooleanLimit struct {
	Value bool
}

func (c *BooleanLimit) AsciiDocString(dataType *matter.DataType) string {
	return strconv.FormatBool(c.Value)
}

func (c *BooleanLimit) Equal(o matter.ConstraintLimit) bool {
	if oc, ok := o.(*BooleanLimit); ok {
		return oc.Value == c.Value
	}
	return false
}

func (c *BooleanLimit) Min(cc *matter.ConstraintContext) (min matter.DataTypeExtreme) {
	var val uint64
	if c.Value {
		val = 1
	}
	return matter.DataTypeExtreme{
		Type:   matter.DataTypeExtremeTypeUInt64,
		Format: matter.NumberFormatInt,
		UInt64: val,
	}
}

func (c *BooleanLimit) Max(cc *matter.ConstraintContext) (max matter.DataTypeExtreme) {
	return c.Min(cc)
}

func (c *BooleanLimit) Default(cc *matter.ConstraintContext) (max matter.DataTypeExtreme) {
	return c.Min(cc)
}

func (c *BooleanLimit) Clone() matter.ConstraintLimit {
	return &BooleanLimit{Value: c.Value}
}
