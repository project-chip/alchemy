package constraint

import (
	"encoding/json"

	"github.com/project-chip/alchemy/matter/types"
)

type StatusCodeLimit struct {
	StatusCode types.StatusCode
}

func (c *StatusCodeLimit) ASCIIDocString(dataType *types.DataType) string {
	return c.StatusCode.String()
}

func (c *StatusCodeLimit) DataModelString(dataType *types.DataType) string {
	return c.ASCIIDocString(dataType)
}

func (c *StatusCodeLimit) Equal(o Limit) bool {
	_, ok := o.(*NullLimit)
	return ok
}

func (c *StatusCodeLimit) Min(cc Context) (min types.DataTypeExtreme) {
	return types.DataTypeExtreme{UInt64: uint64(c.StatusCode), Type: types.DataTypeExtremeTypeUInt64, Format: types.NumberFormatAuto}
}

func (c *StatusCodeLimit) Max(cc Context) (max types.DataTypeExtreme) {
	return c.Min(cc)
}

func (c *StatusCodeLimit) Fallback(cc Context) (max types.DataTypeExtreme) {
	return c.Min(cc)
}

func (c *StatusCodeLimit) NeedsParens(topLevel bool) bool {
	return false
}

func (c *StatusCodeLimit) Clone() Limit {
	return &NullLimit{}
}

func (c *StatusCodeLimit) MarshalJSON() ([]byte, error) {
	js := map[string]any{
		"type":       "statusCode",
		"statusCode": c.StatusCode,
	}
	return json.Marshal(js)
}
