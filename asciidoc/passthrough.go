package asciidoc

type InlinePassthrough struct {
	position
	raw

	AttributeList
	Elements
}

func NewInlinePassthrough() *InlinePassthrough {
	return &InlinePassthrough{}
}

func (InlinePassthrough) Type() ElementType {
	return ElementTypeInline
}

func (ip *InlinePassthrough) Equals(o Element) bool {
	oa, ok := o.(*InlinePassthrough)
	if !ok {
		return false
	}
	if !ip.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return ip.Elements.Equals(oa.Elements)
}

func (ip *InlinePassthrough) Clone() Element {
	return &InlinePassthrough{position: ip.position, raw: ip.raw, AttributeList: ip.AttributeList.Clone(), Elements: ip.Elements.Clone()}
}

type InlineDoublePassthrough struct {
	position
	raw

	AttributeList
	Elements
}

func NewInlineDoublePassthrough() *InlineDoublePassthrough {
	return &InlineDoublePassthrough{}
}

func (InlineDoublePassthrough) Type() ElementType {
	return ElementTypeInline
}

func (idp *InlineDoublePassthrough) Equals(o Element) bool {
	oa, ok := o.(*InlineDoublePassthrough)
	if !ok {
		return false
	}
	if !idp.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return idp.Elements.Equals(oa.Elements)
}

func (idp *InlineDoublePassthrough) Clone() Element {
	return &InlineDoublePassthrough{position: idp.position, raw: idp.raw, AttributeList: idp.AttributeList.Clone(), Elements: idp.Elements.Clone()}
}
