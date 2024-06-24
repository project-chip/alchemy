package asciidoc

import (
	"strings"
)

type Section struct {
	position
	raw

	AttributeList
	Set

	Title Set
	Level int

	parent *Section
}

func NewSection(title Set, level int) *Section {
	return &Section{Title: title, Level: level}
}

func (Section) Type() ElementType {
	return ElementTypeBlock
}

func (s *Section) Parent() *Section {
	return s.parent
}

func (s *Section) AddChild(c *Section) {
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
	return a.Set.Equals(oa.Set)
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
