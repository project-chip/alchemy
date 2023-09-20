package ascii

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
)

type Doc struct {
	Base *types.Document
	Root *Section
}

func NewDoc(d *types.Document) *Doc {
	doc := &Doc{
		Base: d,
		Root: &Section{},
	}
	for _, e := range d.BodyElements() {
		switch el := e.(type) {
		case *types.Section:
			doc.AddSection(doc.Root, el)
		default:
			doc.Root.Elements = append(doc.Root.Elements, e)
		}
	}
	return doc
}

func (d *Doc) parseSectionElements(parent *Section, s *types.Section) {
	for _, e := range s.Elements {
		switch el := e.(type) {
		case *types.Section:
			d.AddSection(parent, el)
		default:
			parent.Elements = append(parent.Elements, &Element{Base: e})
			//fmt.Printf("parsing unknown section element type: %T\n", e)

		}
	}
}

func (d *Doc) AddSection(parent *Section, s *types.Section) {
	ss := &Section{Base: s}

	for _, te := range s.GetTitle() {
		switch tel := te.(type) {
		case *types.StringElement:
			fmt.Printf("section title string: %s\n", tel.Content)
			ss.Name = tel.Content

		default:
			fmt.Printf("unknown section title element type: %T\n", te)
			ss.Elements = append(ss.Elements, te)
		}
	}
	switch s.Level {
	case 1:
		fmt.Printf("Adding top level section %s...\n", ss.Name)
		ss.SecType = DocSectionTypePreface
	case 2:

	}
	parent.Elements = append(parent.Elements, ss)
	d.parseSectionElements(ss, s)
}
