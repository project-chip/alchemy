package render

import (
	"strings"

	"github.com/hasty/alchemy/asciidoc"
)

func renderBlock(cxt *Context, block asciidoc.Element, delimiter string) (err error) {
	ha, ok := block.(asciidoc.Attributable)
	if ok {
		err = renderAttributes(cxt, block, ha.Attributes(), false)

	}
	if err != nil {
		return
	}
	cxt.EnsureNewLine()
	cxt.WriteString(delimiter)
	cxt.EnsureNewLine()
	he, ok := block.(asciidoc.HasElements)
	if ok {
		err = Elements(cxt, "", he.Elements()...)
	}
	cxt.EnsureNewLine()
	cxt.WriteString(delimiter)
	cxt.EnsureNewLine()
	return
}

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
