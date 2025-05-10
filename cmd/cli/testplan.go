package cli

import (
	"context"
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc/render"
	"github.com/project-chip/alchemy/cmd/common"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	testplanRender "github.com/project-chip/alchemy/testplan/render"
	"github.com/project-chip/alchemy/zap"
)

type TestPlan struct {
	testplanRender.RendererOptions `embed:""`

	common.ASCIIDocAttributes  `embed:""`
	spec.ParserOptions         `embed:""`
	spec.FilterOptions         `embed:""`
	pipeline.ProcessingOptions `embed:""`
	files.OutputOptions        `embed:""`
}

func (c *TestPlan) Run(cc *Context) (err error) {

	specParser, err := spec.NewParser(c.ASCIIDocAttributes.ToList(), c.ParserOptions)
	if err != nil {
		return err
	}

	err = errata.LoadErrataConfig(c.ParserOptions.Root)
	if err != nil {
		return
	}

	specFiles, err := pipeline.Start(cc, specParser.Targets)
	if err != nil {
		return err
	}

	specDocs, err := pipeline.Parallel(cc, c.ProcessingOptions, specParser, specFiles)
	if err != nil {
		return err
	}

	specBuilder := spec.NewBuilder(c.ParserOptions.Root)
	specDocs, err = pipeline.Collective(cc, c.ProcessingOptions, &specBuilder, specDocs)
	if err != nil {
		return err
	}

	var appClusterIndexes spec.DocSet
	appClusterIndexes, err = pipeline.Collective(cc, c.ProcessingOptions, common.NewDocTypeFilter(matter.DocTypeAppClusterIndex), specDocs)

	if err != nil {
		return err
	}

	domainIndexer := func(cxt context.Context, input *pipeline.Data[*spec.Doc], index, total int32) (outputs []*pipeline.Data[*spec.Doc], extra []*pipeline.Data[*spec.Doc], err error) {
		doc := input.Content
		top := parse.FindFirst[*spec.Section](doc)
		if top != nil {
			doc.Domain = zap.StringToDomain(top.Name)
			slog.DebugContext(cxt, "Assigned domain", "file", top.Name, "domain", doc.Domain)
		}
		return
	}

	_, err = pipeline.Parallel(cc, c.ProcessingOptions, pipeline.ParallelFunc("Assigning index domains", domainIndexer), appClusterIndexes)
	if err != nil {
		return err
	}

	specDocs, err = filterSpecDocs(cc, specDocs, specBuilder.Spec, c.FilterOptions, c.ProcessingOptions)
	if err != nil {
		return
	}

	generator := testplanRender.NewRenderer(c.RendererOptions)
	var testplans pipeline.StringSet
	testplans, err = pipeline.Parallel(cc, c.ProcessingOptions, generator, specDocs)
	if err != nil {
		return err
	}

	docReader, err := spec.NewStringReader("Reading test plans", c.TestRoot)
	if err != nil {
		return err
	}
	testplanDocs, err := pipeline.Parallel(cc, c.ProcessingOptions, docReader, testplans)
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
	renders, err = pipeline.Parallel(cc, c.ProcessingOptions, renderer, ids)
	if err != nil {
		return err
	}

	writer := files.NewWriter[string]("Writing test plans", c.OutputOptions)
	err = writer.Write(cc, renders, c.ProcessingOptions)

	return
}
