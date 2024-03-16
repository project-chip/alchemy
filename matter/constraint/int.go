package constraint

import (
	"encoding/json"
	"strconv"

	"github.com/hasty/alchemy/matter/types"
)

type IntLimit struct {
	Value int64
}

func (c *IntLimit) AsciiDocString(dataType *types.DataType) string {
	return strconv.FormatInt(c.Value, 10)
}

func (c *IntLimit) DataModelString(dataType *types.DataType) string {
	e := c.value()
	return e.DataModelString(dataType)
}

func (c *IntLimit) Equal(o ConstraintLimit) bool {
	if oc, ok := o.(*IntLimit); ok {
		return oc.Value == c.Value
	}
	return false
}

func (c *IntLimit) value() types.DataTypeExtreme {
	return types.DataTypeExtreme{
		Type:   types.DataTypeExtremeTypeInt64,
		Format: types.NumberFormatInt,
		Int64:  c.Value,
	}
}

func (c *IntLimit) Min(cc Context) (min types.DataTypeExtreme) {
	return c.value()
}

func (c *IntLimit) Max(cc Context) (max types.DataTypeExtreme) {
	return c.value()
}

func (c *IntLimit) Default(cc Context) (max types.DataTypeExtreme) {
	return c.value()
}

func (c *IntLimit) Clone() ConstraintLimit {
	return &IntLimit{Value: c.Value}
}

func (c *IntLimit) MarshalJSON() ([]byte, error) {
	js := map[string]any{
		"type":  "integer",
		"value": c.Value,
	}
	return json.Marshal(js)
}
