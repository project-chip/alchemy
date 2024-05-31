package asciidoc

type Checklist int

const (
	ChecklistUnknown Checklist = iota
	ChecklistUnchecked
	ChecklistChecked
)

type UnorderedList struct {
	position
	raw

	Set
	AttributeList
}

func (UnorderedList) Type() ElementType {
	return ElementTypeBlock
}

type OrderedListItem struct {
	position
	raw

	Set
	AttributeList

	Indent string
	Marker string
}

func NewOrderedListItem(indent string, marker string) *OrderedListItem {
	return &OrderedListItem{
		Indent: indent,
		Marker: marker,
	}
}

func (OrderedListItem) Type() ElementType {
	return ElementTypeBlock
}

func (a *OrderedListItem) Equals(o Element) bool {
	oa, ok := o.(*OrderedListItem)
	if !ok {
		return false
	}
	if a.Marker != oa.Marker {
		return false
	}
	if !a.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return a.Set.Equals(oa.Set)
}

type UnorderedListItem struct {
	position
	raw

	Set
	AttributeList

	Indent    string
	Marker    string
	Checklist Checklist
}

func NewUnorderedListItem(indent string, marker string, checklist Checklist) *UnorderedListItem {
	return &UnorderedListItem{
		Indent:    indent,
		Marker:    marker,
		Checklist: checklist,
	}
}

func (UnorderedListItem) Type() ElementType {
	return ElementTypeBlock
}

func (a *UnorderedListItem) Equals(o Element) bool {
	oa, ok := o.(*UnorderedListItem)
	if !ok {
		return false
	}
	if a.Marker != oa.Marker || a.Checklist != oa.Checklist {
		return false
	}
	if !a.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return a.Set.Equals(oa.Set)
}

type DescriptionListItem struct {
	position
	raw

	Set
	AttributeList

	Marker string
	Term   Set
}

func NewDescriptionListItem(marker string, term Set) *DescriptionListItem {
	return &DescriptionListItem{
		Marker: marker,
		Term:   term,
	}
}

func (DescriptionListItem) Type() ElementType {
	return ElementTypeBlock
}

func (a *DescriptionListItem) Equals(o Element) bool {
	oa, ok := o.(*DescriptionListItem)
	if !ok {
		return false
	}
	if a.Marker != oa.Marker || !a.Term.Equals(oa.Term) {
		return false
	}
	if !a.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return a.Set.Equals(oa.Set)
}

type AttachedBlock struct {
	position
	raw

	ChildElement Element
}

func (AttachedBlock) Type() ElementType {
	return ElementTypeBlock
}

func NewAttachedBlock(child Element) *AttachedBlock {
	return &AttachedBlock{
		ChildElement: child,
	}
}

func (a *AttachedBlock) Equals(o Element) bool {
	oa, ok := o.(*AttachedBlock)
	if !ok {
		return false
	}

	return a.ChildElement.Equals(oa.ChildElement)
}

func (s *AttachedBlock) Child() Element {
	return s.ChildElement
}

type ListContinuation struct {
	position
	raw

	ChildElement Element
}

func (ListContinuation) Type() ElementType {
	return ElementTypeBlock
}

func (s *ListContinuation) Child() Element {
	return s.ChildElement
}

func NewListContinuation(child Element) *ListContinuation {
	return &ListContinuation{
		ChildElement: child,
	}
}

func (a *ListContinuation) Equals(o Element) bool {
	oa, ok := o.(*ListContinuation)
	if !ok {
		return false
	}
	return a.ChildElement.Equals(oa.ChildElement)
}
