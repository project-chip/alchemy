package main

import (
	"fmt"
	"io"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

type doc struct {
	base *types.Document
	root *section

	preamble []interface{}

	orderedListDepth   int
	unorderedListDepth int
}

func (d *doc) parseSectionElements(parent *section, s *types.Section) {
	for _, e := range s.Elements {
		switch el := e.(type) {
		case *types.Section:
			d.addSection(parent, el)
		default:
			parent.elements = append(parent.elements, &element{base: e})
			//fmt.Printf("parsing unknown section element type: %T\n", e)

		}
	}
}

func (d *doc) addSection(parent *section, s *types.Section) {
	ss := &section{base: s}

	for _, te := range s.GetTitle() {
		switch tel := te.(type) {
		case *types.StringElement:
			fmt.Printf("section title string: %s\n", tel.Content)
			ss.name = tel.Content

		default:
			fmt.Printf("unknown section title element type: %T\n", te)
			ss.elements = append(ss.elements, te)
		}
	}
	switch s.Level {
	case 1:
		fmt.Printf("Adding top level section %s...\n", ss.name)
		ss.secType = docSectionTypePreface
	case 2:

	}
	parent.elements = append(parent.elements, ss)
	d.parseSectionElements(ss, s)
}

func getSectionTitle(s *types.Section) string {
	for _, te := range s.GetTitle() {
		switch tel := te.(type) {
		case *types.StringElement:
			return tel.Content
		}
	}
	return ""
}

// ParseDocument parses the content of the reader identitied by the filename and applies all the substitutions and arrangements
func parseDocument(r io.Reader, config *configuration.Configuration, opts ...parser.Option) (*types.Document, error) {
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
