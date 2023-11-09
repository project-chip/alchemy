package constraint

import (
	"math/big"
	"strings"

	"github.com/hasty/alchemy/matter"
)

func ParseConstraint(constraint string) matter.Constraint {
	c, err := ParseReader("", strings.NewReader(constraint))
	if err != nil {
		return &matter.GenericConstraint{Value: constraint}
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
