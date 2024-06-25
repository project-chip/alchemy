package constraint

import (
	"encoding/json"

	"github.com/project-chip/alchemy/matter/types"
)

type EmptyLimit struct {
}

func (c *EmptyLimit) ASCIIDocString(dataType *types.DataType) string {
	return "empty"
}

func (c *EmptyLimit) DataModelString(dataType *types.DataType) string {
	return "empty"
}

func (c *EmptyLimit) Equal(o Limit) bool {
	_, ok := o.(*EmptyLimit)
	return ok
}

func (c *EmptyLimit) Min(cc Context) (min types.DataTypeExtreme) {
	return types.DataTypeExtreme{Type: types.DataTypeExtremeTypeEmpty, Format: types.NumberFormatHex}
}

func (c *EmptyLimit) Max(cc Context) (max types.DataTypeExtreme) {
	return c.Min(cc)
}

func (c *EmptyLimit) Default(cc Context) (max types.DataTypeExtreme) {
	return c.Min(cc)
}

func (c *EmptyLimit) Clone() Limit {
	return &EmptyLimit{}
}

func (c *EmptyLimit) MarshalJSON() ([]byte, error) {
	js := map[string]any{
		"type": "empty",
	}
	return json.Marshal(js)
}
