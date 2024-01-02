package constraint

import (
	"github.com/hasty/alchemy/matter/types"
)

type DescribedConstraint struct {
}

func (c *DescribedConstraint) Type() ConstraintType {
	return ConstraintTypeDescribed
}

func (c *DescribedConstraint) AsciiDocString(dataType *types.DataType) string {
	return "desc"
}

func (c *DescribedConstraint) Equal(o Constraint) bool {
	_, ok := o.(*DescribedConstraint)
	return ok
}

func (c *DescribedConstraint) Min(cc Context) (min types.DataTypeExtreme) {
	return
}

func (c *DescribedConstraint) Max(cc Context) (max types.DataTypeExtreme) {
	return
}

func (c *DescribedConstraint) Default(cc Context) (max types.DataTypeExtreme) {
	return
}

func (c *DescribedConstraint) Clone() Constraint {
	return &DescribedConstraint{}
}
