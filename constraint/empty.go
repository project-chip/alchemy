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

func (c *EmptyLimit) Min(cc *matter.ConstraintContext) (min matter.ConstraintExtreme) {
	return matter.ConstraintExtreme{Type: matter.ConstraintExtremeTypeEmpty, Format: matter.NumberFormatHex}
}

func (c *EmptyLimit) Max(cc *matter.ConstraintContext) (max matter.ConstraintExtreme) {
	return c.Min(cc)
}

func (c *EmptyLimit) Default(cc *matter.ConstraintContext) (max matter.ConstraintExtreme) {
	return c.Min(cc)
}
