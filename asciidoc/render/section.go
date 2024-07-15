package render

import (
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
)

func renderSection(cxt Target, s *asciidoc.Section) (err error) {
	cxt.FlushWrap()
	err = renderAttributes(cxt, s.Attributes(), false)
	if err != nil {
		return
	}
	err = renderSectionTitle(cxt, s.Title, s.Level+1)
	return
}

func renderSectionTitle(cxt Target, title asciidoc.Set, level int) (err error) {
	cxt.DisableWrap()
	cxt.WriteString(strings.Repeat("=", level))
	cxt.WriteRune(' ')
	err = Elements(cxt, "", title...)
	cxt.EnsureNewLine()
	cxt.EnableWrap()
	return
}
