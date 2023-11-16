package constraint

import (
	"fmt"

	"github.com/hasty/alchemy/matter"
)

type HexLimit struct {
	Value uint64
}

func (c *HexLimit) AsciiDocString(dataType *matter.DataType) string {
	val := c.Value
	if dataType != nil {
		switch dataType.Size() {
		case 1:
			return fmt.Sprintf("0x%02X", uint8(val))
		case 2:
			return fmt.Sprintf("0x%04X", uint16(val))
		case 4:
			return fmt.Sprintf("0x%08X", uint32(val))
		case 8:
			return fmt.Sprintf("0x%16X", uint64(val))
		}
	}
	return fmt.Sprintf("0x%X", uint64(val))
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
			Format: matter.ConstraintExtremeFormatHex,
			UInt64: c.Value},
		matter.ConstraintExtreme{
			Type:   matter.ConstraintExtremeTypeUInt64,
			Format: matter.ConstraintExtremeFormatHex,
			UInt64: c.Value,
		}
}
