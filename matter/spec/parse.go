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

func InlineParse(path asciidoc.Path, specRoot string, attributes ...asciidoc.AttributeName) (*Doc, error) {
	ac := &preparseContext{
		docPath:  path,
		rootPath: specRoot,
	}

	for _, a := range attributes {
		ac.Set(string(a), nil)
	}

	contents, err := os.Open(path.Absolute)
	if err != nil {
		return nil, err
	}
	defer contents.Close()

	d, err := parse.Inline(ac, path.Relative, contents)
	if err != nil {
		return nil, fmt.Errorf("parse error in %s: %w", path, err)
	}
	doc, err := newDoc(d, path)
	if err != nil {
		return nil, err
	}
	doc.parsed = true
	return doc, nil
}

func ParseFile(path asciidoc.Path, specRoot string, attributes ...asciidoc.AttributeName) (*Doc, error) {

	contents, err := os.Open(path.Absolute)
	if err != nil {
		return nil, err
	}
	defer contents.Close()
	return parseReader(contents, path, specRoot, attributes...)
}

func parseReader(r io.Reader, path asciidoc.Path, specRoot string, attributes ...asciidoc.AttributeName) (doc *Doc, err error) {
	var d *asciidoc.Document

	d, err = parseDocument(r, path, specRoot, attributes...)

	if err != nil {
		return nil, fmt.Errorf("parse error in %s: %w", path, err)
	}

	doc, err = newDoc(d, path)
	if err != nil {
		return nil, err
	}
	doc.parsed = true
	return doc, nil
}

func parseDocument(r io.Reader, path asciidoc.Path, specRoot string, attributes ...asciidoc.AttributeName) (*asciidoc.Document, error) {
	ac := &preparseContext{
		docPath:  path,
		rootPath: specRoot,
	}

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

type Parser struct {
	Root       string
	attributes []asciidoc.AttributeName

	inline bool
}

func NewParser(attributes []asciidoc.AttributeName, parserOptions ...ParserOption) (Parser, error) {
	p := Parser{attributes: attributes}
	for _, opt := range parserOptions {
		opt(&p)
	}
	if !filepath.IsAbs(p.Root) {
		var err error
		p.Root, err = filepath.Abs(p.Root)
		if err != nil {
			return Parser{}, err
		}
	}
	return p, nil
}

func (p Parser) Name() string {
	return "Parsing documents"
}

func (p Parser) Targets(cxt context.Context) ([]string, error) {
	return getSpecPaths(p.Root)
}

func (p Parser) Process(cxt context.Context, input *pipeline.Data[struct{}], index int32, total int32) (outputs []*pipeline.Data[*Doc], extras []*pipeline.Data[struct{}], err error) {

	var path asciidoc.Path
	path, err = NewSpecPath(input.Path, p.Root)
	if err != nil {
		return
	}
	var doc *Doc
	if p.inline {
		doc, err = InlineParse(path, p.Root, p.attributes...)
	} else {
		doc, err = ParseFile(path, p.Root, p.attributes...)
	}
	if err != nil {
		return
	}
	outputs = append(outputs, &pipeline.Data[*Doc]{Path: input.Path, Content: doc})
	return
}
