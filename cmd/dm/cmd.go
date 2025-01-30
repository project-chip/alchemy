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

	errata.LoadErrataConfig(specRoot)

	asciiSettings := common.ASCIIDocAttributes(cmd)
	fileOptions := files.Flags(cmd)
	pipelineOptions := pipeline.Flags(cmd)

	docParser, err := spec.NewParser(specRoot, asciiSettings)
	if err != nil {
		return err
	}

	specBuilder := spec.NewBuilder(spec.IgnoreHierarchy(true))

	specFiles, err := pipeline.Start(cxt, spec.Targeter(specRoot))
	if err != nil {
		return err
	}

	specDocs, err := pipeline.Parallel(cxt, pipelineOptions, docParser, specFiles)
	if err != nil {
		return err
	}
	specDocs, err = pipeline.Collective(cxt, pipelineOptions, &specBuilder, specDocs)
	if err != nil {
		return err
	}

	if len(args) > 0 {
		filter := files.NewPathFilter[*spec.Doc](args)
		specDocs, err = pipeline.Collective(cxt, pipelineOptions, filter, specDocs)
		if err != nil {
			return err
		}
	}

	dataModelRenderer := dm.NewRenderer(dmRoot, specBuilder.Spec)

	dataModelDocs, err := pipeline.Parallel(cxt, pipelineOptions, dataModelRenderer, specDocs)
	if err != nil {
		return err
	}

	clusterIDJSON, err := dataModelRenderer.GenerateClusterIDsJson()
	if err != nil {
		return err
	}
	dataModelDocs.Store(clusterIDJSON.Path, clusterIDJSON)

	writer := files.NewWriter[string]("Writing data model", fileOptions)
	err = writer.Write(cxt, dataModelDocs, pipelineOptions)
	if err != nil {
		return err
	}
	return
}

func init() {
	Command.Flags().String("specRoot", "connectedhomeip-spec", "the src root of your clone of CHIP-Specifications/connectedhomeip-spec")
	Command.Flags().String("dmRoot", "connectedhomeip/data_model/master", "where to place the data model files")
}
