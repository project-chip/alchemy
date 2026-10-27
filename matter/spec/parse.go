package spec

import (
	"context"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/pipeline"
)

type Reader struct {
	options ParserOptions
}

func NewReader(parserOptions ParserOptions) (Reader, error) {
	return Reader{options: parserOptions}, nil
}

func (p Reader) Name() string {
	return "Parsing documents"
}

func (p Reader) Targets(cxt context.Context) ([]string, error) {
	return getSpecPaths(p.options.Root)
}

func (p Reader) Process(cxt context.Context, input *pipeline.Data[struct{}], index int32, total int32) (outputs []*pipeline.Data[*asciidoc.Document], extras []*pipeline.Data[struct{}], err error) {

	var path asciidoc.Path
	path, err = NewSpecPath(input.Path, p.options.Root)
	if err != nil {
		return
	}
	var doc *asciidoc.Document
	doc, err = readFile(path.Absolute, p.options.Root)
	if err != nil {
		return
	}
	outputs = append(outputs, &pipeline.Data[*asciidoc.Document]{Path: input.Path, Content: doc})
	return
}
