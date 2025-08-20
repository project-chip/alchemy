package asciidoc

type ThematicBreak struct {
	position
	raw

	AttributeList
}

func NewThematicBreak() *ThematicBreak {
	return &ThematicBreak{}
}

func (ThematicBreak) Type() ElementType {
	return ElementTypeBlock
}

func (tb *ThematicBreak) Equals(e Element) bool {
	otb, ok := e.(*ThematicBreak)
	if !ok {
		return false
	}
	return tb.AttributeList.Equals(otb.AttributeList)
}

func (tb *ThematicBreak) Clone() Element {
	return &ThematicBreak{position: tb.position, raw: tb.raw, AttributeList: tb.AttributeList.Clone()}
}

type PageBreak struct {
	position
	raw

	AttributeList
}

func NewPageBreak() *PageBreak {
	return &PageBreak{}
}

func (PageBreak) Type() ElementType {
	return ElementTypeBlock
}

func (pb *PageBreak) Equals(e Element) bool {
	opb, ok := e.(*PageBreak)
	if !ok {
		return false
	}
	return pb.AttributeList.Equals(opb.AttributeList)
}

func (pb *PageBreak) Clone() Element {
	return &PageBreak{position: pb.position, raw: pb.raw, AttributeList: pb.AttributeList.Clone()}
}
