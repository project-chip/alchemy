package render

import "github.com/hasty/adoc/elements"

func renderParagraph(cxt *Context, p *elements.Paragraph, previous *any) (err error) {
	err = renderAttributes(cxt, p, p.Attributes(), false)
	if err != nil {
		return
	}
	renderAdmonition(cxt, p.Admonition)
	err = Elements(cxt, "", p.Elements()...)
	return
}

func renderAdmonition(cxt *Context, a elements.AdmonitionType) {
	switch a {
	case elements.AdmonitionTypeNote:
		cxt.WriteString("NOTE: ")
	case elements.AdmonitionTypeTip:
		cxt.WriteString("TIP: ")
	case elements.AdmonitionTypeImportant:
		cxt.WriteString("IMPORTANT: ")
	case elements.AdmonitionTypeCaution:
		cxt.WriteString("CAUTION: ")
	case elements.AdmonitionTypeWarning:
		cxt.WriteString("WARNING: ")
	}
}
