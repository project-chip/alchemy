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

func (ba *BlockAttributes) Equals(e Element) bool {
	ouar, ok := e.(*BlockAttributes)
	if !ok {
		return false
	}
	return ba.AttributeList.Equals(ouar.AttributeList)
}

func (ba *BlockAttributes) Clone() Element {
	return &BlockAttributes{AttributeList: ba.AttributeList.Clone()}

}
