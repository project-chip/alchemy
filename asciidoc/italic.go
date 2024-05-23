package asciidoc

type Italic struct {
	position
	raw

	AttributeList
	Set
}

func NewItalic(elements Set) *Italic {
	return &Italic{
		Set: elements,
	}
}

func (Italic) Type() ElementType {
	return ElementTypeInline
}

func (Italic) TextFormat() TextFormat {
	return TextFormatItalic
}

func (a *Italic) Equals(o Element) bool {
	oa, ok := o.(*Italic)
	if !ok {
		return false
	}
	if !a.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return a.Set.Equals(oa.Set)
}

type DoubleItalic struct {
	position
	raw

	AttributeList
	Set
}

func NewDoubleItalic(elements Set) *DoubleItalic {
	return &DoubleItalic{
		Set: elements,
	}
}

func (DoubleItalic) Type() ElementType {
	return ElementTypeInline
}

func (DoubleItalic) TextFormat() TextFormat {
	return TextFormatItalic
}

func (a *DoubleItalic) Equals(o Element) bool {
	oa, ok := o.(*DoubleItalic)
	if !ok {
		return false
	}
	if !a.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return a.Set.Equals(oa.Set)
}
