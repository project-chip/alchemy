package render

import (
	"context"
	"regexp"
	"strings"

	"github.com/project-chip/alchemy/internal/pipeline"
)

type Option func(r *Renderer)

type Renderer struct {
	wordWrapLength int
}

func NewRenderer(options ...Option) *Renderer {
	r := &Renderer{}
	for _, o := range options {
		o(r)
	}
	return r
}

func (p Renderer) Name() string {
	return "Rendering Asciidoc"
}

func (p Renderer) Type() pipeline.ProcessorType {
	return pipeline.ProcessorTypeIndividual
}

func (p Renderer) Process(cxt context.Context, input *pipeline.Data[InputDocument], index int32, total int32) (outputs []*pipeline.Data[string], extra []*pipeline.Data[InputDocument], err error) {
	doc := input.Content
	var renderContext Target
	if p.wordWrapLength > 0 {
		renderContext = NewWrappedTarget(cxt, p.wordWrapLength)
	} else {
		renderContext = NewUnwrappedTarget(cxt)
	}

	err = Elements(renderContext, "", doc.Elements()...)
	if err != nil {
		return
	}
	renderContext.EnsureNewLine()
	output := strings.TrimSpace(renderContext.String())
	output = postProcess(output)
	outputs = append(outputs, &pipeline.Data[string]{Path: input.Path, Content: output})
	return
}

var eolWhitespacePattern = regexp.MustCompile(`(?m)[ ]+\n`)

func postProcess(s string) string {
	return eolWhitespacePattern.ReplaceAllString(s, "\n") + "\n"
}
