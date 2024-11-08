package asciidoc

type AlchemyEscape struct {
	position
	raw

	Set
	LineList
}

func NewAlchemyEscape() *AlchemyEscape {
	return &AlchemyEscape{}
}

func (AlchemyEscape) Type() ElementType {
	return ElementTypeBlock
}

func (mlc *AlchemyEscape) Equals(e Element) bool {
	omlc, ok := e.(*AlchemyEscape)
	if !ok {
		return false
	}
	return omlc.LineList.Equals(mlc.LineList)
}
