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

func (bi *BlockImage) Equals(o Element) bool {
	oa, ok := o.(*BlockImage)
	if !ok {
		return false
	}
	if !bi.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return bi.ImagePath.Equals(oa.ImagePath)
}

func (bi *BlockImage) Clone() Element {
	return &BlockImage{position: bi.position, raw: bi.raw, AttributeList: bi.AttributeList.Clone(), ImagePath: bi.ImagePath.Clone()}
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

func (ii *InlineImage) Equals(o Element) bool {
	oa, ok := o.(*InlineImage)
	if !ok {
		return false
	}
	if !ii.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return ii.ImagePath.Equals(oa.ImagePath)
}

func (ii *InlineImage) Clone() Element {
	return &InlineImage{AttributeList: ii.AttributeList.Clone(), ImagePath: ii.ImagePath.Clone()}
}
