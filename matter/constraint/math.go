package constraint

import (
	"encoding/json"
	"math/big"
	"strings"

	"github.com/project-chip/alchemy/matter/types"
)

type MathExpressionLimit struct {
	Operand string
	Left    Limit
	Right   Limit
}

func (c *MathExpressionLimit) ASCIIDocString(dataType *types.DataType) string {
	var s strings.Builder
	leftRequiresParens := c.Left.NeedsParens(false)
	if leftRequiresParens {
		s.WriteString("(")
	}
	s.WriteString(c.Left.ASCIIDocString(dataType))
	if leftRequiresParens {
		s.WriteString(")")
	}
	s.WriteString(" ")
	s.WriteString(c.Operand)
	s.WriteString(" ")
	rightRequiresParens := c.Right.NeedsParens(false)
	if rightRequiresParens {
		s.WriteString("(")
	}
	s.WriteString(c.Right.ASCIIDocString(dataType))
	if rightRequiresParens {
		s.WriteString(")")
	}
	return s.String()
}

func (c *MathExpressionLimit) DataModelString(dataType *types.DataType) string {
	return c.ASCIIDocString(dataType)
}

func (c *MathExpressionLimit) Equal(o Limit) bool {
	if oc, ok := o.(*MathExpressionLimit); ok {
		return oc.Operand == c.Operand && oc.Left.Equal(c.Left) && oc.Right.Equal(c.Right)
	}
	return false
}

type Number interface {
	int64 | uint64
}

func operate(operand string, left, right *big.Int) (val *big.Int) {
	val = new(big.Int)
	switch operand {
	case "+":
		val.Add(left, right)
	case "-":
		val.Sub(left, right)
	case "*":
		val.Mul(left, right)
	case "/":
		val.Div(left, right)
	}
	return
}

func (c *MathExpressionLimit) Min(cc Context) types.DataTypeExtreme {
	var leftMin, rightMin types.DataTypeExtreme
	leftMin = c.Left.Min(cc)
	rightMin = c.Right.Min(cc)
	var m types.DataTypeExtreme
	switch c.Operand {
	case "+":
		m = c.operate(leftMin, rightMin)
	case "*":
		if leftMin.Constant {
			if rightMin.Constant {
				m = c.operate(leftMin, rightMin)
			} else if leftMin.IsNegative() {
				// if left is a negative constant, multiply by the largest number on the right
				m = c.operate(leftMin, c.Right.Max(cc))
			} else {
				// if left is a positive constant, multiply by the smallest number on the right
				m = c.operate(leftMin, rightMin)
			}
		} else if rightMin.Constant {
			if rightMin.IsNegative() {
				m = c.operate(c.Left.Max(cc), rightMin)
			} else {
				m = c.operate(leftMin, rightMin)
			}
		} else {
			// Both sides are non-constant
			if leftMin.IsNegative() && rightMin.IsNegative() {
				// Multiply both sides and take the smallest
				m = types.MinExtreme(c.operate(leftMin, c.Right.Max(cc)), c.operate(c.Left.Max(cc), rightMin))
			} else {
				m = c.operate(leftMin, rightMin)
			}
		}
	case "-":
		m = c.operate(leftMin, c.Right.Max(cc))
	case "/":
		rightMax := c.Right.Max(cc)
		if leftMin.Constant {
			if rightMax.Constant {
				m = c.operate(leftMin, rightMax)
			} else if leftMin.IsNegative() {
				// If left is a negative constant, divide by the smallest number on the right
				m = c.operate(leftMin, rightMin)
			} else {
				// If left is a positive constant, divide by the largest number on the right
				m = c.operate(leftMin, rightMax)
			}
		} else {
			if leftMin.IsNegative() && rightMax.IsNegative() {
				m = types.MaxExtreme(c.operate(leftMin, rightMax), c.operate(c.Left.Max(cc), rightMin))
			} else {
				m = c.operate(leftMin, rightMax)
			}

		}
	}
	return m
}

