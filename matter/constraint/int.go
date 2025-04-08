package constraint

import (
	"encoding/json"
	"strconv"

	"github.com/project-chip/alchemy/matter/types"
)

type IntLimit struct {
	Value int64 `json:"value"`
}

func (c *IntLimit) ASCIIDocString(dataType *types.DataType) string {
	return strconv.FormatInt(c.Value, 10)
}

func (c *IntLimit) DataModelString(dataType *types.DataType) string {
	e := c.value(dataType)
	return e.DataModelString(dataType)
}

func (c *IntLimit) Equal(o Limit) bool {
	if oc, ok := o.(*IntLimit); ok {
		return oc.Value == c.Value
	}
	return false
}

func (c *IntLimit) value(dataType *types.DataType) types.DataTypeExtreme {
	if dataType != nil && dataType.BaseType.IsUnsigned() {
		return types.DataTypeExtreme{
			Type:     types.DataTypeExtremeTypeUInt64,
			Format:   types.NumberFormatInt,
			UInt64:   uint64(c.Value),
			Constant: true,
		}
	}
	return types.DataTypeExtreme{
		Type:     types.DataTypeExtremeTypeInt64,
		Format:   types.NumberFormatInt,
		Int64:    c.Value,
		Constant: true,
	}
}

func (c *IntLimit) Min(cc Context) (min types.DataTypeExtreme) {
	return c.value(cc.DataType())
}

func (c *IntLimit) Max(cc Context) (max types.DataTypeExtreme) {
	return c.value(cc.DataType())
}

func (c *IntLimit) Fallback(cc Context) (max types.DataTypeExtreme) {
	return c.value(cc.DataType())
}

func (c *IntLimit) NeedsParens(topLevel bool) bool {
	return false
}

func (c *IntLimit) Clone() Limit {
	return &IntLimit{Value: c.Value}
}

func (c *IntLimit) MarshalJSON() ([]byte, error) {
	js := map[string]any{
		"type":  "integer",
		"value": c.Value,
	}
	return json.Marshal(js)
}
