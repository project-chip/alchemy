package constraint

import (
	"encoding/json"

	"github.com/hasty/alchemy/matter/types"
)

type ManufacturerLimit struct {
	Value string `json:"value,omitempty"`
}

func (c *ManufacturerLimit) ASCIIDocString(dataType *types.DataType) string {
	return c.Value
}

func (c *ManufacturerLimit) DataModelString(dataType *types.DataType) string {
	return c.Value
}

func (c *ManufacturerLimit) Equal(o Limit) bool {
	if oc, ok := o.(*ManufacturerLimit); ok {
		return oc.Value == c.Value
	}
	return false
}

func (c *ManufacturerLimit) Min(cc Context) (min types.DataTypeExtreme) {
	return
}

func (c *ManufacturerLimit) Max(cc Context) (max types.DataTypeExtreme) {
	return
}

func (c *ManufacturerLimit) Default(cc Context) (max types.DataTypeExtreme) {
	return
}

func (c *ManufacturerLimit) Clone() Limit {
	return &ManufacturerLimit{Value: c.Value}
}

func (c *ManufacturerLimit) MarshalJSON() ([]byte, error) {
	js := map[string]any{
		"type": "manufacturerDefined",
	}
	if len(c.Value) > 0 {
		js["value"] = c.Value
	}
	return json.Marshal(js)
}
