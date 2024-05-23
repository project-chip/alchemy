package render

import (
	"fmt"
	"strings"

	"github.com/hasty/adoc/asciidoc"
	"github.com/hasty/alchemy/internal/parse"
)

type Section interface {
	GetASCIISection() *asciidoc.Section
}

func Elements(cxt *Context, prefix string, elementList ...asciidoc.Element) (err error) {
	var previous any
	for _, e := range elementList {
		if he, ok := e.(Section); ok {
			e = he.GetASCIISection()
		}
		if hb, ok := e.(parse.HasBase); ok {
			e = hb.GetBase()
		}
		switch el := e.(type) {
		case asciidoc.EmptyLine:
			cxt.WriteNewline()
			cxt.WriteRune('\n')
		case *asciidoc.NewLine:
			cxt.WriteRune('\n')
		case *asciidoc.Section:
			err = renderSection(cxt, el)
			if err == nil {
				err = Elements(cxt, "", el.Elements()...)
			}
		case *asciidoc.Paragraph:
			err = renderParagraph(cxt, el, &previous)
			if err != nil {
				return
			}
		case *asciidoc.Table:
			err = renderTable(cxt, el)
		case *asciidoc.EmptyLine:
			cxt.WriteNewline()
			cxt.WriteRune('\n')
		case *asciidoc.CrossReference:
			err = renderInternalCrossReference(cxt, el)
		case *asciidoc.AttributeEntry:
			err = renderAttributeEntry(cxt, el)
		case *asciidoc.String:
			text := el.Value
			if strings.HasPrefix(text, "ifdef::") || strings.HasPrefix(text, "ifndef::") || strings.HasPrefix(text, "endif::[]") {
				cxt.WriteNewline()
			}
			cxt.WriteString(text)
		case *asciidoc.SingleLineComment:
			cxt.WriteNewline()
			cxt.WriteString("//")
			cxt.WriteString(el.Value)
			cxt.WriteNewline()
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
			cxt.WriteNewline()
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
			renderDelimiter(cxt, el.Delimiter)
			for _, l := range el.Lines() {
				cxt.WriteString(l)
				cxt.WriteRune('\n')
			}
			renderDelimiter(cxt, el.Delimiter)
		case *asciidoc.DescriptionListItem:
			renderAttributes(cxt, el, el.Attributes(), false)
			Elements(cxt, "", el.Term...)
			cxt.WriteString(el.Marker)
			cxt.WriteRune(' ')
			Elements(cxt, "", el.Elements()...)
			cxt.WriteNewline()
		case *asciidoc.LiteralBlock:
			renderAttributes(cxt, el, el.Attributes(), false)
			renderDelimiter(cxt, el.Delimiter)
			for _, l := range el.Lines() {
				cxt.WriteString(l)
				cxt.WriteRune('\n')
			}
			renderDelimiter(cxt, el.Delimiter)
		case *asciidoc.SidebarBlock:
			renderAttributes(cxt, el, el.Attributes(), false)
			renderDelimiter(cxt, el.Delimiter)
			Elements(cxt, "", el.Elements()...)
			renderDelimiter(cxt, el.Delimiter)
		case *asciidoc.Listing:
			renderAttributes(cxt, el, el.Attributes(), false)
			renderDelimiter(cxt, el.Delimiter)
			for _, l := range el.Lines() {
				cxt.WriteString(l)
				cxt.WriteRune('\n')
			}
			renderDelimiter(cxt, el.Delimiter)
		case *asciidoc.ExampleBlock:
			renderAttributes(cxt, el, el.Attributes(), false)
			renderDelimiter(cxt, el.Delimiter)
			Elements(cxt, "", el.Elements()...)
			renderDelimiter(cxt, el.Delimiter)
		case *asciidoc.StemBlock:
			renderAttributes(cxt, el, el.Attributes(), false)
			renderDelimiter(cxt, el.Delimiter)
			for _, l := range el.Lines() {
				cxt.WriteString(l)
				cxt.WriteRune('\n')
			}
			renderDelimiter(cxt, el.Delimiter)
		case *asciidoc.OpenBlock:
			renderAttributes(cxt, el, el.Attributes(), false)
			renderDelimiter(cxt, el.Delimiter)
			Elements(cxt, "", el.Elements()...)
			renderDelimiter(cxt, el.Delimiter)
		case *asciidoc.FileInclude:
			cxt.WriteString("include::")
			Elements(cxt, "", el.Elements()...)
			attributes := el.Attributes()
			if len(attributes) == 0 {
				cxt.WriteString("[]\n")
			} else {
				renderAttributes(cxt, el, el.Attributes(), true)
				cxt.WriteRune('\n')
			}
		case *asciidoc.Anchor:
			cxt.WriteString("[[")
			cxt.WriteString(el.ID)
			anchorElements := el.Elements()
			if len(anchorElements) > 0 {
				cxt.WriteString(",")
				Elements(cxt, "", anchorElements...)
			}
			cxt.WriteString("]]")
		case *asciidoc.Admonition:
			renderAdmonition(cxt, el.AdmonitionType)
		case *asciidoc.AttachedBlock:
			cxt.WriteString("+\n")
			err = Elements(cxt, "", el.Child())
		case *asciidoc.LineBreak:
			cxt.WriteString("+")
		case *asciidoc.Counter:
			cxt.WriteString("{counter")
			if !el.Display {
				cxt.WriteRune('2')
			}
			cxt.WriteRune(':')
			cxt.WriteString(el.Name)
			if len(el.InitialValue) > 0 {
				cxt.WriteRune(':')
				cxt.WriteString(el.InitialValue)
			}
			cxt.WriteString("}")
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
