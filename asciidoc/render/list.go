package render

import "github.com/project-chip/alchemy/asciidoc"

func renderOrderedListElement(cxt *Context, el *asciidoc.OrderedListItem) (err error) {

	cxt.EnsureNewLine()

	err = renderAttributes(cxt, el.Attributes(), false)
	if err != nil {
		return
	}
	cxt.WriteString(el.Indent)
	cxt.WriteString(el.Marker)
	cxt.WriteString(" ")
	err = Elements(cxt, "", el.Elements()...)
	return
}

func renderUnorderedListElement(cxt *Context, el *asciidoc.UnorderedListItem) (err error) {
	cxt.EnsureNewLine()

	err = renderAttributes(cxt, el.Attributes(), false)
	if err != nil {
		return
	}
	cxt.WriteString(el.Indent)
	cxt.WriteString(el.Marker)
	cxt.WriteString(" ")
	err = Elements(cxt, "", el.Elements()...)
	return
}

func renderLabeledListElement(cxt *Context, el *asciidoc.DescriptionListItem) error {
	cxt.EnsureNewLine()
	err := renderAttributes(cxt, el.Attributes(), false)
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

func renderDescriptionListItem(cxt *Context, el *asciidoc.DescriptionListItem) {
	renderAttributes(cxt, el.Attributes(), false)
	Elements(cxt, "", el.Term...)
	cxt.WriteString(el.Marker)
	cxt.WriteRune(' ')
	Elements(cxt, "", el.Elements()...)
	cxt.EnsureNewLine()
}
