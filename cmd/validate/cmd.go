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
	spec.ParserFlags(Command.Flags())
}

func validateSpec(cmd *cobra.Command, args []string) (err error) {

	cxt := cmd.Context()
	flags := cmd.Flags()

	parserOptions := spec.ParserOptions(flags)
	asciiSettings := common.ASCIIDocAttributes(flags)

	specParser, err := spec.NewParser(asciiSettings, parserOptions...)
	if err != nil {
		return err
	}

	err = errata.LoadErrataConfig(specParser.Root)
	if err != nil {
		return
	}

	pipelineOptions := pipeline.PipelineOptions(flags)

	specFiles, err := pipeline.Start(cxt, specParser.Targets)
	if err != nil {
		return err
	}

	specDocs, err := pipeline.Parallel(cxt, pipelineOptions, specParser, specFiles)
	if err != nil {
		return err
	}

	specBuilder := spec.NewBuilder(specParser.Root)
	_, err = pipeline.Collective(cxt, pipelineOptions, &specBuilder, specDocs)
	if err != nil {
		return err
	}

	validate.Validate(specBuilder.Spec)
	return

}
