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

	err = testscript.Pipeline(cxt, specRoot, sdkRoot, pipelineOptions, asciiSettings, generatorOptions, fileOptions, args)

	return
}
