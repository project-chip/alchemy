package parse

import "github.com/project-chip/alchemy/asciidoc"

type SearchShould uint8

const (
	SearchShouldContinue SearchShould = iota
	SearchShouldStop
	SearchShouldSkip
)

type ElementSearchCallback[T any] func(el T, parent asciidoc.Parent, index int) SearchShould

func Search[T any](reader asciidoc.Reader, parent asciidoc.Parent, els asciidoc.Elements, callback ElementSearchCallback[T]) {
	search(reader, parent, els, callback)
}

func search[T any](reader asciidoc.Reader, parent asciidoc.Parent, els asciidoc.Elements, callback ElementSearchCallback[T]) SearchShould {
	var i int
	for el := range reader.Iterate(parent, els) {
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
		case asciidoc.Parent:
			shortCircuit = search(reader, el, el.Children(), callback)
			if shortCircuit == SearchShouldStop {
				return shortCircuit
			}
		}
	}
	return SearchShouldContinue
}
