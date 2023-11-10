package constraint

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/hasty/alchemy/matter"
)

func ParseConstraint(constraint string) matter.Constraint {
	c, err := ParseReader("", strings.NewReader(constraint))
	if err != nil {
		return &GenericConstraint{Value: constraint}
	}
	return c.(matter.Constraint)
}

func parseFloat(s string) (*big.Float, error) {
	var z big.Float
	f, _, err := (&z).Parse(s, 10)
	if err != nil {
		return nil, err
	}
	return f, nil
}

type AllConstraint struct {
	Field *matter.Field
	Value string
}

func (c *AllConstraint) AsciiDocString() string {
	return c.Value
}

func (c *AllConstraint) Equal(o matter.Constraint) bool {
	_, ok := o.(*AllConstraint)
	return ok
}

func (c *AllConstraint) MinMax(cc *matter.ConstraintContext) (min matter.ConstraintExtreme, max matter.ConstraintExtreme) {
	return
}

type DescribedConstraint struct {
}

func (c *DescribedConstraint) AsciiDocString() string {
	return "desc"
}

func (c *DescribedConstraint) Equal(o matter.Constraint) bool {
	_, ok := o.(*DescribedConstraint)
	return ok
}

func (c *DescribedConstraint) MinMax(cc *matter.ConstraintContext) (min matter.ConstraintExtreme, max matter.ConstraintExtreme) {
	return
}

type MinConstraint struct {
	Min matter.ConstraintLimit `json:"min"`
}

func (c *MinConstraint) AsciiDocString() string {
	return fmt.Sprintf("min %s", c.Min.AsciiDocString())
}

func (c *MinConstraint) Equal(o matter.Constraint) bool {
	if oc, ok := o.(*MinConstraint); ok {
		return oc.Min.Equal(c.Min)
	}
	return false
}

func (c *MinConstraint) MinMax(cc *matter.ConstraintContext) (min matter.ConstraintExtreme, max matter.ConstraintExtreme) {
	min, _ = c.Min.MinMax(cc)
	return
}

type MaxConstraint struct {
	Max matter.ConstraintLimit
}

func (c *MaxConstraint) AsciiDocString() string {
	return fmt.Sprintf("max %s", c.Max.AsciiDocString())
}

func (c *MaxConstraint) Equal(o matter.Constraint) bool {
	if oc, ok := o.(*MaxConstraint); ok {
		return oc.Max.Equal(c.Max)
	}
	return false
}

func (c *MaxConstraint) MinMax(cc *matter.ConstraintContext) (min matter.ConstraintExtreme, max matter.ConstraintExtreme) {
	_, max = c.Max.MinMax(cc)
	return
}

type RangeConstraint struct {
	Min matter.ConstraintLimit `json:"min"`
	Max matter.ConstraintLimit `json:"max"`
}

func (c *RangeConstraint) AsciiDocString() string {
	return fmt.Sprintf("%s to %s", c.Min.AsciiDocString(), c.Max.AsciiDocString())
}

func (c *RangeConstraint) Equal(o matter.Constraint) bool {
	if oc, ok := o.(*RangeConstraint); ok {
		return oc.Min.Equal(c.Min) && oc.Max.Equal(c.Max)
	}
	return false
}

func (c *RangeConstraint) MinMax(cc *matter.ConstraintContext) (from matter.ConstraintExtreme, to matter.ConstraintExtreme) {
	fromMin, fromMax := c.Min.MinMax(cc)
	toMin, toMax := c.Min.MinMax(cc)
	if fromMin.Defined() || toMin.Defined() {
		from.Type = matter.ConstraintExtremeTypeInt64
		if fromMin.Defined() && toMin.Defined() {
			from.Int64 = min(fromMin.Int64, toMin.Int64)
		} else if fromMin.Defined() {
			from.Int64 = fromMin.Int64
		} else {
			from.Int64 = toMin.Int64
		}
	}
	if fromMax.Defined() || toMax.Defined() {
		to.Type = matter.ConstraintExtremeTypeInt64
		if fromMax.Defined() && toMax.Defined() {
			to.Int64 = min(fromMax.Int64, toMax.Int64)
		} else if fromMax.Defined() {
			to.Int64 = fromMax.Int64
		} else {
			to.Int64 = toMax.Int64
		}
	}
	return
}

type ListConstraint struct {
	Constraint      matter.Constraint
	EntryConstraint matter.Constraint
}

func (c *ListConstraint) AsciiDocString() string {
	return fmt.Sprintf("%s[%s]", c.Constraint.AsciiDocString(), c.EntryConstraint.AsciiDocString())
}

func (c *ListConstraint) Equal(o matter.Constraint) bool {
	if oc, ok := o.(*ListConstraint); ok {
		return oc.Constraint.Equal(c.Constraint) && oc.EntryConstraint.Equal(c.EntryConstraint)
	}
	return false
}

