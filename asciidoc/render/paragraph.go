package render

import "github.com/project-chip/alchemy/asciidoc"

func renderParagraph(cxt Target, p *asciidoc.Paragraph) (err error) {
	cxt.FlushWrap()
	err = renderAttributes(cxt, p.Attributes(), attributeRenderTypeBlock)
	if err != nil {
		return
	}
	renderAdmonition(cxt, p.Admonition)
	err = Elements(cxt, "", p.Children()...)
	return
}

func renderAdmonition(cxt Target, a asciidoc.AdmonitionType) {
	cxt.StartBlock()
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
	cxt.EndBlock()
}
