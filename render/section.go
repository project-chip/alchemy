package render

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/output"
)

func renderSection(cxt *output.Context, s *types.Section) {
	renderAttributes(cxt, s, s.Attributes)
	renderSectionTitle(cxt, s.Title, s.Level+1)
}

func renderSectionTitle(cxt *output.Context, title []interface{}, level int) {
	cxt.WriteString(strings.Repeat("=", level))
	cxt.WriteRune(' ')
	RenderElements(cxt, "", title)
	cxt.WriteRune('\n')
}
