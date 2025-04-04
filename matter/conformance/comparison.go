package conformance

import (
	"encoding/json"
	"fmt"
	"math"
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

// Does its best to compare two values
func compare(context Context, op ComparisonOperator, a ComparisonValue, b ComparisonValue) (bool, error) {
	av, err := a.Value(context)
	if err != nil {
		return false, err
	}
	bv, err := b.Value(context)
	if err != nil {
		return false, err
	}
	ai64, aIsI64 := toInt64(av)
	if aIsI64 {
		bi64, bIsI64 := toInt64(bv)
		if bIsI64 { // They can both safely be int64
			return compareNumbers(op, ai64, bi64)
		}
		switch bv := bv.(type) {
		case uint64:
			if bv <= math.MaxInt64 { // If b is less than the largest possible int64, we can compare them normally
				return compareNumbers(op, ai64, int64(bv))
			}
			// b is definitely larger than a
			switch op {
			case ComparisonOperatorEqual, ComparisonOperatorGreaterThan, ComparisonOperatorGreaterThanOrEqual:
				return false, nil
			case ComparisonOperatorNotEqual, ComparisonOperatorLessThan, ComparisonOperatorLessThanOrEqual:
				return true, nil
			default:
				return false, fmt.Errorf("invalid op on comparison: %s", op.String())
			}
		}
	}
	aUI64, aIsUI64 := toUint64(av)
	if aIsUI64 {
		bui64, bIsUI64 := toUint64(bv)
		if bIsUI64 {
			return compareNumbers(op, aUI64, bui64)
		}
		switch bv.(type) {
		case int64, int32, int16, int8, int: // Must've been less than zero, so b is definitely less than a
			switch op {
			case ComparisonOperatorNotEqual, ComparisonOperatorGreaterThan, ComparisonOperatorGreaterThanOrEqual:
				return true, nil
			case ComparisonOperatorEqual, ComparisonOperatorLessThan, ComparisonOperatorLessThanOrEqual:
				return false, nil
			default:
				return false, fmt.Errorf("invalid op on comparison: %s", op.String())
			}
		}
	}
	switch av := av.(type) {
	case float32:
		switch bv := bv.(type) {
		case float32:
			return compareNumbers(op, av, bv)
		case float64:
			return compareNumbers(op, av, float32(bv))
		default:
			return false, fmt.Errorf("can't compare float32 to %T", bv)
		}
	case float64:
		switch bv := bv.(type) {
		case float32:
			return compareNumbers(op, av, float64(bv))
		case float64:
			return compareNumbers(op, av, bv)
		default:
			return false, fmt.Errorf("can't compare float64 to %T", bv)
		}
	case bool:
		switch bv := bv.(type) {
		case bool:
			switch op {
			case ComparisonOperatorNotEqual:
				return av != bv, nil
			case ComparisonOperatorEqual:
				return av == bv, nil
			default:
				return false, fmt.Errorf("invalid op on comparison: %s", op.String())
			}
		default:
			return false, fmt.Errorf("invalid type on comparison with bool: %T", bv)
		}
	default:
		return false, fmt.Errorf("invalid type on comparison: %T", av)
	}
}

func compareNumbers[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](op ComparisonOperator, a, b T) (bool, error) {
	switch op {
	case ComparisonOperatorEqual:
		return a == b, nil
	case ComparisonOperatorNotEqual:
		return a != b, nil
	case ComparisonOperatorGreaterThan:
		return a > b, nil
	case ComparisonOperatorGreaterThanOrEqual:
		return a >= b, nil
	case ComparisonOperatorLessThan:
		return a < b, nil
	case ComparisonOperatorLessThanOrEqual:
		return a <= b, nil
	default:
		return false, fmt.Errorf("invalid op on comparison: %s", op.String())
	}
}

func toInt64(i any) (int64, bool) {
	switch i := i.(type) {
	case int64:
		return i, true
	case int32:
		return int64(i), true
	case int16:
		return int64(i), true
	case int8:
		return int64(i), true
	case int:
		return int64(i), true
	case uint:
		return int64(i), true
	case uint32:
		return int64(i), true
	case uint16:
		return int64(i), true
	case uint8:
		return int64(i), true
	case uint64:
		if i <= math.MaxInt64 {
			return int64(i), true
		}
		return 0, false
	default:
		return 0, false
	}
}

func toUint64(i any) (uint64, bool) {
	switch i := i.(type) {
	case int64:
		if i >= 0 {
			return uint64(i), true
		}
	case int32:
		if i >= 0 {
			return uint64(i), true
		}
	case int16:
		if i >= 0 {
			return uint64(i), true
		}
	case int8:
		if i >= 0 {
			return uint64(i), true
		}
	case int:
		if i >= 0 {
			return uint64(i), true
		}
	case uint64:
		return i, true
	case uint32:
		return uint64(i), true
	case uint16:
		return uint64(i), true
	case uint8:
		return uint64(i), true
	case uint:
		return uint64(i), true
	}
	return 0, false
}
