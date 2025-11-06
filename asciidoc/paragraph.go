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

func (at AdmonitionType) Clone() Element {
	return at
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

func (a *Admonition) Equals(e Element) bool {
	oat, ok := e.(*Admonition)
	if !ok {
		return false
	}
	if a.AdmonitionType != oat.AdmonitionType {
		return false
	}
	return a.AttributeList.Equals(oat.AttributeList)
}

func (a *Admonition) Clone() Element {
	return &Admonition{
		AdmonitionType: a.AdmonitionType,
		AttributeList:  a.AttributeList.Clone(),
	}
}

type Paragraph struct {
	position

	AttributeList
	Elements

	Admonition AdmonitionType
}

func NewParagraph() *Paragraph {
	return &Paragraph{}
}

func (Paragraph) Type() ElementType {
	return ElementTypeBlock
}

func (p *Paragraph) Equals(o Element) bool {
	oa, ok := o.(*Paragraph)
	if !ok {
		return false
	}
	if p.Admonition != oa.Admonition {
		return false
	}
	if !p.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return p.Elements.Equals(oa.Elements)
}

func (p *Paragraph) Clone() Element {
	return &Paragraph{
		position:      p.position,
		AttributeList: p.AttributeList.Clone(),
		Elements:      p.Elements.Clone(),
		Admonition:    p.Admonition,
	}
}
