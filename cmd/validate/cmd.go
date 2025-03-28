package validate

import (
	"github.com/project-chip/alchemy/cmd/common"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/spec/validate"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "validate",
	Short: "validate the Matter specification object model",
	RunE:  validateSpec,
}

func init() {
	Command.Flags().String("specRoot", "connectedhomeip-spec", "the src root of your clone of CHIP-Specifications/connectedhomeip-spec")
}

func validateSpec(cmd *cobra.Command, args []string) (err error) {

	cxt := cmd.Context()
	flags := cmd.Flags()

	specRoot, _ := flags.GetString("specRoot")

	err = errata.LoadErrataConfig(specRoot)
	if err != nil {
		return
	}

	asciiSettings := common.ASCIIDocAttributes(flags)
	pipelineOptions := pipeline.Flags(flags)

	specFiles, err := pipeline.Start(cxt, spec.Targeter(specRoot))
	if err != nil {
		return err
	}

	docParser, err := spec.NewParser(specRoot, asciiSettings)
	if err != nil {
		return err
	}
	specDocs, err := pipeline.Parallel(cxt, pipelineOptions, docParser, specFiles)
	if err != nil {
		return err
	}

	specBuilder := spec.NewBuilder()
	_, err = pipeline.Collective(cxt, pipelineOptions, &specBuilder, specDocs)
	if err != nil {
		return err
	}

	validate.Validate(specBuilder.Spec)
	return

}
