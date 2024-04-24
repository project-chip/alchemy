package render

import (
	"fmt"
	"strings"

	"github.com/hasty/adoc/elements"
	"github.com/hasty/alchemy/internal/parse"
)

type Section interface {
	GetASCIISection() *elements.Section
}

func Elements(cxt *Context, prefix string, elementList ...elements.Element) (err error) {
	var previous any
	for _, e := range elementList {
		if he, ok := e.(Section); ok {
			e = he.GetASCIISection()
		}
		if hb, ok := e.(parse.HasBase); ok {
			e = hb.GetBase()
		}
		switch el := e.(type) {
		case *elements.Section:
			err = renderSection(cxt, el)
			if err == nil {
				err = Elements(cxt, "", el.Elements()...)
			}
		case *elements.DelimitedBlock:
			err = renderDelimitedBlock(cxt, el)
		case *elements.Paragraph:
			cxt.WriteString(prefix)
			err = renderParagraph(cxt, el, &previous)
			if err != nil {
				return
			}
		case *elements.Table:
			err = renderTable(cxt, el)
		case *elements.EmptyLine:
			cxt.WriteNewline()
			cxt.WriteRune('\n')
		case *elements.CrossReference:
			err = renderInternalCrossReference(cxt, el)
		case *elements.List:
			err = renderList(cxt, el)
		case *elements.AttributeEntry:
			err = renderAttributeEntry(cxt, el)
		case elements.String:
			text := string(el)
			if strings.HasPrefix(text, "ifdef::") || strings.HasPrefix(text, "ifndef::") || strings.HasPrefix(text, "endif::[]") {
				cxt.WriteNewline()
			}
			cxt.WriteString(text)
		case *elements.SingleLineComment:
			cxt.WriteNewline()
			cxt.WriteString("//")
			cxt.WriteString(el.Value)
			cxt.WriteNewline()
		case *elements.BlockImage:
			err = renderImageBlock(cxt, el)
		case *elements.Link:
			err = renderInlineLink(cxt, el)
		case *elements.SpecialCharacter:
			err = renderSpecialCharacter(cxt, el)
		case *elements.QuotedText:
			err = renderQuotedText(cxt, el)
		case *elements.Preamble:
			err = renderPreamble(cxt, el)
		case *elements.Symbol:
			err = renderSymbol(cxt, el)
		case *elements.LineContinuation:
			cxt.WriteString(" +")
		case *elements.DocumentHeader:
			err = renderAttributes(cxt, el, el.Attributes, false)
			if err != nil {
				return
			}
			err = renderSectionTitle(cxt, el.Title, 1)
			if err != nil {
				return
			}
			err = Elements(cxt, "", el.Elements)
			cxt.WriteNewline()
		case elements.AttributeReference:
			cxt.WriteString(fmt.Sprintf("{%s}", el.Name()))
		case *elements.InlineImage:
			err = renderInlineImage(cxt, el)
		case *elements.FootnoteReference:
			err = renderFootnoteReference(cxt, el)
		case *elements.InlinePassthrough:
			switch el.Kind {
			case elements.SinglePlusPassthrough, elements.TriplePlusPassthrough:
				cxt.WriteString(string(el.Kind))
				err = Elements(cxt, "", el.Elements)
				cxt.WriteString(string(el.Kind))
			case elements.PassthroughMacro:
				cxt.WriteString("pass:[")
				err = Elements(cxt, "", el.Elements)
				cxt.WriteRune(']')
			}
		case *elements.AttributeReset:
			err = renderAttributes(cxt, el, el.Attributes, false)
			if err != nil {
				return
			}
			renderAttributeReset(cxt, el)
		case *elements.RawLine:
			cxt.WriteString(el.Content)
			if el.EOL {
				cxt.WriteRune('\n')
			}
		case *elements.ListElements:
			err = renderListElements(cxt, el)
		case *elements.UnorderedListElement:
			err = renderUnorderedListElement(cxt, el)
		case *elements.OrderedListElement:
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
