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

func (a *Link) Equals(o Element) bool {
	oa, ok := o.(*Link)
	if !ok {
		return false
	}
	if !a.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return a.URL.Equals(oa.URL)
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

func (a *LinkMacro) Equals(o Element) bool {
	oa, ok := o.(*LinkMacro)
	if !ok {
		return false
	}
	if !a.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return a.URL.Equals(oa.URL)
}
