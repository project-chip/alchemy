package asciidoc

import "iter"

type Iterator interface {
	Iterate(elements Elements) iter.Seq[Element]
}

type RawIterator struct {
}

func NewRawIterator() *RawIterator {
	return &RawIterator{}
}

func (rr *RawIterator) Iterate(elements Elements) iter.Seq[Element] {
	return func(yield func(Element) bool) {
		for _, e := range elements {
			if !yield(e) {
				return
			}
		}
	}
}

var _ Iterator = &RawIterator{}

type Traverser interface {
	Parent
	Traverse() iter.Seq[Parent]
}
