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

func (c *MathExpressionLimit) AsciiDocString() string {
	return fmt.Sprintf("(%s %s %s)", c.Left.AsciiDocString(), c.Operand, c.Right.AsciiDocString())
}

func (c *MathExpressionLimit) Equal(o matter.ConstraintLimit) bool {
	if oc, ok := o.(*MathExpressionLimit); ok {
		return oc.Operand == c.Operand && oc.Left.Equal(c.Left) && oc.Right.Equal(c.Right)
	}
	return false
}

func add(left matter.ConstraintExtreme, right matter.ConstraintExtreme) (result matter.ConstraintExtreme) {

	switch left.Type {
	case matter.ConstraintExtremeTypeInt64:
		switch right.Type {
		case matter.ConstraintExtremeTypeInt64:
			if left.Int64 > (math.MaxInt64 - right.Int64) {
				// We'd overflow if we added without casting
			}
		}
	case matter.ConstraintExtremeTypeUInt64:
	case matter.ConstraintExtremeTypeUndefined:
		return
	}
	return
}

func (c *MathExpressionLimit) MinMax(cc *matter.ConstraintContext) (matter.ConstraintExtreme, matter.ConstraintExtreme) {
	leftMin, leftMax := c.Left.MinMax(cc)
	rightMin, rightMax := c.Right.MinMax(cc)
	var min, max int64
	switch c.Operand {
	case "+":
		min = leftMin.Int64 + rightMin.Int64
		max = leftMax.Int64 + rightMax.Int64
	case "-":
		min = leftMin.Int64 - rightMin.Int64
		max = leftMax.Int64 - rightMax.Int64
	case "*":
		min = leftMin.Int64 * rightMin.Int64
		max = leftMax.Int64 * rightMax.Int64
	case "/":
		min = leftMin.Int64 / rightMin.Int64
		max = leftMax.Int64 / rightMax.Int64
	}
	var minExtreme matter.ConstraintExtreme
	var maxExtreme matter.ConstraintExtreme
	if leftMin.Defined() && rightMin.Defined() {
		minExtreme.Type = matter.ConstraintExtremeTypeInt64
		minExtreme.Int64 = min
	}
	if leftMax.Defined() && rightMax.Defined() {
		maxExtreme.Type = matter.ConstraintExtremeTypeInt64
		maxExtreme.Int64 = max
	}
	return minExtreme, maxExtreme
}
