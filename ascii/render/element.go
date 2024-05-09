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
		case elements.EmptyLine:
			cxt.WriteNewline()
			cxt.WriteRune('\n')
		case *elements.NewLine:
			cxt.WriteRune('\n')
		case *elements.Section:
			err = renderSection(cxt, el)
			if err == nil {
				err = Elements(cxt, "", el.Elements()...)
			}
		case *elements.Paragraph:
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
		case *elements.AttributeEntry:
			err = renderAttributeEntry(cxt, el)
		case *elements.String:
			text := el.Value
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
			err = renderLink(cxt, el)
		case elements.SpecialCharacter:
			err = renderSpecialCharacter(cxt, el)
		case *elements.Bold:
			err = renderFormattedText(cxt, el, "*")
		case *elements.DoubleBold:
			err = renderFormattedText(cxt, el, "**")
		case *elements.Monospace:
			err = renderFormattedText(cxt, el, "`")
		case *elements.DoubleMonospace:
			err = renderFormattedText(cxt, el, "``")
		case *elements.Superscript:
			err = renderFormattedText(cxt, el, "^")
		case *elements.Subscript:
			err = renderFormattedText(cxt, el, "~")
		case *elements.Italic:
			err = renderFormattedText(cxt, el, "_")
		case *elements.DoubleItalic:
			err = renderFormattedText(cxt, el, "__")
		case *elements.Marked:
			err = renderFormattedText(cxt, el, "#")
		case *elements.DoubleMarked:
			err = renderFormattedText(cxt, el, "##")
		case *elements.LineContinuation:
			cxt.WriteString(" +")
		case elements.AttributeReference:
			cxt.WriteString(fmt.Sprintf("{%s}", el.Name()))
		case *elements.InlineImage:
			err = renderInlineImage(cxt, el)
		case elements.SingleLineComment:
			cxt.WriteString("//")
			cxt.WriteString(el.Value)
			cxt.WriteNewline()
		//case *elements.FootnoteReference:
		//	err = renderFootnoteReference(cxt, el)
		/*case *elements.InlinePassthrough:
		switch el.Kind {
		case elements.SinglePlusPassthrough, elements.TriplePlusPassthrough:
			cxt.WriteString(string(el.Kind))
			err = Elements(cxt, "", el.Elements)
			cxt.WriteString(string(el.Kind))
		case elements.PassthroughMacro:
			cxt.WriteString("pass:[")
			err = Elements(cxt, "", el.Elements)
			cxt.WriteRune(']')
		}*/
		case *elements.InlinePassthrough:
			cxt.WriteString("+")
			err = Elements(cxt, "", el.Elements()...)
			cxt.WriteString("+")
		case *elements.InlineDoublePassthrough:
			cxt.WriteString("++")
			err = Elements(cxt, "", el.Elements()...)
			cxt.WriteString("++")
		case *elements.AttributeReset:
			renderAttributeReset(cxt, el)
		case *elements.UnorderedListItem:
			err = renderUnorderedListElement(cxt, el)
		case *elements.OrderedListItem:
			err = renderOrderedListElement(cxt, el)
		case *elements.ListContinuation:
			cxt.WriteNewline()
			cxt.WriteString("+\n")
			err = Elements(cxt, "", el.Child())
		case *elements.IfDef:
			cxt.WriteString("ifdef::")
			for i, a := range el.Attributes {
				if i > 0 {
					cxt.WriteString(",")
				}
				cxt.WriteString(string(a))
			}
			cxt.WriteString("[]\n")
		case *elements.IfNDef:
			cxt.WriteString("ifndef::")
			for i, a := range el.Attributes {
				if i > 0 {
					cxt.WriteString(",")
				}
				cxt.WriteString(string(a))
			}
			cxt.WriteString("[]\n")
		case *elements.IfEval:
			cxt.WriteString("ifeval::[")
			cxt.WriteString(el.Left)
			cxt.WriteRune(' ')
			cxt.WriteString(el.Operator.String())
			cxt.WriteRune(' ')
			cxt.WriteString(el.Right)
			cxt.WriteString("]\n")
		case *elements.EndIf:
			cxt.WriteString("endif::")
			for i, a := range el.Attributes {
				if i > 0 {
					cxt.WriteString(",")
				}
				cxt.WriteString(string(a))
			}
			cxt.WriteString("[]\n")
		case *elements.MultiLineComment:
			renderDelimiter(cxt, el.Delimiter)
			for _, l := range el.Lines() {
				cxt.WriteString(l)
				cxt.WriteRune('\n')
			}
			renderDelimiter(cxt, el.Delimiter)
		case *elements.DescriptionListItem:
			renderAttributes(cxt, el, el.Attributes(), false)
			Elements(cxt, "", el.Term...)
			cxt.WriteString(el.Marker)
			cxt.WriteRune(' ')
			Elements(cxt, "", el.Elements()...)
			cxt.WriteNewline()
		case *elements.LiteralBlock:
			renderAttributes(cxt, el, el.Attributes(), false)
			renderDelimiter(cxt, el.Delimiter)
			for _, l := range el.Lines() {
				cxt.WriteString(l)
				cxt.WriteRune('\n')
			}
			renderDelimiter(cxt, el.Delimiter)
		case *elements.SidebarBlock:
			renderAttributes(cxt, el, el.Attributes(), false)
			renderDelimiter(cxt, el.Delimiter)
			Elements(cxt, "", el.Elements()...)
			renderDelimiter(cxt, el.Delimiter)
		case *elements.Listing:
			renderAttributes(cxt, el, el.Attributes(), false)
			renderDelimiter(cxt, el.Delimiter)
			for _, l := range el.Lines() {
				cxt.WriteString(l)
				cxt.WriteRune('\n')
			}
			renderDelimiter(cxt, el.Delimiter)
		case *elements.ExampleBlock:
			renderAttributes(cxt, el, el.Attributes(), false)
			renderDelimiter(cxt, el.Delimiter)
			Elements(cxt, "", el.Elements()...)
			renderDelimiter(cxt, el.Delimiter)
		case *elements.PassthroughBlock:
			renderAttributes(cxt, el, el.Attributes(), false)
			renderDelimiter(cxt, el.Delimiter)
			for _, l := range el.Lines() {
				cxt.WriteString(l)
				cxt.WriteRune('\n')
			}
			renderDelimiter(cxt, el.Delimiter)
		case *elements.OpenBlock:
			renderAttributes(cxt, el, el.Attributes(), false)
			renderDelimiter(cxt, el.Delimiter)
			Elements(cxt, "", el.Elements()...)
			renderDelimiter(cxt, el.Delimiter)
		case *elements.FileInclude:
			cxt.WriteString("include::")
			Elements(cxt, "", el.Elements()...)
			attributes := el.Attributes()
			if len(attributes) == 0 {
				cxt.WriteString("[]\n")
			} else {
				renderAttributes(cxt, el, el.Attributes(), true)
			}
		case *elements.Anchor:
			cxt.WriteString("[[")
			cxt.WriteString(el.ID)
			anchorElements := el.Elements()
			if len(anchorElements) > 0 {
				cxt.WriteString(", ")
				Elements(cxt, "", anchorElements...)
			}
			cxt.WriteString("]]")
		case *elements.Admonition:
			renderAdmonition(cxt, el.AdmonitionType)
		case *elements.AttachedBlock:
			cxt.WriteString("+\n")
			err = Elements(cxt, "", el.Child())
		case *elements.LineBreak:
			cxt.WriteString("+")
		case *elements.Counter:
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
