package ascii

import (
	"io"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

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
