package asciidoc

type Listing struct {
	position
	raw

	AttributeList

	Delimiter Delimiter
	LineList
}

func NewListing(delimiter Delimiter) *Listing {
	return &Listing{Delimiter: delimiter}
}

func (Listing) Type() ElementType {
	return ElementTypeBlock
}

func (l *Listing) Equals(o Element) bool {
	oa, ok := o.(*Listing)
	if !ok {
		return false
	}
	if !l.Delimiter.Equals(oa.Delimiter) {
		return false
	}
	if !l.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return l.LineList.Equals(oa.LineList)
}

func (l *Listing) Clone() Element {
	return &Listing{
		position:      l.position,
		raw:           l.raw,
		AttributeList: l.AttributeList.Clone(),
		Delimiter:     l.Delimiter,
		LineList:      l.LineList.Clone(),
	}
}
