package disco

import (
	"context"

	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/matter/spec"
)

type discoContext struct {
	context.Context

	doc    *spec.Doc
	errata *errata.Disco
	parsed *docParse

	potentialDataTypes map[string][]*DataTypeEntry
}

func newContext(parent context.Context, doc *spec.Doc) *discoContext {
	return &discoContext{
		Context:            parent,
		doc:                doc,
		errata:             errata.GetDisco(doc.Path.Relative),
		potentialDataTypes: make(map[string][]*DataTypeEntry),
	}
}
