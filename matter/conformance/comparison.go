package conformance

import (
	"encoding/json"
	"fmt"
)

type ComparisonOperator uint8

const (
	ComparisonOperatorNone ComparisonOperator = iota
	ComparisonOperatorEqual
	ComparisonOperatorNotEqual
	ComparisonOperatorLessThan
	ComparisonOperatorLessThanOrEqual
	ComparisonOperatorGreaterThan
	ComparisonOperatorGreaterThanOrEqual
)

func (co ComparisonOperator) String() string {
	switch co {
	case ComparisonOperatorEqual:
		return "=="
	case ComparisonOperatorNotEqual:
		return "!="
	case ComparisonOperatorLessThan:
		return "<"
	case ComparisonOperatorLessThanOrEqual:
		return "<="
	case ComparisonOperatorGreaterThan:
		return ">"
	case ComparisonOperatorGreaterThanOrEqual:
		return ">="
	}
	return ""
}

type ComparisonValue interface {
	ASCIIDocString() string
	Description() string

	Compare(context Context, other ComparisonValue, op ComparisonOperator) (bool, error)
	Value(context Context) (any, error)
	Equal(e ComparisonValue) bool
	Clone() ComparisonValue
}

type ComparisonExpression struct {
	Op    ComparisonOperator
	Left  ComparisonValue
	Right ComparisonValue
}

func (ce *ComparisonExpression) ASCIIDocString() string {
	return fmt.Sprintf("(%s %s %s)", ce.Left.ASCIIDocString(), ce.Op.String(), ce.Right.ASCIIDocString())
}

func (ce *ComparisonExpression) Description() string {
	return fmt.Sprintf("(%s %s %s)", ce.Left.ASCIIDocString(), ce.Op.String(), ce.Right.ASCIIDocString())
}

func (ce *ComparisonExpression) Eval(context Context) (bool, error) {
	return ce.Left.Compare(context, ce.Right, ce.Op)
}

func (ce *ComparisonExpression) Equal(e Expression) bool {
	if ce == nil {
		return e == nil
	} else if e == nil {
		return false
	}
	oee, ok := e.(*ComparisonExpression)
	if !ok {
		return false
	}
	if ce.Op != oee.Op {
		return false
	}
	if !ce.Left.Equal(oee.Left) {
		return false
	}
	if !ce.Right.Equal(oee.Right) {
		return false
	}
	return true
}

func (ce *ComparisonExpression) MarshalJSON() ([]byte, error) {
	js := map[string]any{
		"type":  "equality",
		"left":  ce.Left,
		"right": ce.Right,
	}
	js["op"] = ce.Op.String()
	return json.Marshal(js)
}

func (ce *ComparisonExpression) Clone() Expression {
	return &ComparisonExpression{Op: ce.Op, Left: ce.Left.Clone(), Right: ce.Right.Clone()}
}

func compare(context Context, op ComparisonOperator, a ComparisonValue, b ComparisonValue) (bool, error) {
	av, err := a.Value(context)
	if err != nil {
		return false, err
	}
	bv, err := b.Value(context)
	if err != nil {
		return false, err
	}
	switch av := av.(type) {
	case int64:
		switch bv := bv.(type) {
		case int64:
			switch op {
			case ComparisonOperatorEqual:
				return av == bv, nil
			case ComparisonOperatorNotEqual:
				return av != bv, nil
			case ComparisonOperatorGreaterThan:
				return av > bv, nil
			case ComparisonOperatorGreaterThanOrEqual:
				return av >= bv, nil
			case ComparisonOperatorLessThan:
				return av < bv, nil
			case ComparisonOperatorLessThanOrEqual:
				return av <= bv, nil
			default:
				return false, fmt.Errorf("invalid op on comparison: %s", op.String())
			}
		default:
			return false, fmt.Errorf("can not compare int to %T", bv)
		}
	case float64:
		switch bv := bv.(type) {
		case float64:
			switch op {
			case ComparisonOperatorEqual:
				return av == bv, nil
			case ComparisonOperatorNotEqual:
				return av != bv, nil
			case ComparisonOperatorGreaterThan:
				return av > bv, nil
			case ComparisonOperatorGreaterThanOrEqual:
				return av >= bv, nil
			case ComparisonOperatorLessThan:
				return av < bv, nil
			case ComparisonOperatorLessThanOrEqual:
				return av <= bv, nil
			default:
				return false, fmt.Errorf("invalid op on comparison: %s", op.String())
			}
		default:
			return false, fmt.Errorf("can not compare int to %T", bv)
		}
	default:
		return false, fmt.Errorf("invalid type on comparison: %T", av)
	}
}
