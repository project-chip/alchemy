package asciidoc

type Writer struct {
	set Set
}

func NewWriter(set Set) *Writer {
	return &Writer{
		set: set,
	}
}

func (r *Writer) Write(el Element) {
	r.set = append(r.set, el)
}

func (r *Writer) WriteSet(el Set) {
	r.set = append(r.set, el...)
}

func (r *Writer) Set() Set {
	return r.set
}
