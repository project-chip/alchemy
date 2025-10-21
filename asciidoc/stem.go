package asciidoc

type StemBlock struct {
	position
	raw

	AttributeList

	Delimiter Delimiter
	LineList
}

func NewStemBlock(delimiter Delimiter) *StemBlock {
	return &StemBlock{Delimiter: delimiter}
}

func (StemBlock) Type() ElementType {
	return ElementTypeBlock
}

func (a *StemBlock) Equals(o Element) bool {
	oa, ok := o.(*StemBlock)
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

func (a *StemBlock) Clone() Element {
	return &StemBlock{position: a.position, raw: a.raw, AttributeList: a.AttributeList.Clone(), Delimiter: a.Delimiter, LineList: a.LineList.Clone()}
}
