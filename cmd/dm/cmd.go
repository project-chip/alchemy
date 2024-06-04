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

	asciiSettings := common.ASCIIDocAttributes(cmd)
	fileOptions := files.Flags(cmd)
	pipelineOptions := pipeline.Flags(cmd)

	asciiSettings = append(ascii.GithubSettings(), asciiSettings...)

	specFiles, err := pipeline.Start[struct{}](cxt, files.SpecTargeter(specRoot))
	if err != nil {
		return err
	}

	docParser := ascii.NewParser(asciiSettings)
	specDocs, err := pipeline.Process[struct{}, *ascii.Doc](cxt, pipelineOptions, docParser, specFiles)
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
		if err != nil {
			return err
		}
	}

	renderer := dm.NewRenderer(sdkRoot)
	dataModelDocs, err := pipeline.Process[*ascii.Doc, string](cxt, pipelineOptions, renderer, specDocs)
	if err != nil {
		return err
	}

	clusterIDJSON, err := renderer.GenerateClusterIDsJson()
	if err != nil {
		return err
	}
	dataModelDocs.Store(clusterIDJSON.Path, clusterIDJSON)

	writer := files.NewWriter[string]("Writing data model", fileOptions)
	_, err = pipeline.Process[string, struct{}](cxt, pipelineOptions, writer, dataModelDocs)
	if err != nil {
		return err
	}
	return
}

func init() {
	Command.Flags().String("specRoot", "connectedhomeip-spec", "the src root of your clone of CHIP-Specifications/connectedhomeip-spec")
	Command.Flags().String("sdkRoot", "connectedhomeip", "the root of your clone of project-chip/connectedhomeip")
}
