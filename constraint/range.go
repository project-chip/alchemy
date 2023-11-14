package constraint

import (
	"fmt"

	"github.com/hasty/alchemy/matter"
)

type RangeConstraint struct {
	Min matter.ConstraintLimit `json:"min"`
	Max matter.ConstraintLimit `json:"max"`
}

func (c *RangeConstraint) AsciiDocString() string {
	return fmt.Sprintf("%s to %s", c.Min.AsciiDocString(), c.Max.AsciiDocString())
}

func (c *RangeConstraint) Equal(o matter.Constraint) bool {
	if oc, ok := o.(*RangeConstraint); ok {
		return oc.Min.Equal(c.Min) && oc.Max.Equal(c.Max)
	}
	return false
}

func (c *RangeConstraint) MinMax(cc *matter.ConstraintContext) (from matter.ConstraintExtreme, to matter.ConstraintExtreme) {
	from, _ = c.Min.MinMax(cc)
	_, to = c.Max.MinMax(cc)

	return
}
