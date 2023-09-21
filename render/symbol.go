package render

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/output"
)

func renderQuotedText(cxt *output.Context, qt *types.QuotedText) {
	renderAttributes(cxt, qt, qt.Attributes)
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
	cxt.WriteString(wrapper)
	RenderElements(cxt, "", qt.Elements)
	cxt.WriteString(wrapper)
}

func renderSpecialCharacter(cxt *output.Context, s *types.SpecialCharacter) {
	switch s.Name {
	case "<":
		cxt.WriteRune('<')
	case ">":
		cxt.WriteRune('>')
	case "&":
		cxt.WriteRune('&')
	default:
		panic(fmt.Errorf("unknown special character: %s", s.Name))
	}
}

func renderSymbol(cxt *output.Context, s *types.Symbol) {
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
	default:
		panic(fmt.Errorf("unknown symbol: \"%s\"", s.Name))
	}
}
