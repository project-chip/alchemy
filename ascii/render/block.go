package render

import (
	"strings"

	"github.com/hasty/adoc/elements"
)

func renderBlock(cxt *Context, block elements.Element, delimiter string) (err error) {
	ha, ok := block.(elements.Attributable)
	if ok {
		err = renderAttributes(cxt, block, ha.Attributes(), false)

	}
	if err != nil {
		return
	}
	cxt.WriteNewline()
	cxt.WriteString(delimiter)
	cxt.WriteNewline()
	he, ok := block.(elements.HasElements)
	if ok {
		err = Elements(cxt, "", he.Elements()...)
	}
	cxt.WriteNewline()
	cxt.WriteString(delimiter)
	cxt.WriteNewline()
	return
}

func renderDelimiter(cxt *Context, delimiter elements.Delimiter) {
	var char string
	switch delimiter.Type {
	case elements.DelimitedBlockTypeComment:
		char = "/"
	case elements.DelimitedBlockTypeExample:
		char = "="
	case elements.DelimitedBlockTypeFenced:
		char = "`"
	case elements.DelimitedBlockTypeListing:
		char = "-"
	case elements.DelimitedBlockTypeLiteral:
		char = "."
	case elements.DelimitedBlockTypeNone:
		return
	case elements.DelimitedBlockTypeOpen:
		char = "-"
	case elements.DelimitedBlockTypeStem:
		char = "+"
	case elements.DelimitedBlockTypeQuote:
		char = "_"
	case elements.DelimitedBlockTypeSidebar:
		char = "*"
	case elements.DelimitedBlockTypeMultiLineComment:
		char = "/"
	}
	cxt.WriteNewline()
	cxt.WriteString(strings.Repeat(char, delimiter.Length))
	cxt.WriteNewline()
}
