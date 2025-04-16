package files

import (
	"context"
	"fmt"
	"os"

	"github.com/project-chip/alchemy/internal/pipeline"
)

type DryRun[T string | []byte] struct {
	writer
}

func (sp *DryRun[T]) Write(cxt context.Context, data pipeline.Map[string, *pipeline.Data[T]], pipelineOptions pipeline.ProcessingOptions) (err error) {
	_, err = pipeline.Collective(cxt, pipelineOptions, sp, data)
	return
}

func (sp *DryRun[T]) Process(cxt context.Context, inputs []*pipeline.Data[T]) (outputs []*pipeline.Data[struct{}], err error) {
	pipeline.SortData(inputs)
	for _, i := range inputs {
		fmt.Fprintf(os.Stderr, "Skipping %s...\n", i.Path)
	}
	return
}
