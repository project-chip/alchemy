package yaml2python

import (
	"github.com/project-chip/alchemy/cmd/common"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/pipeline"
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
	flags.String("specRoot", "connectedhomeip-spec", "the src root of your clone of CHIP-Specifications/connectedhomeip-spec")
	flags.String("sdkRoot", "connectedhomeip", "the root of your clone of project-chip/connectedhomeip")
	flags.String("templateRoot", "", "the root of your local template files; if not specified, Alchemy will use an internal copy")
	flags.Bool("overwrite", true, "overwrite existing test scripts")
}

func tp(cmd *cobra.Command, args []string) (err error) {

	cxt := cmd.Context()
	flags := cmd.Flags()

	specRoot, _ := flags.GetString("specRoot")
	sdkRoot, _ := flags.GetString("sdkRoot")

	asciiSettings := common.ASCIIDocAttributes(flags)
	fileOptions := files.Flags(flags)
	pipelineOptions := pipeline.Flags(flags)

	overwrite, _ := flags.GetBool("overwrite")
	templateRoot, _ := flags.GetString("templateRoot")
	generatorOptions := []python.GeneratorOption{
		python.Overwrite(overwrite),
		python.TemplateRoot(templateRoot),
	}

	err = yaml.Pipeline(cxt, specRoot, sdkRoot, pipelineOptions, asciiSettings, generatorOptions, fileOptions, args)

	return
}
