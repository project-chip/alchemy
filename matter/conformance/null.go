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

func (fv *NullValue) Compare(context Context, other ComparisonValue, op ComparisonOperator) (bool, error) {
	switch op {
	case ComparisonOperatorEqual:
		switch other.(type) {
		case nil, *NullValue:
			return true, nil
		default:
			return false, nil
		}
	case ComparisonOperatorNotEqual:
		switch other.(type) {
		case nil, *NullValue:
			return false, nil
		default:
			return true, nil
		}
	default:
		return false, fmt.Errorf("invalid operator with null value: %s", op.String())
	}
}

func (fv NullValue) Value(context Context) (any, error) {
	return nil, nil
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
