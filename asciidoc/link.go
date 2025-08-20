package asciidoc

type Link struct {
	position
	raw

	AttributeList

	URL URL
}

func NewLink(url URL) *Link {
	return &Link{URL: url}
}

func (Link) Type() ElementType {
	return ElementTypeInline
}

func (l *Link) Equals(o Element) bool {
	oa, ok := o.(*Link)
	if !ok {
		return false
	}
	if !l.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return l.URL.Equals(oa.URL)
}

func (l *Link) Clone() Element {
	return &Link{position: l.position, raw: l.raw, AttributeList: l.AttributeList.Clone(), URL: l.URL.Clone().(URL)}
}

type LinkMacro struct {
	position
	raw

	AttributeList

	URL URL
}

func NewLinkMacro(url URL) *LinkMacro {
	return &LinkMacro{URL: url}
}

func (LinkMacro) Type() ElementType {
	return ElementTypeInline
}

func (lm *LinkMacro) Equals(o Element) bool {
	oa, ok := o.(*LinkMacro)
	if !ok {
		return false
	}
	if !lm.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return lm.URL.Equals(oa.URL)
}

func (lm *LinkMacro) Clone() Element {
	return &LinkMacro{position: lm.position, raw: lm.raw, AttributeList: lm.AttributeList.Clone(), URL: lm.URL.Clone().(URL)}
}
