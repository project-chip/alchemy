package main

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func (d *doc) render() string {
	var out output
	d.renderElements("", d.root.elements, &out)
	out.WriteNewline()
	return out.String()
}

func (d *doc) renderSection(s *types.Section, out *output) {
	d.renderAttributes(s, s.Attributes, out)
	//	out.WriteRune('\n')
	out.WriteString(strings.Repeat("=", s.Level+1))
	out.WriteRune(' ')
	out.WriteString(getSectionTitle(s))
	out.WriteRune('\n')
}

func (d *doc) renderElements(prefix string, elements []interface{}, out *output) {
	var previous interface{}
	for _, e := range elements {
		if section, ok := e.(*section); ok {
			d.renderSection(section.base, out)
			d.renderElements("", section.elements, out)
			continue
		}
		if el, ok := e.(*element); ok {
			e = el.base
		}
		switch el := e.(type) {
		case *types.DelimitedBlock:
			switch el.Kind {
			case "comment":
				d.renderComment(el, out)
			case "sidebar":
				d.renderSidebar(el, out)
			case "example":
				d.renderExample(el, out)
			default:
				fmt.Printf("unknown block kind: %s\n", el.Kind)
			}
		case *types.Paragraph:
			out.WriteString(prefix)
			d.renderParagraph(el, out, &previous)
			continue
		case *types.Table:
			d.renderTable(el, out)
		case *types.BlankLine:
			switch previous.(type) {
			case *types.StringElement:
				out.WriteRune('\n')
			default:

			}
			out.WriteRune('\n')
		case *types.InternalCrossReference:
			d.renderInternalCrossReference(el, out)
		case *types.List:
			d.renderList(el, out)
		case *types.AttributeDeclaration:
			out.WriteString(el.RawText())
		case *types.StringElement:
			text, _ := el.RawText()
			out.WriteString(text)
		case *types.SinglelineComment:
			out.WriteString("//")
			out.WriteString(el.Content)
			out.WriteNewline()
		case *types.ImageBlock:
			text := el.Location.Scheme
			out.WriteString(text)
			switch p := el.Location.Path.(type) {
			case string:
				out.WriteString(p)
			default:
				fmt.Printf("unknown image location path type: %T\n", p)
			}
			d.renderAttributes(el, el.Attributes, out)
		default:
			fmt.Printf("unknown element type: %T\n", el)
		}
		previous = e
	}
}

func (d *doc) renderParagraph(p *types.Paragraph, out *output, previous *interface{}) {
	d.renderAttributes(p, p.Attributes, out)

	for _, e := range p.Elements {
		switch el := e.(type) {
		case *types.StringElement:
			text, _ := el.RawText()
			out.WriteString(text)
		case *types.InternalCrossReference:
			d.renderInternalCrossReference(el, out)
		case *types.Symbol:
			d.renderSymbol(el, out)
		case *types.SpecialCharacter:
			d.renderSpecialCharacter(el, out)
		case *types.QuotedText:
			d.renderQuotedText(el, out)
		case *types.InlineLink:
			text := el.Location.Scheme
			out.WriteString(text)
			switch p := el.Location.Path.(type) {
			case string:
				out.WriteString(p)
			default:
				fmt.Printf("unknown location path type: %T\n", p)
			}
		case *types.LineBreak:
			out.WriteString(" +")
		case *types.PredefinedAttribute:
			out.WriteString(fmt.Sprintf("{%s}", el.Name))
		default:
			fmt.Printf("unknown paragraph element type: %T\n", el)
		}
		*previous = e
	}
}

func (d *doc) renderQuotedText(qt *types.QuotedText, out *output) {
	d.renderAttributes(qt, qt.Attributes, out)
	var wrapper string
	switch qt.Kind {
	case types.SingleQuoteBold:
		wrapper = "*"
	case types.DoubleQuoteBold:
		wrapper = "**"
	case types.SingleQuoteMonospace:
		wrapper = "`"
	case types.SingleQuoteSuperscript:
		wrapper = "^"
	case types.SingleQuoteSubscript:
		wrapper = "~"
	case types.SingleQuoteItalic:
		wrapper = "_"
	case types.DoubleQuoteItalic:
		wrapper = "__"
	default:
		fmt.Printf("unsupported quoted text kind: %s\n", qt.Kind)
		return
	}
	out.WriteString(wrapper)
	d.renderElements("", qt.Elements, out)
	out.WriteString(wrapper)
}

func (d *doc) renderSpecialCharacter(s *types.SpecialCharacter, out *output) {
	switch s.Name {
	case "<":
		out.WriteRune('<')
	case ">":
		out.WriteRune('>')
	case "&":
		out.WriteRune('&')
	default:
		fmt.Printf("unknown special character: %s\n", s.Name)
	}
}

