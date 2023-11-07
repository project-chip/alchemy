package conformance

import (
	"fmt"

	"github.com/hasty/alchemy/matter"
)

type ConformanceExpression interface {
	fmt.Stringer

	Eval(context matter.ConformanceContext) (bool, error)
}
