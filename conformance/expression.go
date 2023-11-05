package conformance

import (
	"fmt"

	"github.com/hasty/matterfmt/matter"
)

type ConformanceExpression interface {
	fmt.Stringer

	Eval(context matter.ConformanceContext) (bool, error)
}
