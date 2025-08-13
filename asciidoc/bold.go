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

func (b *Bold) Traverse(parent Parent) iter.Seq2[Parent, Parent] {
	return func(yield func(Parent, Parent) bool) {
		if !b.AttributeList.traverse(b, yield) {
			return
		}
	}
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

func (db *DoubleBold) Traverse(parent Parent) iter.Seq2[Parent, Parent] {
	return func(yield func(Parent, Parent) bool) {
		if !db.AttributeList.traverse(db, yield) {
			return
		}
	}
}
