package constraint

import (
	"encoding/json"
	"strings"

	"github.com/project-chip/alchemy/matter/types"
)

type ReferenceLimit struct {
	Reference string `json:"reference"`
	Entity    types.Entity
	Label     string `json:"label"`
	Field     Limit  `json:"field,omitempty"`
}

func (c *ReferenceLimit) ASCIIDocString(dataType *types.DataType) string {
	var s strings.Builder
	s.WriteString("<<")
	s.WriteString(c.Reference)
	if len(c.Label) > 0 {
		s.WriteString(", ")
		s.WriteString(c.Label)
	}
	s.WriteString(">>")
	if c.Field != nil {
		s.WriteRune('.')
		s.WriteString(c.Field.ASCIIDocString(dataType))
	}
	return s.String()
}

func (c *ReferenceLimit) DataModelString(dataType *types.DataType) string {
	return c.ASCIIDocString(dataType)
}

func (c *ReferenceLimit) Equal(o Limit) bool {
	if oc, ok := o.(*ReferenceLimit); ok {
		return oc.Reference == c.Reference
	}
	return false
}

func (c *ReferenceLimit) Min(cc Context) (min types.DataTypeExtreme) {
	min = cc.MinEntityValue(c.Entity, c.Field)
	return
}

func (c *ReferenceLimit) Max(cc Context) (max types.DataTypeExtreme) {
	max = cc.MaxEntityValue(c.Entity, c.Field)
	return
}

func (c *ReferenceLimit) Fallback(cc Context) (def types.DataTypeExtreme) {
	return cc.Fallback(c.Entity, c.Field)
}

func (c *ReferenceLimit) Clone() Limit {
	return &ReferenceLimit{Reference: c.Reference}
}

func (c *ReferenceLimit) MarshalJSON() ([]byte, error) {
	js := map[string]any{
		"type":      "reference",
		"reference": c.Reference,
	}
	return json.Marshal(js)
}
