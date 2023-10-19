package adoc

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/output"
)

func renderPreamble(cxt *output.Context, p *types.Preamble) error {
	return RenderElements(cxt, "", p.Elements)
}
