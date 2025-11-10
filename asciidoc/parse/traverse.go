package parse

import (
	"github.com/project-chip/alchemy/asciidoc"
)

type ElementTraverseCallback[T any] func(doc *asciidoc.Document, el T, parent asciidoc.ParentElement, index int) SearchShould

func Traverse[T any](doc *asciidoc.Document, reader asciidoc.Reader, parent asciidoc.ParentElement, els asciidoc.Elements, callback ElementSearchCallback[T]) {
	traverse(doc, reader, parent, els, callback)
}

func traverse[T any](doc *asciidoc.Document, reader asciidoc.Reader, parent asciidoc.ParentElement, els asciidoc.Elements, callback ElementSearchCallback[T]) SearchShould {
	var i int
	for el := range reader.Iterate(parent, els) {
		doc := doc
		if de, ok := el.(asciidoc.DocumentElement); ok {
			doc = de.Document()
		}
		var shortCircuit SearchShould
		if et, ok := el.(T); ok {
			switch el := el.(type) {
			case asciidoc.Traverser:
				for parent, els := range el.Traverse(parent) {
					shortCircuit = traverse(parent.Document(), reader, parent, els.Children(), callback)
					if shortCircuit == SearchShouldStop {
						return shortCircuit
					}
				}
			}
			shortCircuit = callback(doc, et, parent, i)
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
			shortCircuit = traverse(el.Document(), reader, el, reader.Children(el), callback)
			if shortCircuit == SearchShouldStop {
				return shortCircuit
			}
		case asciidoc.Parent:
			shortCircuit = traverse(doc, reader, parent, el.Children(), callback)
			if shortCircuit == SearchShouldStop {
				return shortCircuit
			}
		}
	}
	return SearchShouldContinue
}
