package asciidoc

import (
	"fmt"
	"strings"
)

type AnchorAttribute struct {
	attribute

	ID    *String `json:"id"`
	Label Elements
}

func NewAnchorAttribute(id *String, label Elements) *AnchorAttribute {
	return &AnchorAttribute{ID: id, Label: label}
}

func (a *AnchorAttribute) Equals(o Attribute) bool {
	oa, ok := o.(*AnchorAttribute)
	if !ok {
		return false
	}
	if oa.ID.Value != a.ID.Value {
		return false
	}
	return a.Label.Equals(oa.Label)
}

func (a *AnchorAttribute) Value() any {
	return a.ID
}

func (na *AnchorAttribute) SetValue(v any) error {
	if v, ok := v.(*String); ok {
		na.ID = v
		return nil
	}
	return fmt.Errorf("invalid type for AnchorAttribute: %T", v)
}

func (AnchorAttribute) AttributeType() AttributeType {
	return AttributeTypeID
}

func (AnchorAttribute) QuoteType() AttributeQuoteType {
	return AttributeQuoteTypeNone
}

func (na *AnchorAttribute) AsciiDocString() string {
	if na.ID == nil {
		return ""
	}
	var sb strings.Builder
	sb.WriteString(na.ID.Value)
	if len(na.Label) > 0 {
		sb.WriteString(",")
		for _, e := range na.Label {
			attributeAsciiDocStringElement(&sb, e)
		}
	}
	return sb.String()
}

type Anchor struct {
	position

	ID string
	Elements
}

func (Anchor) Type() ElementType {
	return ElementTypeInline
}

func (Anchor) AttributeType() AttributeType {
	return AttributeTypeID
}

func (Anchor) QuoteType() AttributeQuoteType {
	return AttributeQuoteTypeNone
}

func (a *Anchor) Equals(o Element) bool {
	oa, ok := o.(*Anchor)
	if !ok {
		return false
	}
	if oa.ID != a.ID {
		return false
	}
	return a.Elements.Equals(oa.Elements)
}

func NewAnchor(id string, label Elements) *Anchor {
	return &Anchor{ID: id, Elements: label}
}
