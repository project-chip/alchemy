package common

import (
	"context"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
)

type DocTypeFilter struct {
	docType      matter.DocType
	specificaion *spec.Specification
}

func (sp *DocTypeFilter) Name() string {
	return ""
}

func (sp *DocTypeFilter) Process(cxt context.Context, inputs []*pipeline.Data[*asciidoc.Document]) (outputs []*pipeline.Data[*asciidoc.Document], err error) {
	for _, i := range inputs {
		lib, ok := sp.specificaion.LibraryForDocument(i.Content)
		if ok {
			var docType matter.DocType
			docType, err = lib.DocType(i.Content)
			if err != nil {
				return
			}
			if docType == sp.docType {
				outputs = append(outputs, i)
			}

		}
	}
	return
}

func NewDocTypeFilter(specificaion *spec.Specification, docType matter.DocType) *DocTypeFilter {
	return &DocTypeFilter{specificaion: specificaion, docType: docType}
}
