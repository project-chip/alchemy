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

func (slc *SingleLineComment) Clone() Element {
	return &SingleLineComment{position: slc.position, raw: slc.raw, Value: slc.Value}
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

func (mlc *MultiLineComment) Clone() Element {
	return &MultiLineComment{position: mlc.position, raw: mlc.raw, Delimiter: mlc.Delimiter, LineList: mlc.LineList.Clone()}
}
