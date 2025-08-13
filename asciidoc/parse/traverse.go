package parse

import (
	"github.com/project-chip/alchemy/asciidoc"
)

func Traverse[T any](reader asciidoc.Reader, parent asciidoc.Parent, els asciidoc.Elements, callback ElementSearchCallback[T]) {
	traverse(reader, parent, els, callback)
}

func traverse[T any](reader asciidoc.Reader, parent asciidoc.Parent, els asciidoc.Elements, callback ElementSearchCallback[T]) SearchShould {
	var i int
	for el := range reader.Iterate(parent, els) {

		var shortCircuit SearchShould
		if et, ok := el.(T); ok {
			switch el := el.(type) {
			case asciidoc.Traverser:
				for parent, els := range el.Traverse(parent) {
					shortCircuit = traverse(reader, parent, els.Children(), callback)
					if shortCircuit == SearchShouldStop {
						return shortCircuit
					}
				}
			}
			shortCircuit = callback(et, parent, i)
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
		case asciidoc.Parent:
			shortCircuit = traverse(reader, el, el.Children(), callback)
			if shortCircuit == SearchShouldStop {
				return shortCircuit
			}
		}
	}
	return SearchShouldContinue
}
