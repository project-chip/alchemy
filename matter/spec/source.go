package spec

import (
	"github.com/hasty/alchemy/asciidoc"
)

type source struct {
	doc     *Doc
	element asciidoc.Element
}

func newSource(d *Doc, e asciidoc.Element) *source {
	return &source{doc: d, element: e}
}

func (s *source) Origin() (path string, line int) {
	if hp, ok := s.element.(asciidoc.HasPosition); ok {
		line, _, _ = hp.Position()
	} else {
		line = -1
	}
	path = s.doc.Path
	return
}
