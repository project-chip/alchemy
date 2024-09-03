package spec

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/internal/pipeline"
)

func ParseFile(path string, attributes ...asciidoc.AttributeName) (*Doc, error) {

	contents, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer contents.Close()
	return parseReader(contents, path, attributes...)
}

func Parse(contents string, path string, attributes ...asciidoc.AttributeName) (doc *Doc, err error) {
	return parseReader(strings.NewReader(contents), path, attributes...)

}

func parseReader(r io.Reader, path string, attributes ...asciidoc.AttributeName) (doc *Doc, err error) {
	var d *asciidoc.Document

	d, err = ParseDocument(r, path, attributes...)

	if err != nil {
		return nil, fmt.Errorf("parse error in %s: %w", path, err)
	}

	var p Path
	p, err = NewPath(path)
	if err != nil {
		return nil, err
	}

	doc, err = NewDoc(d, p)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func ParseDocument(r io.Reader, path string, attributes ...asciidoc.AttributeName) (*asciidoc.Document, error) {
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

	if filepath.Base(path) == "DoorLock.adoc" { // Craptastic workaround for very weird table cell
		var doorLockPattern = regexp.MustCompile(`\n+\s*[^&\n]+&#8224;\s+`)
		parsed = doorLockPattern.ReplaceAllString(parsed, "\n")
	}

	return parse.String(path, parsed)
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
