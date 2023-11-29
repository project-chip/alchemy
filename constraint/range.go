package constraint

import (
	"fmt"

	"github.com/hasty/alchemy/matter"
)

type RangeConstraint struct {
	Minimum matter.ConstraintLimit `json:"min"`
	Maximum matter.ConstraintLimit `json:"max"`
}

func (c *RangeConstraint) AsciiDocString(dataType *matter.DataType) string {
	return fmt.Sprintf("%s to %s", c.Minimum.AsciiDocString(dataType), c.Maximum.AsciiDocString(dataType))
}

func (c *RangeConstraint) Equal(o matter.Constraint) bool {
	if oc, ok := o.(*RangeConstraint); ok {
		return oc.Minimum.Equal(c.Minimum) && oc.Maximum.Equal(c.Maximum)
	}
	return false
}

func (c *RangeConstraint) Min(cc *matter.ConstraintContext) (from matter.ConstraintExtreme) {
	return c.Minimum.Min(cc)
}

func (c *RangeConstraint) Max(cc *matter.ConstraintContext) (to matter.ConstraintExtreme) {
	return c.Maximum.Max(cc)

}