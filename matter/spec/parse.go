package spec

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/internal/pipeline"
)

func ParseFile(path Path, attributes ...asciidoc.AttributeName) (*Doc, error) {

	contents, err := os.Open(path.Absolute)
	if err != nil {
		return nil, err
	}
	defer contents.Close()
	return parseReader(contents, path, attributes...)
}

func parseReader(r io.Reader, path Path, attributes ...asciidoc.AttributeName) (doc *Doc, err error) {
	var d *asciidoc.Document

	d, err = parseDocument(r, path, attributes...)

	if err != nil {
		return nil, fmt.Errorf("parse error in %s: %w", path, err)
	}

	doc, err = newDoc(d, path)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func parseDocument(r io.Reader, path Path, attributes ...asciidoc.AttributeName) (*asciidoc.Document, error) {
	ac := &parse.AttributeContext{}

	for _, a := range attributes {
		ac.Set(string(a), nil)
	}

	parsed, err := parse.PreParseReader(ac, path.Relative, r)
	if err != nil {
		return nil, err
	}

	if len(parsed) == 0 {
		return &asciidoc.Document{}, nil
	}

	if filepath.Base(path.Absolute) == "DoorLock.adoc" { // Craptastic workaround for very weird table cell
		var doorLockPattern = regexp.MustCompile(`\n+\s*[^&\n]+&#8224;\s+`)
		parsed = doorLockPattern.ReplaceAllString(parsed, "\n")
	}

	return parse.String(path.Relative, parsed)
}

type parseOption func(p *Parser)

type Parser struct {
	rootPath   string
	attributes []asciidoc.AttributeName
}

func NewParser(rootPath string, attributes []asciidoc.AttributeName) (Parser, error) {
	if !filepath.IsAbs(rootPath) {
		var err error
		rootPath, err = filepath.Abs(rootPath)
		if err != nil {
			return Parser{}, err
		}
	}
	return Parser{rootPath: rootPath, attributes: attributes}, nil
}

func (p Parser) Name() string {
	return "Parsing documents"
}

func (p Parser) Type() pipeline.ProcessorType {
	return pipeline.ProcessorTypeIndividual
}

func (p Parser) Process(cxt context.Context, input *pipeline.Data[struct{}], index int32, total int32) (outputs []*pipeline.Data[*Doc], extras []*pipeline.Data[struct{}], err error) {

	var path Path
	path, err = NewSpecPath(input.Path, p.rootPath)
	if err != nil {
		return
	}
	var doc *Doc
	doc, err = ParseFile(path, p.attributes...)
	if err != nil {
		return
	}
	outputs = append(outputs, &pipeline.Data[*Doc]{Path: input.Path, Content: doc})
	return
}
