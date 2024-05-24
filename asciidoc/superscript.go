package asciidoc

type Superscript struct {
	position
	raw

	AttributeList
	Set
}

func NewSuperscript() *Superscript {
	return &Superscript{}
}

func (Superscript) Type() ElementType {
	return ElementTypeInline
}

func (a *Superscript) Equals(o Element) bool {
	oa, ok := o.(*Superscript)
	if !ok {
		return false
	}
	if !a.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return a.Set.Equals(oa.Set)
}
