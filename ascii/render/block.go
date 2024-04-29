package render

import (
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

func renderMultiLineComment(cxt *Context, comment *elements.MultiLineComment) (err error) {
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
