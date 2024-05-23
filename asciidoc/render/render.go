package render

import (
	"context"
	"regexp"
	"strings"

	"github.com/hasty/alchemy/internal/pipeline"
)

func Render(cxt context.Context, doc InputDocument) (string, error) {
	renderContext := NewContext(cxt, doc)
	err := Elements(renderContext, "", doc.Elements()...)
	if err != nil {
		return "", err
	}
	renderContext.WriteNewline()
	output := strings.TrimSpace(renderContext.String())
	return postProcess(output), nil
}

var eolWhitespacePattern = regexp.MustCompile(`(?m)[ ]+\n`)

func postProcess(s string) string {
	return eolWhitespacePattern.ReplaceAllString(s, "\n") + "\n"
}

type Renderer struct {
}

func NewRenderer() *Renderer {
	return &Renderer{}
}

func (p Renderer) Name() string {
	return "Rendering Asciidoc"
}

func (p Renderer) Type() pipeline.ProcessorType {
	return pipeline.ProcessorTypeIndividual
}

func (p Renderer) Process(cxt context.Context, input *pipeline.Data[InputDocument], index int32, total int32) (outputs []*pipeline.Data[string], extra []*pipeline.Data[InputDocument], err error) {
	doc := input.Content
	var out string
	out, err = Render(cxt, doc)
	if err != nil {
		return
	}
	outputs = append(outputs, &pipeline.Data[string]{Path: input.Path, Content: out})
	return
}
