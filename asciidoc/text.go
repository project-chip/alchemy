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

func NewStringElements(s ...string) (elements Elements) {
	for _, s := range s {
		elements = append(elements, NewString(s))
	}
	return
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

func (s *String) Clone() Element {
	return &String{Value: s.Value}
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

func (s SpecialCharacter) Clone() Element {
	return SpecialCharacter{Character: s.Character}
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

func (lc LineContinuation) Clone() Element {
	return LineContinuation{position: lc.position}
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

func (lb LineBreak) Clone() Element {
	return &LineBreak{position: lb.position}
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

func (nl NewLine) Clone() Element {
	return &NewLine{position: nl.position}
}

func StringValue(el Element) string {
	switch el := el.(type) {
	case *String:
		return el.Value
	}
	return ""
}
