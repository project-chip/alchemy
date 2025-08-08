package asciidoc

type Reader struct {
	index int
	set   Elements
}

func NewReader(set Elements) *Reader {
	return &Reader{
		set: set,
	}
}

func (r *Reader) Read() Element {
	if r.index >= len(r.set) {
		return nil
	}
	e := r.set[r.index]
	r.index++
	return e
}
