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

	var docs pipeline.Map[string, *pipeline.Data[*spec.Doc]]

	if specRoot == "" {
		allPaths, e := files.PathsTargeter(docPaths...)(cxt)
		if e == nil {
			specRoot = spec.DeriveSpecPathFromPaths(allPaths)

		}
	}

	errata.LoadErrataConfig(specRoot)

	if specRoot != "" {

		specTargeter := spec.Targeter(specRoot)

		var inputs pipeline.Map[string, *pipeline.Data[struct{}]]
		inputs, err = pipeline.Start[struct{}](cxt, specTargeter)
		if err != nil {
			return err
		}

		docReader := spec.NewReader("Reading spec docs")
		docs, err = pipeline.Process[struct{}, *spec.Doc](cxt, pipelineOptions, docReader, inputs)
		if err != nil {
			return err
		}

		specBuilder := spec.NewBuilder()
		docs, err = pipeline.Process[*spec.Doc, *spec.Doc](cxt, pipelineOptions, &specBuilder, docs)
		if err != nil {
			return err
		}
		if len(docPaths) > 0 {
			filter := files.NewPathFilter[*spec.Doc](docPaths)
			docs, err = pipeline.Process[*spec.Doc, *spec.Doc](cxt, pipelineOptions, filter, docs)
			if err != nil {
				return err
			}
		}
	} else if len(docPaths) > 0 {
		var inputs pipeline.Map[string, *pipeline.Data[struct{}]]
		inputs, err = pipeline.Start[struct{}](cxt, files.PathsTargeter(docPaths...))
		if err != nil {
			return err
		}

		docReader := spec.NewReader("Reading docs")
		docs, err = pipeline.Process[struct{}, *spec.Doc](cxt, pipelineOptions, docReader, inputs)
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

	baller := NewBaller(discoOptions, pipelineOptions)

	var balledDocs pipeline.Map[string, *pipeline.Data[*spec.Doc]]
	balledDocs, err = pipeline.Process[*spec.Doc, *spec.Doc](cxt, pipelineOptions, baller, docs)
	if err != nil {
		return err
	}

	anchorNormalizer := newAnchorNormalizer(discoOptions)
	var normalizedDocs pipeline.Map[string, *pipeline.Data[render.InputDocument]]
	normalizedDocs, err = pipeline.Process[*spec.Doc, render.InputDocument](cxt, pipelineOptions, anchorNormalizer, balledDocs)
	if err != nil {
		return err
	}

	renderer := render.NewRenderer(renderOptions...)
	var renders pipeline.Map[string, *pipeline.Data[string]]
	renders, err = pipeline.Process[render.InputDocument, string](cxt, pipelineOptions, renderer, normalizedDocs)
	if err != nil {
		return err
	}

	_, err = pipeline.Process[string, struct{}](cxt, pipelineOptions, writer, renders)
	return
}
