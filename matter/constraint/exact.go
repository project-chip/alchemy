package constraint

import (
	"github.com/hasty/alchemy/matter/types"
)

type ExactConstraint struct {
	Value ConstraintLimit
}

func (c *ExactConstraint) Type() ConstraintType {
	return ConstraintTypeExact
}

func (c *ExactConstraint) AsciiDocString(dataType *types.DataType) string {
	return c.Value.AsciiDocString(dataType)
}

func (c *ExactConstraint) Equal(o Constraint) bool {
	if oc, ok := o.(*ExactConstraint); ok {
		return oc.Value.Equal(c.Value)
	}
	return false
}

func (c *ExactConstraint) Min(cc Context) (min types.DataTypeExtreme) {
	return c.Value.Min(cc)
}

func (c *ExactConstraint) Max(cc Context) (max types.DataTypeExtreme) {
	return c.Value.Max(cc)
}

func (c *ExactConstraint) Default(cc Context) (max types.DataTypeExtreme) {
	return c.Value.Default(cc)
}

func (c *ExactConstraint) Clone() Constraint {
	return &ExactConstraint{Value: c.Value.Clone()}
}