func (d *doc) renderSymbol(s *types.Symbol, out *output) {
	switch s.Name {
	case "'":
		out.WriteRune('\'')
	case "=>":
		out.WriteString("=>")
	case "->":
		out.WriteString("->")
	case "<=":
		out.WriteString("<=")
	case "<-":
		out.WriteString("<-")
	default:
		fmt.Printf("unknown symbol: %s\n", s.Name)
	}
}

func (d *doc) renderInternalCrossReference(cf *types.InternalCrossReference, out *output) {
	fmt.Printf("icf ID type: %T -> %v\n", cf.ID, cf.ID)
	fmt.Printf("icf Label type: %T -> %v\n", cf.Label, cf.Label)

	switch el := cf.ID.(type) {
	case string:
		id := el
		//fmt.Printf("id %s\n", id)
		ref, ok := d.base.ElementReferences[id]
		if ok {
			fmt.Printf("id %s => icf ref type: %T\n", id, ref)
			switch idref := ref.(type) {
			case []interface{}:
				for _, i := range idref {
					fmt.Printf("\ticf ref child type: %T\n", i)
					switch iv := i.(type) {
					case *types.StringElement:
						t, _ := iv.RawText()

						//fmt.Printf("val: %s\n", t)
						el = t
					}
				}
			}
		}
		out.WriteString("<<")
		out.WriteString(el)
		out.WriteString(">>")
	default:
		fmt.Printf("unknown internal cross reference ID type: %T\n", el)
	}
}

func (d *doc) renderComment(comment *types.DelimitedBlock, out *output) {
	for _, e := range comment.Elements {
		switch el := e.(type) {
		case *types.StringElement:
			out.WriteRune('\n')
			out.WriteString("////")
			out.WriteNewline()
			text, _ := el.RawText()
			out.WriteString(text)
			out.WriteNewline()
			out.WriteString("////")
			out.WriteNewline()
		default:
			fmt.Printf("unknown comment element type: %T\n", el)
		}
	}
}

func (d *doc) renderExample(comment *types.DelimitedBlock, out *output) {
	d.renderAttributes(comment, comment.Attributes, out)
	out.WriteNewline()
	out.WriteString("====")
	out.WriteNewline()
	d.renderElements("", comment.Elements, out)
	out.WriteNewline()
	out.WriteString("====")
	out.WriteNewline()

}

func (d *doc) renderSidebar(comment *types.DelimitedBlock, out *output) {
	var previous interface{}
	for _, e := range comment.Elements {
		switch el := e.(type) {
		case *types.Paragraph:
			out.WriteNewline()
			out.WriteString("****")
			out.WriteNewline()
			d.renderParagraph(el, out, &previous)
			out.WriteNewline()
			out.WriteString("****")
			out.WriteNewline()
		case *types.StringElement:
			out.WriteRune('\n')
			out.WriteString("****")
			out.WriteNewline()
			text, _ := el.RawText()
			out.WriteString(text)
			out.WriteNewline()
			out.WriteString("****")
			out.WriteNewline()
		default:
			fmt.Printf("unknown sidebar element type: %T\n", el)
		}
	}
}

func (d *doc) renderList(l *types.List, out *output) {
	d.renderAttributes(l, l.Attributes, out)
	switch l.Kind {
	case types.OrderedListKind:
		d.renderOrderedList(l, out)
	case types.UnorderedListKind:
		d.renderUnorderedList(l, out)
	default:
		fmt.Printf("unsupported list type: %s", l.Kind)
	}

}

func (d *doc) renderOrderedList(l *types.List, out *output) {
	d.orderedListDepth++
	for _, e := range l.Elements {
		switch el := e.(type) {
		case *types.OrderedListElement:
			d.renderAttributes(el, el.Attributes, out)
			d.renderElements(strings.Repeat(".", d.orderedListDepth)+" ", el.Elements, out)
			out.WriteRune('\n')
		default:
			fmt.Printf("unknown ordered list element type: %T\n", el)
		}
	}
	d.orderedListDepth--
}

func (d *doc) renderUnorderedList(l *types.List, out *output) {
	out.WriteNewline()
	for _, e := range l.Elements {
		switch el := e.(type) {
		case *types.UnorderedListElement:
			var bullet string
			switch el.BulletStyle {
			case types.Dash:
				bullet = "-"
			case types.OneAsterisk:
				bullet = "*"
			case types.TwoAsterisks:
				bullet = "**"
			case types.ThreeAsterisks:
				bullet = "***"
			case types.FourAsterisks:
				bullet = "****"
			case types.FiveAsterisks:
				bullet = "*****"
			}
			d.renderAttributes(el, el.Attributes, out)
			d.renderElements(bullet+" ", el.Elements, out)
			out.WriteNewline()
		default:
			fmt.Printf("unknown unordered list element type: %T\n", el)
		}
	}
}
