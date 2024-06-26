package render

import "github.com/project-chip/alchemy/asciidoc"

func renderDelimitedLines(cxt *Context, el asciidoc.HasLines, delimiter asciidoc.Delimiter) {
	if ae, ok := el.(asciidoc.Attributable); ok {
		renderAttributes(cxt, ae.Attributes(), false)
	}
	renderDelimiter(cxt, delimiter)
	for _, l := range el.Lines() {
		cxt.WriteString(l)
		cxt.WriteRune('\n')
	}
	renderDelimiter(cxt, delimiter)
}

func renderDelimitedElements(cxt *Context, el asciidoc.HasElements, delimiter asciidoc.Delimiter) {
	if ae, ok := el.(asciidoc.Attributable); ok {
		renderAttributes(cxt, ae.Attributes(), false)
	}
	renderDelimiter(cxt, delimiter)
	Elements(cxt, "", el.Elements()...)
	renderDelimiter(cxt, delimiter)
}
