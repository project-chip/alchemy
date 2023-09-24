package render

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/output"
)

func RenderElements(cxt *output.Context, prefix string, elements []interface{}) (err error) {
	var previous interface{}
	for _, e := range elements {
		if section, ok := e.(*ascii.Section); ok {
			err = renderSection(cxt, section.Base)
			if err != nil {
				return
			}
			err = RenderElements(cxt, "", section.Elements)
			if err != nil {
				return
			}
			continue
		}
		if el, ok := e.(*ascii.Element); ok {
			e = el.Base
		}
		switch el := e.(type) {
		case *types.DelimitedBlock:
			err = renderDelimitedBlock(cxt, el)
		case *types.Paragraph:
			cxt.WriteString(prefix)
			err = renderParagraph(cxt, el, &previous)
			if err != nil {
				return
			}
			continue // skip setting previous to the paragraph itself
		case *types.Table:
			err = renderTable(cxt, el)
		case *types.BlankLine:
			if _, ok := previous.(*types.StringElement); ok {
				cxt.WriteRune('\n')
			}
			cxt.WriteRune('\n')
		case *types.InternalCrossReference:
			err = renderInternalCrossReference(cxt, el)
		case *types.List:
			err = renderList(cxt, el)
		case *types.AttributeDeclaration:
			err = renderAttributeDeclaration(cxt, el)
		case *types.StringElement:
			text, _ := el.RawText()
			cxt.WriteString(text)
		case *types.SinglelineComment:
			cxt.WriteString("//")
			cxt.WriteString(el.Content)
			cxt.WriteNewline()
		case *types.ImageBlock:
			err = renderImageBlock(cxt, el)
		case *types.InlineLink:
			err = renderInlineLink(cxt, el)
		case *types.SpecialCharacter:
			err = renderSpecialCharacter(cxt, el)
		case *types.QuotedText:
			err = renderQuotedText(cxt, el)
		case *types.Preamble:
			err = renderPreamble(cxt, el)
		case *types.Symbol:
			err = renderSymbol(cxt, el)
		case *types.DocumentHeader:
			err = renderAttributes(cxt, el, el.Attributes)
			if err != nil {
				return
			}
			err = renderSectionTitle(cxt, el.Title, 1)
			if err != nil {
				return
			}
			err = RenderElements(cxt, "", el.Elements)
			cxt.WriteRune('\n')
		default:
			err = fmt.Errorf("unknown element type: %T", el)
		}
		if err != nil {
			return
		}
		previous = e
	}
	return
}
