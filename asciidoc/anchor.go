package asciidoc

import (
	"fmt"
	"iter"
	"strings"
)

type AnchorAttribute struct {
	attribute

	ID    Elements `json:"id"`
	Label Elements
}

func NewAnchorAttribute(id Elements, label Elements) *AnchorAttribute {
	return &AnchorAttribute{ID: id, Label: label}
}

func (aa *AnchorAttribute) Equals(o Attribute) bool {
	oa, ok := o.(*AnchorAttribute)
	if !ok {
		return false
	}
	if !oa.ID.Equals(aa.ID) {
		return false
	}
	return aa.Label.Equals(oa.Label)
}

func (aa *AnchorAttribute) Value() any {
	return aa.ID
}

func (aa *AnchorAttribute) SetValue(v any) error {
	if v, ok := v.(Elements); ok {
		aa.ID = v
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

func (aa *AnchorAttribute) AsciiDocString() string {
	if aa.ID == nil {
		return ""
	}
	var sb strings.Builder
	for _, e := range aa.ID {
		attributeAsciiDocStringElement(&sb, e)
	}
	if len(aa.Label) > 0 {
		sb.WriteString(",")
		for _, e := range aa.Label {
			attributeAsciiDocStringElement(&sb, e)
		}
	}
	return sb.String()
}

func (aa *AnchorAttribute) Traverse(parent Parent) iter.Seq2[Parent, Parent] {
	return func(yield func(Parent, Parent) bool) {
		if !yield(parent, &aa.Label) {
			return
		}
	}
}

type Anchor struct {
	position

	ID Elements
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
	if !oa.ID.Equals(a.ID) {
		return false
	}
	return a.Elements.Equals(oa.Elements)
}

func NewAnchor(id Elements, label Elements) *Anchor {
	return &Anchor{ID: id, Elements: label}
}
