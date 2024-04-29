package render

import "github.com/hasty/adoc/elements"

func renderParagraph(cxt *Context, p *elements.Paragraph, previous *any) (err error) {
	err = renderAttributes(cxt, p, p.Attributes(), false)
	if err != nil {
		return
	}
	err = Elements(cxt, "", p.Elements()...)
	return
}
