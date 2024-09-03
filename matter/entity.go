package matter

import (
	"github.com/project-chip/alchemy/asciidoc"
)

type entity struct {
	source asciidoc.Element
}

func (e entity) Source() asciidoc.Element {
	return e.source
}

func (e entity) Origin() (path string, line int) {
	switch s := e.source.(type) {
	case Source:
		return s.Origin()
	default:
		return "", -1
	}
}
