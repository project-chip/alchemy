package asciidoc

type CrossReferenceFormat uint8

const (
	CrossReferenceFormatNatural CrossReferenceFormat = iota
	CrossReferenceFormatMacro
)

type CrossReference struct {
	position
	raw

	AttributeList
	Elements

	ID     string
	Format CrossReferenceFormat
}

func NewCrossReference(id string, format CrossReferenceFormat) *CrossReference {
	return &CrossReference{ID: id, Format: format}
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
	return a.Elements.Equals(oa.Elements)
}

func NewCrossReferenceMacro(path Elements) AttributableElement {
	if len(path) == 1 {
		if s, ok := path[0].(*String); ok {
			return NewCrossReference(s.Value, CrossReferenceFormatMacro)
		}
	}
	return NewDocumentCrossReference(path)
}

type DocumentCrossReference struct {
	position
	raw

	AttributeList

	ReferencePath Elements
}

func NewDocumentCrossReference(path Elements) *DocumentCrossReference {
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
