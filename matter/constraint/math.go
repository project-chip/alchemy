package constraint

import (
	"encoding/json"
	"fmt"
	"math"

	"github.com/project-chip/alchemy/matter/types"
)

type MathExpressionLimit struct {
	Operand string
	Left    Limit
	Right   Limit
}

func (c *MathExpressionLimit) ASCIIDocString(dataType *types.DataType) string {
	return fmt.Sprintf("(%s %s %s)", c.Left.ASCIIDocString(dataType), c.Operand, c.Right.ASCIIDocString(dataType))
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

func (c *MathExpressionLimit) Min(cc Context) types.DataTypeExtreme {
	leftMin := c.Left.Min(cc)
	rightMin := c.Right.Min(cc)
	return c.operate(leftMin, rightMin)
}

func (c *MathExpressionLimit) Max(cc Context) types.DataTypeExtreme {
	leftMax := c.Left.Max(cc)
	rightMax := c.Right.Max(cc)
	return c.operate(leftMax, rightMax)
}

func (c *MathExpressionLimit) operate(left types.DataTypeExtreme, right types.DataTypeExtreme) (extreme types.DataTypeExtreme) {
	switch left.Type {
	case types.DataTypeExtremeTypeInt64:
		switch right.Type {
		case types.DataTypeExtremeTypeInt64:
			extreme.Int64 = operate(c.Operand, left.Int64, right.Int64)
			extreme.Type = types.DataTypeExtremeTypeInt64
		case types.DataTypeExtremeTypeUInt64:
			if right.UInt64 > math.MaxInt64 {
				break
			}
			extreme.Int64 = operate(c.Operand, left.Int64, int64(right.UInt64))
			extreme.Type = types.DataTypeExtremeTypeInt64
		default:
		}
	case types.DataTypeExtremeTypeUInt64:
		switch right.Type {
		case types.DataTypeExtremeTypeInt64:
			if right.Int64 < 0 {
				break
			}
			extreme.UInt64 = operate(c.Operand, left.UInt64, uint64(right.Int64))
			extreme.Type = types.DataTypeExtremeTypeUInt64
		case types.DataTypeExtremeTypeUInt64:
			extreme.UInt64 = operate(c.Operand, left.UInt64, right.UInt64)
			extreme.Type = types.DataTypeExtremeTypeUInt64
		default:
		}
	default:
	}
	if extreme.Type != types.DataTypeExtremeTypeUndefined {
		if left.Format == right.Format {
			extreme.Format = left.Format
		} else {
			extreme.Format = types.NumberFormatAuto
		}
	}
	return
}

func (c *MathExpressionLimit) Fallback(cc Context) (max types.DataTypeExtreme) {
	return c.Min(cc)
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
