package asciidoc

import (
	"fmt"
	"strings"
)

type AttributeEntry struct {
	position
	raw

	Name AttributeName
	Set
}

func NewAttributeEntry(name string) *AttributeEntry {
	return &AttributeEntry{Name: AttributeName(name)}
}

func (AttributeEntry) Type() ElementType {
	return ElementTypeBlock
}

func (uar *AttributeEntry) Equals(e Element) bool {
	ouar, ok := e.(*AttributeEntry)
	if !ok {
		return false
	}
	if uar.Name != ouar.Name {
		return false
	}
	return uar.Set.Equals(ouar.Set)
}

type AttributeReset struct {
	position
	raw

	Name AttributeName
}

func NewAttributeReset(name string) *AttributeReset {
	return &AttributeReset{Name: AttributeName(name)}
}

func (AttributeReset) Type() ElementType {
	return ElementTypeBlock
}

func (uar *AttributeReset) Equals(e Element) bool {
	ouar, ok := e.(*AttributeReset)
	if !ok {
		return false
	}
	return uar.Name == ouar.Name
}

type Attribute interface {
	HasPosition
	Value() any
	SetValue(any) error
	AsciiDocString() string
	Equals(o Attribute) bool
	AttributeType() AttributeType
	QuoteType() AttributeQuoteType
}

type attribute struct {
	position
	raw
}

func (ae attribute) Type() ElementType {
	return ElementTypeAttribute
}

type NamedAttribute struct {
	attribute

	Name AttributeName

	Val Set

	Quote AttributeQuoteType
}

func NewNamedAttribute(name string, value Set, quoteType AttributeQuoteType) Attribute {
	return &NamedAttribute{Name: AttributeName(name), Val: value, Quote: quoteType}
}

func (na *NamedAttribute) Value() any {
	return na.Val
}

func (na *NamedAttribute) SetValue(v any) error {
	if v, ok := v.(Set); ok {
		na.Val = v
		return nil
	}
	return fmt.Errorf("invalid type for NamedAttribute: %T", v)
}

func (na *NamedAttribute) AttributeType() AttributeType {
	return attributeNameToType(na.Name)
}

func (na *NamedAttribute) QuoteType() AttributeQuoteType {
	return na.Quote
}

func (na *NamedAttribute) AsciiDocString() string {
	return AttributeAsciiDocString(na.Val)
}

func (na *NamedAttribute) Equals(oa Attribute) bool {
	ona, ok := oa.(*NamedAttribute)
	if !ok {
		return false
	}
	if ona.Name != na.Name {
		return false
	}
	if !na.Val.Equals(ona.Val) {
		return false
	}
	if na.Quote != ona.Quote {
		return false
	}
	return true
}

type PositionalAttribute struct {
	attribute

	Offset      int
	ImpliedName AttributeName

	Val Set
}

func NewPositionalAttribute(value Set) Attribute {
	return &PositionalAttribute{Val: value}
}

func (pa *PositionalAttribute) Value() any {
	return pa.Val
}

func (na *PositionalAttribute) SetValue(v any) error {
	if v, ok := v.(Set); ok {
		na.Val = v
		return nil
	}
	return fmt.Errorf("invalid type for PositionalAttribute: %T", v)
}

func (pa *PositionalAttribute) Equals(oa Attribute) bool {
	opa, ok := oa.(*PositionalAttribute)
	if !ok {
		return false
	}
	if opa.Offset != pa.Offset {
		return false
	}
	return pa.Val.Equals(opa.Val)
}

func (pa *PositionalAttribute) AttributeType() AttributeType {
	return attributeNameToType(pa.ImpliedName)
}

func (PositionalAttribute) QuoteType() AttributeQuoteType {
	return AttributeQuoteTypeNone
}

func (na *PositionalAttribute) AsciiDocString() string {
	return AttributeAsciiDocString(na.Val)
}

type TitleAttribute struct {
	attribute

	Val Set
}

func NewTitleAttribute(value Set) Attribute {
	return &TitleAttribute{Val: value}
}

func (ta *TitleAttribute) Value() any {
	return ta.Val
}

