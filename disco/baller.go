package disco

import (
	"context"
	"log/slog"

	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
)

type Baller struct {
	discoOptions []Option
}

func NewBaller(discoOptions []Option, pipelineOptions pipeline.Options) Baller {
	return Baller{discoOptions: discoOptions}
}

func (r Baller) Name() string {
	return "Disco balling"
}

func (r Baller) Type() pipeline.ProcessorType {
	return pipeline.ProcessorTypeIndividual
}

func (r Baller) Process(cxt context.Context, input *pipeline.Data[*spec.Doc], index int32, total int32) (outputs []*pipeline.Data[*spec.Doc], extras []*pipeline.Data[*spec.Doc], err error) {
	b := NewBall(input.Content)
	for _, option := range r.discoOptions {
		option(b)
	}
	err = b.disco(cxt)
	if err != nil {
		if err == ErrEmptyDoc {
			err = nil
			return
		}
		slog.Warn("Error disco balling document", "path", input.Path, "error", err)
		err = nil
		return
	}
	outputs = append(outputs, pipeline.NewData[*spec.Doc](input.Path, input.Content))
	return
}
