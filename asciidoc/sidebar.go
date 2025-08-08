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

func (a *SidebarBlock) Equals(o Element) bool {
	oa, ok := o.(*SidebarBlock)
	if !ok {
		return false
	}
	if !a.Delimiter.Equals(oa.Delimiter) {
		return false
	}
	if !a.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return a.Elements.Equals(oa.Elements)
}
