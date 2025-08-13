package spec

import (
	"context"
	"fmt"
	"os"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/pipeline"
)

func ParseFile(path asciidoc.Path, specRoot string, attributes ...asciidoc.AttributeName) (*Doc, error) {

	contents, err := os.Open(path.Absolute)
	if err != nil {
		return nil, err
	}
	defer contents.Close()

	d, err := ReadFile(path.Absolute, specRoot)
	if err != nil {
		return nil, fmt.Errorf("parse error in \"%s\": %w", path, err)
	}

	d.parsed = true
	d.reader, err = preparse(nil, d, specRoot, attributes)
	if err != nil {
		return nil, err
	}
	return d, nil
}

type Parser struct {
	options ParserOptions
}

func NewParser(parserOptions ParserOptions) (Parser, error) {
	return Parser{options: parserOptions}, nil
}

func (p Parser) Name() string {
	return "Parsing documents"
}

func (p Parser) Targets(cxt context.Context) ([]string, error) {
	return getSpecPaths(p.options.Root)
}

func (p Parser) Process(cxt context.Context, input *pipeline.Data[struct{}], index int32, total int32) (outputs []*pipeline.Data[*Doc], extras []*pipeline.Data[struct{}], err error) {

	var path asciidoc.Path
	path, err = NewSpecPath(input.Path, p.options.Root)
	if err != nil {
		return
	}
	var doc *Doc
	doc, err = ReadFile(path.Absolute, p.options.Root)
	if err != nil {
		return
	}
	outputs = append(outputs, &pipeline.Data[*Doc]{Path: input.Path, Content: doc})
	return
}
