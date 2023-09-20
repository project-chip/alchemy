package render

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/output"
)

func renderList(cxt *output.Context, l *types.List) {
	renderAttributes(cxt, l, l.Attributes)
	switch l.Kind {
	case types.OrderedListKind:
		renderOrderedList(cxt, l)
	case types.UnorderedListKind:
		renderUnorderedList(cxt, l)
	default:
		fmt.Printf("unsupported list type: %s", l.Kind)
	}

}

func renderOrderedList(cxt *output.Context, l *types.List) {
	cxt.OrderedListDepth++
	for _, e := range l.Elements {
		switch el := e.(type) {
		case *types.OrderedListElement:
			renderAttributes(cxt, el, el.Attributes)
			RenderElements(cxt, strings.Repeat(".", cxt.OrderedListDepth)+" ", el.Elements)
			cxt.WriteRune('\n')
		default:
			fmt.Printf("unknown ordered list element type: %T\n", el)
		}
	}
	cxt.OrderedListDepth--
}

func renderUnorderedList(cxt *output.Context, l *types.List) {
	cxt.WriteNewline()
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
			renderAttributes(cxt, el, el.Attributes)
			RenderElements(cxt, bullet+" ", el.Elements)
			cxt.WriteNewline()
		default:
			fmt.Printf("unknown unordered list element type: %T\n", el)
		}
	}
}
