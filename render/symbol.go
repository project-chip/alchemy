package render

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/output"
)

func renderQuotedText(cxt *output.Context, qt *types.QuotedText) (err error) {
	renderAttributes(cxt, qt, qt.Attributes, true)
	var wrapper string
	switch qt.Kind {
	case types.SingleQuoteBold:
		wrapper = "*"
	case types.DoubleQuoteBold:
		wrapper = "**"
	case types.SingleQuoteMonospace:
		wrapper = "`"
	case types.DoubleQuoteMonospace:
		wrapper = "``"
	case types.SingleQuoteSuperscript:
		wrapper = "^"
	case types.SingleQuoteSubscript:
		wrapper = "~"
	case types.SingleQuoteItalic:
		wrapper = "_"
	case types.DoubleQuoteItalic:
		wrapper = "__"
	case types.SingleQuoteMarked:
		wrapper = "#"
	case types.DoubleQuoteMarked:
		wrapper = "##"
	default:
		err = fmt.Errorf("unsupported quoted text kind: %s", qt.Kind)
		return
	}
	cxt.WriteString(wrapper)
	err = RenderElements(cxt, "", qt.Elements)
	cxt.WriteString(wrapper)
	return
}

func renderSpecialCharacter(cxt *output.Context, s *types.SpecialCharacter) error {
	switch s.Name {
	case "<":
		cxt.WriteRune('<')
	case ">":
		cxt.WriteRune('>')
	case "&":
		cxt.WriteRune('&')
	default:
		return fmt.Errorf("unknown special character: %s", s.Name)
	}
	return nil
}

func renderSymbol(cxt *output.Context, s *types.Symbol) error {
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
	default:
		return fmt.Errorf("unknown symbol: \"%s\"", s.Name)
	}
	return nil
}
