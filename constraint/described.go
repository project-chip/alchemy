package constraint

import "github.com/hasty/alchemy/matter"

type DescribedConstraint struct {
}

func (c *DescribedConstraint) AsciiDocString(dataType *matter.DataType) string {
	return "desc"
}

func (c *DescribedConstraint) Equal(o matter.Constraint) bool {
	_, ok := o.(*DescribedConstraint)
	return ok
}

func (c *DescribedConstraint) Min(cc *matter.ConstraintContext) (min matter.ConstraintExtreme) {
	return
}

func (c *DescribedConstraint) Max(cc *matter.ConstraintContext) (max matter.ConstraintExtreme) {
	return
}
