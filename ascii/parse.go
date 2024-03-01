package ascii

import (
	"fmt"
	"io"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/internal/text"
)

func ParseFile(path string, settings ...configuration.Setting) (*Doc, error) {

	contents, err := readFile(path)
	if err != nil {
		return nil, err
	}
	return Parse(contents, path, settings...)
}

func Parse(contents string, path string, settings ...configuration.Setting) (doc *Doc, err error) {
	baseConfig := make([]configuration.Setting, 0, len(settings)+1)
	baseConfig = append(baseConfig, configuration.WithFilename(path))
	baseConfig = append(baseConfig, settings...)

	config := configuration.NewConfiguration(baseConfig...)
	config.IgnoreIncludes = true

	contents = text.RemoveComments(contents)

	contents, err = parser.Preprocess(strings.NewReader(contents), config)
	if err != nil {
		return nil, err
	}

	var d *types.Document

	d, err = ParseDocument(strings.NewReader(contents), config, parser.MaxExpressions(2000000))

	if err != nil {
		return nil, fmt.Errorf("failed parse: %w", err)
	}

	doc, err = NewDoc(d)
	if err != nil {
		return nil, err
	}
	doc.Path = path

	PatchUnrecognizedReferences(doc)

	return doc, nil
}

func ParseDocument(r io.Reader, config *configuration.Configuration, opts ...parser.Option) (*types.Document, error) {
	done := make(chan interface{})
	defer close(done)

	newContext := func() *parser.ParseContext {
		c := parser.NewParseContext(config, opts...)
		//c.IgnoreColumnDefs(true)
		c.SuppressAttributeSubstitution(true)
		return c
	}

	footnotes := types.NewFootnotes()
	doc, err := parser.Aggregate(newContext(),
		// SplitHeader(done,
		//parser.ArrangeLists(done,
		parser.CollectFootnotes(footnotes, done,
			parser.ApplySubstitutions(newContext(), done, // needs to be before 'ArrangeLists'
				parser.RefineFragments(newContext(), r, done,
					parser.ParseDocumentFragments(newContext(), r, done),
				),
			),
		),
		//),

		// ),
	)
	if err != nil {
		return nil, err
	}
	if len(footnotes.Notes) > 0 {
		doc.Footnotes = footnotes.Notes
	}
	return doc, nil
}
