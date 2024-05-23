package common

import (
	"context"

	"github.com/hasty/alchemy/internal/pipeline"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/spec"
)

type DocTypeFilter struct {
	docType matter.DocType
}

func (sp *DocTypeFilter) Name() string {
	return ""
}

func (sp *DocTypeFilter) Type() pipeline.ProcessorType {
	return pipeline.ProcessorTypeCollective
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
