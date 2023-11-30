package constraint

import (
	"strings"

	"github.com/hasty/alchemy/matter"
)

func ParseConstraint(constraint string) matter.Constraint {
	c, err := ParseReader("", strings.NewReader(constraint))
	if err != nil {
		return &GenericConstraint{Value: constraint}
	}
	return c.(matter.Constraint)
}

type ManufacturerLimit struct {
	Value string
}

func (c *ManufacturerLimit) AsciiDocString(dataType *matter.DataType) string {
	return c.Value
}

func (c *ManufacturerLimit) Equal(o matter.ConstraintLimit) bool {
	if oc, ok := o.(*ManufacturerLimit); ok {
		return oc.Value == c.Value
	}
	return false
}

func (c *ManufacturerLimit) Min(cc *matter.ConstraintContext) (min matter.ConstraintExtreme) {
	return
}

func (c *ManufacturerLimit) Max(cc *matter.ConstraintContext) (max matter.ConstraintExtreme) {
	return
}

func (c *ManufacturerLimit) Default(cc *matter.ConstraintContext) (max matter.ConstraintExtreme) {
	return
}
