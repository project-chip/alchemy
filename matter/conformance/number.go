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

func (fv *FloatValue) Compare(context Context, other ComparisonValue, op ComparisonOperator) (ExpressionResult, error) {
	if fv == nil {
		return nil, fmt.Errorf("can not compare nil value")
	}
	if other == nil {
		return nil, fmt.Errorf("can not compare to nil value")
	}
	val, err := other.Value(context)
	if err != nil {
		return nil, err
	}
	ore, ok := val.Result().(*FloatValue)
	if !ok {
		return nil, fmt.Errorf("can not compare to non-float value")
	}
	switch op {
	case ComparisonOperatorGreaterThan:
		return &expressionResult{value: fv.Float.GreaterThan(ore.Float), confidence: coalesceConfidences(ConfidenceDefinite, val.Confidence())}, nil
	case ComparisonOperatorGreaterThanOrEqual:
		return &expressionResult{value: fv.Float.GreaterThanOrEqual(ore.Float), confidence: coalesceConfidences(ConfidenceDefinite, val.Confidence())}, nil
	case ComparisonOperatorLessThan:
		return &expressionResult{value: fv.Float.LessThan(ore.Float), confidence: coalesceConfidences(ConfidenceDefinite, val.Confidence())}, nil
	case ComparisonOperatorLessThanOrEqual:
		return &expressionResult{value: fv.Float.LessThanOrEqual(ore.Float), confidence: coalesceConfidences(ConfidenceDefinite, val.Confidence())}, nil
	default:
		return nil, fmt.Errorf("invalid operator: %s", op.String())
	}
}

func (fv *FloatValue) Value(context Context) (ExpressionResult, error) {
	return &expressionResult{value: fv, confidence: ConfidenceDefinite}, nil
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

func (fv *IntValue) Compare(context Context, other ComparisonValue, op ComparisonOperator) (ExpressionResult, error) {
	if fv == nil {
		return nil, fmt.Errorf("can not compare nil value")
	}
	if other == nil {
		return nil, fmt.Errorf("can not compare to nil value")
	}
	val, err := other.Value(context)
	if err != nil {
		return nil, err
	}
	ore, ok := val.Result().(*IntValue)
	if !ok {
		return nil, fmt.Errorf("can not compare to non-int value")
	}
	switch op {
	case ComparisonOperatorGreaterThan:
		return &expressionResult{value: fv.Int > ore.Int, confidence: coalesceConfidences(ConfidenceDefinite, val.Confidence())}, nil
	case ComparisonOperatorGreaterThanOrEqual:
		return &expressionResult{value: fv.Int >= ore.Int, confidence: coalesceConfidences(ConfidenceDefinite, val.Confidence())}, nil
	case ComparisonOperatorLessThan:
		return &expressionResult{value: fv.Int < ore.Int, confidence: coalesceConfidences(ConfidenceDefinite, val.Confidence())}, nil
	case ComparisonOperatorLessThanOrEqual:
		return &expressionResult{value: fv.Int <= ore.Int, confidence: coalesceConfidences(ConfidenceDefinite, val.Confidence())}, nil
	case ComparisonOperatorEqual:
		return &expressionResult{value: fv.Int == ore.Int, confidence: coalesceConfidences(ConfidenceDefinite, val.Confidence())}, nil
	case ComparisonOperatorNotEqual:
		return &expressionResult{value: fv.Int != ore.Int, confidence: coalesceConfidences(ConfidenceDefinite, val.Confidence())}, nil
	default:
		return nil, fmt.Errorf("invalid operator: %s", op.String())
	}
}

func (fv *IntValue) Value(context Context) (ExpressionResult, error) {
	return &expressionResult{value: fv, confidence: ConfidenceDefinite}, nil
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

func (fv *HexValue) Compare(context Context, other ComparisonValue, op ComparisonOperator) (ExpressionResult, error) {
	if fv == nil {
		return nil, fmt.Errorf("can not compare nil value")
	}
	if other == nil {
		return nil, fmt.Errorf("can not compare to nil value")
	}
	val, err := other.Value(context)
	if err != nil {
		return nil, err
	}
	ore, ok := val.Result().(*HexValue)
	if !ok {
		return nil, fmt.Errorf("can not compare to non-int value")
	}
	switch op {
	case ComparisonOperatorGreaterThan:
		return &expressionResult{value: fv.Hex > ore.Hex, confidence: coalesceConfidences(ConfidenceDefinite, val.Confidence())}, nil
	case ComparisonOperatorGreaterThanOrEqual:
		return &expressionResult{value: fv.Hex >= ore.Hex, confidence: coalesceConfidences(ConfidenceDefinite, val.Confidence())}, nil
	case ComparisonOperatorLessThan:
		return &expressionResult{value: fv.Hex < ore.Hex, confidence: coalesceConfidences(ConfidenceDefinite, val.Confidence())}, nil
	case ComparisonOperatorLessThanOrEqual:
		return &expressionResult{value: fv.Hex <= ore.Hex, confidence: coalesceConfidences(ConfidenceDefinite, val.Confidence())}, nil
	case ComparisonOperatorEqual:
		return &expressionResult{value: fv.Hex == ore.Hex, confidence: coalesceConfidences(ConfidenceDefinite, val.Confidence())}, nil
	case ComparisonOperatorNotEqual:
		return &expressionResult{value: fv.Hex != ore.Hex, confidence: coalesceConfidences(ConfidenceDefinite, val.Confidence())}, nil

	default:
		return nil, fmt.Errorf("invalid operator: %s", op.String())
	}
}

func (fv *HexValue) Value(context Context) (ExpressionResult, error) {
	return &expressionResult{value: fv, confidence: ConfidenceDefinite}, nil
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
