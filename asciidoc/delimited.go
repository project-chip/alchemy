package asciidoc

type DelimitedBlockType uint8

const (
	DelimitedBlockTypeNone DelimitedBlockType = iota
	DelimitedBlockTypeComment
	DelimitedBlockTypeMultiLineComment
	DelimitedBlockTypeExample
	DelimitedBlockTypeFenced
	DelimitedBlockTypeListing
	DelimitedBlockTypeLiteral
	DelimitedBlockTypeOpen
	DelimitedBlockTypeSidebar
	DelimitedBlockTypeStem
	DelimitedBlockTypeTable
	DelimitedBlockTypeQuote
)

type Delimiter struct {
	position
	raw

	Type   DelimitedBlockType
	Length int
}

func (d Delimiter) Equals(od Delimiter) bool {
	return d.Type == od.Type && d.Length == od.Length
}
