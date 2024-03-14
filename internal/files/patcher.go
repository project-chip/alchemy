package files

import (
	"context"
	"fmt"
	"os"

	"github.com/hasty/alchemy/internal/pipeline"
	"github.com/hexops/gotextdiff"
	"github.com/hexops/gotextdiff/myers"
	"github.com/hexops/gotextdiff/span"
)

type Patcher[T string | []byte] struct {
	writer
}

func (sp *Patcher[T]) Type() pipeline.ProcessorType {
	return pipeline.ProcessorTypeCollective
}

func (sp *Patcher[T]) Process(cxt context.Context, inputs []*pipeline.Data[T]) (outputs []*pipeline.Data[struct{}], err error) {
	for _, i := range inputs {
		var exists bool
		exists, err = Exists(i.Path)
		if err != nil {
			return
		}
		var existing string
		if exists {
			var eb []byte
			eb, err = os.ReadFile(i.Path)
			if err != nil {
				return
			}
			existing = string(eb)
		}
		edits := myers.ComputeEdits(span.URIFromPath(i.Path), existing, string(i.Content))
		fmt.Print(gotextdiff.ToUnified(i.Path, i.Path, existing, edits))
	}
	return
}
