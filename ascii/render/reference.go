package render

import (
	"strings"

	"github.com/hasty/adoc/elements"
)

func renderInternalCrossReference(cxt *Context, cf *elements.CrossReference) (err error) {
	id := cf.ID

	if strings.HasPrefix(id, "_") {
		return
	}
	cxt.WriteString("<<")
	cxt.WriteString(id)
	if len(cf.Set) > 0 {
		cxt.WriteString(",")
		Elements(cxt, "", cf.Elements()...)
	}
	cxt.WriteString(">>")
	return
}
