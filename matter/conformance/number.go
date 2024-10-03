package conformance

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type FloatValue struct {
	raw   string
	Float decimal.Decimal `json:"value"`
}

func NewFloatValue(d decimal.Decimal, raw string) *FloatValue {
	return &FloatValue{raw: raw, Float: d}
}

func (fv FloatValue) ASCIIDocString() string {
	return fv.raw
}

func (fv FloatValue) Description() string {
	return fv.raw
}

func (fv *FloatValue) Compare(context Context, other ComparisonValue, op ComparisonOperator) (bool, error) {
	if fv == nil {
		return false, fmt.Errorf("can not compare nil value")
	}
	if other == nil {
		return false, fmt.Errorf("can not compare to nil value")
	}
	ore, ok := other.(*FloatValue)
	if !ok {
		return false, fmt.Errorf("can not compare to non-float value")
	}
	switch op {
	case ComparisonOperatorGreaterThan:
		return fv.Float.GreaterThan(ore.Float), nil
	case ComparisonOperatorGreaterThanOrEqual:
		return fv.Float.GreaterThanOrEqual(ore.Float), nil
	case ComparisonOperatorLessThan:
		return fv.Float.LessThan(ore.Float), nil
	case ComparisonOperatorLessThanOrEqual:
		return fv.Float.LessThanOrEqual(ore.Float), nil
	default:
		return false, fmt.Errorf("invalid operator: %s", op.String())
	}
}

func (fv FloatValue) Value(context Context) (any, error) {
	return fv.Float, nil
}

func (fv *FloatValue) Equal(ofv ComparisonValue) bool {
	if ofv == nil {
		return fv == nil
	} else if fv == nil {
		return false
	}
	ore, ok := ofv.(*FloatValue)
	if !ok {
		return false
	}
	if !fv.Float.Equal(ore.Float) {
		return false
	}
	return true
}

func (fv FloatValue) Clone() ComparisonValue {
	return &FloatValue{raw: fv.raw, Float: fv.Float.Copy()}
}

type IntValue struct {
	raw string
	Int int64 `json:"value"`
}

func NewIntValue(value int64, raw string) *IntValue {
	return &IntValue{raw: raw, Int: value}
}

func (fv IntValue) ASCIIDocString() string {
	return fv.raw
}

func (fv IntValue) Description() string {
	return fv.raw
}

func (fv *IntValue) Compare(context Context, other ComparisonValue, op ComparisonOperator) (bool, error) {
	if fv == nil {
		return false, fmt.Errorf("can not compare nil value")
	}
	if other == nil {
		return false, fmt.Errorf("can not compare to nil value")
	}
	ore, ok := other.(*IntValue)
	if !ok {
		return false, fmt.Errorf("can not compare to non-int value")
	}
	switch op {
	case ComparisonOperatorGreaterThan:
		return fv.Int > ore.Int, nil
	case ComparisonOperatorGreaterThanOrEqual:
		return fv.Int >= ore.Int, nil
	case ComparisonOperatorLessThan:
		return fv.Int < ore.Int, nil
	case ComparisonOperatorLessThanOrEqual:
		return fv.Int <= ore.Int, nil
	default:
		return false, fmt.Errorf("invalid operator: %s", op.String())
	}
}

func (fv IntValue) Value(context Context) (any, error) {
	return fv.Int, nil
}

func (fv *IntValue) Equal(ofv ComparisonValue) bool {
	if ofv == nil {
		return fv == nil
	} else if fv == nil {
		return false
	}
	ore, ok := ofv.(*IntValue)
	if !ok {
		return false
	}
	if fv.Int != ore.Int {
		return false
	}
	return true
}

func (fv IntValue) Clone() ComparisonValue {
	return &IntValue{raw: fv.raw, Int: fv.Int}
}

type HexValue struct {
	raw string
	Hex uint64 `json:"value"`
}

func NewHexValue(value uint64, raw string) *HexValue {
	return &HexValue{raw: raw, Hex: value}
}

func (fv HexValue) ASCIIDocString() string {
	return fv.raw
}

func (fv HexValue) Description() string {
	return fv.raw
}

func (fv *HexValue) Compare(context Context, other ComparisonValue, op ComparisonOperator) (bool, error) {
	if fv == nil {
		return false, fmt.Errorf("can not compare nil value")
	}
	if other == nil {
		return false, fmt.Errorf("can not compare to nil value")
	}
	ore, ok := other.(*HexValue)
	if !ok {
		return false, fmt.Errorf("can not compare to non-int value")
	}
	switch op {
	case ComparisonOperatorGreaterThan:
		return fv.Hex > ore.Hex, nil
	case ComparisonOperatorGreaterThanOrEqual:
		return fv.Hex >= ore.Hex, nil
	case ComparisonOperatorLessThan:
		return fv.Hex < ore.Hex, nil
	case ComparisonOperatorLessThanOrEqual:
		return fv.Hex <= ore.Hex, nil
	default:
		return false, fmt.Errorf("invalid operator: %s", op.String())
	}
}

func (fv HexValue) Value(context Context) (any, error) {
	return fv.Hex, nil
}

func (fv *HexValue) Equal(ofv ComparisonValue) bool {
	if ofv == nil {
		return fv == nil
	} else if fv == nil {
		return false
	}
	ore, ok := ofv.(*HexValue)
	if !ok {
		return false
	}
	if fv.Hex != ore.Hex {
		return false
	}
	return true
}

func (fv HexValue) Clone() ComparisonValue {
	return &HexValue{raw: fv.raw, Hex: fv.Hex}
}
