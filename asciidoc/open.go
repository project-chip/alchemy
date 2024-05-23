package asciidoc

type OpenBlock struct {
	position
	raw

	AttributeList

	Delimiter Delimiter
	Set
}

func NewOpenBlock(delimiter Delimiter) *OpenBlock {
	return &OpenBlock{Delimiter: delimiter}
}

func (OpenBlock) Type() ElementType {
	return ElementTypeBlock
}

func (a *OpenBlock) Equals(o Element) bool {
	oa, ok := o.(*OpenBlock)
	if !ok {
		return false
	}
	if !a.Delimiter.Equals(oa.Delimiter) {
		return false
	}
	if !a.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return a.Set.Equals(oa.Set)
}
