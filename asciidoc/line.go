package asciidoc

type LineList []string

func (s LineList) Lines() []string {
	return s
}

func (s *LineList) Append(e string) {
	*s = append(*s, e)
}

func (s *LineList) SetLines(els []string) {
	*s = els
}

func (s LineList) Equals(oll LineList) bool {
	if len(s) != len(oll) {
		return false
	}
	for i, l := range s {
		ol := oll[i]
		if l != ol {
			return false
		}
	}
	return true
}

type HasLines interface {
	Element
	Lines() []string
	Append(e string)
	SetLines(e []string)
}

type EmptyLine struct {
	Text string
}

func NewEmptyLine(text string) EmptyLine {
	return EmptyLine{Text: text}
}

func (EmptyLine) Type() ElementType {
	return ElementTypeBlock
}

func (a EmptyLine) Equals(o Element) bool {
	oa, ok := o.(EmptyLine)
	if !ok {
		return false
	}

	return a.Text == oa.Text
}

type ParagraphLine struct {
	Text string
}

func (ParagraphLine) Type() ElementType {
	return ElementTypeInlineLiteral
}

func NewParagraphLine(text string) ParagraphLine {
	return ParagraphLine{Text: text}
}

func (a ParagraphLine) Equals(o Element) bool {
	oa, ok := o.(ParagraphLine)
	if !ok {
		return false
	}

	return a.Text == oa.Text
}
