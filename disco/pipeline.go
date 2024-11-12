package disco

import (
	"context"
	"fmt"

	"github.com/project-chip/alchemy/asciidoc/render"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
)

func Pipeline(cxt context.Context, specRoot string, docPaths []string, pipelineOptions pipeline.Options, discoOptions []Option, renderOptions []render.Option, writer files.Writer[string]) (err error) {

	var docs spec.DocSet

	if specRoot == "" {
		allPaths, e := files.PathsTargeter(docPaths...)(cxt)
		if e == nil {
			specRoot = spec.DeriveSpecPathFromPaths(allPaths)
		}
	}

	errata.LoadErrataConfig(specRoot)

	if specRoot != "" {

		specTargeter := spec.Targeter(specRoot)

		var inputs pipeline.Paths
		inputs, err = pipeline.Start(cxt, specTargeter)
		if err != nil {
			return err
		}

		docReader, err := spec.NewReader("Reading spec docs", specRoot)
		if err != nil {
			return err
		}
		docs, err = pipeline.Parallel(cxt, pipelineOptions, docReader, inputs)
		if err != nil {
			return err
		}

		specBuilder := spec.NewBuilder()
		docs, err = pipeline.Collective(cxt, pipelineOptions, &specBuilder, docs)
		if err != nil {
			return err
		}
		if len(docPaths) > 0 {
			filter := files.NewPathFilter[*spec.Doc](docPaths)
			docs, err = pipeline.Collective(cxt, pipelineOptions, filter, docs)
			if err != nil {
				return err
			}
		}
	} else if len(docPaths) > 0 {
		var inputs pipeline.Paths
		inputs, err = pipeline.Start(cxt, files.PathsTargeter(docPaths...))
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
