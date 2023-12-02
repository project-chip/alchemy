package conformance

import (
	"fmt"

	"github.com/hasty/alchemy/matter"
)

type EqualityExpression struct {
	Not   bool
	Left  matter.ConformanceExpression
	Right matter.ConformanceExpression
}

func (o *EqualityExpression) String() string {
	if o.Not {
		return fmt.Sprintf("(%s != %s)", o.Left, o.Right)
	}

	return fmt.Sprintf("(%s == %s)", o.Left, o.Right)
}

func (ee *EqualityExpression) Eval(context matter.ConformanceContext) (bool, error) {
	l, err := ee.Left.Eval(context)
	if err != nil {
		return false, err
	}
	r, err := ee.Right.Eval(context)
	if err != nil {
		return false, err
	}
	if ee.Not {
		return l != r, nil
	}
	return l == r, nil
}
