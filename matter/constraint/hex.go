package constraint

import (
	"fmt"

	"github.com/hasty/alchemy/matter/types"
)

type HexLimit struct {
	Value uint64
}

func (c *HexLimit) AsciiDocString(dataType *types.DataType) string {
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

func (c *HexLimit) Equal(o ConstraintLimit) bool {
	if oc, ok := o.(*HexLimit); ok {
		return oc.Value == c.Value
	}
	return false
}

func (c *HexLimit) Min(cc Context) (min types.DataTypeExtreme) {
	return types.DataTypeExtreme{
		Type:   types.DataTypeExtremeTypeUInt64,
		Format: types.NumberFormatHex,
		UInt64: c.Value,
	}
}

func (c *HexLimit) Max(cc Context) (max types.DataTypeExtreme) {
	return c.Min(cc)
}

func (c *HexLimit) Default(cc Context) (max types.DataTypeExtreme) {
	return c.Min(cc)
}

func (c *HexLimit) Clone() ConstraintLimit {
	return &HexLimit{Value: c.Value}
}
