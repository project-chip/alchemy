package ascii

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/matter"
)

type Section struct {
	Name string

	Parent interface{}
	Base   *types.Section

	SecType matter.Section

	Elements []interface{}
}

func NewSection(parent interface{}, s *types.Section) (*Section, error) {
	ss := &Section{Parent: parent, Base: s}

	switch name := types.Reduce(s.Title).(type) {
	case string:
		ss.Name = name
	case []interface{}:
		var complexName strings.Builder
		for _, e := range name {
			switch v := e.(type) {
			case *types.StringElement:
				complexName.WriteString(v.Content)
			case string:
				complexName.WriteString(v)
			case *types.Symbol:
				complexName.WriteString(v.Name)
			case *types.SpecialCharacter:
				complexName.WriteString(v.Name)
			case *types.InlineLink:
			default:
				return nil, fmt.Errorf("unknown section title component type: %T", e)
			}
		}
		ss.Name = complexName.String()
	default:
		return nil, fmt.Errorf("unknown section title type: %T", name)
	}
	for _, e := range s.Elements {
		switch el := e.(type) {
		case *types.Section:
			s, err := NewSection(ss, el)
			if err != nil {
				return nil, err
			}
			ss.Elements = append(ss.Elements, s)
		default:
			ss.Elements = append(ss.Elements, NewElement(ss, e))
		}
	}
	return ss, nil
}

func (s *Section) AppendSection(ns *Section) error {
	s.Elements = append(s.Elements, ns)
	return nil
}

func (s *Section) GetElements() []interface{} {
	return s.Elements
}

func (s *Section) SetElements(elements []interface{}) error {
	s.Elements = elements
	return nil
}
