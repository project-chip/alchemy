package asciidoc

type Writer struct {
	elements Elements
}

func NewWriter(set Elements) *Writer {
	return &Writer{
		elements: set,
	}
}

func (r *Writer) Write(el Element) {
	r.elements = append(r.elements, el)
}

func (r *Writer) WriteSet(el Elements) {
	r.elements = append(r.elements, el...)
}

func (r *Writer) Set() Elements {
	return r.elements
}
