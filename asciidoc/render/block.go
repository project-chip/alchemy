package render

import (
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
)

func renderDelimiter(cxt *Context, delimiter asciidoc.Delimiter) {
	var char string
	switch delimiter.Type {
	case asciidoc.DelimitedBlockTypeComment:
		char = "/"
	case asciidoc.DelimitedBlockTypeExample:
		char = "="
	case asciidoc.DelimitedBlockTypeFenced:
		char = "`"
	case asciidoc.DelimitedBlockTypeListing:
		char = "-"
	case asciidoc.DelimitedBlockTypeLiteral:
		char = "."
	case asciidoc.DelimitedBlockTypeNone:
		return
	case asciidoc.DelimitedBlockTypeOpen:
		char = "-"
	case asciidoc.DelimitedBlockTypeStem:
		char = "+"
	case asciidoc.DelimitedBlockTypeQuote:
		char = "_"
	case asciidoc.DelimitedBlockTypeSidebar:
		char = "*"
	case asciidoc.DelimitedBlockTypeMultiLineComment:
		char = "/"
	}
	cxt.EnsureNewLine()
	cxt.WriteString(strings.Repeat(char, delimiter.Length))
	cxt.EnsureNewLine()
}
