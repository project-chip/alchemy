package asciidoc

type FencedDelimiter struct {
	Delimiter

	Language Elements
}

func (fd FencedDelimiter) Equals(ofd FencedDelimiter) bool {
	if !fd.Delimiter.Equals(ofd.Delimiter) {
		return false
	}
	return fd.Language.Equals(ofd.Language)
}

func NewFencedDelimiter(length int, language Elements) FencedDelimiter {
	return FencedDelimiter{Delimiter: Delimiter{Type: DelimitedBlockTypeFenced, Length: length}, Language: language}
}

type FencedBlock struct {
	position
	raw

	Delimiter FencedDelimiter
	AttributeList
	Elements
}

func NewFencedBlock(delimiter FencedDelimiter) *FencedBlock {
	return &FencedBlock{Delimiter: delimiter}
}

func (FencedBlock) Type() ElementType {
	return ElementTypeBlock
}

func (a *FencedBlock) Equals(o Element) bool {
	oa, ok := o.(*FencedBlock)
	if !ok {
		return false
	}
	if !a.Delimiter.Equals(oa.Delimiter) {
		return false
	}
	if !a.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return a.Elements.Equals(oa.Elements)
}
