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

func (bv *BooleanValue) Compare(context Context, other ComparisonValue, op ComparisonOperator) (bool, error) {
	if bv == nil {
		return false, fmt.Errorf("can not compare nil value")
	}
	if other == nil {
		return false, fmt.Errorf("can not compare to nil value")
	}
	ore, ok := other.(*BooleanValue)
	if !ok {
		return false, fmt.Errorf("can not compare to non-boolean value")
	}
	switch op {
	case ComparisonOperatorEqual:
		return bv.Boolean == ore.Boolean, nil
	case ComparisonOperatorNotEqual:
		return bv.Boolean != ore.Boolean, nil
	default:
		return false, fmt.Errorf("invalid operator: %s", op.String())
	}
}

func (bv BooleanValue) Value(context Context) (any, error) {
	return bv.Boolean, nil
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
