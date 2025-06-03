package conformance

import "fmt"

type BooleanValue struct {
	raw     string
	Boolean bool `json:"value"`
}

func NewBooleanValue(value bool, raw string) *BooleanValue {
	return &BooleanValue{raw: raw, Boolean: value}
}

func (bv BooleanValue) ASCIIDocString() string {
	return bv.raw
}

func (bv BooleanValue) Description() string {
	return bv.raw
}

func (bv *BooleanValue) Compare(context Context, other ComparisonValue, op ComparisonOperator) (result ExpressionResult, err error) {
	if bv == nil {
		err = fmt.Errorf("can not compare nil value")
		return
	}
	if other == nil {
		err = fmt.Errorf("can not compare to nil value")
		return
	}
	var val ExpressionResult
	val, err = other.Value(context)
	if err != nil {
		return
	}
	return compareResults(&expressionResult{value: bv.Boolean, confidence: ConfidenceDefinite}, val, op)
}

func (bv *BooleanValue) Value(context Context) (ExpressionResult, error) {
	return &expressionResult{value: bv.Boolean, confidence: ConfidenceDefinite}, nil
}

func (bv *BooleanValue) Result() ComparisonValue {
	return bv
}

func (bv *BooleanValue) Equal(ofv ComparisonValue) bool {
	if ofv == nil {
		return bv == nil
	} else if bv == nil {
		return false
	}
	ore, ok := ofv.(*BooleanValue)
	if !ok {
		return false
	}
	if bv.Boolean != ore.Boolean {
		return false
	}
	return true
}

func (bv BooleanValue) Clone() ComparisonValue {
	return &BooleanValue{raw: bv.raw, Boolean: bv.Boolean}
}
