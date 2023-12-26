package constraint

import (
	"github.com/hasty/alchemy/matter"
)

type ExactConstraint struct {
	Value matter.ConstraintLimit
}

func (c *ExactConstraint) Type() matter.ConstraintType {
	return matter.ConstraintTypeExact
}

func (c *ExactConstraint) AsciiDocString(dataType *matter.DataType) string {
	return c.Value.AsciiDocString(dataType)
}

func (c *ExactConstraint) Equal(o matter.Constraint) bool {
	if oc, ok := o.(*ExactConstraint); ok {
		return oc.Value.Equal(c.Value)
	}
	return false
}

func (c *ExactConstraint) Min(cc *matter.ConstraintContext) (min matter.ConstraintExtreme) {
	return c.Value.Min(cc)
}

func (c *ExactConstraint) Max(cc *matter.ConstraintContext) (max matter.ConstraintExtreme) {
	return c.Value.Max(cc)
}

func (c *ExactConstraint) Default(cc *matter.ConstraintContext) (max matter.ConstraintExtreme) {
	return c.Value.Default(cc)
}

func (c *ExactConstraint) Clone() matter.Constraint {
	return &ExactConstraint{Value: c.Value.Clone()}
}
