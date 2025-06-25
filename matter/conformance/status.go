package conformance

import (
	"fmt"

	"github.com/project-chip/alchemy/matter/types"
)

type StatusCodeValue struct {
	raw        string
	StatusCode types.StatusCode `json:"statusCode"`
}

func NewStatusCodeValue(statusCode types.StatusCode, raw string) *StatusCodeValue {
	return &StatusCodeValue{raw: raw, StatusCode: statusCode}
}

func (scv StatusCodeValue) ASCIIDocString() string {
	return scv.raw
}

func (scv StatusCodeValue) Description() string {
	return scv.raw
}

func (scv *StatusCodeValue) Compare(context Context, other ComparisonValue, op ComparisonOperator) (ExpressionResult, error) {
	ov, err := other.Value(context)
	if err != nil {
		return nil, err
	}
	switch op {
	case ComparisonOperatorEqual:
		switch other := ov.Result().(type) {
		case *StatusCodeValue:
			return &expressionResult{value: other.StatusCode == scv.StatusCode, confidence: coalesceConfidences(ConfidenceDefinite, ov.Confidence())}, nil
		default:
			return nil, fmt.Errorf("comparing with non-status code value: %s", op.String())
		}
	case ComparisonOperatorNotEqual:
		switch other := ov.Result().(type) {
		case *StatusCodeValue:
			return &expressionResult{value: other.StatusCode != scv.StatusCode, confidence: coalesceConfidences(ConfidenceDefinite, ov.Confidence())}, nil
		default:
			return nil, fmt.Errorf("comparing with non-status code value: %s", op.String())
		}
	default:
		return nil, fmt.Errorf("invalid operator with status code value: %s", op.String())
	}
}

func (scv *StatusCodeValue) Value(context Context) (ExpressionResult, error) {
	return &expressionResult{value: scv, confidence: ConfidenceDefinite}, nil
}

func (scv *StatusCodeValue) Equal(ofv ComparisonValue) bool {
	if ofv == nil {
		return scv == nil
	} else if scv == nil {
		return false
	}
	osc, ok := ofv.(*StatusCodeValue)
	return ok && osc.StatusCode == scv.StatusCode
}

func (scv StatusCodeValue) Clone() ComparisonValue {
	return &StatusCodeValue{raw: scv.raw, StatusCode: scv.StatusCode}
}
