package asciidoc

import (
	"strings"
)

type String struct {
	Value string
}

func NewString(s string) *String {
	return &String{Value: s}
}

func JoinStrings(ss []*String) *String {
	var sb strings.Builder
	for _, s := range ss {
		sb.WriteString(s.Value)
	}
	return &String{sb.String()}
}

func (String) Type() ElementType {
	return ElementTypeInlineLiteral
}

func (s *String) Equals(e Element) bool {
	os, ok := e.(*String)
	if !ok {
		return false
	}
	return os.Value == s.Value
}

type SpecialCharacter struct {
	Character string
}

func NewSpecialCharacter(character string) SpecialCharacter {
	return SpecialCharacter{Character: character}
}

func (SpecialCharacter) Type() ElementType {
	return ElementTypeInlineLiteral
}

func (s SpecialCharacter) Equals(e Element) bool {
	os, ok := e.(SpecialCharacter)
	if !ok {
		return false
	}
	return os.Character == s.Character
}

type LineContinuation struct {
	position
	raw
}

func (LineContinuation) Type() ElementType {
	return ElementTypeInlineLiteral
}

func (LineContinuation) Equals(e Element) bool {
	_, ok := e.(LineContinuation)
	return ok
}

type LineBreak struct {
	position
	raw
}

func (LineBreak) Type() ElementType {
	return ElementTypeInlineLiteral
}

func (*LineBreak) Equals(e Element) bool {
	_, ok := e.(*LineBreak)
	return ok
}

type NewLine struct {
	position
	raw
}

func (NewLine) Type() ElementType {
	return ElementTypeInlineLiteral
}

func (*NewLine) Equals(e Element) bool {
	_, ok := e.(*NewLine)
	return ok
}
