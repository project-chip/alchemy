package render

import "github.com/hasty/alchemy/asciidoc"

func renderSingleLineComment(cxt *Context, el *asciidoc.SingleLineComment) {
	cxt.EnsureNewLine()
	cxt.WriteString("//")
	cxt.WriteString(el.Value)
	cxt.EnsureNewLine()
}
