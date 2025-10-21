package asciidoc

type LineList []string

func (ll LineList) Lines() []string {
	return ll
}

func (ll *LineList) AppendLine(e string) {
	*ll = append(*ll, e)
}

func (ll *LineList) SetLines(els []string) {
	*ll = els
}

func (ll LineList) Equals(oll LineList) bool {
	if len(ll) != len(oll) {
		return false
	}
	for i, l := range ll {
		ol := oll[i]
		if l != ol {
			return false
		}
	}
	return true
}

func (ll LineList) Clone() []string {
	var els []string
	for _, e := range ll {
		els = append(els, e)
	}
	return els
}

type HasLines interface {
	Element
	Lines() []string
	AppendLine(e string)
	SetLines(e []string)
}

type EmptyLine struct {
	position

	Text string
}

func NewEmptyLine(text string) *EmptyLine {
	return &EmptyLine{Text: text}
}

func (*EmptyLine) Type() ElementType {
	return ElementTypeBlock
}

func (el *EmptyLine) Equals(o Element) bool {
	oa, ok := o.(*EmptyLine)
	if !ok {
		return false
	}

	return el.Text == oa.Text
}

func (el *EmptyLine) Clone() Element {
	return &EmptyLine{position: el.position, Text: el.Text}
}
