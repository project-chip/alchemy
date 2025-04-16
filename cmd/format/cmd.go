package format

import (
	"github.com/project-chip/alchemy/asciidoc/render"
	"github.com/project-chip/alchemy/cmd/cli"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/paths"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
)

type Command struct {
	pipeline.ProcessingOptions `embed:""`
	files.OutputOptions        `embed:""`
	render.RenderOptions       `embed:""`

	Paths []string `arg:"" help:"Paths of AsciiDoc files to format" required:""`
}

func (f *Command) Run(alchemy *cli.Alchemy) (err error) {
	var inputs pipeline.Paths

	inputs, err = pipeline.Start(alchemy, paths.NewTargeter(f.Paths...))

	if err != nil {
		return err
	}

	docReader, err := spec.NewReader("Reading docs", "")
	if err != nil {
		return err
	}
	docs, err := pipeline.Parallel(alchemy, f.ProcessingOptions, docReader, inputs)
	if err != nil {
		return err
	}

	ids := pipeline.NewConcurrentMapPresized[string, *pipeline.Data[render.InputDocument]](docs.Size())
	err = pipeline.Cast(docs, ids)
	if err != nil {
		return err
	}

	renderer := render.NewRenderer(f.RenderOptions.ToOptions()...)
	var renders pipeline.StringSet
	renders, err = pipeline.Parallel(alchemy, f.ProcessingOptions, renderer, ids)
	if err != nil {
		return err
	}

	writer := files.NewWriter[string]("Formatting docs", f.OutputOptions)
	err = writer.Write(alchemy, renders, f.ProcessingOptions)
	return
}
