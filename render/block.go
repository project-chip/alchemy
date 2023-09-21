package render

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/output"
)

func renderDelimitedBlock(cxt *output.Context, db *types.DelimitedBlock) {
	switch db.Kind {
	case "comment":
		renderComment(cxt, db)
	case "sidebar":
		renderSidebar(cxt, db)
	case "example":
		renderExample(cxt, db)
	case "listing":
		renderListing(cxt, db)
	case "literal":
		renderLiteral(cxt, db)
	default:
		panic(fmt.Errorf("unknown delimited block kind: %s", db.Kind))
	}
}

func renderComment(cxt *output.Context, comment *types.DelimitedBlock) {
	for _, e := range comment.Elements {
		switch el := e.(type) {
		case *types.StringElement:
			cxt.WriteRune('\n')
			cxt.WriteString("////")
			cxt.WriteNewline()
			text, _ := el.RawText()
			cxt.WriteString(text)
			cxt.WriteNewline()
			cxt.WriteString("////")
			cxt.WriteNewline()
		default:
			panic(fmt.Errorf("unknown comment element type: %T", el))
		}
	}
}

func renderExample(cxt *output.Context, comment *types.DelimitedBlock) {
	renderAttributes(cxt, comment, comment.Attributes)
	cxt.WriteNewline()
	cxt.WriteString("====")
	cxt.WriteNewline()
	RenderElements(cxt, "", comment.Elements)
	cxt.WriteNewline()
	cxt.WriteString("====")
	cxt.WriteNewline()

}

func renderListing(cxt *output.Context, comment *types.DelimitedBlock) {
	renderAttributes(cxt, comment, comment.Attributes)
	cxt.WriteNewline()
	cxt.WriteString("----")
	cxt.WriteNewline()
	RenderElements(cxt, "", comment.Elements)
	cxt.WriteNewline()
	cxt.WriteString("----")
	cxt.WriteNewline()

}

func renderLiteral(cxt *output.Context, comment *types.DelimitedBlock) {
	renderAttributes(cxt, comment, comment.Attributes)
	cxt.WriteNewline()
	cxt.WriteString("----")
	cxt.WriteNewline()
	RenderElements(cxt, "", comment.Elements)
	cxt.WriteNewline()
	cxt.WriteString("----")
	cxt.WriteNewline()

}

func renderSidebar(cxt *output.Context, comment *types.DelimitedBlock) {
	var previous interface{}
	for _, e := range comment.Elements {
		switch el := e.(type) {
		case *types.Paragraph:
			cxt.WriteNewline()
			cxt.WriteString("****")
			cxt.WriteNewline()
			renderParagraph(cxt, el, &previous)
			cxt.WriteNewline()
			cxt.WriteString("****")
			cxt.WriteNewline()
		case *types.StringElement:
			cxt.WriteRune('\n')
			cxt.WriteString("****")
			cxt.WriteNewline()
			text, _ := el.RawText()
			cxt.WriteString(text)
			cxt.WriteNewline()
			cxt.WriteString("****")
			cxt.WriteNewline()
		default:
			panic(fmt.Errorf("unknown sidebar element type: %T", el))
		}
	}
}
