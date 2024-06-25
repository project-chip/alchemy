package render

import (
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
)

func renderSection(cxt *Context, s *asciidoc.Section) (err error) {
	err = renderAttributes(cxt, s.Attributes(), false)
	if err != nil {
		return
	}
	err = renderSectionTitle(cxt, s.Title, s.Level+1)
	return
}

func renderSectionTitle(cxt *Context, title asciidoc.Set, level int) (err error) {
	cxt.WriteString(strings.Repeat("=", level))
	cxt.WriteRune(' ')
	err = Elements(cxt, "", title...)
	cxt.EnsureNewLine()
	return
}
