package render

import (
	"fmt"
)

func renderList(cxt *Context, l *elements.List) (err error) {
	err = renderAttributes(cxt, l, l.Attributes, false)
	if err != nil {
		return
	}
	switch l.Kind {
	case elements.OrderedListKind:
		err = renderOrderedList(cxt, l)
	case elements.UnorderedListKind:
		err = renderUnorderedList(cxt, l)
	case elements.LabeledListKind:
		err = renderLabeledList(cxt, l)
	default:
		err = fmt.Errorf("unsupported list type: %s", l.Kind)
	}
	return
}

func renderOrderedList(cxt *Context, l *elements.List) (err error) {
	for _, e := range l.Elements {
		switch el := e.(type) {
		case *elements.OrderedListElement:
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

func renderOrderedListElement(cxt *Context, el *elements.OrderedListElement) (err error) {
	cxt.WriteNewline()
	var bullet string
	switch el.Style {
	case elements.Arabic:
		bullet = "."
	case elements.LowerAlpha:
		bullet = ".."
	case elements.LowerRoman:
		bullet = "..."
	case elements.UpperAlpha:
		bullet = "...."
	case elements.UpperRoman:
		bullet = "....."
	}
	err = renderAttributes(cxt, el, el.Attributes, false)
	if err != nil {
		return
	}
	err = Elements(cxt, bullet+" ", el.Elements)
	return
}

func renderUnorderedList(cxt *Context, l *elements.List) (err error) {
	cxt.WriteNewline()
	for _, e := range l.Elements {
		switch el := e.(type) {
		case *elements.UnorderedListElement:
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

func renderUnorderedListElement(cxt *Context, el *elements.UnorderedListElement) (err error) {
	cxt.WriteNewline()
	var bullet string
	switch el.BulletStyle {
	case elements.Dash:
		bullet = "-"
	case elements.OneAsterisk:
		bullet = "*"
	case elements.TwoAsterisks:
		bullet = "**"
	case elements.ThreeAsterisks:
		bullet = "***"
	case elements.FourAsterisks:
		bullet = "****"
	case elements.FiveAsterisks:
		bullet = "*****"
	}
	err = renderAttributes(cxt, el, el.Attributes, false)
	if err != nil {
		return
	}
	err = Elements(cxt, bullet+" ", el.Elements)
	return
}

func renderLabeledList(cxt *Context, l *elements.List) (err error) {
	for _, e := range l.Elements {
		switch el := e.(type) {
		case *elements.LabeledListElement:
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

func renderLabeledListElement(cxt *Context, el *elements.LabeledListElement) error {
	cxt.WriteNewline()
	err := renderAttributes(cxt, el, el.Attributes, false)
	if err != nil {
		return err
	}
	err = Elements(cxt, "", el.Term)
	if err != nil {
		return err
	}
	cxt.WriteString(string(el.Style))
	cxt.WriteRune(' ')
	err = Elements(cxt, "", el.Elements)
	if err != nil {
		return err
	}
	return nil
}

func renderListElements(cxt *Context, les *elements.ListElements) (err error) {
	for _, le := range les.Elements {
		switch el := le.(type) {
		case *elements.OrderedListElement:
			err = renderOrderedListElement(cxt, el)
		case *elements.UnorderedListElement:
			err = renderUnorderedListElement(cxt, el)
		case *elements.ListContinuation:
			cxt.WriteNewline()
			cxt.WriteString("+\n")
			err = Elements(cxt, "", []any{el.Element})
		case *elements.LabeledListElement:
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
