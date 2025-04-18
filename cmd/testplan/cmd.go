package testplan

import (
	"context"
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc/render"
	"github.com/project-chip/alchemy/cmd/common"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/internal/paths"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	testplanRender "github.com/project-chip/alchemy/testplan/render"
	"github.com/project-chip/alchemy/zap"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "testplan [filename_pattern]",
	Short: "create an initial test plan from the spec, optionally filtered to the files specified in filename_pattern",
	RunE:  tp,
}

func init() {
	flags := Command.Flags()

	flags.String("specRoot", "connectedhomeip-spec", "the src root of your clone of CHIP-Specifications/connectedhomeip-spec")
	flags.String("testRoot", "chip-test-plans", "the root of your clone of CHIP-Specifications/chip-test-plans")
	flags.String("templateRoot", "", "the root of your local template files; if not specified, Alchemy will use an internal copy")
	flags.Bool("overwrite", false, "overwrite existing test plans")
}

func tp(cmd *cobra.Command, args []string) (err error) {

	cxt := cmd.Context()

	flags := cmd.Flags()

	testRoot, _ := flags.GetString("testRoot")
	overwrite, _ := flags.GetBool("overwrite")
	templateRoot, _ := flags.GetString("templateRoot")

	asciiSettings := common.ASCIIDocAttributes(flags)
	fileOptions := files.OutputOptions(flags)
	pipelineOptions := pipeline.PipelineOptions(flags)
	parserOptions := spec.ParserOptions(flags)
	var testplanGeneratorOptions []testplanRender.GeneratorOption

	specParser, err := spec.NewParser(asciiSettings, parserOptions...)
	if err != nil {
		return err
	}

	if templateRoot != "" {
		testplanGeneratorOptions = append(testplanGeneratorOptions, testplanRender.TemplateRoot(templateRoot))
	}

	err = errata.LoadErrataConfig(specParser.Root)
	if err != nil {
		return
	}

	specFiles, err := pipeline.Start(cxt, specParser.Targets)
	if err != nil {
		return err
	}

	specDocs, err := pipeline.Parallel(cxt, pipelineOptions, specParser, specFiles)
	if err != nil {
		return err
	}

	specBuilder := spec.NewBuilder(specParser.Root)
	specDocs, err = pipeline.Collective(cxt, pipelineOptions, &specBuilder, specDocs)
	if err != nil {
		return err
	}

	var appClusterIndexes spec.DocSet
	appClusterIndexes, err = pipeline.Collective(cxt, pipelineOptions, common.NewDocTypeFilter(matter.DocTypeAppClusterIndex), specDocs)

	if err != nil {
		return err
	}

	domainIndexer := func(cxt context.Context, input *pipeline.Data[*spec.Doc], index, total int32) (outputs []*pipeline.Data[*spec.Doc], extra []*pipeline.Data[*spec.Doc], err error) {
		doc := input.Content
		top := parse.FindFirst[*spec.Section](doc.Elements())
		if top != nil {
			doc.Domain = zap.StringToDomain(top.Name)
			slog.DebugContext(cxt, "Assigned domain", "file", top.Name, "domain", doc.Domain)
		}
		return
	}

	_, err = pipeline.Parallel(cxt, pipelineOptions, pipeline.ParallelFunc("Assigning index domains", domainIndexer), appClusterIndexes)
	if err != nil {
		return err
	}

	if len(args) > 0 { // Filter the spec by whatever extra args were passed
		filter := paths.NewFilter[*spec.Doc](specParser.Root, args)
		specDocs, err = pipeline.Collective(cxt, pipelineOptions, filter, specDocs)
		if err != nil {
			return err
		}
	}

	generator := testplanRender.NewRenderer(testRoot, overwrite, testplanGeneratorOptions...)
	var testplans pipeline.StringSet
	testplans, err = pipeline.Parallel(cxt, pipelineOptions, generator, specDocs)
	if err != nil {
		return err
	}

	docReader, err := spec.NewStringReader("Reading test plans", testRoot)
	if err != nil {
		return err
	}
	testplanDocs, err := pipeline.Parallel(cxt, pipelineOptions, docReader, testplans)
	if err != nil {
		return err
	}

	ids := pipeline.NewConcurrentMapPresized[string, *pipeline.Data[render.InputDocument]](testplanDocs.Size())
	err = pipeline.Cast(testplanDocs, ids)
	if err != nil {
		return err
	}

	renderer := render.NewRenderer()
	var renders pipeline.StringSet
	renders, err = pipeline.Parallel(cxt, pipelineOptions, renderer, ids)
	if err != nil {
		return err
	}

	writer := files.NewWriter[string]("Writing test plans", fileOptions)
	err = writer.Write(cxt, renders, pipelineOptions)

	return
}
