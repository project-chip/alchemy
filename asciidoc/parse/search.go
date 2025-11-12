package parse

import (
	"github.com/project-chip/alchemy/asciidoc"
)

type SearchShould uint8

const (
	SearchShouldContinue SearchShould = iota
	SearchShouldStop
	SearchShouldSkip
)

type ElementSearchCallback[T any] func(doc *asciidoc.Document, el T, parent asciidoc.ParentElement, index int) SearchShould

func Search[T any](doc *asciidoc.Document, reader asciidoc.Reader, parent asciidoc.ParentElement, els asciidoc.Elements, callback ElementSearchCallback[T]) {
	search(doc, reader, parent, els, callback)
}

func search[T any](doc *asciidoc.Document, reader asciidoc.Reader, parent asciidoc.ParentElement, els asciidoc.Elements, callback ElementSearchCallback[T]) SearchShould {
	var i int
	for el := range reader.Iterate(parent, els) {
		doc := doc
		if de, ok := el.(asciidoc.DocumentElement); ok {
			doc = de.Document()
		}
		p := parent
		if ce, ok := el.(asciidoc.ChildElement); ok {
			if pe, ok := reader.Parent(ce).(asciidoc.ParentElement); ok {
				p = pe
			}
		}
		var shortCircuit SearchShould
		if el, ok := el.(T); ok {
			shortCircuit = callback(doc, el, p, i)
		}
		i++
		switch shortCircuit {
		case SearchShouldStop:
			return shortCircuit
		case SearchShouldSkip:
			continue
		case SearchShouldContinue:
		}
		switch el := el.(type) {
		case asciidoc.ParentElement:
			shortCircuit = search(el.Document(), reader, el, reader.Children(el), callback)
			if shortCircuit == SearchShouldStop {
				return shortCircuit
			}
		case asciidoc.Parent:
			shortCircuit = search(doc, reader, p, el.Children(), callback)
			if shortCircuit == SearchShouldStop {
				return shortCircuit
			}
		}
	}
	return SearchShouldContinue
}
