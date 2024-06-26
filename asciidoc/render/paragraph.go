package render

import "github.com/project-chip/alchemy/asciidoc"

func renderParagraph(cxt *Context, p *asciidoc.Paragraph) (err error) {
	err = renderAttributes(cxt, p.Attributes(), false)
	if err != nil {
		return
	}
	renderAdmonition(cxt, p.Admonition)
	err = Elements(cxt, "", p.Elements()...)
	return
}

func renderAdmonition(cxt *Context, a asciidoc.AdmonitionType) {
	switch a {
	case asciidoc.AdmonitionTypeNote:
		cxt.WriteString("NOTE: ")
	case asciidoc.AdmonitionTypeTip:
		cxt.WriteString("TIP: ")
	case asciidoc.AdmonitionTypeImportant:
		cxt.WriteString("IMPORTANT: ")
	case asciidoc.AdmonitionTypeCaution:
		cxt.WriteString("CAUTION: ")
	case asciidoc.AdmonitionTypeWarning:
		cxt.WriteString("WARNING: ")
	}
}
