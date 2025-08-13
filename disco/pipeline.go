package disco

import (
	"context"
	"fmt"

	"github.com/project-chip/alchemy/asciidoc/render"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/paths"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
)

func Pipeline(cxt context.Context, specRoot string, docPaths []string, pipelineOptions pipeline.ProcessingOptions, discoOptions DiscoOptions, renderOptions []render.Option, writer files.Writer[string]) (err error) {

	var docs spec.DocSet

	if specRoot == "" {
		allPaths, e := paths.NewTargeter(docPaths...)(cxt)
		if e == nil {
			specRoot = spec.DeriveSpecPathFromPaths(allPaths)
		}
	}

	err = errata.LoadErrataConfig(specRoot)
	if err != nil {
		return
	}

	if specRoot != "" {

		_, docs, err = spec.Read(cxt, pipelineOptions, nil, specRoot, docPaths)
		if err != nil {
			return err
		}

	} else if len(docPaths) > 0 {
		var inputs pipeline.Paths
		inputs, err = pipeline.Start(cxt, paths.NewTargeter(docPaths...))
		if err != nil {
			return err
		}

		docReader, err := spec.NewReader("Reading docs", specRoot)
		if err != nil {
			return err
		}
		docs, err = pipeline.Parallel(cxt, pipelineOptions, docReader, inputs)
		if err != nil {
			return err
		}
	} else {
		err = fmt.Errorf("disco ball requires spec root or document paths")
		return
	}

	if err != nil {
		return err
	}

	baller := NewBaller(discoOptions)

	var balledDocs spec.DocSet
	balledDocs, err = pipeline.Parallel(cxt, pipelineOptions, baller, docs)
	if err != nil {
		return err
	}

	anchorNormalizer := newAnchorNormalizer(discoOptions)
	var normalizedDocs pipeline.Map[string, *pipeline.Data[render.InputDocument]]
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
