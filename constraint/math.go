package constraint

import (
	"fmt"
	"math"

	"github.com/hasty/alchemy/matter"
)

type MathExpressionLimit struct {
	Operand string
	Left    matter.ConstraintLimit
	Right   matter.ConstraintLimit
}

func (c *MathExpressionLimit) AsciiDocString(dataType *matter.DataType) string {
	return fmt.Sprintf("(%s %s %s)", c.Left.AsciiDocString(dataType), c.Operand, c.Right.AsciiDocString(dataType))
}

func (c *MathExpressionLimit) Equal(o matter.ConstraintLimit) bool {
	if oc, ok := o.(*MathExpressionLimit); ok {
		return oc.Operand == c.Operand && oc.Left.Equal(c.Left) && oc.Right.Equal(c.Right)
	}
	return false
}

type Number interface {
	int64 | uint64
}

func add[T Number](a, b T) T {
	return a + b
}

func operate[T Number](operand string, left, right T) (val T) {
	switch operand {
	case "+":
		val = left + right
	case "-":
		val = left - right
	case "*":
		val = left * right
	case "/":
		val = left / right
	}
	return
}

func (c *MathExpressionLimit) MinMax(cc *matter.ConstraintContext) (matter.ConstraintExtreme, matter.ConstraintExtreme) {
	leftMin, leftMax := c.Left.MinMax(cc)
	rightMin, rightMax := c.Right.MinMax(cc)

	var minExtreme matter.ConstraintExtreme
	var maxExtreme matter.ConstraintExtreme
	minExtreme = c.operate(leftMin, rightMin)
	maxExtreme = c.operate(leftMax, rightMax)
	return minExtreme, maxExtreme
}

func (c *MathExpressionLimit) operate(left matter.ConstraintExtreme, right matter.ConstraintExtreme) (extreme matter.ConstraintExtreme) {
	switch left.Type {
	case matter.ConstraintExtremeTypeInt64:
		switch right.Type {
		case matter.ConstraintExtremeTypeInt64:
			extreme.Int64 = operate(c.Operand, left.Int64, right.Int64)
			extreme.Type = matter.ConstraintExtremeTypeInt64
		case matter.ConstraintExtremeTypeUInt64:
			if right.UInt64 > math.MaxInt64 {
				break
			}
			extreme.Int64 = operate(c.Operand, left.Int64, int64(right.UInt64))
			extreme.Type = matter.ConstraintExtremeTypeInt64
		default:
		}
	case matter.ConstraintExtremeTypeUInt64:
		switch right.Type {
		case matter.ConstraintExtremeTypeInt64:
			if right.Int64 < 0 {
				break
			}
			extreme.UInt64 = operate(c.Operand, left.UInt64, uint64(right.Int64))
			extreme.Type = matter.ConstraintExtremeTypeUInt64
		case matter.ConstraintExtremeTypeUInt64:
			extreme.UInt64 = operate(c.Operand, left.UInt64, right.UInt64)
			extreme.Type = matter.ConstraintExtremeTypeUInt64
		default:
		}
	default:
	}
	if extreme.Type != matter.ConstraintExtremeTypeUndefined {
		if left.Format == right.Format {
			extreme.Format = left.Format
		} else {
			extreme.Format = matter.ConstraintExtremeFormatAuto
		}
	}
	return
}
