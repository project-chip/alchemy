package render

import (
	"fmt"
)

func renderFootnoteReference(cxt *Context, fr *elements.FootnoteReference) (err error) {
	var fn *elements.Footnote
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
