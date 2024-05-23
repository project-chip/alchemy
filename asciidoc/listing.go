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

func (a *Listing) Equals(o Element) bool {
	oa, ok := o.(*Listing)
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
