package asciidoc

type SourceBlock struct {
	position
	raw

	Delimiter Delimiter
	AttributeList
	Set
}

func NewSourceBlock(delimiter Delimiter) *SourceBlock {
	return &SourceBlock{Delimiter: delimiter}
}

func (SourceBlock) Type() ElementType {
	return ElementTypeBlock
}

func (a *SourceBlock) Equals(o Element) bool {
	oa, ok := o.(*SourceBlock)
	if !ok {
		return false
	}
	if a.Delimiter != oa.Delimiter {
		return false
	}
	if !a.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return a.Set.Equals(oa.Set)
}
