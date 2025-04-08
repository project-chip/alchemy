package constraint

import (
	"encoding/json"
	"strings"

	"github.com/project-chip/alchemy/matter/types"
)

type TagListConstraint struct {
	Tags Limit
}

func (tlc *TagListConstraint) Type() Type {
	return ConstraintTypeTagList
}

func (tlc *TagListConstraint) ASCIIDocString(dataType *types.DataType) string {
	var s strings.Builder
	s.WriteString("Includes ")
	requiresParens := tlc.Tags.NeedsParens(false)
	if requiresParens {
		s.WriteString("(")
	}
	s.WriteString(tlc.Tags.ASCIIDocString(dataType))
	if requiresParens {
		s.WriteString(")")
	}
	s.WriteString(" tags")
	return s.String()
}

func (tlc *TagListConstraint) Equal(o Constraint) bool {
	otlc, ok := o.(*TagListConstraint)
	if !ok {
		return false
	}
	return tlc.Tags.Equal(otlc.Tags)
}

func (tlc *TagListConstraint) Min(c Context) (min types.DataTypeExtreme) {
	return types.DataTypeExtreme{Type: types.DataTypeExtremeTypeUndefined, Format: types.NumberFormatAuto}
}

func (tlc *TagListConstraint) Max(c Context) (max types.DataTypeExtreme) {
	return tlc.Min(c)
}

func (tlc *TagListConstraint) NeedsParens(topLevel bool) bool {
	return false
}

func (tlc *TagListConstraint) Clone() Constraint {
	return &TagListConstraint{Tags: tlc.Tags.Clone()}
}

type TagIdentifierLimit struct {
	Tag    string `json:"tag"`
	Entity types.Entity
}

func (c *TagIdentifierLimit) ASCIIDocString(dataType *types.DataType) string {
	var s strings.Builder
	s.WriteString("`")
	s.WriteString(c.Tag)
	s.WriteString("`")
	return s.String()
}

func (c *TagIdentifierLimit) DataModelString(dataType *types.DataType) string {
	return c.Tag
}

func (c *TagIdentifierLimit) Equal(o Limit) bool {
	if oc, ok := o.(*TagIdentifierLimit); ok {
		return oc.Tag == c.Tag
	}
	return false
}

func (c *TagIdentifierLimit) Min(cc Context) (min types.DataTypeExtreme) {
	return
}

func (c *TagIdentifierLimit) Max(cc Context) (max types.DataTypeExtreme) {
	return
}

func (c *TagIdentifierLimit) Fallback(cc Context) (def types.DataTypeExtreme) {
	return
}

func (c *TagIdentifierLimit) NeedsParens(topLevel bool) bool {
	return false
}

func (c *TagIdentifierLimit) Clone() Limit {
	return &TagIdentifierLimit{Tag: c.Tag}
}

func (c *TagIdentifierLimit) MarshalJSON() ([]byte, error) {
	js := map[string]any{
		"type": "tag",
		"tag":  c.Tag,
	}
	return json.Marshal(js)
}
