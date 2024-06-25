package parse

import (
	"encoding/xml"

	"github.com/project-chip/alchemy/matter"
)

type source struct {
	path string
	line int
}

func newSource(path string, d *xml.Decoder) matter.Source {
	l, _ := d.InputPos()
	return &source{path: path, line: l}
}

func (s *source) Origin() (path string, line int) {
	return s.path, s.line
}
