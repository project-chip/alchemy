package dm

import (
	"context"

	"github.com/project-chip/alchemy/cmd/common"
	"github.com/project-chip/alchemy/dm"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
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
	dmRoot, _ := cmd.Flags().GetString("dmRoot")

	errata.OverlayErrataConfig(specRoot)

	asciiSettings := common.ASCIIDocAttributes(cmd)
	fileOptions := files.Flags(cmd)
	pipelineOptions := pipeline.Flags(cmd)

	specFiles, err := pipeline.Start[struct{}](cxt, spec.Targeter(specRoot))
	if err != nil {
		return err
	}

	docParser := spec.NewParser(asciiSettings)
	specDocs, err := pipeline.Process[struct{}, *spec.Doc](cxt, pipelineOptions, docParser, specFiles)
	if err != nil {
		return err
	}
	specBuilder := spec.NewBuilder()
	specBuilder.IgnoreHierarchy = true
	specDocs, err = pipeline.Process[*spec.Doc, *spec.Doc](cxt, pipelineOptions, &specBuilder, specDocs)
	if err != nil {
		return err
	}

	if len(args) > 0 {
		filter := files.NewPathFilter[*spec.Doc](args)
		specDocs, err = pipeline.Process[*spec.Doc, *spec.Doc](cxt, pipelineOptions, filter, specDocs)
		if err != nil {
			return err
		}
	}

	renderer := dm.NewRenderer(dmRoot)
	dataModelDocs, err := pipeline.Process[*spec.Doc, string](cxt, pipelineOptions, renderer, specDocs)
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
	Command.Flags().String("dmRoot", "connectedhomeip/data_model/master", "where to place the data model files")
}
