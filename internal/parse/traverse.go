package parse

import (
	"iter"
	"slices"

	"github.com/project-chip/alchemy/asciidoc"
)

type SearchShould uint8

const (
	SearchShouldContinue SearchShould = iota
	SearchShouldStop
	SearchShouldSkip
)

func FindAll[T any](parent asciidoc.HasElements) iter.Seq[T] {
	return func(yield func(T) bool) {
		traverse(parent, parent.Elements(), func(el T, parent HasElements, index int) SearchShould {
			if !yield(el) {
				return SearchShouldStop
			}
			return SearchShouldContinue
		})
	}
}

func FindFirst[T any](parent asciidoc.HasElements) T {
	var found T
	traverse(parent, parent.Elements(), func(el T, parent HasElements, index int) SearchShould {
		found = el
		return SearchShouldStop
	})
	return found
}

func Skim[T any](elements asciidoc.Set) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, e := range elements {
			if el, ok := e.(T); ok {
				if !yield(el) {
					break
				}
			}
		}
	}
}

func SkimList[T any](elements asciidoc.Set) []T {
	var list []T
	for e := range Skim[T](elements) {
		list = append(list, e)
	}
	return list
}

func SkimFunc[T any](elements asciidoc.Set, callback func(t T) bool) bool {
	for _, e := range elements {
		var shortCircuit bool
		if el, ok := e.(T); ok {
			shortCircuit = callback(el)
		}
		if shortCircuit {
			return true
		}
	}
	return false
}

func Filter(parent HasElements, callback func(parent HasElements, el asciidoc.Element) (remove bool, replace asciidoc.Set, shortCircuit bool)) (shortCircuit bool) {
	i := 0
	els := parent.Elements()
	var altered bool
	for i < len(els) {
		e := els[i]
		switch el := e.(type) {
		case HasElements:
			shortCircuit = Filter(el, callback)
		}
		if shortCircuit {
			break
		}
		var remove bool
		var replace asciidoc.Set
		remove, replace, shortCircuit = callback(parent, e)
		var empty asciidoc.Set
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
		parent.SetElements(els)
	}
	return
}

type TraverseCallback[T any] func(el T, parent HasElements, index int) SearchShould

func Traverse[T any](parent HasElements, els asciidoc.Set, callback TraverseCallback[T]) {
	traverse(parent, els, callback)
}

func traverse[T any](parent HasElements, els asciidoc.Set, callback TraverseCallback[T]) SearchShould {

	for i, e := range els {
		var shortCircuit SearchShould
		if el, ok := e.(T); ok {
			shortCircuit = callback(el, parent, i)
		}
		switch shortCircuit {
		case SearchShouldStop:
			return shortCircuit
		case SearchShouldSkip:
			continue
		case SearchShouldContinue:
		}

		if he, ok := e.(HasElements); ok {
			shortCircuit = traverse(he, he.Elements(), callback)
		}
		if shortCircuit == SearchShouldStop {
			return shortCircuit
		}

	}
	return SearchShouldContinue
}
