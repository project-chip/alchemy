package render

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/output"
)

func renderParagraph(cxt *output.Context, p *types.Paragraph, previous *interface{}) (err error) {
	renderAttributes(cxt, p, p.Attributes)

	for _, e := range p.Elements {
		switch el := e.(type) {
		case *types.StringElement:
			text, _ := el.RawText()
			cxt.WriteString(text)
		case *types.InternalCrossReference:
			renderInternalCrossReference(cxt, el)
		case *types.Symbol:
			err = renderSymbol(cxt, el)
		case *types.SpecialCharacter:
			err = renderSpecialCharacter(cxt, el)
		case *types.QuotedText:
			err = renderQuotedText(cxt, el)
		case *types.InlineLink:
			err = renderInlineLink(cxt, el)
		case *types.LineBreak:
			cxt.WriteString(" +")
		case *types.PredefinedAttribute:
			cxt.WriteString(fmt.Sprintf("{%s}", el.Name))
		case *types.InlineImage:
			err = renderInlineImage(cxt, el)
		case *types.SinglelineComment:
			cxt.WriteString("//")
			cxt.WriteString(el.Content)
			cxt.WriteNewline()
		default:
			err = fmt.Errorf("unknown paragraph element type: %T", el)
		}
		if err != nil {
			return
		}
		*previous = e
	}
	return
}
