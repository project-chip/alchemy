package asciidoc

import (
	"iter"
	"strings"
)

type Section struct {
	position
	raw
	child

	AttributeList
	Elements

	Title Elements
	Level int
}

func NewSection(title Elements, level int) *Section {
	return &Section{Title: title, Level: level}
}

func (Section) Type() ElementType {
	return ElementTypeBlock
}

func (s *Section) ParentSection() *Section {
	if ps, ok := s.parent.(*Section); ok {
		return ps
	}
	return nil
}

func (s *Section) AddChildSection(c *Section) {
	c.parent = s
}

func (a *Section) Equals(o Element) bool {
	oa, ok := o.(*Section)
	if !ok {
		return false
	}
	if a.Level != oa.Level {
		return false
	}
	if !a.Title.Equals(oa.Title) {
		return false
	}
	if !a.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return a.Elements.Equals(oa.Elements)
}

func (s *Section) Clone() Element {
	return &Section{
		position:      s.position,
		raw:           s.raw,
		AttributeList: s.AttributeList.Clone(),
		Elements:      s.Elements.Clone(),
		Title:         s.Title.Clone(),
		Level:         s.Level,
	}
}

func (s Section) Name() string {
	var sb strings.Builder
	for _, t := range s.Title {
		switch t := t.(type) {
		case *String:
			sb.WriteString(t.Value)
		}
	}
	return sb.String()
}

func (s *Section) Traverse(parent ParentElement) iter.Seq2[ParentElement, Parent] {
	return func(yield func(ParentElement, Parent) bool) {
		if !s.AttributeList.traverse(s, yield) {
			return
		}
		if !yield(s, &s.Title) {
			return
		}
	}
}
