package render

import "github.com/project-chip/alchemy/asciidoc"

func renderSingleLineComment(cxt Target, el *asciidoc.SingleLineComment) {
	cxt.FlushWrap()
	cxt.DisableWrap()
	cxt.EnsureNewLine()
	cxt.WriteString("//")
	cxt.WriteString(el.Value)
	cxt.EnsureNewLine()
	cxt.EnableWrap()
}
