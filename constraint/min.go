package constraint

import (
	"fmt"

	"github.com/hasty/alchemy/matter"
)

type MinConstraint struct {
	Min matter.ConstraintLimit `json:"min"`
}

func (c *MinConstraint) AsciiDocString(dataType *matter.DataType) string {
	return fmt.Sprintf("min %s", c.Min.AsciiDocString(dataType))
}

func (c *MinConstraint) Equal(o matter.Constraint) bool {
	if oc, ok := o.(*MinConstraint); ok {
		return oc.Min.Equal(c.Min)
	}
	return false
}

func (c *MinConstraint) MinMax(cc *matter.ConstraintContext) (min matter.ConstraintExtreme, max matter.ConstraintExtreme) {
	min, _ = c.Min.MinMax(cc)
	return
}
