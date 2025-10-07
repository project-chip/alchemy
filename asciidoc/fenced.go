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

func (fb *FencedBlock) Equals(o Element) bool {
	oa, ok := o.(*FencedBlock)
	if !ok {
		return false
	}
	if !fb.Delimiter.Equals(oa.Delimiter) {
		return false
	}
	if !fb.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return fb.Elements.Equals(oa.Elements)
}

func (fb *FencedBlock) Clone() Element {
	return &FencedBlock{position: fb.position, raw: fb.raw, Delimiter: fb.Delimiter, AttributeList: fb.AttributeList.Clone(), Elements: fb.Elements.Clone()}
}
