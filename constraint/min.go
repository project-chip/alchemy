package constraint

import (
	"fmt"

	"github.com/hasty/alchemy/matter"
)

type MinConstraint struct {
	Minimum matter.ConstraintLimit `json:"min"`
}

func (c *MinConstraint) AsciiDocString(dataType *matter.DataType) string {
	return fmt.Sprintf("min %s", c.Minimum.AsciiDocString(dataType))
}

func (c *MinConstraint) Equal(o matter.Constraint) bool {
	if oc, ok := o.(*MinConstraint); ok {
		return oc.Minimum.Equal(c.Minimum)
	}
	return false
}

func (c *MinConstraint) Min(cc *matter.ConstraintContext) (min matter.ConstraintExtreme) {
	return c.Minimum.Min(cc)
}

func (c *MinConstraint) Max(cc *matter.ConstraintContext) (max matter.ConstraintExtreme) {
	return
}
