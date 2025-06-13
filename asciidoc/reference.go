package asciidoc

type CrossReference struct {
	position
	raw

	Set

	ID string
}

func NewCrossReference(id string) *CrossReference {
	return &CrossReference{ID: id}
}

func (CrossReference) Type() ElementType {
	return ElementTypeInline
}

func (a *CrossReference) Equals(o Element) bool {
	oa, ok := o.(*CrossReference)
	if !ok {
		return false
	}
	if a.ID != oa.ID {
		return false
	}
	return a.Set.Equals(oa.Set)
}

type DocumentCrossReference struct {
	position
	raw

	AttributeList

	ReferencePath Set
}

func NewDocumentCrossReference(path Set) *DocumentCrossReference {
	return &DocumentCrossReference{ReferencePath: path}
}

func (DocumentCrossReference) Type() ElementType {
	return ElementTypeInline
}

func (a *DocumentCrossReference) Equals(o Element) bool {
	oa, ok := o.(*DocumentCrossReference)
	if !ok {
		return false
	}
	if !a.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return a.ReferencePath.Equals(oa.ReferencePath)
}
