package adoc

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/output"
)

func renderDelimitedBlock(cxt *output.Context, db *types.DelimitedBlock) (err error) {
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

func renderComment(cxt *output.Context, comment *types.DelimitedBlock) (err error) {
	for _, e := range comment.Elements {
		switch el := e.(type) {
		case *types.StringElement:
			cxt.WriteNewline()
			cxt.WriteString("////")
			cxt.WriteNewline()
			text, _ := el.RawText()
			cxt.WriteString(text)
			cxt.WriteNewline()
			cxt.WriteString("////")
			cxt.WriteNewline()
		default:
			err = fmt.Errorf("unknown comment element type: %T", el)
		}
		if err != nil {
			return
		}
	}
	return
}

func renderBlock(cxt *output.Context, block *types.DelimitedBlock, delimiter string) (err error) {
	err = renderAttributes(cxt, block, block.Attributes, false)
	if err != nil {
		return
	}
	cxt.WriteNewline()
	cxt.WriteString(delimiter)
	cxt.WriteNewline()
	err = RenderElements(cxt, "", block.Elements)
	cxt.WriteNewline()
	cxt.WriteString(delimiter)
	cxt.WriteNewline()
	return
}
