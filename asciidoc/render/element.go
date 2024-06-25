package render

import (
	"fmt"
	"strings"

	"github.com/hasty/alchemy/asciidoc"
	"github.com/hasty/alchemy/internal/parse"
)

type Section interface {
	GetASCIISection() *asciidoc.Section
}

func Elements(cxt *Context, prefix string, elementList ...asciidoc.Element) (err error) {
	for _, e := range elementList {
		if he, ok := e.(Section); ok {
			e = he.GetASCIISection()
		}
		if hb, ok := e.(parse.HasBase); ok {
			e = hb.GetBase()
		}
		switch el := e.(type) {
		case asciidoc.EmptyLine:
			cxt.EnsureNewLine()
			cxt.WriteRune('\n')
		case *asciidoc.NewLine:
			cxt.WriteRune('\n')
		case *asciidoc.Section:
			err = renderSection(cxt, el)
			if err == nil {
				err = Elements(cxt, "", el.Elements()...)
			}
		case *asciidoc.Paragraph:
			err = renderParagraph(cxt, el)
			if err != nil {
				return
			}
		case *asciidoc.Table:
			err = renderTable(cxt, el)
		case *asciidoc.EmptyLine:
			cxt.EnsureNewLine()
			cxt.WriteRune('\n')
		case *asciidoc.CrossReference:
			err = renderInternalCrossReference(cxt, el)
		case *asciidoc.AttributeEntry:
			err = renderAttributeEntry(cxt, el)
		case *asciidoc.String:
			text := el.Value
			if strings.HasPrefix(text, "ifdef::") || strings.HasPrefix(text, "ifndef::") || strings.HasPrefix(text, "endif::[]") {
				cxt.EnsureNewLine()
			}
			cxt.WriteString(text)
		case *asciidoc.SingleLineComment:
			renderSingleLineComment(cxt, el)
		case *asciidoc.BlockImage:
			err = renderImageBlock(cxt, el)
		case *asciidoc.Link:
			err = renderLink(cxt, el)
		case asciidoc.SpecialCharacter:
			err = renderSpecialCharacter(cxt, el)
		case *asciidoc.Bold:
			err = renderFormattedText(cxt, el, "*")
		case *asciidoc.DoubleBold:
			err = renderFormattedText(cxt, el, "**")
		case *asciidoc.Monospace:
			err = renderFormattedText(cxt, el, "`")
		case *asciidoc.DoubleMonospace:
			err = renderFormattedText(cxt, el, "``")
		case *asciidoc.Superscript:
			err = renderFormattedText(cxt, el, "^")
		case *asciidoc.Subscript:
			err = renderFormattedText(cxt, el, "~")
		case *asciidoc.Italic:
			err = renderFormattedText(cxt, el, "_")
		case *asciidoc.DoubleItalic:
			err = renderFormattedText(cxt, el, "__")
		case *asciidoc.Marked:
			err = renderFormattedText(cxt, el, "#")
		case *asciidoc.DoubleMarked:
			err = renderFormattedText(cxt, el, "##")
		case *asciidoc.LineContinuation:
			cxt.WriteString(" +")
		case asciidoc.AttributeReference:
			cxt.WriteString(fmt.Sprintf("{%s}", el.Name()))
		case *asciidoc.InlineImage:
			err = renderInlineImage(cxt, el)
		case *asciidoc.InlinePassthrough:
			cxt.WriteString("+")
			err = Elements(cxt, "", el.Elements()...)
			cxt.WriteString("+")
		case *asciidoc.InlineDoublePassthrough:
			cxt.WriteString("++")
			err = Elements(cxt, "", el.Elements()...)
			cxt.WriteString("++")
		case *asciidoc.AttributeReset:
			renderAttributeReset(cxt, el)
		case *asciidoc.UnorderedListItem:
			err = renderUnorderedListElement(cxt, el)
		case *asciidoc.OrderedListItem:
			err = renderOrderedListElement(cxt, el)
		case *asciidoc.ListContinuation:
			cxt.EnsureNewLine()
			cxt.WriteString("+\n")
			err = Elements(cxt, "", el.Child())
		case *asciidoc.IfDef:
			renderConditional(cxt, "ifdef::", el.Attributes, el.Union)
		case *asciidoc.IfNDef:
			renderConditional(cxt, "ifndef::", el.Attributes, el.Union)
		case *asciidoc.IfEval:
			renderIfEval(cxt, el)
		case *asciidoc.EndIf:
			renderConditional(cxt, "endif::", el.Attributes, el.Union)
		case *asciidoc.MultiLineComment:
			renderDelimitedLines(cxt, el, el.Delimiter)
		case *asciidoc.DescriptionListItem:
			renderDescriptionListItem(cxt, el)
		case *asciidoc.LiteralBlock:
			renderDelimitedLines(cxt, el, el.Delimiter)
		case *asciidoc.SidebarBlock:
			renderDelimitedElements(cxt, el, el.Delimiter)
		case *asciidoc.Listing:
			renderDelimitedLines(cxt, el, el.Delimiter)
		case *asciidoc.ExampleBlock:
			renderDelimitedElements(cxt, el, el.Delimiter)
		case *asciidoc.StemBlock:
			renderDelimitedLines(cxt, el, el.Delimiter)
		case *asciidoc.OpenBlock:
			renderDelimitedElements(cxt, el, el.Delimiter)
		case *asciidoc.FileInclude:
			renderFileInclude(cxt, el)
		case *asciidoc.Anchor:
			err = renderAnchor(cxt, el)
		case *asciidoc.Admonition:
			renderAdmonition(cxt, el.AdmonitionType)
		case *asciidoc.AttachedBlock:
			cxt.WriteString("+\n")
			err = Elements(cxt, "", el.Child())
		case *asciidoc.LineBreak:
			cxt.WriteString("+")
		case *asciidoc.Counter:
			renderCounter(cxt, el)
		case *asciidoc.ThematicBreak:
			cxt.WriteString("'''\n")
		case nil:
		default:
			err = fmt.Errorf("unknown render element type: %T", el)
		}
		if err != nil {
			return
		}
	}
	return
}
