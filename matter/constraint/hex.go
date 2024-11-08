package constraint

import (
	"encoding/json"
	"fmt"

	"github.com/project-chip/alchemy/matter/types"
)

type HexLimit struct {
	Value uint64 `json:"value"`
}

func (c *HexLimit) ASCIIDocString(dataType *types.DataType) string {
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

func (c *HexLimit) DataModelString(dataType *types.DataType) string {
	e := c.value()
	return e.DataModelString(dataType)
}

func (c *HexLimit) Equal(o Limit) bool {
	if oc, ok := o.(*HexLimit); ok {
		return oc.Value == c.Value
	}
	return false
}

func (c *HexLimit) value() types.DataTypeExtreme {
	return types.DataTypeExtreme{
		Type:   types.DataTypeExtremeTypeUInt64,
		Format: types.NumberFormatHex,
		UInt64: c.Value,
	}
}

func (c *HexLimit) Min(cc Context) (min types.DataTypeExtreme) {
	return c.value()
}

func (c *HexLimit) Max(cc Context) (max types.DataTypeExtreme) {
	return c.value()
}

func (c *HexLimit) Fallback(cc Context) (max types.DataTypeExtreme) {
	return c.value()
}

func (c *HexLimit) Clone() Limit {
	return &HexLimit{Value: c.Value}
}

func (c *HexLimit) MarshalJSON() ([]byte, error) {
	js := map[string]any{
		"type":  "hex",
		"value": c.Value,
	}
	return json.Marshal(js)
}
