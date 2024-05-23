package asciidoc

type FileInclude struct {
	AttributeList
	position
	raw

	Set
}

func NewFileInclude() *FileInclude {
	return &FileInclude{}
}

func (FileInclude) Type() ElementType {
	return ElementTypeBlock
}

func (a *FileInclude) Equals(o Element) bool {
	oa, ok := o.(*FileInclude)
	if !ok {
		return false
	}
	if !a.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return a.Set.Equals(oa.Set)
}
