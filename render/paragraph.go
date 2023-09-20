package render

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/output"
)

func renderParagraph(cxt *output.Context, p *types.Paragraph, previous *interface{}) {
	renderAttributes(cxt, p, p.Attributes)

	for _, e := range p.Elements {
		switch el := e.(type) {
		case *types.StringElement:
			text, _ := el.RawText()
			cxt.WriteString(text)
		case *types.InternalCrossReference:
			renderInternalCrossReference(cxt, el)
		case *types.Symbol:
			renderSymbol(cxt, el)
		case *types.SpecialCharacter:
			renderSpecialCharacter(cxt, el)
		case *types.QuotedText:
			renderQuotedText(cxt, el)
		case *types.InlineLink:
			text := el.Location.Scheme
			cxt.WriteString(text)
			switch p := el.Location.Path.(type) {
			case string:
				cxt.WriteString(p)
			default:
				fmt.Printf("unknown location path type: %T\n", p)
			}
		case *types.LineBreak:
			cxt.WriteString(" +")
		case *types.PredefinedAttribute:
			cxt.WriteString(fmt.Sprintf("{%s}", el.Name))
		default:
			fmt.Printf("unknown paragraph element type: %T\n", el)
		}
		*previous = e
	}
}
