package render

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func renderFootnoteReference(cxt *Context, fr *types.FootnoteReference) (err error) {
	var fn *types.Footnote
	for _, f := range cxt.Doc.Footnotes() {
		if f.ID == fr.ID {
			fn = f
			break
		}
	}
	if fn == nil {
		return fmt.Errorf("missing footnote ID %d", fr.ID)
	}
	cxt.WriteString("footnote:[")
	err = Elements(cxt, "", fn.Elements)
	cxt.WriteString("]")
	return
}
