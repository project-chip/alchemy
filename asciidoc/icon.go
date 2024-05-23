package asciidoc

type Icon struct {
	position
	raw

	AttributeList

	Path string
}

func NewIcon(path string) *Icon {
	return &Icon{Path: path}
}

func (Icon) Type() ElementType {
	return ElementTypeInline
}

func (a *Icon) Equals(o Element) bool {
	oa, ok := o.(*Icon)
	if !ok {
		return false
	}
	if !a.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return a.Path == oa.Path
}
