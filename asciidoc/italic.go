package asciidoc

type Italic struct {
	position
	raw

	AttributeList
	Elements
}

func NewItalic(elements Elements) *Italic {
	return &Italic{
		Elements: elements,
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
	return a.Elements.Equals(oa.Elements)
}

type DoubleItalic struct {
	position
	raw

	AttributeList
	Elements
}

func NewDoubleItalic(elements Elements) *DoubleItalic {
	return &DoubleItalic{
		Elements: elements,
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
	return a.Elements.Equals(oa.Elements)
}
