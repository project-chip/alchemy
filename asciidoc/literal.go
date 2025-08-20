package asciidoc

type LiteralBlock struct {
	position
	raw

	AttributeList

	Delimiter Delimiter
	LineList
}

func NewLiteralBlock(delimiter Delimiter) *LiteralBlock {
	return &LiteralBlock{Delimiter: delimiter}
}

func (LiteralBlock) Type() ElementType {
	return ElementTypeBlock
}

func (a *LiteralBlock) Equals(o Element) bool {
	oa, ok := o.(*LiteralBlock)
	if !ok {
		return false
	}
	if !a.Delimiter.Equals(oa.Delimiter) {
		return false
	}
	if !a.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return a.LineList.Equals(oa.LineList)
}

func (a *LiteralBlock) Clone() Element {
	return &LiteralBlock{position: a.position, raw: a.raw, AttributeList: a.AttributeList.Clone(), Delimiter: a.Delimiter, LineList: a.LineList.Clone()}
}
