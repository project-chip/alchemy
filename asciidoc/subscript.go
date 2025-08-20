package asciidoc

type Subscript struct {
	position
	raw

	AttributeList
	Elements
}

func NewSubscript() *Subscript {
	return &Subscript{}
}

func (Subscript) Type() ElementType {
	return ElementTypeInline
}

func (a *Subscript) Equals(o Element) bool {
	oa, ok := o.(*Subscript)
	if !ok {
		return false
	}
	if !a.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return a.Elements.Equals(oa.Elements)
}

func (a *Subscript) Clone() Element {
	return &Subscript{position: a.position, raw: a.raw, AttributeList: a.AttributeList.Clone(), Elements: a.Elements.Clone()}
}
