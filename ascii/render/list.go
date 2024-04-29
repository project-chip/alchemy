package render

import (
	"github.com/hasty/adoc/elements"
)

func renderOrderedListElement(cxt *Context, el *elements.OrderedListItem) (err error) {
	err = renderAttributes(cxt, el, el.Attributes(), false)

	cxt.WriteNewline()

	err = renderAttributes(cxt, el, el.Attributes(), false)
	if err != nil {
		return
	}
	err = Elements(cxt, el.Marker+" ", el.Elements()...)
	return
}

func renderUnorderedListElement(cxt *Context, el *elements.UnorderedListItem) (err error) {
	cxt.WriteNewline()

	err = renderAttributes(cxt, el, el.Attributes(), false)
	if err != nil {
		return
	}
	err = Elements(cxt, el.Marker+" ", el.Elements()...)
	return
}

func renderLabeledListElement(cxt *Context, el *elements.DescriptionListItem) error {
	cxt.WriteNewline()
	err := renderAttributes(cxt, el, el.Attributes(), false)
	if err != nil {
		return err
	}
	err = Elements(cxt, "", el.Term...)
	if err != nil {
		return err
	}
	cxt.WriteString(string(el.Marker))
	cxt.WriteRune(' ')
	err = Elements(cxt, "", el.Elements()...)
	if err != nil {
		return err
	}
	return nil
}
