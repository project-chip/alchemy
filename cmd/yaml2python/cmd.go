package yaml2python

import (
	"github.com/project-chip/alchemy/cmd/common"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/testscript/python"
	"github.com/project-chip/alchemy/yaml2python"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "yaml2python",
	Short: "create a shell python script from a test YAML",
	RunE:  tp,
}

func init() {
	Command.Flags().String("specRoot", "connectedhomeip-spec", "the src root of your clone of CHIP-Specifications/connectedhomeip-spec")
	Command.Flags().String("sdkRoot", "connectedhomeip", "the root of your clone of project-chip/connectedhomeip")
	Command.Flags().String("templateRoot", "", "the root of your local template files; if not specified, Alchemy will use an internal copy")
	Command.Flags().Bool("overwrite", true, "overwrite existing test scripts")
}

func tp(cmd *cobra.Command, args []string) (err error) {

	cxt := cmd.Context()

	specRoot, _ := cmd.Flags().GetString("specRoot")
	sdkRoot, _ := cmd.Flags().GetString("sdkRoot")

	asciiSettings := common.ASCIIDocAttributes(cmd)
	fileOptions := files.Flags(cmd)
	pipelineOptions := pipeline.Flags(cmd)

	overwrite, _ := cmd.Flags().GetBool("overwrite")
	templateRoot, _ := cmd.Flags().GetString("templateRoot")
	generatorOptions := []python.GeneratorOption{
		python.Overwrite(overwrite),
		python.TemplateRoot(templateRoot),
	}

	err = yaml2python.Pipeline(cxt, specRoot, sdkRoot, pipelineOptions, asciiSettings, generatorOptions, fileOptions, args)

	return
}
