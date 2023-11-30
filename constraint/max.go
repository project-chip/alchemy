package constraint

import (
	"fmt"

	"github.com/hasty/alchemy/matter"
)

type MaxConstraint struct {
	Maximum matter.ConstraintLimit
}

func (c *MaxConstraint) AsciiDocString(dataType *matter.DataType) string {
	return fmt.Sprintf("max %s", c.Maximum.AsciiDocString(dataType))
}

func (c *MaxConstraint) Equal(o matter.Constraint) bool {
	if oc, ok := o.(*MaxConstraint); ok {
		return oc.Maximum.Equal(c.Maximum)
	}
	return false
}

func (c *MaxConstraint) Min(cc *matter.ConstraintContext) (min matter.ConstraintExtreme) {
	return
}

func (c *MaxConstraint) Max(cc *matter.ConstraintContext) (max matter.ConstraintExtreme) {
	return c.Maximum.Max(cc)
}

func (c *MaxConstraint) Default(cc *matter.ConstraintContext) (max matter.ConstraintExtreme) {
	return
}
