package spec

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/hasty/alchemy/asciidoc"
	"github.com/hasty/alchemy/asciidoc/parse"
	"github.com/hasty/alchemy/internal/pipeline"
	"github.com/hasty/alchemy/internal/text"
)

func ParseFile(path string, attributes ...asciidoc.AttributeName) (*Doc, error) {

	contents, err := readFile(path)
	if err != nil {
		return nil, err
	}
	return Parse(contents, path, attributes...)
}

func Parse(contents string, path string, attributes ...asciidoc.AttributeName) (doc *Doc, err error) {

	contents = text.RemoveComments(contents)

	var d *asciidoc.Document

	d, err = ParseDocument(strings.NewReader(contents), path, attributes...)

	if err != nil {
		return nil, fmt.Errorf("parse error in %s: %w", path, err)
	}

	doc, err = NewDoc(d)
	if err != nil {
		return nil, err
	}
	doc.Path = path

	return doc, nil
}

func ParseDocument(r io.Reader, path string, attributes ...asciidoc.AttributeName) (*asciidoc.Document, error) {
	done := make(chan any)
	defer close(done)

	if len(attributes) == 0 {
		return parse.Reader(path, r)
	}

	ac := &parse.AttributeContext{}

	for _, a := range attributes {
		ac.Set(string(a), nil)
	}

	parsed, err := parse.PreParseReader(ac, path, r)
	if err != nil {
		return nil, err
	}

	if len(parsed) == 0 {
		return &asciidoc.Document{}, nil
	}

	return parse.Reader(path, strings.NewReader(parsed))
}

type Parser struct {
	attributes []asciidoc.AttributeName
}

func NewParser(attributes []asciidoc.AttributeName) Parser {
	return Parser{attributes: attributes}
}

func (p Parser) Name() string {
	return "Parsing documents"
}

func (p Parser) Type() pipeline.ProcessorType {
	return pipeline.ProcessorTypeIndividual
}

func (p Parser) Process(cxt context.Context, input *pipeline.Data[struct{}], index int32, total int32) (outputs []*pipeline.Data[*Doc], extras []*pipeline.Data[struct{}], err error) {
	var doc *Doc
	doc, err = ParseFile(input.Path, p.attributes...)
	if err != nil {
		return
	}
	outputs = append(outputs, &pipeline.Data[*Doc]{Path: input.Path, Content: doc})
	return
}
