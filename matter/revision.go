package matter

import (
	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/matter/types"
)

type Revision struct {
	entity
	Number      *Number `json:"number,omitempty"`
	Description string  `json:"description,omitempty"`
}

func NewRevision(parent types.Entity, source asciidoc.Element) *Revision {
	return &Revision{
		entity: entity{parent: parent, source: source},
	}
}

type Revisions []*Revision

func (r Revisions) MostRecent() *Revision {
	var lastRevision *Revision
	var lastRevisionNumber uint64
	for _, rev := range r {
		if rev.Number.Valid() && rev.Number.Value() > lastRevisionNumber {
			lastRevision = rev
			lastRevisionNumber = rev.Number.Value()
		}
	}
	return lastRevision
}
