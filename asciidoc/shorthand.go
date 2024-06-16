package asciidoc

import (
	"fmt"
	"regexp"
	"strings"
)

type ShorthandStyle struct {
	attribute

	Val Set
}

func NewShorthandStyle(value Set) *ShorthandStyle {
	return &ShorthandStyle{Val: value}
}

func (sa *ShorthandStyle) Equals(osa *ShorthandStyle) bool {
	if sa == nil {
		return osa == nil
	}
	if osa == nil {
		return false
	}
	return sa.Val.Equals(osa.Val)
}

type ShorthandID struct {
	attribute

	Val Set
}

func NewShorthandID(value Set) *ShorthandID {
	return &ShorthandID{Val: value}
}

func (sa *ShorthandID) Equals(osa *ShorthandID) bool {
	if sa == nil {
		return osa == nil
	}
	if osa == nil {
		return false
	}
	return sa.Val.Equals(osa.Val)
}

type ShorthandRole struct {
	attribute

	Val Set
}

func NewShorthandRole(value Set) *ShorthandRole {
	return &ShorthandRole{Val: value}
}

func (sa *ShorthandRole) Equals(osa *ShorthandRole) bool {
	if sa == nil {
		return osa == nil
	}
	if osa == nil {
		return false
	}
	return sa.Val.Equals(osa.Val)
}

type ShorthandOption struct {
	attribute

	Val Set
}

func NewShorthandOption(value Set) *ShorthandOption {
	return &ShorthandOption{Val: value}
}

func (sa *ShorthandOption) Equals(osa *ShorthandOption) bool {
	if sa == nil {
		return osa == nil
	}
	if osa == nil {
		return false
	}
	return sa.Val.Equals(osa.Val)
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
		s.WriteString(AttributeAsciiDocString(na.Style.Val))
	}
	if na.ID != nil {
		s.WriteRune('#')
		s.WriteString(AttributeAsciiDocString(na.ID.Val))
	}
	if len(na.Roles) > 0 {
		for _, r := range na.Roles {
			s.WriteRune('.')
			s.WriteString(AttributeAsciiDocString(r.Val))
		}
	}
	if len(na.Options) > 0 {
		for _, o := range na.Options {
			s.WriteRune('%')
			s.WriteString(AttributeAsciiDocString(o.Val))
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

var shorthandAttributePattern = regexp.MustCompile(`^(?P<Style>[^#\.%\n]+)?(?P<Elements>(?:[#\.%][^#\.%\n]+)*)$`)
var shorthandElementPattern = regexp.MustCompile(`[#\.%][^#\.%\n]+`)

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
	val := AttributeAsciiDocString(els)
	matches := shorthandAttributePattern.FindStringSubmatch(val)
	if matches == nil {
		return
	}
	ss := matches[1]
	if len(ss) > 0 {
		style = NewShorthandStyle(Set{NewString(ss)})
	}
	elements := matches[2]
	if len(elements) > 0 {
		ems := shorthandElementPattern.FindAllString(elements, -1)
		if ems == nil {
			return
		}
		for _, em := range ems {
			head := em[0]
			val := em[1:]
			switch head {
			case '#':
				id = NewShorthandID(Set{NewString(val)})
			case '.':
				roles = append(roles, NewShorthandRole(Set{NewString(val)}))
			case '%':
				options = append(options, NewShorthandOption(Set{NewString(val)}))
			}
		}
	}
	return
}
