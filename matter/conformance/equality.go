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

func (ee *EqualityExpression) ASCIIDocString() string {
	if ee.Not {
		return fmt.Sprintf("(%s != %s)", ee.Left.ASCIIDocString(), ee.Right.ASCIIDocString())
	}

	return fmt.Sprintf("(%s == %s)", ee.Left.ASCIIDocString(), ee.Right.ASCIIDocString())
}

func (ee *EqualityExpression) Description() string {
	if ee.Not {
		return fmt.Sprintf("(%s != %s)", ee.Left.Description(), ee.Right.Description())
	}

	return fmt.Sprintf("(%s == %s)", ee.Left.Description(), ee.Right.Description())
}

func (ee *EqualityExpression) Eval(context Context) (result ExpressionResult, err error) {
	var l, r ExpressionResult
	l, err = ee.Left.Eval(context)
	if err != nil {
		return
	}
	r, err = ee.Right.Eval(context)
	if err != nil {
		return
	}
	if ee.Not {
		return compareResults(l, r, ComparisonOperatorNotEqual)
	}
	return compareResults(l, r, ComparisonOperatorEqual)
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
