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

func (b *Monospace) Equals(e Element) bool {
	ob, ok := e.(*Monospace)
	if !ok {
		return false
	}
	if !b.AttributeList.Equals(ob.AttributeList) {
		return false
	}
	return b.Elements.Equals(ob.Elements)
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

func (b *DoubleMonospace) Equals(e Element) bool {
	ob, ok := e.(*DoubleMonospace)
	if !ok {
		return false
	}
	if !b.AttributeList.Equals(ob.AttributeList) {
		return false
	}
	return b.Elements.Equals(ob.Elements)
}
