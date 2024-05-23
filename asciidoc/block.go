package asciidoc

type BlockAttributes struct {
	AttributeList
}

func NewBlockAttributes(a AttributeList) *BlockAttributes {
	return &BlockAttributes{AttributeList: a}
}

func (BlockAttributes) Type() ElementType {
	return ElementTypeAttributes
}

func (uar *BlockAttributes) Equals(e Element) bool {
	ouar, ok := e.(*BlockAttributes)
	if !ok {
		return false
	}
	return uar.AttributeList.Equals(ouar.AttributeList)
}
