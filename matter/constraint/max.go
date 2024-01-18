package constraint

import (
	"fmt"

	"github.com/hasty/alchemy/matter/types"
)

type MaxConstraint struct {
	Maximum ConstraintLimit
}

func (c *MaxConstraint) Type() ConstraintType {
	return ConstraintTypeMax
}

func (c *MaxConstraint) AsciiDocString(dataType *types.DataType) string {
	return fmt.Sprintf("max %s", c.Maximum.AsciiDocString(dataType))
}

func (c *MaxConstraint) Equal(o Constraint) bool {
	if oc, ok := o.(*MaxConstraint); ok {
		return oc.Maximum.Equal(c.Maximum)
	}
	return false
}

func (c *MaxConstraint) Min(cc Context) (min types.DataTypeExtreme) {
	return
}

func (c *MaxConstraint) Max(cc Context) (max types.DataTypeExtreme) {
	return c.Maximum.Max(cc)
}

func (c *MaxConstraint) Default(cc Context) (max types.DataTypeExtreme) {
	return
}

func (c *MaxConstraint) Clone() Constraint {
	return &MaxConstraint{Maximum: c.Maximum.Clone()}
}