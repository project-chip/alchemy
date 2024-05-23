package asciidoc

type Bold struct {
	position
	raw

	AttributeList
	Set
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
	return b.Set.Equals(ob.Set)
}

type DoubleBold struct {
	position
	raw

	AttributeList
	Set
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
	return b.Set.Equals(ob.Set)
}
