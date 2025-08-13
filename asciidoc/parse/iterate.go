package parse

import (
	"slices"

	"github.com/project-chip/alchemy/asciidoc"
)

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

		if he, ok := el.(asciidoc.Parent); ok {
			shortCircuit = iterate(reader, he, he.Children(), callback)
		}
		if shortCircuit == SearchShouldStop {
			return shortCircuit
		}
	}
	return SearchShouldContinue
}

func Filter(parent asciidoc.Parent, callback func(parent asciidoc.Parent, el asciidoc.Element) (remove bool, replace asciidoc.Elements, shortCircuit bool)) (shortCircuit bool) {
	i := 0
	els := parent.Children()
	var altered bool
	for i < len(els) {
		e := els[i]
		switch el := e.(type) {
		case asciidoc.Parent:
			shortCircuit = Filter(el, callback)
		}
		if shortCircuit {
			break
		}
		var remove bool
		var replace asciidoc.Elements
		remove, replace, shortCircuit = callback(parent, e)
		var empty asciidoc.Elements
		if len(replace) > 0 {
			els = slices.Delete(els, i, i+1)
			els = slices.Insert(els, i, replace...)
			altered = true
		} else if remove {
			els = slices.Replace(els, i, i+1, empty...)
			altered = true
		} else {
			i++
		}
		if shortCircuit {
			break
		}
	}
	if altered {
		parent.SetChildren(els)
	}
	return
}
