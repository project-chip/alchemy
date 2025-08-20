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

func (qb *QuoteBlock) Equals(o Element) bool {
	oa, ok := o.(*QuoteBlock)
	if !ok {
		return false
	}
	if !qb.Delimiter.Equals(oa.Delimiter) {
		return false
	}
	if !qb.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return qb.Elements.Equals(oa.Elements)
}

func (qb *QuoteBlock) Clone() Element {
	return &QuoteBlock{
		position:      qb.position,
		raw:           qb.raw,
		Delimiter:     qb.Delimiter,
		AttributeList: qb.AttributeList.Clone(),
		Elements:      qb.Elements.Clone(),
	}
}
