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

	Elements
	AttributeList
}

func (UnorderedList) Type() ElementType {
	return ElementTypeBlock
}

type OrderedListItem struct {
	position
	raw

	Elements
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

func (oli *OrderedListItem) Equals(o Element) bool {
	oa, ok := o.(*OrderedListItem)
	if !ok {
		return false
	}
	if oli.Marker != oa.Marker {
		return false
	}
	if !oli.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return oli.Elements.Equals(oa.Elements)
}

func (oli *OrderedListItem) Clone() Element {
	return &OrderedListItem{
		position:      oli.position,
		raw:           oli.raw,
		Elements:      oli.Elements.Clone(),
		AttributeList: oli.AttributeList.Clone(),
		Indent:        oli.Indent,
		Marker:        oli.Marker,
	}
}

type UnorderedListItem struct {
	position
	raw

	Elements
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

func (uoli *UnorderedListItem) Equals(o Element) bool {
	oa, ok := o.(*UnorderedListItem)
	if !ok {
		return false
	}
	if uoli.Marker != oa.Marker || uoli.Checklist != oa.Checklist {
		return false
	}
	if !uoli.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return uoli.Elements.Equals(oa.Elements)
}

func (uoli *UnorderedListItem) Clone() Element {
	return &UnorderedListItem{
		position:      uoli.position,
		raw:           uoli.raw,
		Elements:      uoli.Elements.Clone(),
		AttributeList: uoli.AttributeList.Clone(),
		Indent:        uoli.Indent,
		Marker:        uoli.Marker,
		Checklist:     uoli.Checklist,
	}
}

type DescriptionListItem struct {
	position
	raw

	Elements
	AttributeList

	Marker string
	Term   Elements
}

func NewDescriptionListItem(marker string, term Elements) *DescriptionListItem {
	return &DescriptionListItem{
		Marker: marker,
		Term:   term,
	}
}

func (DescriptionListItem) Type() ElementType {
	return ElementTypeBlock
}

func (dli *DescriptionListItem) Equals(o Element) bool {
	oa, ok := o.(*DescriptionListItem)
	if !ok {
		return false
	}
	if dli.Marker != oa.Marker || !dli.Term.Equals(oa.Term) {
		return false
	}
	if !dli.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return dli.Elements.Equals(oa.Elements)
}

func (dli *DescriptionListItem) Clone() Element {
	return &DescriptionListItem{
		position:      dli.position,
		raw:           dli.raw,
		Elements:      dli.Elements.Clone(),
		AttributeList: dli.AttributeList.Clone(),
		Marker:        dli.Marker,
		Term:          dli.Term.Clone(),
	}
}

type ListContinuation struct {
	position
	raw

	NewLineCount int
	ChildElement Element
}

func (ListContinuation) Type() ElementType {
	return ElementTypeBlock
}

func (lc *ListContinuation) Child() Element {
	return lc.ChildElement
}

func NewListContinuation(child Element, newLineCount int) *ListContinuation {
	return &ListContinuation{
		ChildElement: child,
		NewLineCount: newLineCount,
	}
}

func (lc *ListContinuation) Equals(o Element) bool {
	oa, ok := o.(*ListContinuation)
	if !ok {
		return false
	}
	return lc.ChildElement.Equals(oa.ChildElement)
}

func (lc *ListContinuation) Clone() Element {
	return &ListContinuation{
		position:     lc.position,
		raw:          lc.raw,
		ChildElement: lc.ChildElement.Clone(),
		NewLineCount: lc.NewLineCount,
	}
}
