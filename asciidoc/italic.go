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

func (i *Italic) Equals(o Element) bool {
	oa, ok := o.(*Italic)
	if !ok {
		return false
	}
	if !i.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return i.Elements.Equals(oa.Elements)
}

func (i *Italic) Clone() Element {
	return &Italic{position: i.position, raw: i.raw, AttributeList: i.AttributeList.Clone(), Elements: i.Elements.Clone()}
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

func (di *DoubleItalic) Equals(o Element) bool {
	oa, ok := o.(*DoubleItalic)
	if !ok {
		return false
	}
	if !di.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return di.Elements.Equals(oa.Elements)
}

func (di *DoubleItalic) Clone() Element {
	return &DoubleItalic{position: di.position, raw: di.raw, AttributeList: di.AttributeList.Clone(), Elements: di.Elements.Clone()}
}
