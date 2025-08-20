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

func (fb *FencedBlock) Equals(o Element) bool {
	oa, ok := o.(*FencedBlock)
	if !ok {
		return false
	}
	if fb.Delimiter != oa.Delimiter {
		return false
	}
	if !fb.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return fb.Elements.Equals(oa.Elements)
}

func (fb *FencedBlock) Clone() Element {
	return &FencedBlock{position: fb.position, raw: fb.raw, Delimiter: fb.Delimiter, AttributeList: fb.AttributeList.Clone(), Elements: fb.Elements.Clone()}
}
