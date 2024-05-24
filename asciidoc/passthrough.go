package asciidoc

type InlinePassthrough struct {
	position
	raw

	AttributeList
	Set
}

func NewInlinePassthrough() *InlinePassthrough {
	return &InlinePassthrough{}
}

func (InlinePassthrough) Type() ElementType {
	return ElementTypeInline
}

func (a *InlinePassthrough) Equals(o Element) bool {
	oa, ok := o.(*InlinePassthrough)
	if !ok {
		return false
	}
	if !a.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return a.Set.Equals(oa.Set)
}

type InlineDoublePassthrough struct {
	position
	raw

	AttributeList
	Set
}

func NewInlineDoublePassthrough() *InlineDoublePassthrough {
	return &InlineDoublePassthrough{}
}

func (InlineDoublePassthrough) Type() ElementType {
	return ElementTypeInline
}

func (a *InlineDoublePassthrough) Equals(o Element) bool {
	oa, ok := o.(*InlineDoublePassthrough)
	if !ok {
		return false
	}
	if !a.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return a.Set.Equals(oa.Set)
}
