package asciidoc

import (
	"fmt"
	"strings"
)

type ShorthandStyle struct {
	attribute

	Set
}

func NewShorthandStyle(value ...Element) *ShorthandStyle {
	return &ShorthandStyle{Set: value}
}

func (sa *ShorthandStyle) Equals(osa Element) bool {
	if osa, ok := osa.(*ShorthandStyle); ok {
		if osa == nil {
			return sa == nil
		}
		return sa.Set.Equals(osa.Set)
	}
	return false
}

type ShorthandID struct {
	attribute

	Set
}

func NewShorthandID(value ...Element) *ShorthandID {
	return &ShorthandID{Set: value}
}

func (sa *ShorthandID) Equals(osa Element) bool {
	if osa, ok := osa.(*ShorthandID); ok {
		if osa == nil {
			return sa == nil
		}
		return sa.Set.Equals(osa.Set)
	}
	return false
}

type ShorthandRole struct {
	attribute

	Set
}

func NewShorthandRole(value ...Element) *ShorthandRole {
	return &ShorthandRole{Set: value}
}

func (sa *ShorthandRole) Equals(osa Element) bool {
	if osa, ok := osa.(*ShorthandRole); ok {
		if osa == nil {
			return sa == nil
		}
		return sa.Set.Equals(osa.Set)
	}
	return false
}

type ShorthandOption struct {
	attribute

	Set
}

func NewShorthandOption(value ...Element) *ShorthandOption {
	return &ShorthandOption{Set: value}
}

func (sa *ShorthandOption) Equals(osa Element) bool {
	if osa, ok := osa.(*ShorthandOption); ok {
		if osa == nil {
			return sa == nil
		}
		eq := sa.Set.Equals(osa.Set)
		return eq
	}
	return false
}

type ShorthandAttribute struct {
	attribute

	Style   *ShorthandStyle
	ID      *ShorthandID
	Roles   []*ShorthandRole
	Options []*ShorthandOption
}

func (ae ShorthandAttribute) Type() ElementType {
	return ElementTypeAttribute
}

func (ae *ShorthandAttribute) Value() any {
	var sb strings.Builder

	return sb.String()
}

func (ShorthandAttribute) AttributeType() AttributeType {
	return AttributeTypeTitle
}

func (ShorthandAttribute) QuoteType() AttributeQuoteType {
	return AttributeQuoteTypeNone
}

func (ta *ShorthandAttribute) Equals(oa Attribute) bool {
	ota, ok := oa.(*ShorthandAttribute)
	if !ok {
		return false
	}
	if !ta.Style.Equals(ota.Style) {
		return false
	}
	if !ta.ID.Equals(ota.ID) {
		return false
	}
	if len(ta.Roles) != len(ota.Roles) {
		return false
	}
	for i, r := range ta.Roles {
		or := ota.Roles[i]
		if !r.Equals(or) {
			return false
		}
	}
	if len(ta.Options) != len(ota.Options) {
		return false
	}
	for i, r := range ta.Options {
		or := ota.Options[i]
		if !r.Equals(or) {
			return false
		}
	}
	return true
}

func (na *ShorthandAttribute) SetValue(v any) error {
	if _, ok := v.(Set); ok {
		return nil
	}
	return fmt.Errorf("invalid type for ShorthandAttribute: %T", v)
}

func (na *ShorthandAttribute) AsciiDocString() string {
	var s strings.Builder
	if na.Style != nil {
		s.WriteString(AttributeAsciiDocString(na.Style.Set))
	}
	if na.ID != nil {
		s.WriteRune('#')
		s.WriteString(AttributeAsciiDocString(na.ID.Set))
	}
	if len(na.Roles) > 0 {
		for _, r := range na.Roles {
			s.WriteRune('.')
			s.WriteString(AttributeAsciiDocString(r.Set))
		}
	}
	if len(na.Options) > 0 {
		for _, o := range na.Options {
			s.WriteRune('%')
			s.WriteString(AttributeAsciiDocString(o.Set))
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

func parseShorthandAttributeValues(els Set) (style *ShorthandStyle, id *ShorthandID, roles []*ShorthandRole, options []*ShorthandOption) {
	if len(els) == 0 {
		return
	}
	var currentShorthand HasElements
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
					if len(cs.Set) > 0 {
						style = cs
					}
				case *ShorthandID:
					if len(cs.Set) > 0 {
						id = cs
					}
				case *ShorthandRole:
					if len(cs.Set) > 0 {
						roles = append(roles, cs)
					}
				case *ShorthandOption:
					if len(cs.Set) > 0 {
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
		if len(cs.Set) > 0 {
			style = cs
		}
	case *ShorthandID:
		if len(cs.Set) > 0 {
			id = cs
		}
	case *ShorthandRole:
		if len(cs.Set) > 0 {
			roles = append(roles, cs)
		}
	case *ShorthandOption:
		if len(cs.Set) > 0 {
			options = append(options, cs)
		}
	}
	return
}
