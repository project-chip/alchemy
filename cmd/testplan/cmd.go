package testplan

import (
	"context"
	"log/slog"

	"github.com/hasty/alchemy/cmd/common"
	"github.com/hasty/alchemy/internal/files"
	"github.com/hasty/alchemy/internal/parse"
	"github.com/hasty/alchemy/internal/pipeline"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/spec"
	"github.com/hasty/alchemy/testplan"
	"github.com/hasty/alchemy/zap"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "testplan",
	Short: "create an initial test plan from the spec",
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
	testRoot, _ := cmd.Flags().GetString("testRoot")
	overwrite, _ := cmd.Flags().GetBool("overwrite")

	asciiSettings := common.ASCIIDocAttributes(cmd)
	fileOptions := files.Flags(cmd)
	pipelineOptions := pipeline.Flags(cmd)

	asciiSettings = append(spec.GithubSettings(), asciiSettings...)

	specFiles, err := pipeline.Start[struct{}](cxt, files.SpecTargeter(specRoot))
	if err != nil {
		return err
	}

	docParser := spec.NewParser(asciiSettings)
	specDocs, err := pipeline.Process[struct{}, *spec.Doc](cxt, pipelineOptions, docParser, specFiles)
	if err != nil {
		return err
	}

	var specParser files.SpecParser
	specDocs, err = pipeline.Process[*spec.Doc, *spec.Doc](cxt, pipelineOptions, &specParser, specDocs)
	if err != nil {
		return err
	}

	var appClusterIndexes pipeline.Map[string, *pipeline.Data[*spec.Doc]]
	appClusterIndexes, err = pipeline.Process[*spec.Doc, *spec.Doc](cxt, pipelineOptions, common.NewDocTypeFilter(matter.DocTypeAppClusterIndex), specDocs)

	if err != nil {
		return err
	}

	_, err = pipeline.ProcessSerialFunc[*spec.Doc, *spec.Doc](cxt, pipelineOptions, appClusterIndexes, "Assigning index domains", func(cxt context.Context, input *pipeline.Data[*spec.Doc], index, total int32) (outputs []*pipeline.Data[*spec.Doc], extra []*pipeline.Data[*spec.Doc], err error) {
		doc := input.Content
		top := parse.FindFirst[*spec.Section](doc.Elements())
		if top != nil {
			doc.Domain = zap.StringToDomain(top.Name)
			slog.DebugContext(cxt, "Assigned domain", "file", top.Name, "domain", doc.Domain)
		}
		return
	})
	if err != nil {
		return err
	}

	if len(args) > 0 { // Filter the spec by whatever extra args were passed
		filter := files.NewPathFilter[*spec.Doc](args)
		specDocs, err = pipeline.Process[*spec.Doc, *spec.Doc](cxt, pipelineOptions, filter, specDocs)
		if err != nil {
			return err
		}
	}

	var clusters pipeline.Map[string, *pipeline.Data[*matter.Cluster]]
	clusters, err = pipeline.Process[*spec.Doc, *matter.Cluster](cxt, pipelineOptions, &common.EntityFilter[*spec.Doc, *matter.Cluster]{}, specDocs)
	if err != nil {
		return err
	}

	generator := testplan.NewGenerator(testRoot, overwrite)
	var testplans pipeline.Map[string, *pipeline.Data[string]]
	testplans, err = pipeline.Process[*matter.Cluster, string](cxt, pipelineOptions, generator, clusters)

	writer := files.NewWriter[string]("Writing test plans", fileOptions)
	_, err = pipeline.Process[string, struct{}](cxt, pipelineOptions, writer, testplans)
	if err != nil {
		return err
	}

	return
}
