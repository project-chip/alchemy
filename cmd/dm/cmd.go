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
	dmRoot, _ := flags.GetString("dmRoot")

	asciiSettings := common.ASCIIDocAttributes(flags)
	fileOptions := files.OutputOptions(flags)
	pipelineOptions := pipeline.PipelineOptions(flags)

	parserOptions := spec.ParserOptions(flags)

	specParser, err := spec.NewParser(asciiSettings, parserOptions...)
	if err != nil {
		return err
	}

	err = errata.LoadErrataConfig(specParser.Root)
	if err != nil {
		return
	}

	specBuilder := spec.NewBuilder(specParser.Root, spec.IgnoreHierarchy(true))

	specFiles, err := pipeline.Start(cxt, specParser.Targets)
	if err != nil {
		return err
	}

	specDocs, err := pipeline.Parallel(cxt, pipelineOptions, specParser, specFiles)
	if err != nil {
		return err
	}
	specDocs, err = pipeline.Collective(cxt, pipelineOptions, &specBuilder, specDocs)
	if err != nil {
		return err
	}

	if len(args) > 0 {
		filter := paths.NewFilter[*spec.Doc](specParser.Root, args)
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
	spec.ParserFlags(flags)
	flags.String("dmRoot", "connectedhomeip/data_model/master", "where to place the data model files")
}
