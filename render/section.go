package render

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/output"
)

func renderSection(cxt *output.Context, s *types.Section) {
	renderAttributes(cxt, s, s.Attributes)
	cxt.WriteString(strings.Repeat("=", s.Level+1))
	cxt.WriteRune(' ')
	RenderElements(cxt, "", s.Title)
	cxt.WriteRune('\n')
}
