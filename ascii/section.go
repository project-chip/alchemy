package ascii

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/matter"
)

type Section struct {
	Name string

	Base *types.Section

	SecType matter.Section

	Elements []interface{}
}

func NewSection(s *types.Section) *Section {
	ss := &Section{Base: s}
	for _, te := range s.Title {
		switch tel := te.(type) {
		case *types.StringElement:
			ss.Name = tel.Content
		case *types.InlineLink:

		default:
			//fmt.Printf("unknown section title element type: %T\n", te)
			//ss.Elements = append(ss.Elements, te)
		}
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
