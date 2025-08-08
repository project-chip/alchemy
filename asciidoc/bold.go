package asciidoc

type Bold struct {
	position
	raw

	AttributeList
	Elements
}

func NewBold() *Bold {
	return &Bold{}
}

func (Bold) Type() ElementType {
	return ElementTypeInline
}

func (Bold) TextFormat() TextFormat {
	return TextFormatBold
}

func (b *Bold) Equals(e Element) bool {
	ob, ok := e.(*Bold)
	if !ok {
		return false
	}
	if !b.AttributeList.Equals(ob.AttributeList) {
		return false
	}
	return b.Elements.Equals(ob.Elements)
}

type DoubleBold struct {
	position
	raw

	AttributeList
	Elements
}

func NewDoubleBold() *DoubleBold {
	return &DoubleBold{}
}

func (DoubleBold) Type() ElementType {
	return ElementTypeInline
}

func (DoubleBold) TextFormat() TextFormat {
	return TextFormatBold
}

func (b *DoubleBold) Equals(e Element) bool {
	ob, ok := e.(*DoubleBold)
	if !ok {
		return false
	}
	if !b.AttributeList.Equals(ob.AttributeList) {
		return false
	}
	return b.Elements.Equals(ob.Elements)
}
