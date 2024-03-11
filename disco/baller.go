package disco

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/ascii/render"
	"github.com/hasty/alchemy/internal/pipeline"
)

type Baller struct {
	discoOptions    []Option
	pipelineOptions pipeline.Options
}

func NewBaller(discoOptions []Option, pipelineOptions pipeline.Options) Baller {
	return Baller{discoOptions: discoOptions, pipelineOptions: pipelineOptions}
}

func (r Baller) Name() string {
	return "Disco balling"
}

func (r Baller) Type() pipeline.ProcessorType {
	return r.pipelineOptions.DefaultProcessorType()
}

func (r Baller) Process(cxt context.Context, input *pipeline.Data[*ascii.Doc], index int32, total int32) (outputs []*pipeline.Data[render.InputDocument], extras []*pipeline.Data[*ascii.Doc], err error) {
	var output *pipeline.Data[render.InputDocument]
	output, err = r.process(cxt, input)
	if err != nil {
		return
	}
	if output != nil {
		outputs = append(outputs, output)
	}
	return
}

func (r Baller) ProcessAll(cxt context.Context, inputs []*pipeline.Data[*ascii.Doc]) (outputs []*pipeline.Data[render.InputDocument], err error) {
	for _, input := range inputs {
		fmt.Fprintf(os.Stderr, "Disco-balling %s...\n", input.Path)
		var output *pipeline.Data[render.InputDocument]
		output, err = r.process(cxt, input)
		if err != nil {
			return
		}
		if output != nil {
			outputs = append(outputs, output)
		}
	}
	return
}

func (r Baller) process(cxt context.Context, input *pipeline.Data[*ascii.Doc]) (output *pipeline.Data[render.InputDocument], err error) {
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
	output = &pipeline.Data[render.InputDocument]{Path: input.Path, Content: input.Content}
	return
}
