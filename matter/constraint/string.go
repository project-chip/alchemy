package constraint

import (
	"encoding/json"
	"fmt"

	"github.com/project-chip/alchemy/matter/types"
)

type StringLimit struct {
	Value string `json:"value"`
}

func (c *StringLimit) ASCIIDocString(dataType *types.DataType) string {
	return fmt.Sprintf("\"%s\"", c.Value)
}

func (c *StringLimit) DataModelString(dataType *types.DataType) string {
	return fmt.Sprintf("\"%s\"", c.Value)
}

func (c *StringLimit) Equal(o Limit) bool {
	_, ok := o.(*StringLimit)
	return ok
}

func (c *StringLimit) Min(cc Context) (min types.DataTypeExtreme) {
	return types.DataTypeExtreme{Type: types.DataTypeExtremeTypeUndefined, Format: types.NumberFormatAuto}
}

func (c *StringLimit) Max(cc Context) (max types.DataTypeExtreme) {
	return c.Min(cc)
}

func (c *StringLimit) Fallback(cc Context) (max types.DataTypeExtreme) {
	return c.Min(cc)
}

func (c *StringLimit) NeedsParens(topLevel bool) bool {
	return false
}

func (c *StringLimit) Clone() Limit {
	return &StringLimit{}
}

func (c *StringLimit) MarshalJSON() ([]byte, error) {
	js := map[string]any{
		"type":  "string",
		"value": c.Value,
	}
	return json.Marshal(js)
}
