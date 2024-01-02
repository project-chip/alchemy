package constraint

import (
	"fmt"

	"github.com/hasty/alchemy/matter/types"
)

type MinConstraint struct {
	Minimum ConstraintLimit `json:"min"`
}

func (c *MinConstraint) Type() ConstraintType {
	return ConstraintTypeMin
}

func (c *MinConstraint) AsciiDocString(dataType *types.DataType) string {
	return fmt.Sprintf("min %s", c.Minimum.AsciiDocString(dataType))
}

func (c *MinConstraint) Equal(o Constraint) bool {
	if oc, ok := o.(*MinConstraint); ok {
		return oc.Minimum.Equal(c.Minimum)
	}
	return false
}

func (c *MinConstraint) Min(cc Context) (min types.DataTypeExtreme) {
	return c.Minimum.Min(cc)
}

func (c *MinConstraint) Max(cc Context) (max types.DataTypeExtreme) {
	return
}

func (c *MinConstraint) Default(cc Context) (max types.DataTypeExtreme) {
	return
}

func (c *MinConstraint) Clone() Constraint {
	return &MinConstraint{Minimum: c.Minimum.Clone()}
}
