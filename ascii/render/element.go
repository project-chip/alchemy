package render

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/internal/parse"
)

type Section interface {
	GetAsciiSection() *types.Section
}

func RenderElements(cxt *Context, prefix string, elements []interface{}) (err error) {
	var previous interface{}
	for _, e := range elements {
		if he, ok := e.(Section); ok {
			e = he.GetAsciiSection()
		}
		if hb, ok := e.(parse.HasBase); ok {
			e = hb.GetBase()
		}
		switch el := e.(type) {
		case *types.Section:
			err = renderSection(cxt, el)
			if err == nil {
				err = RenderElements(cxt, "", el.Elements)
			}
		case *types.DelimitedBlock:
			err = renderDelimitedBlock(cxt, el)
		case *types.Paragraph:
			cxt.WriteString(prefix)
			err = renderParagraph(cxt, el, &previous)
			if err != nil {
				return
			}
		case *types.Table:
			err = renderTable(cxt, el)
		case *types.BlankLine:
			switch previous.(type) {
			case *types.StringElement, *types.FootnoteReference, *types.Paragraph:
				//cxt.WriteNewline()
			}
			cxt.WriteNewline()
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
			cxt.WriteNewline()
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
		case *types.LineBreak:
			cxt.WriteString(" +")
		case *types.DocumentHeader:
			err = renderAttributes(cxt, el, el.Attributes, false)
			if err != nil {
				return
			}
			err = renderSectionTitle(cxt, el.Title, 1)
			if err != nil {
				return
			}
			err = RenderElements(cxt, "", el.Elements)
			cxt.WriteNewline()
		case *types.PredefinedAttribute:
			cxt.WriteString(fmt.Sprintf("{%s}", el.Name))
		case *types.InlineImage:
			err = renderInlineImage(cxt, el)
		case *types.FootnoteReference:
			err = renderFootnoteReference(cxt, el)
		case *types.InlinePassthrough:
			switch el.Kind {
			case types.SinglePlusPassthrough, types.TriplePlusPassthrough:
				cxt.WriteString(string(el.Kind))
				err = RenderElements(cxt, "", el.Elements)
				cxt.WriteString(string(el.Kind))
			case types.PassthroughMacro:
				cxt.WriteString("pass:[")
				err = RenderElements(cxt, "", el.Elements)
				cxt.WriteRune(']')
			}
		case *types.AttributeReset:
			err = renderAttributes(cxt, el, el.Attributes, false)
			if err != nil {
				return
			}
			renderAttributeReset(cxt, el)
		case *types.RawLine:
			cxt.WriteString(el.Content)
			if el.EOL {
				cxt.WriteRune('\n')
			}
		case *types.ListElements:
			err = renderListElements(cxt, el)
		case *types.UnorderedListElement:
			err = renderUnorderedListElement(cxt, el)
		case *types.OrderedListElement:
			err = renderOrderedListElement(cxt, el)
		case nil:
		default:
			err = fmt.Errorf("unknown render element type: %T", el)
		}
		if err != nil {
			return
		}
		previous = e
	}
	return
}
