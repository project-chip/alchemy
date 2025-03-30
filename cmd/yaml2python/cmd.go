package yaml2python

import (
	"github.com/project-chip/alchemy/cmd/common"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/testscript/python"
	"github.com/project-chip/alchemy/testscript/yaml"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "yaml2python [filename_pattern]",
	Short: "create a shell python script from a test YAML, optionally filtered to the files specified by filename_pattern",
	RunE:  tp,
}

func init() {
	flags := Command.Flags()
	spec.ParserFlags(flags)
	flags.String("sdkRoot", "connectedhomeip", "the root of your clone of project-chip/connectedhomeip")
	python.Flags(flags)
}

func tp(cmd *cobra.Command, args []string) (err error) {

	cxt := cmd.Context()
	flags := cmd.Flags()

	sdkRoot, _ := flags.GetString("sdkRoot")

	asciiSettings := common.ASCIIDocAttributes(flags)
	fileOptions := files.OutputOptions(flags)
	pipelineOptions := pipeline.PipelineOptions(flags)
	generatorOptions := python.GeneratorOptions(flags)
	parserOptions := spec.ParserOptions(flags)

	err = yaml.Pipeline(cxt, sdkRoot, pipelineOptions, parserOptions, asciiSettings, generatorOptions, fileOptions, args)

	return
}
