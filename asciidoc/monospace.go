package asciidoc

type Monospace struct {
	position
	raw

	AttributeList
	Elements
}

func NewMonospace(elements Elements) *Monospace {
	return &Monospace{
		Elements: elements,
	}
}

func (Monospace) Type() ElementType {
	return ElementTypeInline
}

func (Monospace) TextFormat() TextFormat {
	return TextFormatMonospace
}

func (ms *Monospace) Equals(e Element) bool {
	ob, ok := e.(*Monospace)
	if !ok {
		return false
	}
	if !ms.AttributeList.Equals(ob.AttributeList) {
		return false
	}
	return ms.Elements.Equals(ob.Elements)
}

func (ms *Monospace) Clone() Element {
	return &Monospace{position: ms.position, raw: ms.raw, AttributeList: ms.AttributeList.Clone(), Elements: ms.Elements.Clone()}
}

type DoubleMonospace struct {
	position
	raw

	AttributeList
	Elements
}

func NewDoubleMonospace(elements Elements) *DoubleMonospace {
	return &DoubleMonospace{
		Elements: elements,
	}
}

func (DoubleMonospace) Type() ElementType {
	return ElementTypeInline
}

func (DoubleMonospace) TextFormat() TextFormat {
	return TextFormatMonospace
}

func (dm *DoubleMonospace) Equals(e Element) bool {
	ob, ok := e.(*DoubleMonospace)
	if !ok {
		return false
	}
	if !dm.AttributeList.Equals(ob.AttributeList) {
		return false
	}
	return dm.Elements.Equals(ob.Elements)
}

func (dm *DoubleMonospace) Clone() Element {
	return &DoubleMonospace{position: dm.position, raw: dm.raw, AttributeList: dm.AttributeList.Clone(), Elements: dm.Elements.Clone()}
}
