package constraint

import (
	"encoding/json"

	"github.com/hasty/alchemy/matter/types"
)

type NullLimit struct {
}

func (c *NullLimit) AsciiDocString(dataType *types.DataType) string {
	return "null"
}

func (c *NullLimit) DataModelString(dataType *types.DataType) string {
	return c.AsciiDocString(dataType)
}

func (c *NullLimit) Equal(o Limit) bool {
	_, ok := o.(*NullLimit)
	return ok
}

func (c *NullLimit) Min(cc Context) (min types.DataTypeExtreme) {
	return types.DataTypeExtreme{Type: types.DataTypeExtremeTypeNull, Format: types.NumberFormatAuto}
}

func (c *NullLimit) Max(cc Context) (max types.DataTypeExtreme) {
	return c.Min(cc)
}

func (c *NullLimit) Default(cc Context) (max types.DataTypeExtreme) {
	return c.Min(cc)
}

func (c *NullLimit) Clone() Limit {
	return &NullLimit{}
}

func (c *NullLimit) MarshalJSON() ([]byte, error) {
	js := map[string]any{
		"type": "null",
	}
	return json.Marshal(js)
}
