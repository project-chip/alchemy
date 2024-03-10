package dm

import (
	"context"

	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/cmd/common"
	"github.com/hasty/alchemy/dm"
	"github.com/hasty/alchemy/internal/files"
	"github.com/hasty/alchemy/internal/pipeline"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:     "dm",
	Short:   "transmute the Matter spec into data model XML",
	Aliases: []string{"datamodel", "data-model"},
	RunE:    dataModel,
}

func dataModel(cmd *cobra.Command, args []string) (err error) {
	cxt := context.Background()

	specRoot, _ := cmd.Flags().GetString("specRoot")
	sdkRoot, _ := cmd.Flags().GetString("sdkRoot")

	asciiSettings := common.AsciiDocAttributes(cmd)
	fileOptions := files.Flags(cmd)
	pipelineOptions := pipeline.Flags(cmd)

	specFiles, err := pipeline.Start[struct{}](cxt, files.SpecTargeter(specRoot))
	if err != nil {
		return err
	}

	docReader := ascii.NewParser(pipeline.ProcessorTypeParallel, asciiSettings)
	specDocs, err := pipeline.Process[struct{}, *ascii.Doc](cxt, pipelineOptions, docReader, specFiles)
	if err != nil {
		return err
	}
	var specParser files.SpecParser
	specDocs, err = pipeline.Process[*ascii.Doc, *ascii.Doc](cxt, pipelineOptions, &specParser, specDocs)
	if err != nil {
		return err
	}

	if len(args) > 0 {
		filter := files.NewPathFilter[*ascii.Doc](args)
		specDocs, err = pipeline.Process[*ascii.Doc, *ascii.Doc](cxt, pipelineOptions, filter, specDocs)
	}

	renderer := dm.NewRenderer(sdkRoot)
	dataModelDocs, err := pipeline.Process[*ascii.Doc, string](cxt, pipelineOptions, renderer, specDocs)
	if err != nil {
		return err
	}

	writer := files.NewWriter("Writing data model", fileOptions)
	_, err = pipeline.Process[string, struct{}](cxt, pipelineOptions, writer, dataModelDocs)
	return
}

func init() {
	Command.Flags().String("specRoot", "", "the root of your clone of CHIP-Specifications/connectedhomeip-spec")
	Command.Flags().String("sdkRoot", "", "the root of your clone of project-chip/connectedhomeip")
	_ = Command.MarkFlagRequired("specRoot")
	_ = Command.MarkFlagRequired("sdkRoot")
}
