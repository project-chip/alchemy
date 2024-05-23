package asciidoc

type Reader struct {
	index int
	set   Set
}

func NewReader(set Set) *Reader {
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
