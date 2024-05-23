package asciidoc

type Document struct {
	Set
}

func (Document) Type() ElementType {
	return ElementTypeDocument
}

func (a *Document) Equals(o Element) bool {
	oa, ok := o.(*Document)
	if !ok {
		return false
	}
	return a.Set.Equals(oa.Set)
}
