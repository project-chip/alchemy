package render

import "github.com/project-chip/alchemy/asciidoc"

func renderDelimitedLines(cxt Target, el asciidoc.HasLines, delimiter asciidoc.Delimiter) {
	cxt.FlushWrap()
	if ae, ok := el.(asciidoc.Attributable); ok {
		renderAttributes(cxt, ae.Attributes(), false)
	}
	cxt.DisableWrap()
	renderDelimiter(cxt, delimiter)
	for _, l := range el.Lines() {
		cxt.WriteString(l)
		cxt.WriteRune('\n')
	}
	renderDelimiter(cxt, delimiter)
	cxt.EnableWrap()
}

func renderDelimitedElements(cxt Target, el asciidoc.HasElements, delimiter asciidoc.Delimiter) {
	cxt.FlushWrap()
	if ae, ok := el.(asciidoc.Attributable); ok {
		renderAttributes(cxt, ae.Attributes(), false)
	}
	cxt.DisableWrap()
	renderDelimiter(cxt, delimiter)
	cxt.EnableWrap()
	Elements(cxt, "", el.Elements()...)
	cxt.DisableWrap()
	renderDelimiter(cxt, delimiter)
	cxt.EnableWrap()
}
