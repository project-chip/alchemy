package asciidoc

type Writer struct {
	set Elements
}

func NewWriter(set Elements) *Writer {
	return &Writer{
		set: set,
	}
}

func (r *Writer) Write(el Element) {
	r.set = append(r.set, el)
}

func (r *Writer) WriteSet(el Elements) {
	r.set = append(r.set, el...)
}

func (r *Writer) Set() Elements {
	return r.set
}
