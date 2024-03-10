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

type Writer struct {
	name    string
	options Options
}

func (sp *Writer) Name() string {
	return sp.name
}

func (sp *Writer) Type() pipeline.ProcessorType {
	if sp.options.Serial || sp.options.Patch {
		return pipeline.ProcessorTypeSerial
	}
	return pipeline.ProcessorTypeParallel
}

func (sp *Writer) Process(cxt context.Context, input *pipeline.Data[string], index int32, total int32) (outputs []*pipeline.Data[struct{}], extras []*pipeline.Data[string], err error) {
	if sp.options.DryRun {
		return
	}
	err = os.WriteFile(input.Path, []byte(input.Content), os.ModeAppend|0644)
	return
}

func (sp *Writer) ProcessAll(cxt context.Context, inputs []*pipeline.Data[string]) (outputs []*pipeline.Data[struct{}], err error) {
	if sp.options.DryRun {
		return
	}
	if sp.options.Patch {
		fmt.Fprintf(os.Stderr, "patching %d files\n", len(inputs))
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
			edits := myers.ComputeEdits(span.URIFromPath(i.Path), existing, i.Content)
			fmt.Print(gotextdiff.ToUnified(i.Path, i.Path, existing, edits))
		}
		return
	}
	for _, i := range inputs {
		err = os.WriteFile(i.Path, []byte(i.Content), os.ModeAppend|0644)
		if err != nil {
			err = fmt.Errorf("error writing %s: %w", i.Path, err)
			return
		}
	}
	return
}

func NewWriter(name string, options Options) *Writer {
	return &Writer{name: name, options: options}
}
