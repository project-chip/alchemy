package render

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func renderSection(cxt *Context, s *types.Section) (err error) {
	cxt.WriteNewline()
	err = renderAttributes(cxt, s, s.Attributes, false)
	if err != nil {
		return
	}
	err = renderSectionTitle(cxt, s.Title, s.Level+1)
	return
}

func renderSectionTitle(cxt *Context, title []any, level int) (err error) {
	cxt.WriteString(strings.Repeat("=", level))
	cxt.WriteRune(' ')
	err = Elements(cxt, "", title)
	cxt.WriteNewline()
	return
}
