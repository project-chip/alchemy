package provisional

import (
	"context"
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc/render"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
)

func Pipeline(cxt context.Context, baseRoot string, headRoot string, docPaths []string, pipelineOptions pipeline.ProcessingOptions, renderOptions []render.Option, writer files.Writer[string]) (violations map[string][]spec.Violation, err error) {
	specs, err := spec.LoadSpecPullRequest(cxt, baseRoot, headRoot, docPaths, pipelineOptions, renderOptions)

	if err != nil {
		return nil, err
	}

	return ProcessSpecs(cxt, &specs, pipelineOptions, writer)
}

func ProcessSpecs(cxt context.Context, specs *spec.SpecPullRequest, pipelineOptions pipeline.ProcessingOptions, writer files.Writer[string]) (violations map[string][]spec.Violation, err error) {
	violations = compare(*specs)

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
