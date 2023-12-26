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
