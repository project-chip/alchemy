package constraint

import (
	"fmt"

	"github.com/hasty/alchemy/matter"
)

type MaxConstraint struct {
	Max matter.ConstraintLimit
}

func (c *MaxConstraint) AsciiDocString(dataType *matter.DataType) string {
	return fmt.Sprintf("max %s", c.Max.AsciiDocString(dataType))
}

func (c *MaxConstraint) Equal(o matter.Constraint) bool {
	if oc, ok := o.(*MaxConstraint); ok {
		return oc.Max.Equal(c.Max)
	}
	return false
}

func (c *MaxConstraint) MinMax(cc *matter.ConstraintContext) (min matter.ConstraintExtreme, max matter.ConstraintExtreme) {
	_, max = c.Max.MinMax(cc)
	return
}
