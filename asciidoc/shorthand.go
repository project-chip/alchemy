package asciidoc

import (
	"fmt"
	"strings"
)

type ShorthandStyle struct {
	attribute

	Elements
}

func NewShorthandStyle(value ...Element) *ShorthandStyle {
	return &ShorthandStyle{Elements: value}
}

func (ss *ShorthandStyle) Equals(osa Element) bool {
	if osa, ok := osa.(*ShorthandStyle); ok {
		if osa == nil {
			return ss == nil
		}
		return ss.Elements.Equals(osa.Elements)
	}
	return false
}

func (ss *ShorthandStyle) Clone() Element {
	return &ShorthandStyle{attribute: ss.attribute, Elements: ss.Elements.Clone()}
}

type ShorthandID struct {
	attribute

	Elements
}

func NewShorthandID(value ...Element) *ShorthandID {
	return &ShorthandID{Elements: value}
}

func (sid *ShorthandID) Equals(osa Element) bool {
	if osa, ok := osa.(*ShorthandID); ok {
		if osa == nil {
			return sid == nil
		}
		return sid.Elements.Equals(osa.Elements)
	}
	return false
}

func (sid *ShorthandID) Clone() Element {
	return &ShorthandID{attribute: sid.attribute, Elements: sid.Elements.Clone()}
}

type ShorthandRole struct {
	attribute

	Elements
}

func NewShorthandRole(value ...Element) *ShorthandRole {
	return &ShorthandRole{Elements: value}
}

func (sr *ShorthandRole) Equals(osa Element) bool {
	if osa, ok := osa.(*ShorthandRole); ok {
		if osa == nil {
			return sr == nil
		}
		return sr.Elements.Equals(osa.Elements)
	}
	return false
}

func (sr *ShorthandRole) Clone() Element {
	return &ShorthandRole{attribute: sr.attribute, Elements: sr.Elements.Clone()}
}

type ShorthandOption struct {
	attribute

	Elements
}

func NewShorthandOption(value ...Element) *ShorthandOption {
	return &ShorthandOption{Elements: value}
}

func (so *ShorthandOption) Equals(osa Element) bool {
	if osa, ok := osa.(*ShorthandOption); ok {
		if osa == nil {
			return so == nil
		}
		eq := so.Elements.Equals(osa.Elements)
		return eq
	}
	return false
}

func (so *ShorthandOption) Clone() Element {
	return &ShorthandOption{attribute: so.attribute, Elements: so.Elements.Clone()}
}

type ShorthandAttribute struct {
	attribute

	Style   *ShorthandStyle
	ID      *ShorthandID
	Roles   []*ShorthandRole
	Options []*ShorthandOption
}

func (sa ShorthandAttribute) Type() ElementType {
	return ElementTypeAttribute
}

func (sa *ShorthandAttribute) Value() any {
	var sb strings.Builder

	return sb.String()
}

func (ShorthandAttribute) AttributeType() AttributeType {
	return AttributeTypeTitle
}

func (ShorthandAttribute) QuoteType() AttributeQuoteType {
	return AttributeQuoteTypeNone
}

func (sa *ShorthandAttribute) Equals(oa Attribute) bool {
	ota, ok := oa.(*ShorthandAttribute)
	if !ok {
		return false
	}
	if !sa.Style.Equals(ota.Style) {
		return false
	}
	if !sa.ID.Equals(ota.ID) {
		return false
	}
	if len(sa.Roles) != len(ota.Roles) {
		return false
	}
	for i, r := range sa.Roles {
		or := ota.Roles[i]
		if !r.Equals(or) {
			return false
		}
	}
	if len(sa.Options) != len(ota.Options) {
		return false
	}
	for i, r := range sa.Options {
		or := ota.Options[i]
		if !r.Equals(or) {
			return false
		}
	}
	return true
}

func (sa *ShorthandAttribute) Clone() Attribute {
	csa := &ShorthandAttribute{
		attribute: sa.attribute,
		Style:     sa.Style.Clone().(*ShorthandStyle),
		ID:        sa.ID.Clone().(*ShorthandID),
		Roles:     make([]*ShorthandRole, len(sa.Roles)),
		Options:   make([]*ShorthandOption, len(sa.Options)),
	}
	for i, r := range sa.Roles {
		csa.Roles[i] = r.Clone().(*ShorthandRole)
	}
	for i, o := range sa.Options {
		csa.Options[i] = o.Clone().(*ShorthandOption)
	}
	return csa
}

