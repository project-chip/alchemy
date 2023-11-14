package constraint

import (
	"math/big"
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

func parseFloat(s string) (*big.Float, error) {
	var z big.Float
	f, _, err := (&z).Parse(s, 10)
	if err != nil {
		return nil, err
	}
	return f, nil
}

type ManufacturerLimit struct {
	Value string
}

func (c *ManufacturerLimit) AsciiDocString() string {
	return c.Value
}

func (c *ManufacturerLimit) Equal(o matter.ConstraintLimit) bool {
	if oc, ok := o.(*ManufacturerLimit); ok {
		return oc.Value == c.Value
	}
	return false
}

func (c *ManufacturerLimit) MinMax(cc *matter.ConstraintContext) (min matter.ConstraintExtreme, max matter.ConstraintExtreme) {
	return
}
