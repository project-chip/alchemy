package parse

import "github.com/project-chip/alchemy/asciidoc"

func Traverse[T any](reader asciidoc.Iterator, parent asciidoc.Parent, els asciidoc.Elements, callback ElementSearchCallback[T]) {
	traverse(reader, parent, els, callback)
}

func traverse[T any](reader asciidoc.Iterator, parent asciidoc.Parent, els asciidoc.Elements, callback ElementSearchCallback[T]) SearchShould {
	var i int
	for el := range reader.Iterate(els) {
		var shortCircuit SearchShould
		if el, ok := el.(T); ok {

			shortCircuit = callback(el, parent, i)

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
		case asciidoc.Traverser:
			for els := range el.Traverse() {
				shortCircuit = traverse(reader, el, els.Children(), callback)
				if shortCircuit == SearchShouldStop {
					return shortCircuit
				}
			}
		case asciidoc.Parent:
			shortCircuit = traverse(reader, el, el.Children(), callback)
			if shortCircuit == SearchShouldStop {
				return shortCircuit
			}
		}
	}
	return SearchShouldContinue
}
