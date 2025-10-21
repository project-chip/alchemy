package parse

import (
	"iter"

	"github.com/project-chip/alchemy/asciidoc"
)

func FindAll[T any](doc *asciidoc.Document, reader asciidoc.Reader, parent asciidoc.ParentElement) iter.Seq[T] {
	return func(yield func(T) bool) {
		search(doc, reader, parent, reader.Children(parent), func(doc *asciidoc.Document, el T, parent asciidoc.ParentElement, index int) SearchShould {
			if !yield(el) {
				return SearchShouldStop
			}
			return SearchShouldContinue
		})
	}
}

func FindFirst[T any](doc *asciidoc.Document, reader asciidoc.Reader, parent asciidoc.ParentElement) T {
	var found T
	search(doc, reader, parent, reader.Children(parent), func(doc *asciidoc.Document, el T, parent asciidoc.ParentElement, index int) SearchShould {
		found = el
		return SearchShouldStop
	})
	return found
}

func Skim[T any](reader asciidoc.Reader, parent asciidoc.ParentElement, elements asciidoc.Elements) iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := range reader.Iterate(parent, elements) {
			if el, ok := e.(T); ok {
				if !yield(el) {
					break
				}
			}
		}
	}
}

func SkimList[T any](reader asciidoc.Reader, parent asciidoc.ParentElement, elements asciidoc.Elements) []T {
	var list []T
	for e := range Skim[T](reader, parent, elements) {
		list = append(list, e)
	}
	return list
}

func SkimFunc[T any](reader asciidoc.Reader, parent asciidoc.ParentElement, elements asciidoc.Elements, callback func(t T) bool) bool {
	for e := range reader.Iterate(parent, elements) {
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
