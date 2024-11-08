package constraint

import (
	"encoding/json"

	"github.com/project-chip/alchemy/matter/types"
)

type ReferenceLimit struct {
	Value string `json:"value"`
}

func (c *ReferenceLimit) ASCIIDocString(dataType *types.DataType) string {
	return c.Value
}

func (c *ReferenceLimit) DataModelString(dataType *types.DataType) string {
	return c.ASCIIDocString(dataType)
}

func (c *ReferenceLimit) Equal(o Limit) bool {
	if oc, ok := o.(*ReferenceLimit); ok {
		return oc.Value == c.Value
	}
	return false
}

func (c *ReferenceLimit) Min(cc Context) (min types.DataTypeExtreme) {
	rc := cc.ReferenceConstraint(c.Value)
	if rc == nil {
		return
	}
	return rc.Min(cc)
}

func (c *ReferenceLimit) Max(cc Context) (max types.DataTypeExtreme) {
	rc := cc.ReferenceConstraint(c.Value)
	if rc == nil {
		return
	}
	return rc.Max(cc)
}

func (c *ReferenceLimit) Fallback(cc Context) (def types.DataTypeExtreme) {
	return cc.Fallback(c.Value)
}

func (c *ReferenceLimit) Clone() Limit {
	return &ReferenceLimit{Value: c.Value}
}

func (c *ReferenceLimit) MarshalJSON() ([]byte, error) {
	js := map[string]any{
		"type":  "reference",
		"value": c.Value,
	}
	return json.Marshal(js)
}
