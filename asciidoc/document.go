package asciidoc

type Document struct {
	Elements
}

func (Document) Type() ElementType {
	return ElementTypeDocument
}

func (a *Document) Equals(o Element) bool {
	oa, ok := o.(*Document)
	if !ok {
		return false
	}
	return a.Elements.Equals(oa.Elements)
}
