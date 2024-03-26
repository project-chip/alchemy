package disco

import (
	"context"

	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/ascii/render"
	"github.com/hasty/alchemy/internal/files"
	"github.com/hasty/alchemy/internal/pipeline"
)

func Pipeline(cxt context.Context, targeter pipeline.Targeter, pipelineOptions pipeline.Options, discoOptions []Option, filter *files.PathFilter[*ascii.Doc], writer files.Writer[string]) (err error) {

	var inputs pipeline.Map[string, *pipeline.Data[struct{}]]
	inputs, err = pipeline.Start[struct{}](cxt, targeter)

	if err != nil {
		return err
	}

	docReader := ascii.NewReader("Reading docs")
	docs, err := pipeline.Process[struct{}, *ascii.Doc](cxt, pipelineOptions, docReader, inputs)
	if err != nil {
		return err
	}

	baller := NewBaller(discoOptions, pipelineOptions)
	var balledDocs pipeline.Map[string, *pipeline.Data[render.InputDocument]]
	balledDocs, err = pipeline.Process[*ascii.Doc, render.InputDocument](cxt, pipelineOptions, baller, docs)
	if err != nil {
		return err
	}

	renderer := render.NewRenderer()
	var renders pipeline.Map[string, *pipeline.Data[string]]
	renders, err = pipeline.Process[render.InputDocument, string](cxt, pipelineOptions, renderer, balledDocs)
	if err != nil {
		return err
	}

	_, err = pipeline.Process[string, struct{}](cxt, pipelineOptions, writer, renders)
	return
}