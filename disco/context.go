package disco

import (
	"context"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/errata"
)

type discoContext struct {
	context.Context

	doc    *asciidoc.Document
	errata *errata.Disco
	parsed *docParse

	potentialDataTypes map[string][]*DataTypeEntry
}

func newContext(parent context.Context, doc *asciidoc.Document) *discoContext {
	return &discoContext{
		Context:            parent,
		doc:                doc,
		errata:             errata.GetDisco(doc.Path.Relative),
		potentialDataTypes: make(map[string][]*DataTypeEntry),
	}
}
