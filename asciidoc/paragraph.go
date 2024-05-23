package asciidoc

type AdmonitionType uint8

const (
	AdmonitionTypeNone AdmonitionType = iota
	AdmonitionTypeNote
	AdmonitionTypeTip
	AdmonitionTypeImportant
	AdmonitionTypeCaution
	AdmonitionTypeWarning
)

func (AdmonitionType) Type() ElementType {
	return ElementTypeAttribute
}

func (at AdmonitionType) Equals(e Element) bool {
	oat, ok := e.(AdmonitionType)
	if !ok {
		return false
	}
	return at == oat
}

type Admonition struct {
	AdmonitionType AdmonitionType
	AttributeList
}

func NewAdmonition(admonition AdmonitionType) *Admonition {
	return &Admonition{
		AdmonitionType: admonition,
	}
}

func (Admonition) Type() ElementType {
	return ElementTypeBlock
}

func (at *Admonition) Equals(e Element) bool {
	oat, ok := e.(*Admonition)
	if !ok {
		return false
	}
	if at.AdmonitionType != oat.AdmonitionType {
		return false
	}
	return at.AttributeList.Equals(oat.AttributeList)
}

type Paragraph struct {
	position

	AttributeList
	Set

	Admonition AdmonitionType
}

func NewParagraph() *Paragraph {
	return &Paragraph{}
}

func (Paragraph) Type() ElementType {
	return ElementTypeBlock
}

func (a *Paragraph) Equals(o Element) bool {
	oa, ok := o.(*Paragraph)
	if !ok {
		return false
	}
	if a.Admonition != oa.Admonition {
		return false
	}
	if !a.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return a.Set.Equals(oa.Set)
}
