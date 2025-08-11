package parse

import "github.com/project-chip/alchemy/asciidoc"

type SearchShould uint8

const (
	SearchShouldContinue SearchShould = iota
	SearchShouldStop
	SearchShouldSkip
)

type ElementSearchCallback[T any] func(el T, parent asciidoc.Parent, index int) SearchShould

func Iterate[T any](reader asciidoc.Iterator, parent asciidoc.Parent, els asciidoc.Elements, callback ElementSearchCallback[T]) {
	iterate(reader, parent, els, callback)
}

func iterate[T any](reader asciidoc.Iterator, parent asciidoc.Parent, els asciidoc.Elements, callback ElementSearchCallback[T]) SearchShould {
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

		if he, ok := el.(asciidoc.Parent); ok {
			shortCircuit = iterate(reader, he, he.Children(), callback)
		}
		if shortCircuit == SearchShouldStop {
			return shortCircuit
		}
	}
	return SearchShouldContinue
}
