package conformance

import (
	"fmt"
	"strings"
)

type MathOperand uint8

const (
	MathOperandNone MathOperand = iota
	MathOperandAdd
	MathOperandSubtract
	MathOperandMultiply
	MathOperandDivide
)

func (mo MathOperand) String() string {
	switch mo {
	case MathOperandAdd:
		return "+"
	case MathOperandSubtract:
		return "-"
	case MathOperandMultiply:
		return "*"
	case MathOperandDivide:
		return "/"
	}
	return ""
}

type MathOperation struct {
	Operand MathOperand
	Left    ComparisonValue
	Right   ComparisonValue
}

func (mv *MathOperation) ASCIIDocString() string {
	var sb strings.Builder
	sb.WriteRune('(')
	sb.WriteString(mv.Left.ASCIIDocString())
	sb.WriteRune(' ')
	sb.WriteString(mv.Operand.String())
	sb.WriteRune(' ')
	var parens bool
	switch r := mv.Right.(type) {
	case *MathOperation:
		parens = r.Operand != mv.Operand

	}
	if parens {
		sb.WriteRune('(')
	}
	sb.WriteString(mv.Right.ASCIIDocString())
	if parens {
		sb.WriteRune(')')
	}
	sb.WriteRune(')')
	return sb.String()
}

func (mv *MathOperation) Description() string {
	return mv.ASCIIDocString()
}

func (mv *MathOperation) Value(context Context) (any, error) {
	l, err := mv.Left.Value(context)
	if err != nil {
		return nil, err
	}
	r, err := mv.Right.Value(context)
	if err != nil {
		return nil, err
	}
	switch mv.Operand {
	case MathOperandAdd:
		switch lv := l.(type) {
		case int64:
			switch rv := r.(type) {
			case int64:
				return lv + rv, nil
			default:
				return nil, fmt.Errorf("can not add int64 to %T", rv)
			}
		case float64:
			switch rv := r.(type) {
			case float64:
				return lv + rv, nil
			default:
				return nil, fmt.Errorf("can not add float64 to %T", rv)
			}
		default:
			return nil, fmt.Errorf("can not add %T", lv)
		}
	case MathOperandSubtract:
		switch lv := l.(type) {
		case int64:
			switch rv := r.(type) {
			case int64:
				return lv - rv, nil
			default:
				return nil, fmt.Errorf("can not subtract %T from int64", rv)
			}
		case float64:
			switch rv := r.(type) {
			case float64:
				return lv - rv, nil
			default:
				return nil, fmt.Errorf("can not subtract %T from float64", rv)
			}
		default:
			return nil, fmt.Errorf("can not subtract from %T", lv)
		}
	case MathOperandMultiply:
		switch lv := l.(type) {
		case int64:
			switch rv := r.(type) {
			case int64:
				return lv - rv, nil
			default:
				return nil, fmt.Errorf("can not multiply %T with int64", rv)
			}
		case float64:
			switch rv := r.(type) {
			case float64:
				return lv - rv, nil
			default:
				return nil, fmt.Errorf("can not multiply %T with float64", rv)
			}
		default:
			return nil, fmt.Errorf("can not multiply %T", lv)
		}
	case MathOperandDivide:
		switch lv := l.(type) {
		case int64:
			switch rv := r.(type) {
			case int64:
				return lv / rv, nil
			default:
				return nil, fmt.Errorf("can not divide int64 by %T", rv)
			}
		case float64:
			switch rv := r.(type) {
			case float64:
				return lv / rv, nil
			default:
				return nil, fmt.Errorf("can not divide float64 by %T", rv)
			}
		default:
			return nil, fmt.Errorf("can not divide %T", lv)
		}
	default:
		return nil, fmt.Errorf("unknown operator: %s", mv.Operand.String())
	}
}

func (mv *MathOperation) Compare(context Context, other ComparisonValue, op ComparisonOperator) (bool, error) {
	v, err := mv.Value(context)
	if err != nil {
		return false, err
	}
	cv, ok := v.(ComparisonValue)
	if !ok {
		return false, fmt.Errorf("can not convert math value of type %T to ComparisonValue", v)
	}
	return compare(context, op, cv, other)
}

func (mv *MathOperation) Equal(e ComparisonValue) bool {
	omv, ok := e.(*MathOperation)
	if !ok {
		return false
	}
	if omv.Operand != mv.Operand {
		return false
	}
	if !omv.Left.Equal(mv.Left) {
		return false
	}
	if !omv.Right.Equal(mv.Right) {
		return false
	}
	return true
}

func (mv *MathOperation) Clone() ComparisonValue {
	return &MathOperation{
		Operand: mv.Operand,
		Left:    mv.Left.Clone(),
		Right:   mv.Right.Clone(),
	}
}
