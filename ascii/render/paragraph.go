package render

import "github.com/hasty/adoc/elements"

func renderParagraph(cxt *Context, p *elements.Paragraph, previous *any) (err error) {
	err = renderAttributes(cxt, p, p.Attributes(), false)
	if err != nil {
		return
	}
	switch p.Admonition {
	case elements.AdmonitionNote:
		cxt.WriteString("NOTE: ")
	case elements.AdmonitionTip:
		cxt.WriteString("TIP: ")
	case elements.AdmonitionImportant:
		cxt.WriteString("IMPORTANT: ")
	case elements.AdmonitionCaution:
		cxt.WriteString("CAUTION: ")
	case elements.AdmonitionWarning:
		cxt.WriteString("WARNING: ")
	}
	err = Elements(cxt, "", p.Elements()...)
	return
}
