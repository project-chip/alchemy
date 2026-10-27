package spec

import (
	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
)

type DocumentInfoCache interface {
	DocType(document *asciidoc.Document) (matter.DocType, error)
	SetDocType(document *asciidoc.Document, docType matter.DocType)
	Parents(document *asciidoc.Document) []*asciidoc.Document
}

type DocSet pipeline.Map[string, *pipeline.Data[*asciidoc.Document]]

func NewDocSet() DocSet {
	return pipeline.NewMap[string, *pipeline.Data[*asciidoc.Document]]()
}
