package constraint

import (
	"encoding/json"
	"strconv"

	"github.com/project-chip/alchemy/matter/types"
)

type BooleanLimit struct {
	Value bool `json:"value"`
}

func (c *BooleanLimit) ASCIIDocString(dataType *types.DataType) string {
	return strconv.FormatBool(c.Value)
}

func (c *BooleanLimit) DataModelString(dataType *types.DataType) string {
	return strconv.FormatBool(c.Value)
}

func (c *BooleanLimit) Equal(o Limit) bool {
	if oc, ok := o.(*BooleanLimit); ok {
		return oc.Value == c.Value
	}
	return false
}

func (c *BooleanLimit) Min(cc Context) (min types.DataTypeExtreme) {
	var val uint64
	if c.Value {
		val = 1
	}
	return types.DataTypeExtreme{
		Type:   types.DataTypeExtremeTypeUInt64,
		Format: types.NumberFormatInt,
		UInt64: val,
	}
}

func (c *BooleanLimit) Max(cc Context) (max types.DataTypeExtreme) {
	return c.Min(cc)
}

func (c *BooleanLimit) Fallback(cc Context) (max types.DataTypeExtreme) {
	return c.Min(cc)
}

func (c *BooleanLimit) Clone() Limit {
	return &BooleanLimit{Value: c.Value}
}

func (c *BooleanLimit) MarshalJSON() ([]byte, error) {
	js := map[string]any{
		"type":  "boolean",
		"value": c.Value,
	}
	return json.Marshal(js)
}
