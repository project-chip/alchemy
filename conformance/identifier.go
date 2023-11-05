package conformance

import (
	"fmt"

	"github.com/hasty/matterfmt/matter"
)

type IdentifierExpression struct {
	ID  string
	Not bool
}

func (id *IdentifierExpression) String() string {
	if id.Not {
		return fmt.Sprintf("not %s", id.ID)
	}
	return id.ID
}

func (id *IdentifierExpression) Eval(context matter.ConformanceContext) (bool, error) {
	v, ok := context[id.ID]
	if !ok {
		return id.Not, nil
	}
	if b, ok := v.(bool); ok {
		return b != id.Not, nil
	}
	return false, fmt.Errorf("unexpected value when interpreting identifier %s: %v", id.ID, v)
}
