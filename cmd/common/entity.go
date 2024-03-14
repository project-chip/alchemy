package common

import (
	"context"

	"github.com/hasty/alchemy/internal/pipeline"
	"github.com/hasty/alchemy/matter/types"
)

type EntityFilter[I types.EntityStore, O any] struct {
}

func (sp *EntityFilter[I, O]) Name() string {
	return ""
}

func (sp *EntityFilter[I, O]) Type() pipeline.ProcessorType {
	return pipeline.ProcessorTypeCollective
}

func (sp *EntityFilter[I, O]) Process(cxt context.Context, inputs []*pipeline.Data[I]) (outputs []*pipeline.Data[[]O], err error) {
	for _, i := range inputs {
		var entities []types.Entity
		entities, err = i.Content.Entities()
		if err != nil {
			return
		}
		var matches []O
		for _, e := range entities {
			switch e := e.(type) {
			case O:
				matches = append(matches, e)
			}
		}
		if len(matches) > 0 {
			outputs = append(outputs, pipeline.NewData[[]O](i.Path, matches))
		}
	}
	return
}
