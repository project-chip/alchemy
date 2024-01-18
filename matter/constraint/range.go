package constraint

import (
	"fmt"

	"github.com/hasty/alchemy/matter/types"
)

type RangeConstraint struct {
	Minimum ConstraintLimit `json:"min"`
	Maximum ConstraintLimit `json:"max"`
}

func (c *RangeConstraint) Type() ConstraintType {
	return ConstraintTypeRange
}

func (c *RangeConstraint) AsciiDocString(dataType *types.DataType) string {
	return fmt.Sprintf("%s to %s", c.Minimum.AsciiDocString(dataType), c.Maximum.AsciiDocString(dataType))
}

func (c *RangeConstraint) Equal(o Constraint) bool {
	if oc, ok := o.(*RangeConstraint); ok {
		return oc.Minimum.Equal(c.Minimum) && oc.Maximum.Equal(c.Maximum)
	}
	return false
}

func (c *RangeConstraint) Min(cc Context) (from types.DataTypeExtreme) {
	return c.Minimum.Min(cc)
}

func (c *RangeConstraint) Max(cc Context) (to types.DataTypeExtreme) {
	return c.Maximum.Max(cc)
}

func (c *RangeConstraint) Default(cc Context) (max types.DataTypeExtreme) {
	return
}

func (c *RangeConstraint) Clone() Constraint {
	return &RangeConstraint{Minimum: c.Minimum.Clone(), Maximum: c.Maximum.Clone()}
}