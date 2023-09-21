package ascii

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

type Doc struct {
	Path string

	Base     *types.Document
	Elements []interface{}
}

func NewDoc(d *types.Document) *Doc {
	doc := &Doc{
		Base: d,
	}
	for _, e := range d.Elements {
		switch el := e.(type) {
		case *types.Section:
			doc.Elements = append(doc.Elements, NewSection(el))
		default:
			doc.Elements = append(doc.Elements, e)
		}
	}
	return doc
}
