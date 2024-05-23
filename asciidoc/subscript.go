package asciidoc

type Subscript struct {
	position
	raw

	AttributeList
	Set
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
	return a.Set.Equals(oa.Set)
}
