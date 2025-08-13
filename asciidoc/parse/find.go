package parse

import (
	"iter"

	"github.com/project-chip/alchemy/asciidoc"
)

func FindAll[T any](iterator asciidoc.Iterator, parent asciidoc.ParentElement) iter.Seq[T] {
	return func(yield func(T) bool) {
		search(iterator, parent, parent.Children(), func(el T, parent asciidoc.Parent, index int) SearchShould {
			if !yield(el) {
				return SearchShouldStop
			}
			return SearchShouldContinue
		})
	}
}

func FindFirst[T any](iterator asciidoc.Iterator, parent asciidoc.ParentElement) T {
	var found T
	traverse(iterator, parent, parent.Children(), func(el T, parent asciidoc.Parent, index int) SearchShould {
		found = el
		return SearchShouldStop
	})
	return found
}

func Skim[T any](iterator asciidoc.Iterator, parent asciidoc.ParentElement, elements asciidoc.Elements) iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := range iterator.Iterate(parent, elements) {
			if el, ok := e.(T); ok {
				if !yield(el) {
					break
				}
			}
		}
	}
}

func SkimList[T any](iterator asciidoc.Iterator, parent asciidoc.ParentElement, elements asciidoc.Elements) []T {
	var list []T
	for e := range Skim[T](iterator, parent, elements) {
		list = append(list, e)
	}
	return list
}

func SkimFunc[T any](iterator asciidoc.Iterator, parent asciidoc.ParentElement, elements asciidoc.Elements, callback func(t T) bool) bool {
	for e := range iterator.Iterate(parent, elements) {
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
