package render

import (
	"strings"

	"github.com/hasty/adoc/elements"
)

/*func renderDelimitedBlock(cxt *Context, db *elements.DelimitedBlock) (err error) {
	switch db.Kind {
	case "comment":
		err = renderMultiLineComment(cxt, db)
	case "sidebar":
		err = renderBlock(cxt, db, "****")
	case "example":
		err = renderBlock(cxt, db, "====")
	case "listing":
		err = renderBlock(cxt, db, "----")
	case "literal":
		err = renderBlock(cxt, db, "....")
	case "fenced":
		err = renderBlock(cxt, db, "```")
	case "pass":
		err = renderBlock(cxt, db, "++++")
	case "open":
		err = renderBlock(cxt, db, "--")
	default:
		err = fmt.Errorf("unknown delimited block kind: %s", db.Kind)
	}
	return
}*/

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
	case elements.DelimitedBlockTypeListing:
		char = "-"
	case elements.DelimitedBlockTypeLiteral:
		char = "."
	case elements.DelimitedBlockTypeNone:
		return
	case elements.DelimitedBlockTypeOpen:
		char = "-"
	case elements.DelimitedBlockTypePassthrough:
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
