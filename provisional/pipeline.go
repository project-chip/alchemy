package provisional

import (
	"context"
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/render"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
)

type specs struct {
	Base           *spec.Specification
	BaseInProgress *spec.Specification
	Head           *spec.Specification
	HeadInProgress *spec.Specification
}

func Pipeline(cxt context.Context, baseRoot string, headRoot string, docPaths []string, pipelineOptions pipeline.ProcessingOptions, renderOptions []render.Option, writer files.Writer[string]) (violations map[string][]Violation, err error) {

	var specs specs

	specs.Head, specs.HeadInProgress, err = loadSpecs(cxt, pipelineOptions, headRoot)
	if err != nil {
		return
	}

	slog.Info("cluster count head", "count", len(specs.Head.Clusters))
	slog.Info("cluster count head in-progress", "count", len(specs.HeadInProgress.Clusters))

	specs.Base, specs.BaseInProgress, err = loadSpecs(cxt, pipelineOptions, baseRoot)
	if err != nil {
		return
	}

	slog.Info("cluster count base", "count", len(specs.Base.Clusters))
	slog.Info("cluster count base in-progress", "count", len(specs.BaseInProgress.Clusters))

	violations = compare(specs)
	for path, vs := range violations {
		for _, v := range vs {
			slog.Error("Provisionality violation", slog.String("path", path), matter.LogEntity("entity", v.Entity), slog.String("violationType", v.Type.String()))
		}
	}

	if writer != nil {
		err = patchProvisional(cxt, pipelineOptions, specs.HeadInProgress, violations, writer)
	}
	return
}

func loadSpecs(cxt context.Context, pipelineOptions pipeline.ProcessingOptions, specRoot string) (baseSpec *spec.Specification, inProgressSpec *spec.Specification, err error) {
	parserOptions := spec.ParserOptions{Inline: true, Root: specRoot}
	baseSpec, err = loadSpec(cxt, pipelineOptions, parserOptions)
	if err != nil {
		return
	}
	inProgressSpec, err = loadSpec(cxt, pipelineOptions, parserOptions, "in-progress")
	return
}

func loadSpec(cxt context.Context, pipelineOptions pipeline.ProcessingOptions, parserOptions spec.ParserOptions, attributes ...asciidoc.AttributeName) (s *spec.Specification, err error) {

	var specParser spec.Parser
	specParser, err = spec.NewParser(attributes, parserOptions)
	if err != nil {
		return
	}

	err = errata.LoadErrataConfig(parserOptions.Root)
	if err != nil {
		return
	}

	var specPaths pipeline.Paths
	specPaths, err = pipeline.Start(cxt, specParser.Targets)
	if err != nil {
		return
	}

	var specDocs spec.DocSet
	specDocs, err = pipeline.Parallel(cxt, pipelineOptions, specParser, specPaths)
	if err != nil {
		return
	}

	specBuilder := spec.NewBuilder(parserOptions.Root)
	specDocs, err = pipeline.Collective(cxt, pipelineOptions, &specBuilder, specDocs)
	if err != nil {
		return
	}

	s = specBuilder.Spec
	return
}
