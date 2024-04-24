package render

import (
	"fmt"

	"github.com/hasty/adoc/elements"
)

func renderDelimitedBlock(cxt *Context, db *elements.DelimitedBlock) (err error) {
	switch db.Kind {
	case "comment":
		err = renderComment(cxt, db)
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
}

func renderComment(cxt *Context, comment *elements.MultiLineComment) (err error) {
	for _, e := range comment.Lines {
		cxt.WriteNewline()
		cxt.WriteString("////")
		cxt.WriteNewline()
		cxt.WriteString(e)
		cxt.WriteNewline()
		cxt.WriteString("////")
		cxt.WriteNewline()
	}
	return
}

func renderBlock(cxt *Context, block *elements.DelimitedBlock, delimiter string) (err error) {
	err = renderAttributes(cxt, block, block.Attributes, false)
	if err != nil {
		return
	}
	cxt.WriteNewline()
	cxt.WriteString(delimiter)
	cxt.WriteNewline()
	err = Elements(cxt, "", block.Elements)
	cxt.WriteNewline()
	cxt.WriteString(delimiter)
	cxt.WriteNewline()
	return
}
