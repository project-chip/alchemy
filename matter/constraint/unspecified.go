package constraint

import (
	"encoding/json"

	"github.com/hasty/alchemy/matter/types"
)

type UnspecifiedLimit struct {
}

func (c *UnspecifiedLimit) AsciiDocString(dataType *types.DataType) string {
	return "-"
}

func (c *UnspecifiedLimit) DataModelString(dataType *types.DataType) string {
	return ""
}

func (c *UnspecifiedLimit) Equal(o Limit) bool {
	_, ok := o.(*UnspecifiedLimit)
	return ok
}

func (c *UnspecifiedLimit) Min(cc Context) (min types.DataTypeExtreme) {
	return types.DataTypeExtreme{Type: types.DataTypeExtremeTypeUndefined, Format: types.NumberFormatAuto}
}

func (c *UnspecifiedLimit) Max(cc Context) (max types.DataTypeExtreme) {
	return c.Min(cc)
}

func (c *UnspecifiedLimit) Default(cc Context) (max types.DataTypeExtreme) {
	return c.Min(cc)
}

func (c *UnspecifiedLimit) Clone() Limit {
	return &UnspecifiedLimit{}
}

func (c *UnspecifiedLimit) MarshalJSON() ([]byte, error) {
	js := map[string]any{
		"type": "unspecified",
	}
	return json.Marshal(js)
}
