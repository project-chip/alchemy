package asciidoc

type AlchemyEscape struct {
	position
	raw

	Elements
	LineList
}

func NewAlchemyEscape() *AlchemyEscape {
	return &AlchemyEscape{}
}

func (AlchemyEscape) Type() ElementType {
	return ElementTypeBlock
}

func (ae *AlchemyEscape) Equals(e Element) bool {
	omlc, ok := e.(*AlchemyEscape)
	if !ok {
		return false
	}
	return omlc.LineList.Equals(ae.LineList)
}

func (ae *AlchemyEscape) Clone() Element {
	return &AlchemyEscape{position: ae.position, raw: ae.raw, Elements: ae.Elements.Clone(), LineList: ae.LineList.Clone()}
}
