package common

import (
	"context"

	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
)

type DocTypeFilter struct {
	docType matter.DocType
}

func (sp *DocTypeFilter) Name() string {
	return ""
}

func (sp *DocTypeFilter) Process(cxt context.Context, inputs []*pipeline.Data[*spec.Doc]) (outputs []*pipeline.Data[*spec.Doc], err error) {
	for _, i := range inputs {
		var docType matter.DocType
		docType, err = i.Content.DocType()
		if err != nil {
			return
		}
		if docType == sp.docType {
			outputs = append(outputs, i)
		}
	}
	return
}

func NewDocTypeFilter(docType matter.DocType) *DocTypeFilter {
	return &DocTypeFilter{docType: docType}
}