func (c *ListConstraint) MinMax(cc *matter.ConstraintContext) (min matter.ConstraintExtreme, max matter.ConstraintExtreme) {
	return c.Constraint.MinMax(cc)
}

type GenericConstraint struct {
	Value string
}

func (c *GenericConstraint) AsciiDocString() string {
	return c.Value
}

func (c *GenericConstraint) Equal(o matter.Constraint) bool {
	if oc, ok := o.(*GenericConstraint); ok {
		return oc.Value == c.Value
	}
	return false
}

func (c *GenericConstraint) MinMax(cc *matter.ConstraintContext) (min matter.ConstraintExtreme, max matter.ConstraintExtreme) {
	return
}

type IntLimit struct {
	Value int64
}

func (c *IntLimit) AsciiDocString() string {
	return strconv.FormatInt(c.Value, 10)
}

func (c *IntLimit) Equal(o matter.ConstraintLimit) bool {
	if oc, ok := o.(*IntLimit); ok {
		return oc.Value == c.Value
	}
	return false
}

func (c *IntLimit) MinMax(cc *matter.ConstraintContext) (min matter.ConstraintExtreme, max matter.ConstraintExtreme) {
	return matter.ConstraintExtreme{
			Type:  matter.ConstraintExtremeTypeInt64,
			Int64: c.Value},
		matter.ConstraintExtreme{
			Type:  matter.ConstraintExtremeTypeInt64,
			Int64: c.Value,
		}
}

type HexLimit struct {
	Value uint64
}

func (c *HexLimit) AsciiDocString() string {
	return fmt.Sprintf("0x%X", c.Value)
}

func (c *HexLimit) Equal(o matter.ConstraintLimit) bool {
	if oc, ok := o.(*HexLimit); ok {
		return oc.Value == c.Value
	}
	return false
}

func (c *HexLimit) MinMax(cc *matter.ConstraintContext) (min matter.ConstraintExtreme, max matter.ConstraintExtreme) {
	return matter.ConstraintExtreme{
			Type:   matter.ConstraintExtremeTypeUInt64,
			UInt64: c.Value},
		matter.ConstraintExtreme{
			Type:   matter.ConstraintExtremeTypeUInt64,
			UInt64: c.Value,
		}
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

func (c *PercentLimit) Equal(o matter.ConstraintLimit) bool {
	if oc, ok := o.(*PercentLimit); ok {
		return oc.Value == c.Value && oc.Hundredths == c.Hundredths
	}
	return false
}

func (c *PercentLimit) MinMax(cc *matter.ConstraintContext) (min matter.ConstraintExtreme, max matter.ConstraintExtreme) {
	i, _ := c.Value.Int64()
	return matter.ConstraintExtreme{
			Type:  matter.ConstraintExtremeTypeInt64,
			Int64: i},
		matter.ConstraintExtreme{
			Type:  matter.ConstraintExtremeTypeInt64,
			Int64: i,
		}
}

type TemperatureLimit struct {
	Value *big.Float
}

func (c *TemperatureLimit) AsciiDocString() string {
	return c.Value.String() + "°C"
}

func (c *TemperatureLimit) Equal(o matter.ConstraintLimit) bool {
	if oc, ok := o.(*TemperatureLimit); ok {
		return oc.Value == c.Value
	}
	return false
}

func (c *TemperatureLimit) MinMax(cc *matter.ConstraintContext) (min matter.ConstraintExtreme, max matter.ConstraintExtreme) {
	return
}

type ReferenceLimit struct {
	Value string
}

func (c *ReferenceLimit) AsciiDocString() string {
	return c.Value
}

func (c *ReferenceLimit) Equal(o matter.ConstraintLimit) bool {
	if oc, ok := o.(*ReferenceLimit); ok {
		return oc.Value == c.Value
	}
	return false
}

func (c *ReferenceLimit) MinMax(cc *matter.ConstraintContext) (min matter.ConstraintExtreme, max matter.ConstraintExtreme) {
	r := cc.Fields.GetField(c.Value)
	if cc.VisitedReferences == nil {
		cc.VisitedReferences = make(map[string]struct{})
	}
	if _, ok := cc.VisitedReferences[c.Value]; ok {
		return
	}
	cc.VisitedReferences[c.Value] = struct{}{}
	if r == nil || r.Constraint == nil {
		return
	}
	return r.Constraint.MinMax(cc)
}

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

type ManufacturerLimit struct {
	Value string
}

func (c *ManufacturerLimit) AsciiDocString() string {
	return c.Value
}

func (c *ManufacturerLimit) Equal(o matter.ConstraintLimit) bool {
	if oc, ok := o.(*ManufacturerLimit); ok {
		return oc.Value == c.Value
	}
	return false
}

func (c *ManufacturerLimit) MinMax(cc *matter.ConstraintContext) (min matter.ConstraintExtreme, max matter.ConstraintExtreme) {
	return
}
