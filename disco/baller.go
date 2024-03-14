package disco

import (
	"context"
	"log/slog"

	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/ascii/render"
	"github.com/hasty/alchemy/internal/pipeline"
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

func (r Baller) Process(cxt context.Context, input *pipeline.Data[*ascii.Doc], index int32, total int32) (outputs []*pipeline.Data[render.InputDocument], extras []*pipeline.Data[*ascii.Doc], err error) {
	b := NewBall(input.Content)
	for _, option := range r.discoOptions {
		option(b)
	}
	err = b.Run(cxt)
	if err != nil {
		if err == EmptyDocError {
			err = nil
			return
		}
		slog.Warn("Error disco balling document", "path", input.Path, "error", err)
		err = nil
		return
	}
	outputs = append(outputs, pipeline.NewData[render.InputDocument](input.Path, input.Content))
	return
}
