package constraint

import "github.com/hasty/alchemy/matter"

type DescribedConstraint struct {
}

func (c *DescribedConstraint) AsciiDocString() string {
	return "desc"
}

func (c *DescribedConstraint) Equal(o matter.Constraint) bool {
	_, ok := o.(*DescribedConstraint)
	return ok
}

func (c *DescribedConstraint) MinMax(cc *matter.ConstraintContext) (min matter.ConstraintExtreme, max matter.ConstraintExtreme) {
	return
}
