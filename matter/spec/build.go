package spec

import (
	"context"

	"github.com/hasty/alchemy/internal/pipeline"
	"github.com/hasty/alchemy/matter"
)

type Builder struct {
	Spec *matter.Spec

	IgnoreHierarchy bool
}

func (sp Builder) Name() string {
	return "Building spec"
}

func (sp Builder) Type() pipeline.ProcessorType {
	return pipeline.ProcessorTypeCollective
}

func (sp *Builder) Process(cxt context.Context, inputs []*pipeline.Data[*Doc]) (outputs []*pipeline.Data[*Doc], err error) {
	docs := make([]*Doc, 0, len(inputs))
	for _, i := range inputs {
		docs = append(docs, i.Content)
	}
	sp.Spec, err = BuildSpec(docs, sp.IgnoreHierarchy)
	outputs = inputs
	return
}
