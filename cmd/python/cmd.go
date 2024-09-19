package python

import (
	"context"

	"github.com/project-chip/alchemy/cmd/common"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/testing/generate"
	"github.com/project-chip/alchemy/testing/parse"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "python",
	Short: "create a shell python script from a test plan",
	RunE:  tp,
}

func init() {
	Command.Flags().String("specRoot", "connectedhomeip-spec", "the src root of your clone of CHIP-Specifications/connectedhomeip-spec")
	Command.Flags().String("sdkRoot", "connectedhomeip", "the root of your clone of project-chip/connectedhomeip")
	Command.Flags().Bool("overwrite", false, "overwrite existing test plans")
}

func tp(cmd *cobra.Command, args []string) (err error) {

	cxt := context.Background()

	specRoot, _ := cmd.Flags().GetString("specRoot")
	sdkRoot, _ := cmd.Flags().GetString("sdkRoot")
	//overwrite, _ := cmd.Flags().GetBool("overwrite")
	fileOptions := files.Flags(cmd)
	pipelineOptions := pipeline.Flags(cmd)
	asciiSettings := common.ASCIIDocAttributes(cmd)

	var inputs pipeline.Map[string, *pipeline.Data[struct{}]]
	inputs, err = pipeline.Start[struct{}](cxt, files.PathsTargeter(args...))
	if err != nil {
		return err
	}

	errata.LoadErrataConfig(specRoot)

	var parser parse.TestYamlParser
	parser, err = parse.NewTestYamlParser(sdkRoot)
	if err != nil {
		return err
	}

	var tests pipeline.Map[string, *pipeline.Data[*parse.Test]]
	tests, err = pipeline.Process[struct{}, *parse.Test](cxt, pipelineOptions, parser, inputs)
	if err != nil {
		return err
	}

	docParser, err := spec.NewParser(specRoot, asciiSettings)
	if err != nil {
		return err
	}

	specFiles, err := pipeline.Start[struct{}](cxt, spec.Targeter(specRoot))
	if err != nil {
		return err
	}

	specDocs, err := pipeline.Process[struct{}, *spec.Doc](cxt, pipelineOptions, docParser, specFiles)
	if err != nil {
		return err
	}

	specBuilder := spec.NewBuilder()
	_, err = pipeline.Process[*spec.Doc, *spec.Doc](cxt, pipelineOptions, &specBuilder, specDocs)
	if err != nil {
		return err
	}

	labels, err := parse.LoadPICSLabels(sdkRoot)
	if err != nil {
		return err
	}

	generator := generate.NewPythonTestGenerator(specBuilder.Spec, sdkRoot, true, labels)
	var scripts pipeline.Map[string, *pipeline.Data[string]]
	scripts, err = pipeline.Process[*parse.Test, string](cxt, pipelineOptions, generator, tests)
	if err != nil {
		return err
	}

	writer := files.NewWriter[string]("Writing test scripts", fileOptions)
	_, err = pipeline.Process[string, struct{}](cxt, pipelineOptions, writer, scripts)
	if err != nil {
		return err
	}

	return
}
