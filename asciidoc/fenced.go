package asciidoc

type FencedBlock struct {
	position
	raw

	Delimiter Delimiter
	AttributeList
	Elements
}

func NewFencedBlock(delimiter Delimiter) *FencedBlock {
	return &FencedBlock{Delimiter: delimiter}
}

func (FencedBlock) Type() ElementType {
	return ElementTypeBlock
}

func (a *FencedBlock) Equals(o Element) bool {
	oa, ok := o.(*FencedBlock)
	if !ok {
		return false
	}
	if a.Delimiter != oa.Delimiter {
		return false
	}
	if !a.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return a.Elements.Equals(oa.Elements)
}
