package render

import (
	"github.com/hasty/adoc/elements"
)

func renderOrderedListElement(cxt *Context, el *elements.OrderedListItem) (err error) {

	cxt.WriteNewline()

	err = renderAttributes(cxt, el, el.Attributes(), false)
	if err != nil {
		return
	}
	cxt.WriteString(el.Marker)
	cxt.WriteString(" ")
	err = Elements(cxt, "", el.Elements()...)
	return
}

func renderUnorderedListElement(cxt *Context, el *elements.UnorderedListItem) (err error) {
	cxt.WriteNewline()

	err = renderAttributes(cxt, el, el.Attributes(), false)
	if err != nil {
		return
	}
	cxt.WriteString(el.Marker)
	cxt.WriteString(" ")
	err = Elements(cxt, "", el.Elements()...)
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
