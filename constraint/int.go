package constraint

import (
	"strconv"

	"github.com/hasty/alchemy/matter"
)

type IntLimit struct {
	Value int64
}

func (c *IntLimit) AsciiDocString(dataType *matter.DataType) string {
	return strconv.FormatInt(c.Value, 10)
}

func (c *IntLimit) Equal(o matter.ConstraintLimit) bool {
	if oc, ok := o.(*IntLimit); ok {
		return oc.Value == c.Value
	}
	return false
}

func (c *IntLimit) Min(cc *matter.ConstraintContext) (min matter.ConstraintExtreme) {
	return matter.ConstraintExtreme{
		Type:   matter.ConstraintExtremeTypeInt64,
		Format: matter.NumberFormatInt,
		Int64:  c.Value,
	}
}

func (c *IntLimit) Max(cc *matter.ConstraintContext) (max matter.ConstraintExtreme) {
	return c.Min(cc)
}

func (c *IntLimit) Default(cc *matter.ConstraintContext) (max matter.ConstraintExtreme) {
	return c.Min(cc)
}
