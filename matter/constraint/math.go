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

func (c *MathExpressionLimit) Min(cc *matter.ConstraintContext) matter.DataTypeExtreme {
	leftMin := c.Left.Min(cc)
	rightMin := c.Right.Min(cc)
	return c.operate(leftMin, rightMin)
}

func (c *MathExpressionLimit) Max(cc *matter.ConstraintContext) matter.DataTypeExtreme {
	leftMax := c.Left.Max(cc)
	rightMax := c.Right.Max(cc)
	return c.operate(leftMax, rightMax)
}

func (c *MathExpressionLimit) operate(left matter.DataTypeExtreme, right matter.DataTypeExtreme) (extreme matter.DataTypeExtreme) {
	switch left.Type {
	case matter.DataTypeExtremeTypeInt64:
		switch right.Type {
		case matter.DataTypeExtremeTypeInt64:
			extreme.Int64 = operate(c.Operand, left.Int64, right.Int64)
			extreme.Type = matter.DataTypeExtremeTypeInt64
		case matter.DataTypeExtremeTypeUInt64:
			if right.UInt64 > math.MaxInt64 {
				break
			}
			extreme.Int64 = operate(c.Operand, left.Int64, int64(right.UInt64))
			extreme.Type = matter.DataTypeExtremeTypeInt64
		default:
		}
	case matter.DataTypeExtremeTypeUInt64:
		switch right.Type {
		case matter.DataTypeExtremeTypeInt64:
			if right.Int64 < 0 {
				break
			}
			extreme.UInt64 = operate(c.Operand, left.UInt64, uint64(right.Int64))
			extreme.Type = matter.DataTypeExtremeTypeUInt64
		case matter.DataTypeExtremeTypeUInt64:
			extreme.UInt64 = operate(c.Operand, left.UInt64, right.UInt64)
			extreme.Type = matter.DataTypeExtremeTypeUInt64
		default:
		}
	default:
	}
	if extreme.Type != matter.DataTypeExtremeTypeUndefined {
		if left.Format == right.Format {
			extreme.Format = left.Format
		} else {
			extreme.Format = matter.NumberFormatAuto
		}
	}
	return
}

func (c *MathExpressionLimit) Default(cc *matter.ConstraintContext) (max matter.DataTypeExtreme) {
	return c.Min(cc)
}

func (c *MathExpressionLimit) Clone() matter.ConstraintLimit {
	return &MathExpressionLimit{Operand: c.Operand, Left: c.Left.Clone(), Right: c.Right.Clone()}
}
