package conformance

import "github.com/hasty/alchemy/matter/types"

type ValueStore interface {
	Reference(id string) types.Entity
}

type Context struct {
	Values            map[string]any
	Store             ValueStore
	VisitedReferences map[string]struct{}
}
