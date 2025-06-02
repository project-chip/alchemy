package constraint

import (
	"encoding/json"
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/matter/types"
)

type IdentifierLimit struct {
	ID     string `json:"id"`
	Not    bool   `json:"not,omitempty"`
	Entity types.Entity
	Field  Limit `json:"field,omitempty"`
}

func (c *IdentifierLimit) ASCIIDocString(dataType *types.DataType) string {
	var s strings.Builder
	s.WriteString(c.ID)
	if c.Field != nil {
		s.WriteRune('.')
		s.WriteString(c.Field.ASCIIDocString(dataType))
	}
	return s.String()
}

func (c *IdentifierLimit) DataModelString(dataType *types.DataType) string {
	return c.ASCIIDocString(dataType)
}

func (c *IdentifierLimit) Equal(o Limit) bool {
	if oc, ok := o.(*IdentifierLimit); ok {
		return oc.ID == c.ID
	}
	return false
}

func (c *IdentifierLimit) Min(cc Context) (min types.DataTypeExtreme) {
	if c.Entity == nil {
		slog.Error("Unable to find min value for identifier", slog.String("id", c.ID))
		return
	}

	min = cc.MinEntityValue(c.Entity, c.Field)
	return
}

func (c *IdentifierLimit) Max(cc Context) (max types.DataTypeExtreme) {
	if c.Entity == nil {
		slog.Error("Unable to find max value for identifier", slog.String("id", c.ID))
		return
	}
	max = cc.MaxEntityValue(c.Entity, c.Field)
	return
}

func (c *IdentifierLimit) Fallback(cc Context) (def types.DataTypeExtreme) {
	return cc.Fallback(c.Entity, c.Field)
}

func (c *IdentifierLimit) NeedsParens(topLevel bool) bool {
	return false
}

func (c *IdentifierLimit) Clone() Limit {
	il := &IdentifierLimit{ID: c.ID}
	if c.Field != nil {
		il.Field = c.Field.Clone()
	}
	return il
}

func (c *IdentifierLimit) MarshalJSON() ([]byte, error) {
	js := map[string]any{
		"type": "identifier",
		"id":   c.ID,
		"prop": c.ID,
	}
	return json.Marshal(js)
}
