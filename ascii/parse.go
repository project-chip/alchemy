package ascii

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/hasty/adoc/elements"
	"github.com/hasty/adoc/parse"
	"github.com/hasty/alchemy/internal/pipeline"
	"github.com/hasty/alchemy/internal/text"
)

func ParseFile(path string, attributes ...elements.AttributeName) (*Doc, error) {

	contents, err := readFile(path)
	if err != nil {
		return nil, err
	}
	return Parse(contents, path, attributes...)
}

func Parse(contents string, path string, attributes ...elements.AttributeName) (doc *Doc, err error) {
	/*baseConfig := make([]configuration.Setting, 0, len(settings)+1)
	baseConfig = append(baseConfig, configuration.WithFilename(path))
	baseConfig = append(baseConfig, settings...)

	config := configuration.NewConfiguration(baseConfig...)
	config.IgnoreIncludes = true



	contents, err = parser.Preprocess(strings.NewReader(contents), config)
	if err != nil {
		return nil, err
	}*/

	contents = text.RemoveComments(contents)

	var d *elements.Document

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

func ParseDocument(r io.Reader, path string, attributes ...elements.AttributeName) (*elements.Document, error) {
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
		return &elements.Document{}, nil
	}

	return parse.Reader(path, strings.NewReader(parsed))

	/*newContext := func() *parser.ParseContext {
		c := parser.NewParseContext(config, opts...)
		//c.IgnoreColumnDefs(true)
		c.SuppressAttributeSubstitution(true)
		return c
	}

	footnotes := elements.NewFootnotes()
	doc, err := parser.Aggregate(newContext(),
		parser.CollectFootnotes(footnotes, done,
			parser.ApplySubstitutions(newContext(), done,
				parser.RefineFragments(newContext(), r, done,
					parser.ParseDocumentFragments(newContext(), r, done),
				),
			),
		),
	)
	if err != nil {
		return nil, err
	}
	if len(footnotes.Notes) > 0 {
		doc.Footnotes = footnotes.Notes
	}*/
	//	return doc, nil
}

type Parser struct {
	attributes []elements.AttributeName
}

func NewParser(attributes []elements.AttributeName) Parser {
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
