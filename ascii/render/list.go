package render

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func renderList(cxt *Context, l *types.List) (err error) {
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

func renderOrderedList(cxt *Context, l *types.List) (err error) {
	for _, e := range l.Elements {
		switch el := e.(type) {
		case *types.OrderedListElement:
			err = renderOrderedListElement(cxt, el)
		default:
			err = fmt.Errorf("unknown ordered list element type: %T", el)
		}
		if err != nil {
			break
		}
	}
	return
}

func renderOrderedListElement(cxt *Context, el *types.OrderedListElement) (err error) {
	cxt.WriteNewline()
	var bullet string
	switch el.Style {
	case types.Arabic:
		bullet = "."
	case types.LowerAlpha:
		bullet = ".."
	case types.LowerRoman:
		bullet = "..."
	case types.UpperAlpha:
		bullet = "...."
	case types.UpperRoman:
		bullet = "....."
	}
	renderAttributes(cxt, el, el.Attributes, false)
	err = RenderElements(cxt, bullet+" ", el.Elements)
	return
}

func renderUnorderedList(cxt *Context, l *types.List) (err error) {
	cxt.WriteNewline()
	for _, e := range l.Elements {
		switch el := e.(type) {
		case *types.UnorderedListElement:
			err = renderUnorderedListElement(cxt, el)
		default:
			err = fmt.Errorf("unknown unordered list element type: %T", el)
		}
		if err != nil {
			return
		}
	}
	return
}

func renderUnorderedListElement(cxt *Context, el *types.UnorderedListElement) (err error) {
	cxt.WriteNewline()
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
	return
}

func renderLabeledList(cxt *Context, l *types.List) (err error) {
	for _, e := range l.Elements {
		switch el := e.(type) {
		case *types.LabeledListElement:
			err = renderLabeledListElement(cxt, el)
			if err != nil {
				return
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

func renderLabeledListElement(cxt *Context, el *types.LabeledListElement) error {
	cxt.WriteNewline()
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
	if err != nil {
		return err
	}
	return nil
}

func renderListElements(cxt *Context, les *types.ListElements) (err error) {
	for _, le := range les.Elements {
		switch el := le.(type) {
		case *types.OrderedListElement:
			err = renderOrderedListElement(cxt, el)
		case *types.UnorderedListElement:
			err = renderUnorderedListElement(cxt, el)
		case *types.ListContinuation:
			cxt.WriteNewline()
			cxt.WriteString("+\n")
			err = RenderElements(cxt, "", []interface{}{el.Element})
		case *types.LabeledListElement:
			err = renderLabeledListElement(cxt, el)
		default:
			err = fmt.Errorf("unexpected list element: %T", le)
		}
		if err != nil {
			return
		}
	}
	return
}
