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

func FindAll[T any](elements asciidoc.Set) []T {
	var list []T
	find(elements, func(t T) SearchShould {
		list = append(list, t)
		return SearchShouldContinue
	})
	return list
}

func FindFirst[T any](elements asciidoc.Set) T {
	var found T
	find(elements, func(t T) SearchShould {
		found = t
		return SearchShouldStop
	})
	return found
}

func Search[T any](elements asciidoc.Set, callback func(t T) SearchShould) {
	find(elements, callback)
}

func find[T any](elements asciidoc.Set, callback func(t T) SearchShould) SearchShould {
	for _, e := range elements {
		var shortCircuit SearchShould
		if el, ok := e.(T); ok {
			shortCircuit = callback(el)
		} else if ae, ok := e.(HasBase); ok {
			be := ae.GetBase()
			if el, ok := be.(T); ok {
				shortCircuit = callback(el)
			}
		}
		switch shortCircuit {
		case SearchShouldStop:
			return shortCircuit
		case SearchShouldSkip:
			continue
		case SearchShouldContinue:
		}

		if he, ok := e.(HasElements); ok {
			shortCircuit = find(he.Elements(), callback)
		} else if ae, ok := e.(HasBase); ok {
			be := ae.GetBase()
			if el, ok := be.(HasElements); ok {
				shortCircuit = find(el.Elements(), callback)
			}
		}
		if shortCircuit == SearchShouldStop {
			return shortCircuit
		}
	}
	return SearchShouldContinue
}

func Skim[T any](elements asciidoc.Set) []T {
	var list []T
	for _, e := range elements {
		if ae, ok := e.(HasBase); ok {
			e = ae.GetBase()
		}
		switch el := e.(type) {
		case T:
			list = append(list, el)
		}

	}
	return list
}

func SkimFunc[T any](elements asciidoc.Set, callback func(t T) bool) bool {
	for _, e := range elements {
		var shortCircuit bool
		if el, ok := e.(T); ok {
			shortCircuit = callback(el)
		} else if ae, ok := e.(HasBase); ok {
			be := ae.GetBase()
			if el, ok := be.(T); ok {
				shortCircuit = callback(el)
			}
		}
		if shortCircuit {
			return true
		}
	}
	return false
}

func Filter(parent HasElements, callback func(i any) (remove bool, shortCircuit bool)) (shortCircuit bool) {
	i := 0
	els := parent.Elements()
	var removed bool
	for i < len(els) {
		e := els[i]
		if ae, ok := e.(HasBase); ok {
			e = ae.GetBase()
		}
		switch el := e.(type) {
		case HasElements:
			shortCircuit = Filter(el, callback)
		}
		if shortCircuit {
			break
		}
		remove, shortCircuit := callback(e)
		var empty asciidoc.Set
		if remove {
			els = slices.Replace(els, i, i+1, empty...)
			removed = true
		} else {
			i++
		}
		if shortCircuit {
			break
		}
	}
	if removed {
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
		} else if ae, ok := e.(HasBase); ok {
			be := ae.GetBase()
			if el, ok := be.(T); ok {
				shortCircuit = callback(el, parent, i)
			}
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
		} else if ae, ok := e.(HasBase); ok {
			be := ae.GetBase()
			if el, ok := be.(HasElements); ok {
				shortCircuit = traverse(el, el.Elements(), callback)
			}
		}
		if shortCircuit == SearchShouldStop {
			return shortCircuit
		}

	}
	return SearchShouldContinue
}
