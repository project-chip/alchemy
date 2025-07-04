package files

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/hexops/gotextdiff"
	"github.com/hexops/gotextdiff/myers"
	"github.com/hexops/gotextdiff/span"
	"github.com/project-chip/alchemy/internal/pipeline"
)

type Patcher[T string | []byte] struct {
	writer

	Root string
	out  io.Writer
}

func NewPatcher[T string | []byte](name string, out io.Writer) *Patcher[T] {
	return &Patcher[T]{writer: writer{name: name}, out: out}
}

func (sp *Patcher[T]) Write(cxt context.Context, data pipeline.Map[string, *pipeline.Data[T]], pipelineOptions pipeline.ProcessingOptions) (err error) {
	_, err = pipeline.Collective(cxt, pipelineOptions, sp, data)
	return
}

func (sp *Patcher[T]) Process(cxt context.Context, inputs []*pipeline.Data[T]) (outputs []*pipeline.Data[struct{}], err error) {
	for _, i := range inputs {
		path := i.Path
		var exists bool
		exists, err = Exists(path)
		if err != nil {
			return
		}
		var existing string
		if exists {
			var eb []byte
			eb, err = os.ReadFile(path)
			if err != nil {
				return
			}
			existing = string(eb)
		}

		if sp.Root != "" && filepath.IsAbs(path) {
			path, err = filepath.Rel(sp.Root, path)
			if err != nil {
				err = fmt.Errorf("error getting relative patch path: %w", err)
				return
			}
		}

		edits := myers.ComputeEdits(span.URIFromPath(path), existing, string(i.Content))
		if len(edits) > 0 {
			fmt.Fprintln(sp.out, gotextdiff.ToUnified(path, path, existing, edits))
		}
	}
	return
}
