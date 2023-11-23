package conformance

import (
	"fmt"

	"github.com/hasty/alchemy/matter"
)

type FeatureExpression struct {
	ID  string
	Not bool
}

func (fe *FeatureExpression) String() string {
	if fe.Not {
		return fmt.Sprintf("not %s", fe.ID)
	}
	return fe.ID
}

func (fe *FeatureExpression) Eval(context matter.ConformanceContext) (bool, error) {
	v, ok := context[fe.ID]
	if !ok {
		return fe.Not, nil
	}
	if b, ok := v.(bool); ok {
		return b != fe.Not, nil
	}
	return false, fmt.Errorf("unexpected value when interpreting identifier %s: %v", fe.ID, v)
}
