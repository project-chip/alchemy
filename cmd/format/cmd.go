package format

import (
	"context"

	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/ascii/render"
	"github.com/hasty/alchemy/internal/files"
	"github.com/hasty/alchemy/internal/pipeline"
	"github.com/puzpuzpuz/xsync/v3"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "format",
	Short: "format Matter spec documents",
	RunE:  format,
}

func format(cmd *cobra.Command, args []string) (err error) {
	cxt := context.Background()

	var inputs *xsync.MapOf[string, *pipeline.Data[struct{}]]

	inputs, err = pipeline.Start[struct{}](cxt, files.PathsTargeter(args...))

	if err != nil {
		return err
	}

	pipelineOptions := pipeline.Flags(cmd)
	fileOptions := files.Flags(cmd)

	docReader := ascii.NewReader("Reading docs")
	docs, err := pipeline.Process[struct{}, render.InputDocument](cxt, pipelineOptions, docReader, inputs)
	if err != nil {
		return err
	}

	renderer := render.NewRenderer()
	var renders *xsync.MapOf[string, *pipeline.Data[string]]
	renders, err = pipeline.Process[render.InputDocument, string](cxt, pipelineOptions, renderer, docs)
	if err != nil {
		return err
	}

	writer := files.NewWriter[string]("Formatting docs", fileOptions)
	_, err = pipeline.Process[string, struct{}](cxt, pipelineOptions, writer, renders)
	return
}
