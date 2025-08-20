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

func (i *Icon) Equals(o Element) bool {
	oa, ok := o.(*Icon)
	if !ok {
		return false
	}
	if !i.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return i.Path == oa.Path
}

func (i *Icon) Clone() Element {
	return &Icon{position: i.position, raw: i.raw, AttributeList: i.AttributeList.Clone(), Path: i.Path}
}
