package parse

import (
	"io"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

// ParseDocument parses the content of the reader identitied by the filename and applies all the substitutions and arrangements
func ParseDocument(r io.Reader, config *configuration.Configuration, opts ...parser.Option) (*types.Document, error) {
	done := make(chan interface{})
	defer close(done)

	footnotes := types.NewFootnotes()
	doc, err := parser.Aggregate(parser.NewParseContext(config, opts...),
		// SplitHeader(done,
		parser.ArrangeLists(done,
			parser.CollectFootnotes(footnotes, done,
				parser.ApplySubstitutions(parser.NewParseContext(config, opts...), done, // needs to be before 'ArrangeLists'
					parser.RefineFragments(parser.NewParseContext(config, opts...), r, done,
						parser.ParseDocumentFragments(parser.NewParseContext(config, opts...), r, done),
					),
				),
			),
		),

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
