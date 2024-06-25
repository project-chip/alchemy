package conformance

import "github.com/project-chip/alchemy/matter/types"

type IdentifierStore interface {
	Identifier(id string) (types.Entity, bool)
}

type ReferenceStore interface {
	Reference(ref string) (types.Entity, bool)
}

type Context struct {
	Values            map[string]any
	Identifiers       IdentifierStore
	References        ReferenceStore
	VisitedReferences map[string]struct{}
}
