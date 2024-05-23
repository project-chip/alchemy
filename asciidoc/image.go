package asciidoc

type BlockImage struct {
	position
	raw

	AttributeList

	Path Set
}

func NewBlockImage(path Set) *BlockImage {
	return &BlockImage{Path: path}
}

func (BlockImage) Type() ElementType {
	return ElementTypeBlock
}

func (a *BlockImage) Equals(o Element) bool {
	oa, ok := o.(*BlockImage)
	if !ok {
		return false
	}
	if !a.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return a.Path.Equals(oa.Path)
}

type InlineImage struct {
	AttributeList

	Path Set
}

func NewInlineImage(path Set) *InlineImage {
	return &InlineImage{Path: path}
}

func (InlineImage) Type() ElementType {
	return ElementTypeInline
}

func (a *InlineImage) Equals(o Element) bool {
	oa, ok := o.(*InlineImage)
	if !ok {
		return false
	}
	if !a.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return a.Path.Equals(oa.Path)
}
