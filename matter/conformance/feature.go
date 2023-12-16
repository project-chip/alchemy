package conformance

import (
	"fmt"
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

func (fe *FeatureExpression) Eval(context ConformanceContext) (bool, error) {
	return evalIdentifier(context, fe.ID, fe.Not)
}
