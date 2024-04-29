package render

import (
	"strings"

	"github.com/hasty/adoc/elements"
)

func renderSection(cxt *Context, s *elements.Section) (err error) {
	err = renderAttributes(cxt, s, s.Attributes(), false)
	if err != nil {
		return
	}
	err = renderSectionTitle(cxt, s.Title, s.Level+1)
	return
}

func renderSectionTitle(cxt *Context, title []elements.Element, level int) (err error) {
	cxt.WriteString(strings.Repeat("=", level))
	cxt.WriteRune(' ')
	err = Elements(cxt, "", title...)
	cxt.WriteNewline()
	return
}
