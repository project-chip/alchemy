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

type RawReader struct {
}

func NewRawReader() *RawReader {
	return rawReader
}

func (rr *RawReader) Iterate(parent Parent, elements Elements) ElementIterator {
	return func(yield func(Element) bool) {
		for _, e := range elements {
			if !yield(e) {
				return
			}
		}
	}
}

func (rr *RawReader) Count(elements Elements) int {
	return len(elements)
}

var rawReader *RawReader = &RawReader{}

type TraversalPosition uint8

const (
	TraversalPositionUnknown TraversalPosition = iota
	TraversalPositionTitle
	TraversalPositionValue
	TraversalPositionChild
)

type Traverser interface {
	Traverse(parent Parent) iter.Seq2[Parent, Parent]
}
