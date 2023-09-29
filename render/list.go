package render

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/output"
)

func renderList(cxt *output.Context, l *types.List) (err error) {
	renderAttributes(cxt, l, l.Attributes, false)
	switch l.Kind {
	case types.OrderedListKind:
		err = renderOrderedList(cxt, l)
	case types.UnorderedListKind:
		err = renderUnorderedList(cxt, l)
	case types.LabeledListKind:
		err = renderLabeledList(cxt, l)
	default:
		err = fmt.Errorf("unsupported list type: %s", l.Kind)
	}
	return
}

func renderOrderedList(cxt *output.Context, l *types.List) (err error) {
	cxt.OrderedListDepth++
	cxt.WriteNewline()
	for _, e := range l.Elements {
		switch el := e.(type) {
		case *types.OrderedListElement:
			renderAttributes(cxt, el, el.Attributes, false)
			err = RenderElements(cxt, strings.Repeat(".", cxt.OrderedListDepth)+" ", el.Elements)
			cxt.WriteRune('\n')
		default:
			err = fmt.Errorf("unknown ordered list element type: %T", el)
		}
		if err != nil {
			break
		}
	}
	cxt.OrderedListDepth--
	return
}

func renderUnorderedList(cxt *output.Context, l *types.List) (err error) {
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
			renderAttributes(cxt, el, el.Attributes, false)
			err = RenderElements(cxt, bullet+" ", el.Elements)
			cxt.WriteNewline()
		default:
			err = fmt.Errorf("unknown unordered list element type: %T", el)
		}
		if err != nil {
			return
		}
	}
	return
}

func renderLabeledList(cxt *output.Context, l *types.List) (err error) {
	cxt.WriteNewline()
	for _, e := range l.Elements {
		switch el := e.(type) {
		case *types.LabeledListElement:
			err := renderAttributes(cxt, el, el.Attributes, false)
			if err != nil {
				return err
			}
			err = RenderElements(cxt, "", el.Term)
			if err != nil {
				return err
			}
			cxt.WriteString(string(el.Style))
			cxt.WriteRune(' ')
			err = RenderElements(cxt, "", el.Elements)
			cxt.WriteNewline()
			if err != nil {
				return err
			}
		default:
			err = fmt.Errorf("unknown unordered list element type: %T", el)
		}
		if err != nil {
			return
		}
	}
	return
}
