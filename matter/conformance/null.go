package conformance

import "fmt"

type NullValue struct {
	raw string
}

func NewNullValue(raw string) *NullValue {
	return &NullValue{raw: raw}
}

func (fv NullValue) ASCIIDocString() string {
	return fv.raw
}

func (fv NullValue) Description() string {
	return fv.raw
}

func (fv *NullValue) Compare(context Context, other ComparisonValue, op ComparisonOperator) (ExpressionResult, error) {
	switch op {
	case ComparisonOperatorEqual:
		switch other.(type) {
		case nil, *NullValue:
			return &expressionResult{value: true, confidence: coalesceConfidences(ConfidenceDefinite)}, nil
		default:
			return &expressionResult{value: false, confidence: coalesceConfidences(ConfidenceDefinite)}, nil
		}
	case ComparisonOperatorNotEqual:
		switch other.(type) {
		case nil, *NullValue:
			return &expressionResult{value: false, confidence: coalesceConfidences(ConfidenceDefinite)}, nil
		default:
			return &expressionResult{value: true, confidence: coalesceConfidences(ConfidenceDefinite)}, nil
		}
	default:
		return nil, fmt.Errorf("invalid operator with null value: %s", op.String())
	}
}

func (fv NullValue) Value(context Context) (ExpressionResult, error) {
	return &expressionResult{value: nil, confidence: ConfidenceDefinite}, nil
}

func (fv *NullValue) Equal(ofv ComparisonValue) bool {
	if ofv == nil {
		return fv == nil
	} else if fv == nil {
		return false
	}
	_, ok := ofv.(*NullValue)
	return !ok
}

func (fv NullValue) Clone() ComparisonValue {
	return &NullValue{raw: fv.raw}
}
