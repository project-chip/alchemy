package asciidoc

type SingleLineComment struct {
	position
	raw

	Value string
}

func NewSingleLineComment(value string) *SingleLineComment {
	return &SingleLineComment{Value: value}
}

func (SingleLineComment) Type() ElementType {
	return ElementTypeBlock
}

func (slc *SingleLineComment) Equals(e Element) bool {
	oslc, ok := e.(*SingleLineComment)
	if !ok {
		return false
	}
	return slc.Value == oslc.Value
}

type MultiLineComment struct {
	position
	raw

	Delimiter Delimiter
	LineList
}

func NewMultiLineComment(delimiter Delimiter) *MultiLineComment {
	return &MultiLineComment{Delimiter: delimiter}
}

func (MultiLineComment) Type() ElementType {
	return ElementTypeBlock
}

func (mlc *MultiLineComment) Equals(e Element) bool {
	omlc, ok := e.(*MultiLineComment)
	if !ok {
		return false
	}
	if !omlc.Delimiter.Equals(mlc.Delimiter) {
		return false
	}
	return omlc.LineList.Equals(mlc.LineList)
}
