package common

import (
	"context"
	"log/slog"

	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/types"
)

type EntityFilter[I types.EntityStore, O any] struct {
}

func (sp *EntityFilter[I, O]) Name() string {
	return ""
}

func (sp *EntityFilter[I, O]) Process(cxt context.Context, inputs []*pipeline.Data[I]) (outputs []*pipeline.Data[[]O], err error) {
	type explodable interface {
		Explode() []O
	}
	for _, i := range inputs {
		var entities []types.Entity
		entities, err = i.Content.Entities()
		if err != nil {
			slog.WarnContext(cxt, "error converting to entities", slog.String("path", i.Path), slog.Any("error", err))
			err = nil
			continue
		}
		var matches []O
		for _, e := range entities {
			switch e := e.(type) {
			case explodable:
				matches = append(matches, e.Explode()...)
			case O:
				matches = append(matches, e)
			}
		}
		if len(matches) > 0 {
			outputs = append(outputs, pipeline.NewData(i.Path, matches))
		}
	}
	return
}
