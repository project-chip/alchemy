package cli

import (
	"context"

	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/paths"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/internal/text"
)

type StripComments struct {
	pipeline.ProcessingOptions `embed:""`
	files.OutputOptions        `embed:""`

	Paths []string `arg:"" help:"Paths of files to strip comments from" required:""`
}

func (f *StripComments) Run(cc *Context) (err error) {
	var inputs pipeline.Paths

	inputs, err = pipeline.Start(cc, paths.NewTargeter(f.Paths...))

	if err != nil {
		return err
	}

	reader := files.NewReader("Reading files...")

	var textFiles pipeline.FileSet
	textFiles, err = pipeline.Parallel(cc, f.ProcessingOptions, reader, inputs)

	if err != nil {
		return err
	}

	cr := &commentRemover{}

	var strippedFiles pipeline.FileSet
	strippedFiles, err = pipeline.Parallel(cc, f.ProcessingOptions, cr, textFiles)

	if err != nil {
		return err
	}
	writer := files.NewWriter[[]byte]("Writing files...", f.OutputOptions)
	err = writer.Write(cc, strippedFiles, f.ProcessingOptions)
	return
}

type commentRemover struct {
}

func (p commentRemover) Name() string {
	return "Removing comments"
}

func (p commentRemover) Process(cxt context.Context, input *pipeline.Data[[]byte], index int32, total int32) (outputs []*pipeline.Data[[]byte], extra []*pipeline.Data[[]byte], err error) {
	output := text.RemoveComments(string(input.Content))
	outputs = append(outputs, &pipeline.Data[[]byte]{Path: input.Path, Content: []byte(output)})
	return
}
