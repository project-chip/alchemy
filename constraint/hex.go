package constraint

import (
	"fmt"

	"github.com/hasty/alchemy/matter"
)

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
