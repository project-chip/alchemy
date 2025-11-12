package asciidoc

import "iter"

type Bold struct {
	position
	raw

	AttributeList
	Elements
}

func NewBold() *Bold {
	return &Bold{}
}

func (Bold) Type() ElementType {
	return ElementTypeInline
}

func (Bold) TextFormat() TextFormat {
	return TextFormatBold
}

func (b *Bold) Equals(e Element) bool {
	ob, ok := e.(*Bold)
	if !ok {
		return false
	}
	if !b.AttributeList.Equals(ob.AttributeList) {
		return false
	}
	return b.Elements.Equals(ob.Elements)
}

func (b *Bold) Traverse(parent ParentElement) iter.Seq2[ParentElement, Parent] {
	return func(yield func(ParentElement, Parent) bool) {
		if !b.AttributeList.traverse(b, yield) {
			return
		}
	}
}

var _ ParentElement = &Bold{}
var _ Traverser = &Bold{}

func (b *Bold) Clone() Element {
	return &Bold{position: b.position, raw: b.raw, AttributeList: b.AttributeList.Clone(), Elements: b.Elements.Clone()}
}

type DoubleBold struct {
	position
	raw

	AttributeList
	Elements
}

func NewDoubleBold() *DoubleBold {
	return &DoubleBold{}
}

func (DoubleBold) Type() ElementType {
	return ElementTypeInline
}

func (DoubleBold) TextFormat() TextFormat {
	return TextFormatBold
}

func (db *DoubleBold) Equals(e Element) bool {
	ob, ok := e.(*DoubleBold)
	if !ok {
		return false
	}
	if !db.AttributeList.Equals(ob.AttributeList) {
		return false
	}
	return db.Elements.Equals(ob.Elements)
}

func (db *DoubleBold) Traverse(parent ParentElement) iter.Seq2[ParentElement, Parent] {
	return func(yield func(ParentElement, Parent) bool) {
		if !db.AttributeList.traverse(db, yield) {
			return
		}
	}
}

func (db *DoubleBold) Clone() Element {
	return &DoubleBold{position: db.position, raw: db.raw, AttributeList: db.AttributeList.Clone(), Elements: db.Elements.Clone()}
}

var _ ParentElement = &DoubleBold{}
var _ Traverser = &DoubleBold{}
