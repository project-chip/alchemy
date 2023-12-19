package conformance

import (
	"fmt"
)

type Conformance interface {
	fmt.Stringer

	Type() Type
	Eval(context Context) (State, error)
	Equal(oc Conformance) bool
}

type HasConformance interface {
	GetConformance() Set
}