func (c *MathExpressionLimit) Max(cc Context) types.DataTypeExtreme {
	var leftMax, rightMax types.DataTypeExtreme
	leftMax = c.Left.Max(cc)
	rightMax = c.Right.Max(cc)
	var m types.DataTypeExtreme
	switch c.Operand {
	case "+":
		m = c.operate(leftMax, rightMax)
	case "*":
		if leftMax.Constant {
			if rightMax.Constant {
				m = c.operate(leftMax, rightMax)
			} else if leftMax.IsNegative() {
				// if left is a negative constant, multiply by the smallest number on the right
				m = c.operate(leftMax, c.Right.Min(cc))
			} else {
				// if left is a positive constant, multiply by the largest number on the right
				m = c.operate(c.Left.Min(cc), rightMax)
			}
		} else if rightMax.Constant {
			if rightMax.IsNegative() {
				// if right is a negative constant, multiply by the smallest number on the left
				m = c.operate(c.Left.Min(cc), rightMax)
			} else {
				m = c.operate(leftMax, rightMax)
			}
		} else {
			// Both sides are non-constant
			if leftMax.IsNegative() && rightMax.IsNegative() {
				// Multiply both sides and take the smallest
				m = types.MaxExtreme(c.operate(leftMax, c.Right.Min(cc)), c.operate(c.Left.Min(cc), rightMax))
			} else {
				m = c.operate(leftMax, rightMax)
			}
		}
	case "-":
		m = c.operate(leftMax, c.Right.Min(cc))
	case "/":
		rightMin := c.Right.Min(cc)
		if leftMax.Constant {
			if rightMin.Constant {
				m = c.operate(leftMax, rightMin)
			} else if leftMax.IsNegative() {
				// If left is a negative constant, divide by the largest number on the right
				m = c.operate(leftMax, rightMax)
			} else {
				// If left is a positive constant, divide by the smallest number on the right
				m = c.operate(leftMax, rightMin)
			}
		} else {
			if leftMax.IsNegative() && rightMin.IsNegative() {
				m = types.MaxExtreme(c.operate(leftMax, rightMin), c.operate(c.Left.Min(cc), rightMax))
			} else {
				m = c.operate(leftMax, rightMin)
			}

		}
	}
	return m
}

func (c *MathExpressionLimit) operate(left types.DataTypeExtreme, right types.DataTypeExtreme) (extreme types.DataTypeExtreme) {
	switch left.Type {
	case types.DataTypeExtremeTypeInt64:
		switch right.Type {
		case types.DataTypeExtremeTypeInt64, types.DataTypeExtremeTypeUInt64:
			extreme = types.NewIntegerDataTypeExtreme(operate(c.Operand, left.Big(), right.Big()))
		default:
		}
	case types.DataTypeExtremeTypeUInt64:
		switch right.Type {
		case types.DataTypeExtremeTypeInt64, types.DataTypeExtremeTypeUInt64:
			extreme = types.NewIntegerDataTypeExtreme(operate(c.Operand, left.Big(), right.Big()))
		default:
		}
	default:
	}
	if extreme.Type != types.DataTypeExtremeTypeUndefined {
		switch left.Format {
		case types.NumberFormatHex:
			extreme.Format = types.NumberFormatHex
		case types.NumberFormatInt:
			switch right.Format {
			case types.NumberFormatInt, types.NumberFormatHex:
				extreme.Format = right.Format
			case types.NumberFormatAuto, types.NumberFormatUndefined:
				extreme.Format = left.Format
			}
		case types.NumberFormatAuto:
			switch right.Format {
			case types.NumberFormatInt, types.NumberFormatHex:
				extreme.Format = right.Format
			case types.NumberFormatAuto, types.NumberFormatUndefined:
				extreme.Format = types.NumberFormatAuto
			}
		}
	}
	return
}

func (c *MathExpressionLimit) Fallback(cc Context) (max types.DataTypeExtreme) {
	return c.Min(cc)
}

func (c *MathExpressionLimit) NeedsParens(topLevel bool) bool {
	return !topLevel
}

func (c *MathExpressionLimit) Clone() Limit {
	return &MathExpressionLimit{Operand: c.Operand, Left: c.Left.Clone(), Right: c.Right.Clone()}
}

func (c *MathExpressionLimit) MarshalJSON() ([]byte, error) {
	js := map[string]any{
		"type":  "math",
		"left":  c.Left,
		"right": c.Right,
	}
	return json.Marshal(js)
}

func (c *MathExpressionLimit) UnmarshalJSON(data []byte) (err error) {
	var js struct {
		Left  json.RawMessage `json:"left"`
		Right json.RawMessage `json:"right"`
	}
	err = json.Unmarshal(data, &js)
	if err != nil {
		return
	}

	c.Left, err = UnmarshalLimit(js.Left)
	if err != nil {
		return
	}
	c.Right, err = UnmarshalLimit(js.Right)
	return
}
