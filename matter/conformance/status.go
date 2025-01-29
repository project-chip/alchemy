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

func (scv *StatusCodeValue) Compare(context Context, other ComparisonValue, op ComparisonOperator) (bool, error) {
	switch op {
	case ComparisonOperatorEqual:
		switch other := other.(type) {
		case *StatusCodeValue:
			return other.StatusCode == scv.StatusCode, nil
		default:
			return false, nil
		}
	case ComparisonOperatorNotEqual:
		switch other := other.(type) {
		case *StatusCodeValue:
			return other.StatusCode != scv.StatusCode, nil
		default:
			return true, nil
		}
	default:
		return false, fmt.Errorf("invalid operator with status code value: %s", op.String())
	}
}

func (scv StatusCodeValue) Value(context Context) (any, error) {
	return nil, nil
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
