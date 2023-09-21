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
			renderInlineLink(cxt, el)
		case *types.LineBreak:
			cxt.WriteString(" +")
		case *types.PredefinedAttribute:
			cxt.WriteString(fmt.Sprintf("{%s}", el.Name))
		case *types.InlineImage:
			renderInlineImage(cxt, el)
		case *types.SinglelineComment:
			cxt.WriteString("//")
			cxt.WriteString(el.Content)
			cxt.WriteNewline()
		default:
			panic(fmt.Errorf("unknown paragraph element type: %T", el))
		}
		*previous = e
	}
}
