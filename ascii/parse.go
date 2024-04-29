package ascii

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/hasty/adoc/elements"
	"github.com/hasty/adoc/parse"
	"github.com/hasty/alchemy/internal/pipeline"
)

func ParseFile(path string, attributes ...elements.Attribute) (*Doc, error) {

	contents, err := readFile(path)
	if err != nil {
		return nil, err
	}
	return Parse(contents, path, attributes...)
}

func Parse(contents string, path string, attributes ...elements.Attribute) (doc *Doc, err error) {
	/*baseConfig := make([]configuration.Setting, 0, len(settings)+1)
	baseConfig = append(baseConfig, configuration.WithFilename(path))
	baseConfig = append(baseConfig, settings...)

	config := configuration.NewConfiguration(baseConfig...)
	config.IgnoreIncludes = true

	contents = text.RemoveComments(contents)

	contents, err = parser.Preprocess(strings.NewReader(contents), config)
	if err != nil {
		return nil, err
	}*/

	var d *elements.Document

	d, err = ParseDocument(strings.NewReader(contents), path, attributes...)

	if err != nil {
		return nil, fmt.Errorf("failed parse: %w", err)
	}

	doc, err = NewDoc(d)
	if err != nil {
		return nil, err
	}
	doc.Path = path

	return doc, nil
}

func ParseDocument(r io.Reader, path string, attributes ...elements.Attribute) (*elements.Document, error) {
	done := make(chan any)
	defer close(done)

	return parse.Reader(path, r)

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
	attributes []elements.Attribute
}

func NewParser(attributes []elements.Attribute) Parser {
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
	doc, err = ParseFile(input.Path)
	if err != nil {
		return
	}
	outputs = append(outputs, &pipeline.Data[*Doc]{Path: input.Path, Content: doc})
	return
}
