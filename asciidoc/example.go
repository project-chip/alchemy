package asciidoc

type ExampleBlock struct {
	position
	raw

	Delimiter Delimiter
	AttributeList
	Set
}

func NewExampleBlock(delimiter Delimiter) *ExampleBlock {
	return &ExampleBlock{Delimiter: delimiter}
}

func (ExampleBlock) Type() ElementType {
	return ElementTypeBlock
}

func (a *ExampleBlock) Equals(o Element) bool {
	oa, ok := o.(*ExampleBlock)
	if !ok {
		return false
	}
	if a.Delimiter != oa.Delimiter {
		return false
	}
	if !a.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return a.Set.Equals(oa.Set)
}
