package format

import (
	"context"

	"github.com/hasty/alchemy/asciidoc/render"
	"github.com/hasty/alchemy/internal/files"
	"github.com/hasty/alchemy/internal/pipeline"
	"github.com/hasty/alchemy/matter/spec"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "format",
	Short: "format Matter spec documents",
	RunE:  format,
}

func format(cmd *cobra.Command, args []string) (err error) {
	cxt := context.Background()

	var inputs pipeline.Map[string, *pipeline.Data[struct{}]]

	inputs, err = pipeline.Start[struct{}](cxt, files.PathsTargeter(args...))

	if err != nil {
		return err
	}

	pipelineOptions := pipeline.Flags(cmd)
	fileOptions := files.Flags(cmd)

	docReader := spec.NewReader("Reading docs")
	docs, err := pipeline.Process[struct{}, *spec.Doc](cxt, pipelineOptions, docReader, inputs)
	if err != nil {
		return err
	}

	ids := pipeline.NewConcurrentMapPresized[string, *pipeline.Data[render.InputDocument]](docs.Size())
	err = pipeline.Cast[*spec.Doc, render.InputDocument](docs, ids)
	if err != nil {
		return err
	}

	renderer := render.NewRenderer()
	var renders pipeline.Map[string, *pipeline.Data[string]]
	renders, err = pipeline.Process[render.InputDocument, string](cxt, pipelineOptions, renderer, ids)
	if err != nil {
		return err
	}

	writer := files.NewWriter[string]("Formatting docs", fileOptions)
	_, err = pipeline.Process[string, struct{}](cxt, pipelineOptions, writer, renders)
	return
}
