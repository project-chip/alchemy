package dm

import (
	"github.com/project-chip/alchemy/cmd/common"
	"github.com/project-chip/alchemy/dm"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/paths"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:     "dm [filename_pattern]",
	Short:   "transmute the Matter spec into data model XML; optionally filtered to the files specified in filename_pattern",
	Aliases: []string{"datamodel", "data-model"},
	RunE:    dataModel,
}

func dataModel(cmd *cobra.Command, args []string) (err error) {
	cxt := cmd.Context()

	flags := cmd.Flags()
	specRoot, _ := flags.GetString("specRoot")
	dmRoot, _ := flags.GetString("dmRoot")

	err = errata.LoadErrataConfig(specRoot)
	if err != nil {
		return
	}

	asciiSettings := common.ASCIIDocAttributes(flags)
	fileOptions := files.Flags(flags)
	pipelineOptions := pipeline.Flags(flags)

	docParser, err := spec.NewParser(specRoot, asciiSettings)
	if err != nil {
		return err
	}

	docParser.Inline, _ = flags.GetBool("inline")

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
		filter := paths.NewFilter[*spec.Doc](specRoot, args)
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
	flags := Command.Flags()
	flags.String("specRoot", "connectedhomeip-spec", "the src root of your clone of CHIP-Specifications/connectedhomeip-spec")
	flags.String("dmRoot", "connectedhomeip/data_model/master", "where to place the data model files")
	flags.Bool("inline", false, "use inline parser")
}
