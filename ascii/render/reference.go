package render

import (
	"strings"

	"github.com/hasty/adoc/asciidoc"
)

func renderInternalCrossReference(cxt *Context, cf *asciidoc.CrossReference) (err error) {
	id := cf.ID

	if strings.HasPrefix(id, "_") {
		return
	}
	cxt.WriteString("<<")
	cxt.WriteString(id)
	if !cf.Set.IsWhitespace() {
		cxt.WriteString(",")
		Elements(cxt, "", cf.Elements()...)
	}
	cxt.WriteString(">>")
	return
}
