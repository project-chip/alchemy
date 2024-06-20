package render

import "github.com/hasty/alchemy/asciidoc"

func renderAnchor(cxt *Context, el *asciidoc.Anchor) (err error) {
	cxt.WriteString("[[")
	cxt.WriteString(el.ID)
	anchorElements := el.Elements()
	if len(anchorElements) > 0 {
		cxt.WriteString(",")
		err = Elements(cxt, "", anchorElements...)
	}
	cxt.WriteString("]]")
	return
}
