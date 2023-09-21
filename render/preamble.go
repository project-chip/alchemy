package render

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/output"
)

func renderPreamble(cxt *output.Context, p *types.Preamble) {
	for _, e := range p.Elements {
		fmt.Printf("preamble element: %T\n", e)
	}
	RenderElements(cxt, "", p.Elements)
}