func (sa *ShorthandAttribute) SetValue(v any) error {
	if _, ok := v.(Elements); ok {
		return nil
	}
	return fmt.Errorf("invalid type for ShorthandAttribute: %T", v)
}

func (sa *ShorthandAttribute) AsciiDocString() string {
	var s strings.Builder
	if sa.Style != nil {
		s.WriteString(AttributeAsciiDocString(sa.Style.Elements))
	}
	if sa.ID != nil {
		s.WriteRune('#')
		s.WriteString(AttributeAsciiDocString(sa.ID.Elements))
	}
	if len(sa.Roles) > 0 {
		for _, r := range sa.Roles {
			s.WriteRune('.')
			s.WriteString(AttributeAsciiDocString(r.Elements))
		}
	}
	if len(sa.Options) > 0 {
		for _, o := range sa.Options {
			s.WriteRune('%')
			s.WriteString(AttributeAsciiDocString(o.Elements))
		}
	}
	return s.String()
}

func NewShorthandAttribute(style any, values []any) (*ShorthandAttribute, error) {
	sha := &ShorthandAttribute{}
	if s, ok := style.(*ShorthandStyle); ok {
		sha.Style = s
	}
	for _, v := range values {
		switch v := v.(type) {
		case *ShorthandID:
			sha.ID = v
		case *ShorthandRole:
			sha.Roles = append(sha.Roles, v)
		case *ShorthandOption:
			sha.Options = append(sha.Options, v)
		default:
			return nil, fmt.Errorf("unexpected type in shorthand attribute: %T", v)
		}
	}
	return sha, nil
}

func parseShorthandAttribute(pa *PositionalAttribute) *ShorthandAttribute {
	style, id, roles, options := parseShorthandAttributeValues(pa.Val)
	if style == nil && id == nil && len(roles) == 0 && len(options) == 0 {
		return nil
	}
	return &ShorthandAttribute{
		Style:   style,
		ID:      id,
		Roles:   roles,
		Options: options,
	}
}

func parseShorthandAttributeValues(els Elements) (style *ShorthandStyle, id *ShorthandID, roles []*ShorthandRole, options []*ShorthandOption) {
	if len(els) == 0 {
		return
	}
	var currentShorthand ElementList
	for _, el := range els {
		switch el := el.(type) {
		case *String:
			val := el.Value
			if len(val) == 0 {
				continue
			}
			for {
				hashIndex := strings.IndexAny(val, ".#%")
				if hashIndex < 0 {
					if currentShorthand == nil { // We haven't hit any shorthands yet, so this must be a style
						style = NewShorthandStyle()
						currentShorthand = style
					}
					currentShorthand.Append(NewString(val))
					break
				}
				if hashIndex >= 0 {
					if hashIndex > 0 { // We have some text at the beginning
						if currentShorthand == nil { // We haven't hit any shorthands yet, so this must be a style
							style = NewShorthandStyle()
							currentShorthand = style
						}
						currentShorthand.Append(NewString(val[:hashIndex]))
					}
				}
				switch cs := currentShorthand.(type) {
				case nil:
				case *ShorthandStyle:
					if len(cs.Elements) > 0 {
						style = cs
					}
				case *ShorthandID:
					if len(cs.Elements) > 0 {
						id = cs
					}
				case *ShorthandRole:
					if len(cs.Elements) > 0 {
						roles = append(roles, cs)
					}
				case *ShorthandOption:
					if len(cs.Elements) > 0 {
						options = append(options, cs)
					}
				}
				switch val[hashIndex] {
				case '.':
					currentShorthand = NewShorthandRole()
				case '#':
					currentShorthand = NewShorthandID()
				case '%':
					currentShorthand = NewShorthandOption()
				}
				val = val[hashIndex+1:]
			}
		default:
			if currentShorthand == nil { // Must be a style
				currentShorthand = NewShorthandStyle()
			}
			currentShorthand.Append(el)
		}
	}
	switch cs := currentShorthand.(type) {
	case nil:
	case *ShorthandStyle:
		if len(cs.Elements) > 0 {
			style = cs
		}
	case *ShorthandID:
		if len(cs.Elements) > 0 {
			id = cs
		}
	case *ShorthandRole:
		if len(cs.Elements) > 0 {
			roles = append(roles, cs)
		}
	case *ShorthandOption:
		if len(cs.Elements) > 0 {
			options = append(options, cs)
		}
	}
	return
}
