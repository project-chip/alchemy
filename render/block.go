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
	default:
		fmt.Printf("unknown delimited block kind: %s\n", db.Kind)
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
			fmt.Printf("unknown comment element type: %T\n", el)
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
			fmt.Printf("unknown sidebar element type: %T\n", el)
		}
	}
}
