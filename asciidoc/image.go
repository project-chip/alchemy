package asciidoc

type BlockImage struct {
	position
	raw

	AttributeList

	ImagePath Elements
}

func NewBlockImage(path Elements) *BlockImage {
	return &BlockImage{ImagePath: path}
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
	return a.ImagePath.Equals(oa.ImagePath)
}

type InlineImage struct {
	AttributeList

	ImagePath Elements
}

func NewInlineImage(path Elements) *InlineImage {
	return &InlineImage{ImagePath: path}
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
	return a.ImagePath.Equals(oa.ImagePath)
}
