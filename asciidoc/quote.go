package asciidoc

type QuoteBlock struct {
	position
	raw

	Delimiter Delimiter
	AttributeList

	Elements
}

func NewQuoteBlock(delimiter Delimiter) *QuoteBlock {
	return &QuoteBlock{Delimiter: delimiter}
}

func (QuoteBlock) Type() ElementType {
	return ElementTypeBlock
}

func (a *QuoteBlock) Equals(o Element) bool {
	oa, ok := o.(*QuoteBlock)
	if !ok {
		return false
	}
	if !a.Delimiter.Equals(oa.Delimiter) {
		return false
	}
	if !a.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return a.Elements.Equals(oa.Elements)
}
