package asciidoc

type TextFormat uint8

const (
	TextFormatNone TextFormat = iota
	TextFormatBold
	TextFormatItalic
	TextFormatMonospace
	TextFormatMarked
)

type FormattedTextElement interface {
	Element
	ParentElement
	TextFormat() TextFormat
}
