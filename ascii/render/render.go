package render

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/hasty/alchemy/internal/pipeline"
)

func Render(cxt context.Context, doc InputDocument) (string, error) {
	renderContext := NewContext(cxt, doc)
	err := RenderElements(renderContext, "", doc.GetElements())
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
	pipelineOptions pipeline.Options
}

func NewRenderer(pipelineOptions pipeline.Options) *Renderer {
	return &Renderer{pipelineOptions: pipelineOptions}
}

func (p Renderer) Name() string {
	return "Rendering Asciidoc"
}

func (p Renderer) Type() pipeline.ProcessorType {
	return p.pipelineOptions.DefaultProcessorType()
}

func (p Renderer) Process(cxt context.Context, input *pipeline.Data[InputDocument], index int32, total int32) (outputs []*pipeline.Data[string], extra []*pipeline.Data[InputDocument], err error) {
	outputs, err = p.render(cxt, input)
	return
}

func (p Renderer) ProcessAll(cxt context.Context, inputs []*pipeline.Data[InputDocument]) (outputs []*pipeline.Data[string], err error) {
	for _, input := range inputs {
		var o []*pipeline.Data[string]
		fmt.Fprintf(os.Stderr, "%s...\n", input.Path)
		o, err = p.render(cxt, input)
		if err != nil {
			return
		}
		outputs = append(outputs, o...)
	}
	return
}

func (p Renderer) render(cxt context.Context, input *pipeline.Data[InputDocument]) (outputs []*pipeline.Data[string], err error) {
	doc := input.Content
	var out string
	out, err = Render(cxt, doc)
	if err != nil {
		return
	}
	outputs = append(outputs, &pipeline.Data[string]{Path: input.Path, Content: out})
	return
}
