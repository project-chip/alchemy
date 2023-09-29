package render

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/output"
)

func renderSection(cxt *output.Context, s *types.Section) (err error) {
	err = renderAttributes(cxt, s, s.Attributes, false)
	if err != nil {
		return
	}
	err = renderSectionTitle(cxt, s.Title, s.Level+1)
	return
}

func renderSectionTitle(cxt *output.Context, title []interface{}, level int) (err error) {
	cxt.WriteString(strings.Repeat("=", level))
	cxt.WriteRune(' ')
	err = RenderElements(cxt, "", title)
	cxt.WriteNewline()
	return
}
