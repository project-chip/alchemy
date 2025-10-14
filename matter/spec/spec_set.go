package spec

import (
	"context"
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/render"
	"github.com/project-chip/alchemy/internal/pipeline"
)

type SpecSet struct {
	Base           *Specification
	BaseInProgress *Specification
	Head           *Specification
	HeadInProgress *Specification
}

func loadSpec(cxt context.Context, pipelineOptions pipeline.ProcessingOptions, specRoot string) (baseSpec *Specification, inProgressSpec *Specification, err error) {
	parserOptions := ParserOptions{Root: specRoot}
	baseSpec, _, err = Parse(cxt, parserOptions, pipelineOptions, nil, []asciidoc.AttributeName{})

	if err != nil {
		return
	}
	inProgressSpec, _, err = Parse(cxt, parserOptions, pipelineOptions, nil, []asciidoc.AttributeName{"in-progress"})
	return
}

func LoadSpecSet(cxt context.Context, baseRoot string, headRoot string, docPaths []string, pipelineOptions pipeline.ProcessingOptions, renderOptions []render.Option) (ss SpecSet, err error) {
	ss.Head, ss.HeadInProgress, err = loadSpec(cxt, pipelineOptions, headRoot)
	if err != nil {
		return
	}

	slog.Info("cluster count head", "count", len(ss.Head.Clusters))
	slog.Info("cluster count head in-progress", "count", len(ss.HeadInProgress.Clusters))

	ss.Base, ss.BaseInProgress, err = loadSpec(cxt, pipelineOptions, baseRoot)
	if err != nil {
		return
	}

	slog.Info("cluster count base", "count", len(ss.Base.Clusters))
	slog.Info("cluster count base in-progress", "count", len(ss.BaseInProgress.Clusters))
	return
}
