package asciidoc

type ExampleBlock struct {
	position
	raw

	Delimiter Delimiter
	AttributeList
	Elements
}

func NewExampleBlock(delimiter Delimiter) *ExampleBlock {
	return &ExampleBlock{Delimiter: delimiter}
}

func (ExampleBlock) Type() ElementType {
	return ElementTypeBlock
}

func (eb *ExampleBlock) Equals(o Element) bool {
	oa, ok := o.(*ExampleBlock)
	if !ok {
		return false
	}
	if eb.Delimiter != oa.Delimiter {
		return false
	}
	if !eb.AttributeList.Equals(oa.AttributeList) {
		return false
	}
	return eb.Elements.Equals(oa.Elements)
}

func (eb *ExampleBlock) Clone() Element {
	return &ExampleBlock{position: eb.position, raw: eb.raw, Delimiter: eb.Delimiter, AttributeList: eb.AttributeList.Clone(), Elements: eb.Elements.Clone()}
}
