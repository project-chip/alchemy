package render

import "github.com/project-chip/alchemy/asciidoc"

func renderSingleLineComment(cxt *Context, el *asciidoc.SingleLineComment) {
	cxt.EnsureNewLine()
	cxt.WriteString("//")
	cxt.WriteString(el.Value)
	cxt.EnsureNewLine()
}
