package disco

import (
	"context"
	"fmt"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/render"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/paths"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
)

func Pipeline(cxt context.Context, parserOptions spec.ParserOptions, docPaths []string, pipelineOptions pipeline.ProcessingOptions, discoOptions DiscoOptions, attributes []asciidoc.AttributeName, renderOptions []render.Option, writer files.Writer[string]) (err error) {

	var docs spec.DocSet
	var specification *spec.Specification

	if parserOptions.Root == "" {
		allPaths, e := paths.NewTargeter(docPaths...)(cxt)
		if e == nil {
			parserOptions.Root = spec.DeriveSpecPathFromPaths(allPaths)
		}
	}

	err = errata.LoadErrataConfig(parserOptions.Root)
	if err != nil {
		return
	}

	if parserOptions.Root == "" {
		err = fmt.Errorf("disco ball requires spec root")
		return
	}

	specification, docs, err = spec.Parse(cxt, parserOptions, pipelineOptions, nil, attributes)
	if err != nil {
		return err
	}
	if len(docPaths) > 0 {
		filter := paths.NewIncludeFilter[*asciidoc.Document](parserOptions.Root, docPaths)
		docs, err = pipeline.Collective(cxt, pipelineOptions, filter, docs)
		if err != nil {
			return err
		}
	}

	if err != nil {
		return err
	}

	baller := NewBaller(specification, discoOptions)

	var balledDocs spec.DocSet
	balledDocs, err = pipeline.Parallel(cxt, pipelineOptions, baller, docs)
	if err != nil {
		return err
	}

	anchorNormalizer := newAnchorNormalizer(specification, discoOptions)
	var normalizedDocs pipeline.Map[string, *pipeline.Data[*asciidoc.Document]]
	normalizedDocs, err = pipeline.Collective(cxt, pipelineOptions, anchorNormalizer, balledDocs)
	if err != nil {
		return err
	}

	renderer := render.NewRenderer(renderOptions...)
	var renders pipeline.StringSet
	renders, err = pipeline.Parallel(cxt, pipelineOptions, renderer, normalizedDocs)
	if err != nil {
		return err
	}

	err = writer.Write(cxt, renders, pipelineOptions)
	return
}
