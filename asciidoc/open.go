package asciidoc

type OpenBlock struct {
	position
	raw

	AttributeList

	Delimiter Delimiter
	Elements
}

func NewOpenBlock(delimiter Delimiter) *OpenBlock {
	return &OpenBlock{Delimiter: delimiter}
}

func (OpenBlock) Type() ElementType {
	return ElementTypeBlock
}

func (ob *OpenBlock) Equals(o Element) bool {
	oa, ok := o.(*OpenBlock)
	if !ok {
		return false
	}
	if !ob.Delimiter.Equals(oa.Delimiter) {
		return false
	}
	if !ob.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return ob.Elements.Equals(oa.Elements)
}

func (ob *OpenBlock) Clone() Element {
	return &OpenBlock{position: ob.position, raw: ob.raw, AttributeList: ob.AttributeList.Clone(), Delimiter: ob.Delimiter, Elements: ob.Elements.Clone()}
}
