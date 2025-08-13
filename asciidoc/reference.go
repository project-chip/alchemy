package asciidoc

import "iter"

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

	ID     Elements
	Format CrossReferenceFormat
}

func NewCrossReference(id Elements, format CrossReferenceFormat) *CrossReference {
	return &CrossReference{ID: id, Format: format}
}

func (CrossReference) Type() ElementType {
	return ElementTypeInline
}

func (cr *CrossReference) Equals(o Element) bool {
	oa, ok := o.(*CrossReference)
	if !ok {
		return false
	}
	if !cr.ID.Equals(oa.ID) {
		return false
	}
	return cr.Elements.Equals(oa.Elements)
}

func (cr *CrossReference) Traverse(parent Parent) iter.Seq2[Parent, Parent] {
	return func(yield func(Parent, Parent) bool) {
		if !cr.AttributeList.traverse(cr, yield) {
			return
		}
		if !yield(cr, &cr.ID) {
			return
		}
	}
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

func (dcr *DocumentCrossReference) Equals(o Element) bool {
	oa, ok := o.(*DocumentCrossReference)
	if !ok {
		return false
	}
	if !dcr.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return dcr.ReferencePath.Equals(oa.ReferencePath)
}
