package asciidoc

type Marked struct {
	position
	raw

	AttributeList
	Set
}

func NewMarked() *Marked {
	return &Marked{}
}

func (Marked) Type() ElementType {
	return ElementTypeInline
}

func (Marked) TextFormat() TextFormat {
	return TextFormatMarked
}

func (b *Marked) Equals(e Element) bool {
	ob, ok := e.(*Marked)
	if !ok {
		return false
	}
	if !b.AttributeList.Equals(ob.AttributeList) {
		return false
	}
	return b.Set.Equals(ob.Set)
}

type DoubleMarked struct {
	position
	raw

	AttributeList
	Set
}

func NewDoubleMarked() *DoubleMarked {
	return &DoubleMarked{}
}

func (DoubleMarked) Type() ElementType {
	return ElementTypeInline
}

func (DoubleMarked) TextFormat() TextFormat {
	return TextFormatMarked
}

func (b *DoubleMarked) Equals(e Element) bool {
	ob, ok := e.(*DoubleMarked)
	if !ok {
		return false
	}
	if !b.AttributeList.Equals(ob.AttributeList) {
		return false
	}
	return b.Set.Equals(ob.Set)
}
