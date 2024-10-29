package matter

import (
	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/matter/types"
)

type entity struct {
	parent types.Entity
	source asciidoc.Element
}

func (e entity) Parent() types.Entity {
	return e.parent
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
