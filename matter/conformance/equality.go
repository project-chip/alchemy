package conformance

import (
	"encoding/json"
	"fmt"
)

type EqualityExpression struct {
	Not   bool
	Left  Expression
	Right Expression
}

func (o *EqualityExpression) String() string {
	if o.Not {
		return fmt.Sprintf("(%s != %s)", o.Left, o.Right)
	}

	return fmt.Sprintf("(%s == %s)", o.Left, o.Right)
}

func (ee *EqualityExpression) Eval(context Context) (bool, error) {
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

func (ee *EqualityExpression) Equal(e Expression) bool {
	if ee == nil {
		return e == nil
	} else if e == nil {
		return false
	}
	oee, ok := e.(*EqualityExpression)
	if !ok {
		return false
	}
	if ee.Not != oee.Not {
		return false
	}
	if !ee.Left.Equal(oee.Left) {
		return false
	}
	if !ee.Right.Equal(oee.Right) {
		return false
	}
	return true
}

func (ee *EqualityExpression) MarshalJSON() ([]byte, error) {
	js := map[string]any{
		"type":  "equality",
		"left":  ee.Left,
		"right": ee.Right,
	}
	if ee.Not {
		js["not"] = true
	}
	return json.Marshal(js)
}

func (ee *EqualityExpression) Clone() Expression {
	return &EqualityExpression{Not: ee.Not, Left: ee.Left.Clone(), Right: ee.Right.Clone()}
}
