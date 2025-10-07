package asciidoc

type SidebarBlock struct {
	position
	raw

	Delimiter Delimiter
	AttributeList

	Elements
}

func NewSidebarBlock(delimiter Delimiter) *SidebarBlock {
	return &SidebarBlock{Delimiter: delimiter}
}

func (SidebarBlock) Type() ElementType {
	return ElementTypeBlock
}

func (sbb *SidebarBlock) Equals(o Element) bool {
	oa, ok := o.(*SidebarBlock)
	if !ok {
		return false
	}
	if !sbb.Delimiter.Equals(oa.Delimiter) {
		return false
	}
	if !sbb.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return sbb.Elements.Equals(oa.Elements)
}

func (sbb *SidebarBlock) Clone() Element {
	return &SidebarBlock{position: sbb.position, raw: sbb.raw, Delimiter: sbb.Delimiter, AttributeList: sbb.AttributeList.Clone(), Elements: sbb.Elements.Clone()}
}
