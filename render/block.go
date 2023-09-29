package render

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
		err = renderSidebar(cxt, db)
	case "example":
		err = renderExample(cxt, db)
	case "listing":
		err = renderListing(cxt, db)
	case "literal":
		err = renderLiteral(cxt, db)
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

func renderExample(cxt *output.Context, comment *types.DelimitedBlock) (err error) {
	renderAttributes(cxt, comment, comment.Attributes, false)
	cxt.WriteNewline()
	cxt.WriteString("====")
	cxt.WriteNewline()
	err = RenderElements(cxt, "", comment.Elements)
	cxt.WriteNewline()
	cxt.WriteString("====")
	cxt.WriteNewline()
	return
}

func renderListing(cxt *output.Context, comment *types.DelimitedBlock) (err error) {
	err = renderAttributes(cxt, comment, comment.Attributes, false)
	if err != nil {
		return
	}
	cxt.WriteNewline()
	cxt.WriteString("----")
	cxt.WriteNewline()
	err = RenderElements(cxt, "", comment.Elements)
	if err != nil {
		return
	}
	cxt.WriteNewline()
	cxt.WriteString("----")
	cxt.WriteNewline()
	return
}

func renderLiteral(cxt *output.Context, comment *types.DelimitedBlock) (err error) {
	err = renderAttributes(cxt, comment, comment.Attributes, false)
	if err != nil {
		return
	}
	cxt.WriteNewline()
	cxt.WriteString("----")
	cxt.WriteNewline()
	err = RenderElements(cxt, "", comment.Elements)
	if err != nil {
		return
	}
	cxt.WriteNewline()
	cxt.WriteString("----")
	cxt.WriteNewline()
	return
}

func renderSidebar(cxt *output.Context, comment *types.DelimitedBlock) (err error) {
	err = renderAttributes(cxt, comment, comment.Attributes, false)
	if err != nil {
		return
	}
	cxt.WriteNewline()
	cxt.WriteString("****")
	cxt.WriteNewline()
	err = RenderElements(cxt, "", comment.Elements)
	cxt.WriteNewline()
	cxt.WriteString("****")
	cxt.WriteNewline()
	return
}
