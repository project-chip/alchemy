package matter

import (
	"fmt"
	"strconv"
)

type Constraint interface {
	AsciiDocString() string
	Equal(o Constraint) bool
}

type ConstraintLimit interface {
	AsciiDocString() string
	ZCLString() string
	Equal(o ConstraintLimit) bool
}

type AllConstraint struct {
	Value string
}

func (c *AllConstraint) AsciiDocString() string {
	return c.Value
}

func (c *AllConstraint) ZCLString() string {
	return ""
}

func (c *AllConstraint) Equal(o Constraint) bool {
	_, ok := o.(*AllConstraint)
	return ok
}

type DescribedConstraint struct {
}

func (c *DescribedConstraint) AsciiDocString() string {
	return "desc"
}

func (c *DescribedConstraint) ZCLString() string {
	return ""
}

func (c *DescribedConstraint) Equal(o Constraint) bool {
	_, ok := o.(*DescribedConstraint)
	return ok
}

type IntLimit struct {
	Value int64
}

func (c *IntLimit) AsciiDocString() string {
	return strconv.FormatInt(c.Value, 10)
}

func (c *IntLimit) ZCLString() string {
	return strconv.FormatInt(c.Value, 10)
}

func (c *IntLimit) Equal(o ConstraintLimit) bool {
	if oc, ok := o.(*IntLimit); ok {
		return oc.Value == c.Value
	}
	return false
}

type HexLimit struct {
	Value uint64
}

func (c *HexLimit) AsciiDocString() string {
	return fmt.Sprintf("0x%x", c.Value)
}

func (c *HexLimit) ZCLString() string {
	return fmt.Sprintf("0x%x", c.Value)
}

func (c *HexLimit) Equal(o ConstraintLimit) bool {
	if oc, ok := o.(*HexLimit); ok {
		return oc.Value == c.Value
	}
	return false
}

type PercentLimit struct {
	Value uint64
}

func (c *PercentLimit) AsciiDocString() string {
	return strconv.FormatUint(c.Value, 10)
}

func (c *PercentLimit) ZCLString() string {
	return fmt.Sprintf("0x%x", c.Value)
}

func (c *PercentLimit) Equal(o ConstraintLimit) bool {
	if oc, ok := o.(*PercentLimit); ok {
		return oc.Value == c.Value
	}
	return false
}

type TemperatureLimit struct {
	Value uint64
}

func (c *TemperatureLimit) AsciiDocString() string {
	return strconv.FormatUint(c.Value, 10)
}

func (c *TemperatureLimit) ZCLString() string {
	// TODO: Wrong encoding
	return fmt.Sprintf("0x%x", c.Value)
}

func (c *TemperatureLimit) Equal(o ConstraintLimit) bool {
	if oc, ok := o.(*TemperatureLimit); ok {
		return oc.Value == c.Value
	}
	return false
}

type MaxLengthConstraint struct {
	Length ConstraintLimit
}

func (c *MaxLengthConstraint) AsciiDocString() string {
	return fmt.Sprintf("max %s", c.Length.AsciiDocString())
}

func (c *MaxLengthConstraint) Equal(o Constraint) bool {
	if oc, ok := o.(*MaxLengthConstraint); ok {
		return oc.Length.Equal(c.Length)
	}
	return false
}

type MinLengthConstraint struct {
	Length ConstraintLimit
}

func (c *MinLengthConstraint) AsciiDocString() string {
	return fmt.Sprintf("min %s", c.Length.AsciiDocString())
}

func (c *MinLengthConstraint) Equal(o Constraint) bool {
	if oc, ok := o.(*MinLengthConstraint); ok {
		return oc.Length.Equal(c.Length)
	}
	return false
}

type LengthRangeConstraint struct {
	Min ConstraintLimit
	Max ConstraintLimit
}

func (c *LengthRangeConstraint) AsciiDocString() string {
	return fmt.Sprintf("%s to %s", c.Min.AsciiDocString(), c.Max.AsciiDocString())
}

func (c *LengthRangeConstraint) Equal(o Constraint) bool {
	if oc, ok := o.(*LengthRangeConstraint); ok {
		return oc.Min.Equal(c.Min) && oc.Max.Equal(c.Max)
	}
	return false
}

type MinConstraint struct {
	Min ConstraintLimit
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

type RangeConstraint struct {
	Min ConstraintLimit
	Max ConstraintLimit
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
