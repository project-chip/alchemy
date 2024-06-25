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

func (sp *DryRun[T]) Type() pipeline.ProcessorType {
	return pipeline.ProcessorTypeCollective
}

func (sp *DryRun[T]) Process(cxt context.Context, inputs []*pipeline.Data[T]) (outputs []*pipeline.Data[struct{}], err error) {
	pipeline.SortData[T](inputs)
	for _, i := range inputs {
		fmt.Fprintf(os.Stderr, "Skipping %s...\n", i.Path)
	}
	return
}
