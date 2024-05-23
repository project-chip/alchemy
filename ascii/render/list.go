package render

import "github.com/hasty/adoc/asciidoc"

func renderOrderedListElement(cxt *Context, el *asciidoc.OrderedListItem) (err error) {

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

func renderUnorderedListElement(cxt *Context, el *asciidoc.UnorderedListItem) (err error) {
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

func renderLabeledListElement(cxt *Context, el *asciidoc.DescriptionListItem) error {
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
