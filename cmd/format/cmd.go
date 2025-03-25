package format

import (
	"github.com/project-chip/alchemy/asciidoc/render"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "format filename_pattern",
	Short: "format Matter spec documents specified by filename_pattern",
	RunE:  format,
}

func format(cmd *cobra.Command, args []string) (err error) {
	cxt := cmd.Context()

	var inputs pipeline.Paths

	inputs, err = pipeline.Start(cxt, files.PathsTargeter(args...))

	if err != nil {
		return err
	}

	pipelineOptions := pipeline.Flags(cmd)
	fileOptions := files.Flags(cmd)

	docReader, err := spec.NewReader("Reading docs", "")
	if err != nil {
		return err
	}
	docs, err := pipeline.Parallel(cxt, pipelineOptions, docReader, inputs)
	if err != nil {
		return err
	}

	ids := pipeline.NewConcurrentMapPresized[string, *pipeline.Data[render.InputDocument]](docs.Size())
	err = pipeline.Cast(docs, ids)
	if err != nil {
		return err
	}

	wrap, _ := cmd.Flags().GetInt("wrap")
	renderer := render.NewRenderer(render.Wrap(wrap))
	var renders pipeline.StringSet
	renders, err = pipeline.Parallel(cxt, pipelineOptions, renderer, ids)
	if err != nil {
		return err
	}

	writer := files.NewWriter[string]("Formatting docs", fileOptions)
	err = writer.Write(cxt, renders, pipelineOptions)
	return
}

func init() {
	Command.Flags().Int("wrap", 0, "the maximum length of a line")
}
