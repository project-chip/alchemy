package render

import "github.com/project-chip/alchemy/asciidoc"

func renderOrderedListElement(cxt Target, el *asciidoc.OrderedListItem) (err error) {

	cxt.FlushWrap()
	cxt.DisableWrap()
	cxt.EnsureNewLine()

	err = renderAttributes(cxt, el.Attributes(), false)
	if err != nil {
		return
	}
	cxt.WriteString(el.Indent)
	cxt.WriteString(el.Marker)
	cxt.WriteString(" ")
	cxt.EnableWrap()
	err = Elements(cxt, "", el.Elements()...)
	return
}

func renderUnorderedListElement(cxt Target, el *asciidoc.UnorderedListItem) (err error) {
	cxt.FlushWrap()
	cxt.DisableWrap()
	cxt.EnsureNewLine()

	err = renderAttributes(cxt, el.Attributes(), false)
	if err != nil {
		return
	}
	cxt.WriteString(el.Indent)
	cxt.WriteString(el.Marker)
	cxt.WriteString(" ")
	cxt.EnableWrap()
	err = Elements(cxt, "", el.Elements()...)
	return
}

func renderDescriptionListItem(cxt Target, el *asciidoc.DescriptionListItem) (err error) {
	cxt.FlushWrap()
	renderAttributes(cxt, el.Attributes(), false)
	cxt.DisableWrap()
	err = Elements(cxt, "", el.Term...)
	if err != nil {
		return
	}
	cxt.WriteString(el.Marker)
	cxt.WriteRune(' ')
	cxt.EnableWrap()
	err = Elements(cxt, "", el.Elements()...)
	if err != nil {
		return
	}
	cxt.EnsureNewLine()
	return
}
