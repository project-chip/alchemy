package render

import "github.com/project-chip/alchemy/asciidoc"

func renderSingleLineComment(cxt Target, el *asciidoc.SingleLineComment) {
	cxt.FlushWrap()
	cxt.EnsureNewLine()
	cxt.DisableWrap()
	cxt.WriteString("//")
	cxt.WriteString(el.Value)
	cxt.EnsureNewLine()
	cxt.EnableWrap()
}
