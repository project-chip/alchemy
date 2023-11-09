package matter

import (
	"fmt"
	"math/big"
	"strconv"
)

type Constraint interface {
	AsciiDocString() string
	Equal(o Constraint) bool
	MinMax(fs FieldSet) (min ConstraintExtreme, max ConstraintExtreme)
}

type ConstraintLimit interface {
	AsciiDocString() string
	Equal(o ConstraintLimit) bool
	MinMax(fs FieldSet) (min ConstraintExtreme, max ConstraintExtreme)
}

type ConstraintExtreme struct {
	Defined bool
	Value   int64
}

type AllConstraint struct {
	Value string
}

func (c *AllConstraint) AsciiDocString() string {
	return c.Value
}

func (c *AllConstraint) Equal(o Constraint) bool {
	_, ok := o.(*AllConstraint)
	return ok
}

func (c *AllConstraint) MinMax(fs FieldSet) (min ConstraintExtreme, max ConstraintExtreme) {
	return ConstraintExtreme{}, ConstraintExtreme{}
}

type DescribedConstraint struct {
}

func (c *DescribedConstraint) AsciiDocString() string {
	return "desc"
}

func (c *DescribedConstraint) Equal(o Constraint) bool {
	_, ok := o.(*DescribedConstraint)
	return ok
}

func (c *DescribedConstraint) MinMax(fs FieldSet) (min ConstraintExtreme, max ConstraintExtreme) {
	return ConstraintExtreme{}, ConstraintExtreme{}
}

type MinConstraint struct {
	Min ConstraintLimit `json:"min"`
}

func (c *MinConstraint) AsciiDocString() string {
	return fmt.Sprintf("min %s", c.Min.AsciiDocString())
}

func (c *MinConstraint) Equal(o Constraint) bool {
	if oc, ok := o.(*MinConstraint); ok {
		return oc.Min.Equal(c.Min)
	}
	return false
}

func (c *MinConstraint) MinMax(fs FieldSet) (min ConstraintExtreme, max ConstraintExtreme) {
	return ConstraintExtreme{}, ConstraintExtreme{}
}

type MaxConstraint struct {
	Max ConstraintLimit
}

func (c *MaxConstraint) AsciiDocString() string {
	return fmt.Sprintf("max %s", c.Max.AsciiDocString())
}

func (c *MaxConstraint) Equal(o Constraint) bool {
	if oc, ok := o.(*MaxConstraint); ok {
		return oc.Max.Equal(c.Max)
	}
	return false
}

func (c *MaxConstraint) MinMax(fs FieldSet) (min ConstraintExtreme, max ConstraintExtreme) {
	return ConstraintExtreme{}, ConstraintExtreme{}
}

type RangeConstraint struct {
	Min ConstraintLimit `json:"min"`
	Max ConstraintLimit `json:"max"`
}

func (c *RangeConstraint) AsciiDocString() string {
	return fmt.Sprintf("%s to %s", c.Min.AsciiDocString(), c.Max.AsciiDocString())
}

func (c *RangeConstraint) Equal(o Constraint) bool {
	if oc, ok := o.(*RangeConstraint); ok {
		return oc.Min.Equal(c.Min) && oc.Max.Equal(c.Max)
	}
	return false
}

func (c *RangeConstraint) MinMax(fs FieldSet) (min ConstraintExtreme, max ConstraintExtreme) {
	return ConstraintExtreme{}, ConstraintExtreme{}
}

type ListConstraint struct {
	Constraint      Constraint
	EntryConstraint Constraint
}

func (c *ListConstraint) AsciiDocString() string {
	return fmt.Sprintf("%s[%s]", c.Constraint.AsciiDocString(), c.EntryConstraint.AsciiDocString())
}

func (c *ListConstraint) Equal(o Constraint) bool {
	if oc, ok := o.(*ListConstraint); ok {
		return oc.Constraint.Equal(c.Constraint) && oc.EntryConstraint.Equal(c.EntryConstraint)
	}
	return false
}

func (c *ListConstraint) MinMax(fs FieldSet) (min ConstraintExtreme, max ConstraintExtreme) {
	return ConstraintExtreme{}, ConstraintExtreme{}
}

type GenericConstraint struct {
	Value string
}

func (c *GenericConstraint) AsciiDocString() string {
	return c.Value
}

func (c *GenericConstraint) Equal(o Constraint) bool {
	if oc, ok := o.(*GenericConstraint); ok {
		return oc.Value == c.Value
	}
	return false
}

func (c *GenericConstraint) MinMax(fs FieldSet) (min ConstraintExtreme, max ConstraintExtreme) {
	return ConstraintExtreme{}, ConstraintExtreme{}
}

type IntLimit struct {
	Value int64
}

func (c *IntLimit) AsciiDocString() string {
	return strconv.FormatInt(c.Value, 10)
}

func (c *IntLimit) Equal(o ConstraintLimit) bool {
	if oc, ok := o.(*IntLimit); ok {
		return oc.Value == c.Value
	}
	return false
}

func (c *IntLimit) MinMax(fs FieldSet) (min ConstraintExtreme, max ConstraintExtreme) {
	return ConstraintExtreme{Defined: true, Value: c.Value}, ConstraintExtreme{Defined: true, Value: c.Value}
}

