package asciidoc

type Marked struct {
	position
	raw

	AttributeList
	Elements
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

func (m *Marked) Equals(e Element) bool {
	ob, ok := e.(*Marked)
	if !ok {
		return false
	}
	if !m.AttributeList.Equals(ob.AttributeList) {
		return false
	}
	return m.Elements.Equals(ob.Elements)
}

func (m *Marked) Clone() Element {
	return &Marked{position: m.position, raw: m.raw, AttributeList: m.AttributeList.Clone(), Elements: m.Elements.Clone()}
}

type DoubleMarked struct {
	position
	raw

	AttributeList
	Elements
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

func (dm *DoubleMarked) Equals(e Element) bool {
	ob, ok := e.(*DoubleMarked)
	if !ok {
		return false
	}
	if !dm.AttributeList.Equals(ob.AttributeList) {
		return false
	}
	return dm.Elements.Equals(ob.Elements)
}

func (dm *DoubleMarked) Clone() Element {
	return &DoubleMarked{position: dm.position, raw: dm.raw, AttributeList: dm.AttributeList.Clone(), Elements: dm.Elements.Clone()}
}
