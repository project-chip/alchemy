package ascii

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
)

type Section struct {
	Name string

	Base *types.Section

	SecType DocSectionType

	Elements []interface{}
}

type DocSectionType uint8

const (
	DocSectionTypePreface DocSectionType = 0
	DocSectionTypeAttributes
	DocSectionTypeFeatures
	DocSectionTypeDataTypes
	DocSectionTypeCommands
)

func NewSection(s *types.Section) *Section {
	ss := &Section{Base: s}
	for _, te := range s.Title {
		switch tel := te.(type) {
		case *types.StringElement:
			fmt.Printf("section title string: %s\n", tel.Content)
			ss.Name = tel.Content
		case *types.InlineLink:

		default:
			fmt.Printf("unknown section title element type: %T\n", te)
			//ss.Elements = append(ss.Elements, te)
		}
	}
	switch s.Level {
	case 1:
		fmt.Printf("Adding top level section %s...\n", ss.Name)
		ss.SecType = DocSectionTypePreface
	case 2:

	}
	for _, e := range s.Elements {
		switch el := e.(type) {
		case *types.Section:
			ss.Elements = append(ss.Elements, NewSection(el))
		default:
			ss.Elements = append(ss.Elements, &Element{Base: e})
		}
	}
	return ss
}
