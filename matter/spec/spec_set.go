package spec

import (
	"context"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/pipeline"
)

type SpecPullRequest struct {
	Base           *Specification
	BaseInProgress *Specification
	Head           *Specification
	HeadInProgress *Specification
}

func loadSpecs(cxt context.Context, pipelineOptions pipeline.ProcessingOptions, specRoot string) (baseSpec *Specification, inProgressSpec *Specification, err error) {
	parserOptions := ParserOptions{Root: specRoot}

	var specDocs DocSet
	specDocs, err = LoadSpecDocs(cxt, parserOptions, pipelineOptions)
	if err != nil {
		return
	}

	baseSpec, _, err = Build(cxt, parserOptions, pipelineOptions, nil, specDocs, []asciidoc.AttributeName{})

	if err != nil {
		return
	}
	inProgressSpec, _, err = Build(cxt, parserOptions, pipelineOptions, nil, specDocs, []asciidoc.AttributeName{"in-progress"})
	return
}

func LoadSpecPullRequest(cxt context.Context, baseRoot string, headRoot string, pipelineOptions pipeline.ProcessingOptions) (ss SpecPullRequest, err error) {
	ss.Head, ss.HeadInProgress, err = loadSpecs(cxt, pipelineOptions, headRoot)
	if err != nil {
		return
	}

	ss.Base, ss.BaseInProgress, err = loadSpecs(cxt, pipelineOptions, baseRoot)
	if err != nil {
		return
	}
	return
}