type HexLimit struct {
	Value uint64
}

func (c *HexLimit) AsciiDocString() string {
	return fmt.Sprintf("0x%X", c.Value)
}

func (c *HexLimit) Equal(o ConstraintLimit) bool {
	if oc, ok := o.(*HexLimit); ok {
		return oc.Value == c.Value
	}
	return false
}

func (c *HexLimit) MinMax(fs FieldSet) (min ConstraintExtreme, max ConstraintExtreme) {
	return ConstraintExtreme{Defined: true, Value: int64(c.Value)}, ConstraintExtreme{Defined: true, Value: int64(c.Value)}
}

type PercentLimit struct {
	Value      *big.Float
	Hundredths bool
}

func (c *PercentLimit) AsciiDocString() string {
	if c.Hundredths {
		if c.Value.IsInt() {
			i, _ := c.Value.Uint64()
			return strconv.FormatUint(i/100, 10)
		}
		return fmt.Sprintf("%.2f", c.Value)
	}
	i, _ := c.Value.Uint64()
	return strconv.FormatUint(i, 10)
}

func (c *PercentLimit) Equal(o ConstraintLimit) bool {
	if oc, ok := o.(*PercentLimit); ok {
		return oc.Value == c.Value && oc.Hundredths == c.Hundredths
	}
	return false
}

func (c *PercentLimit) MinMax(fs FieldSet) (min ConstraintExtreme, max ConstraintExtreme) {
	i, _ := c.Value.Int64()
	return ConstraintExtreme{Defined: true, Value: i}, ConstraintExtreme{Defined: true, Value: i}
}

type TemperatureLimit struct {
	Value *big.Float
}

func (c *TemperatureLimit) AsciiDocString() string {
	return c.Value.String() + "°C"
}

func (c *TemperatureLimit) Equal(o ConstraintLimit) bool {
	if oc, ok := o.(*TemperatureLimit); ok {
		return oc.Value == c.Value
	}
	return false
}

func (c *TemperatureLimit) MinMax(fs FieldSet) (min ConstraintExtreme, max ConstraintExtreme) {
	return ConstraintExtreme{}, ConstraintExtreme{}
}

type ReferenceLimit struct {
	Value string
}

func (c *ReferenceLimit) AsciiDocString() string {
	return c.Value
}

func (c *ReferenceLimit) Equal(o ConstraintLimit) bool {
	if oc, ok := o.(*ReferenceLimit); ok {
		return oc.Value == c.Value
	}
	return false
}

func (c *ReferenceLimit) MinMax(fs FieldSet) (min ConstraintExtreme, max ConstraintExtreme) {
	r := fs.GetField(c.Value)
	if r == nil || r.Constraint == nil {
		return ConstraintExtreme{}, ConstraintExtreme{}
	}
	return r.Constraint.MinMax(fs)
}

type MathExpressionLimit struct {
	Operand string
	Left    ConstraintLimit
	Right   ConstraintLimit
}

func (c *MathExpressionLimit) AsciiDocString() string {
	return fmt.Sprintf("(%s %s %s)", c.Left.AsciiDocString(), c.Operand, c.Right.AsciiDocString())
}

func (c *MathExpressionLimit) Equal(o ConstraintLimit) bool {
	if oc, ok := o.(*MathExpressionLimit); ok {
		return oc.Operand == c.Operand && oc.Left.Equal(c.Left) && oc.Right.Equal(c.Right)
	}
	return false
}

func (c *MathExpressionLimit) MinMax(fs FieldSet) (ConstraintExtreme, ConstraintExtreme) {
	leftMin, leftMax := c.Left.MinMax(fs)
	rightMin, rightMax := c.Right.MinMax(fs)
	var min, max int64
	switch c.Operand {
	case "+":
		min = leftMin.Value + rightMin.Value
		max = leftMax.Value + rightMax.Value
	case "-":
		min = leftMin.Value - rightMin.Value
		max = leftMax.Value - rightMax.Value
	case "*":
		min = leftMin.Value * rightMin.Value
		max = leftMax.Value * rightMax.Value
	case "/":
		min = leftMin.Value / rightMin.Value
		max = leftMax.Value / rightMax.Value
	}
	var minExtreme ConstraintExtreme
	var maxExtreme ConstraintExtreme
	if leftMin.Defined && rightMin.Defined {
		minExtreme.Defined = true
		minExtreme.Value = min
	}
	if leftMax.Defined && rightMax.Defined {
		maxExtreme.Defined = true
		maxExtreme.Value = max
	}
	return minExtreme, maxExtreme
}

type ManufacturerLimit struct {
	Value string
}

func (c *ManufacturerLimit) AsciiDocString() string {
	return c.Value
}

func (c *ManufacturerLimit) Equal(o ConstraintLimit) bool {
	if oc, ok := o.(*ManufacturerLimit); ok {
		return oc.Value == c.Value
	}
	return false
}

func (c *ManufacturerLimit) MinMax(fs FieldSet) (ConstraintExtreme, ConstraintExtreme) {
	return ConstraintExtreme{}, ConstraintExtreme{}
}
