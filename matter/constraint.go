package matter

import (
	"fmt"
	"strconv"
)

type Constraint interface {
	AsciiDocString() string
}

type ConstraintLimit interface {
	AsciiDocString() string
	ZCLString() string
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

type DescribedConstraint struct {
}

func (c *DescribedConstraint) AsciiDocString() string {
	return "desc"
}

func (c *DescribedConstraint) ZCLString() string {
	return ""
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

type HexLimit struct {
	Value uint64
}

func (c *HexLimit) AsciiDocString() string {
	return fmt.Sprintf("0x%x", c.Value)
}

func (c *HexLimit) ZCLString() string {
	return fmt.Sprintf("0x%x", c.Value)
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

type TemperatureLimit struct {
	Value uint64
}

func (c *TemperatureLimit) AsciiDocString() string {
	return strconv.FormatUint(c.Value, 10)
}

type MaxLengthConstraint struct {
	Length ConstraintLimit
}

func (c *MaxLengthConstraint) AsciiDocString() string {
	return fmt.Sprintf("max %s", c.Length.AsciiDocString())
}

type MinLengthConstraint struct {
	Length ConstraintLimit
}

func (c *MinLengthConstraint) AsciiDocString() string {
	return fmt.Sprintf("min %s", c.Length.AsciiDocString())
}

type LengthRangeConstraint struct {
	Min ConstraintLimit
	Max ConstraintLimit
}

func (c *LengthRangeConstraint) AsciiDocString() string {
	return fmt.Sprintf("%s to %s", c.Min.AsciiDocString(), c.Max.AsciiDocString())
}

type MinConstraint struct {
	Min ConstraintLimit
}

func (c *MinConstraint) AsciiDocString() string {
	return fmt.Sprintf("min %s", c.Min.AsciiDocString())
}

type MaxConstraint struct {
	Max ConstraintLimit
}

func (c *MaxConstraint) AsciiDocString() string {
	return fmt.Sprintf("max %s", c.Max.AsciiDocString())
}

type RangeConstraint struct {
	Min ConstraintLimit
	Max ConstraintLimit
}

func (c *RangeConstraint) AsciiDocString() string {
	return fmt.Sprintf("%s to %s", c.Min.AsciiDocString(), c.Max.AsciiDocString())
}

type ListConstraint struct {
	Constraint      Constraint
	EntryConstraint Constraint
}

func (c *ListConstraint) AsciiDocString() string {
	return fmt.Sprintf("%s[%s]", c.Constraint.AsciiDocString(), c.EntryConstraint.AsciiDocString())
}

type GenericConstraint struct {
	Value string
}

func (c *GenericConstraint) AsciiDocString() string {
	return c.Value
}
