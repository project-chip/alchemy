package asciidoc

type FileInclude struct {
	AttributeList
	position
	raw

	Elements
}

func NewFileInclude() *FileInclude {
	return &FileInclude{}
}

func (FileInclude) Type() ElementType {
	return ElementTypeBlock
}

func (fi *FileInclude) Equals(o Element) bool {
	oa, ok := o.(*FileInclude)
	if !ok {
		return false
	}
	if !fi.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return fi.Elements.Equals(oa.Elements)
}

func (fi *FileInclude) Clone() Element {
	return &FileInclude{position: fi.position, raw: fi.raw, AttributeList: fi.AttributeList.Clone(), Elements: fi.Elements.Clone()}
}
