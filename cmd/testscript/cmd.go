package testscript

import (
	"github.com/project-chip/alchemy/cmd/common"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/testscript"
	"github.com/project-chip/alchemy/testscript/python"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "testscript [filename_pattern]",
	Short: "create shell python scripts from the spec, optionally filtered to the files specified by filename_pattern",
	RunE:  tp,
}

func init() {
	flags := Command.Flags()
	spec.ParserFlags(flags)
	flags.String("sdkRoot", "connectedhomeip", "the root of your clone of project-chip/connectedhomeip")
	flags.String("templateRoot", "", "the root of your local template files; if not specified, Alchemy will use an internal copy")
	flags.Bool("overwrite", true, "overwrite existing test scripts")
}

func tp(cmd *cobra.Command, args []string) (err error) {

	cxt := cmd.Context()
	flags := cmd.Flags()

	sdkRoot, _ := flags.GetString("sdkRoot")

	asciiSettings := common.ASCIIDocAttributes(flags)
	fileOptions := files.OutputOptions(flags)
	pipelineOptions := pipeline.PipelineOptions(flags)
	parserOptions := spec.ParserOptions(flags)

	overwrite, _ := flags.GetBool("overwrite")
	templateRoot, _ := flags.GetString("templateRoot")
	generatorOptions := []python.GeneratorOption{
		python.Overwrite(overwrite),
		python.TemplateRoot(templateRoot),
	}

	err = testscript.Pipeline(cxt, specRoot, sdkRoot, pipelineOptions, asciiSettings, generatorOptions, fileOptions, args)

	return
}
