package render

import "github.com/project-chip/alchemy/asciidoc"

func renderPageBreak(cxt Target, el *asciidoc.PageBreak) {
	cxt.FlushWrap()
	cxt.DisableWrap()
	cxt.EnsureNewLine()
	cxt.WriteString("<<<\n")
	cxt.EnableWrap()
}

func renderThematicBreak(cxt Target, el *asciidoc.ThematicBreak) {
	cxt.FlushWrap()
	cxt.DisableWrap()
	cxt.EnsureNewLine()
	cxt.WriteString(el.Text)
	cxt.EnsureNewLine()
	cxt.EnableWrap()
}
