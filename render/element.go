package render

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/output"
)

func RenderElements(cxt *output.Context, prefix string, elements []interface{}) {
	var previous interface{}
	for _, e := range elements {
		if section, ok := e.(*ascii.Section); ok {
			renderSection(cxt, section.Base)
			RenderElements(cxt, "", section.Elements)
			continue
		}
		if el, ok := e.(*ascii.Element); ok {
			e = el.Base
		}
		switch el := e.(type) {
		case *types.DelimitedBlock:
			renderDelimitedBlock(cxt, el)
		case *types.Paragraph:
			cxt.WriteString(prefix)
			renderParagraph(cxt, el, &previous)
			continue // skip setting previous to the paragraph itself
		case *types.Table:
			renderTable(cxt, el)
		case *types.BlankLine:
			if _, ok := previous.(*types.StringElement); ok {
				cxt.WriteRune('\n')
			}
			cxt.WriteRune('\n')
		case *types.InternalCrossReference:
			renderInternalCrossReference(cxt, el)
		case *types.List:
			renderList(cxt, el)
		case *types.AttributeDeclaration:
			renderAttributeDeclaration(cxt, el)
		case *types.StringElement:
			text, _ := el.RawText()
			cxt.WriteString(text)
		case *types.SinglelineComment:
			cxt.WriteString("//")
			cxt.WriteString(el.Content)
			cxt.WriteNewline()
		case *types.ImageBlock:
			renderImageBlock(cxt, el)
		case *types.InlineLink:
			renderInlineLink(cxt, el)
		case *types.SpecialCharacter:
			renderSpecialCharacter(cxt, el)
		case *types.QuotedText:
			renderQuotedText(cxt, el)
		case *types.Preamble:
			renderPreamble(cxt, el)
		case *types.Symbol:
			renderSymbol(cxt, el)
		case *types.DocumentHeader:
			renderAttributes(cxt, el, el.Attributes)
			renderSectionTitle(cxt, el.Title, 1)
			RenderElements(cxt, "", el.Elements)
			cxt.WriteRune('\n')
		default:
			panic(fmt.Errorf("unknown element type: %T", el))
		}
		previous = e
	}
}