func (na *TitleAttribute) SetValue(v any) error {
	if v, ok := v.(Set); ok {
		na.Val = v
		return nil
	}
	return fmt.Errorf("invalid type for TitleAttribute: %T", v)
}

func (ta *TitleAttribute) Equals(oa Attribute) bool {
	ota, ok := oa.(*TitleAttribute)
	if !ok {
		return false
	}
	return ta.Val.Equals(ota.Val)
}

func (TitleAttribute) AttributeType() AttributeType {
	return AttributeTypeTitle
}

func (TitleAttribute) QuoteType() AttributeQuoteType {
	return AttributeQuoteTypeNone
}

func (na *TitleAttribute) AsciiDocString() string {
	return AttributeAsciiDocString(na.Val)
}

type AttributeReference interface {
	HasPosition
	Element
	Name() string
}

type UserAttributeReference struct {
	position
	raw

	Value string
}

func (uar *UserAttributeReference) Equals(e Element) bool {
	ouar, ok := e.(*UserAttributeReference)
	if !ok {
		return false
	}
	return ouar.Value == uar.Value
}

func (uar *UserAttributeReference) Name() string {
	return uar.Value
}

func (UserAttributeReference) Type() ElementType {
	return ElementTypeInline
}

func NewAttributeReference(name string) AttributeReference {
	if isCharacterReplacement(name) {
		return &CharacterReplacementReference{Value: name}
	}
	return &UserAttributeReference{Value: name}
}

type Attributable interface {
	Attributes() []Attribute
	GetAttributeByName(name AttributeName) *NamedAttribute
	SetAttribute(name AttributeName, value Set)
	SetAttributes(as ...Attribute)
	DeleteAttribute(name AttributeName)
	AppendAttribute(as ...Attribute)
	ReadAttributes(el Element, as ...Attribute) error
}

type AttributableElement interface {
	Element
	Attributable
}

type BlockElement interface {
	Element
	Attributable
	HasElements
}

var characterReplacementAttributes = map[string]string{
	"blank":          "",
	"empty":          "",
	"sp":             " ",
	"nbsp":           "\u00a0",
	"zwsp":           "\u200b",
	"wj":             "\u2060",
	"apos":           "\u0027",
	"quot":           "\u0022",
	"lsquo":          "\u2018",
	"rsquo":          "\u2019",
	"ldquo":          "\u201c",
	"rdquo":          "\u201d",
	"deg":            "\u00b0",
	"plus":           "\u002b",
	"brvbar":         "\u00a6",
	"vbar":           "|",
	"amp":            "&",
	"lt":             "<",
	"gt":             ">",
	"startsb":        "[",
	"endsb":          "]",
	"caret":          "^",
	"asterisk":       "*",
	"tilde":          "~",
	"backslash":      "\\",
	"backtick":       "`",
	"two-colons":     "::",
	"two-semicolons": ";;",
	"cpp":            "C++",
	"pp":             "",
}

type CharacterReplacementReference UserAttributeReference

func (uar *CharacterReplacementReference) Name() string {
	return uar.Value
}

func (CharacterReplacementReference) Type() ElementType {
	return ElementTypeInline
}

func (uar *CharacterReplacementReference) Equals(e Element) bool {
	ouar, ok := e.(*CharacterReplacementReference)
	if !ok {
		return false
	}
	return ouar.Value == uar.Value
}

func (uar CharacterReplacementReference) ReplacementValue() string {
	return characterReplacementAttributes[uar.Value]
}

func isCharacterReplacement(s string) bool {
	_, ok := characterReplacementAttributes[s]
	return ok
}

func AttributeAsciiDocString(val Set) string {
	var sb strings.Builder
	for _, e := range val {
		attributeAsciiDocStringElement(&sb, e)
	}
	return sb.String()
}

func attributeAsciiDocStringElement(sb *strings.Builder, e Element) {
	switch e := e.(type) {
	case *String:
		sb.WriteString(e.Value)
	case AttributeReference:
		sb.WriteRune('{')
		sb.WriteString(e.Name())
		sb.WriteRune('}')
	default:
		sb.WriteString(fmt.Sprintf("ERROR: unexpected attribute element: %T", e))
	}
}
