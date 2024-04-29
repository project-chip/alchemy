package render

import (
	"github.com/hasty/adoc/elements"
)

/*func renderQuotedText(cxt *Context, qt *elements.QuotedText) (err error) {
	err = renderAttributes(cxt, qt, qt.Attributes, true)
	if err != nil {
		return
	}
	var wrapper string
	switch qt.Kind {
	case elements.SingleQuoteBold:
		wrapper = "*"
	case elements.DoubleQuoteBold:
		wrapper = "**"
	case elements.SingleQuoteMonospace:
		wrapper = "`"
	case elements.DoubleQuoteMonospace:
		wrapper = "``"
	case elements.SingleQuoteSuperscript:
		wrapper = "^"
	case elements.SingleQuoteSubscript:
		wrapper = "~"
	case elements.SingleQuoteItalic:
		wrapper = "_"
	case elements.DoubleQuoteItalic:
		wrapper = "__"
	case elements.SingleQuoteMarked:
		wrapper = "#"
	case elements.DoubleQuoteMarked:
		wrapper = "##"
	default:
		err = fmt.Errorf("unsupported quoted text kind: %s", qt.Kind)
		return
	}
	cxt.WriteString(wrapper)
	err = Elements(cxt, "", qt.Elements)
	cxt.WriteString(wrapper)
	return
}*/

func renderFormattedText(cxt *Context, el elements.BlockElement, wrapper string) (err error) {
	err = renderAttributes(cxt, el, el.Attributes(), true)
	if err != nil {
		return
	}
	cxt.WriteString(wrapper)
	err = Elements(cxt, "", el.Elements()...)
	cxt.WriteString(wrapper)
	return
}

func renderSpecialCharacter(cxt *Context, s elements.SpecialCharacter) error {
	cxt.WriteString(s.Character)
	return nil
}

/*
func renderSymbol(cxt *Context, s *elements.Symbol) error {
	switch s.Name {
	case "'":
		cxt.WriteRune('\'')
	case "=>":
		cxt.WriteString("=>")
	case "->":
		cxt.WriteString("->")
	case "<=":
		cxt.WriteString("<=")
	case "<-":
		cxt.WriteString("<-")
	case "...":
		cxt.WriteString("...")
	case " -- ":
		cxt.WriteString(" -- ")
	case "(C)":
		cxt.WriteString("(C)")
	case "`\"":
		cxt.WriteString("`\"")
	case "\"`":
		cxt.WriteString("\"`")
	case "`'":
		cxt.WriteString("`'")
	case "(TM)":
		cxt.WriteString("(TM)")
	default:
		return fmt.Errorf("unknown symbol: \"%s\"", s.Name)
	}
	return nil
}
*/
