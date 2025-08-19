package asciidoc

import "iter"

type ElementIterator iter.Seq[Element]

func (ei ElementIterator) List() (elements Elements) {
	ei(func(e Element) bool {
		elements = append(elements, e)
		return true
	})
	return
}

func (ei ElementIterator) Count() (count int) {
	ei(func(e Element) bool {
		count++
		return true
	})
	return
}

type Reader interface {
	Iterate(parent Parent, elements Elements) ElementIterator
}

type rawReader struct {
}

func (rr *rawReader) Iterate(parent Parent, elements Elements) ElementIterator {
	return func(yield func(Element) bool) {
		for _, e := range elements {
			if !yield(e) {
				return
			}
		}
	}
}

func (rr *rawReader) Count(elements Elements) int {
	return len(elements)
}

var RawReader Reader = &rawReader{}

type Traverser interface {
	Traverse(parent Parent) iter.Seq2[Parent, Parent]
}
