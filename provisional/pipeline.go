package provisional

import (
	"context"
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/render"
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

	specs.Base, specs.BaseInProgress, err = loadSpecs(cxt, pipelineOptions, baseRoot)
	if err != nil {
		return
	}

	violations = compare(specs)
	for path, vs := range violations {
		for _, v := range vs {
			if v.Type.Has(ViolationTypeNonProvisional) {
				slog.Error("Provisionality violation", slog.String("path", path), matter.LogEntity("entity", v.Entity), slog.String("violationType", v.Type.String()))
			}
		}
	}

	if writer != nil {
		err = patchProvisional(cxt, pipelineOptions, specs.HeadInProgress, violations, writer)
	}
	return
}

func loadSpecs(cxt context.Context, pipelineOptions pipeline.ProcessingOptions, specRoot string) (baseSpec *spec.Specification, inProgressSpec *spec.Specification, err error) {
	parserOptions := spec.ParserOptions{Root: specRoot}

	var specDocs spec.DocSet
	specDocs, err = spec.LoadSpecDocs(cxt, parserOptions, pipelineOptions)
	if err != nil {
		return
	}

	baseSpec, _, err = spec.Build(cxt, parserOptions, pipelineOptions, nil, specDocs, []asciidoc.AttributeName{})

	if err != nil {
		return
	}
	inProgressSpec, _, err = spec.Build(cxt, parserOptions, pipelineOptions, nil, specDocs, []asciidoc.AttributeName{"in-progress"})
	return
}
