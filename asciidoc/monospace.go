package asciidoc

type Monospace struct {
	position
	raw

	AttributeList
	Set
}

func NewMonospace(elements Set) *Monospace {
	return &Monospace{
		Set: elements,
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
	return b.Set.Equals(ob.Set)
}

type DoubleMonospace struct {
	position
	raw

	AttributeList
	Set
}

func NewDoubleMonospace(elements Set) *DoubleMonospace {
	return &DoubleMonospace{
		Set: elements,
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
	return b.Set.Equals(ob.Set)
}
