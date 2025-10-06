package asciidoc

type Document struct {
	Path Path
	Elements
}

func (Document) Type() ElementType {
	return ElementTypeDocument
}

func (doc *Document) Clone() Element {
	return &Document{Path: doc.Path, Elements: doc.Elements.Clone()}
}

func (doc *Document) Equals(o Element) bool {
	oa, ok := o.(*Document)
	if !ok {
		return false
	}
	return doc.Elements.Equals(oa.Elements)
}

func (doc *Document) Document() *Document {
	return doc
}
