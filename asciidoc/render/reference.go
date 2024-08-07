package render

import (
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
)

func renderInternalCrossReference(cxt Target, cf *asciidoc.CrossReference) (err error) {
	id := cf.ID

	if strings.HasPrefix(id, "_") {
		return
	}
	cxt.StartBlock()
	cxt.WriteString("<<")
	cxt.WriteString(id)
	if !cf.Set.IsWhitespace() {
		cxt.WriteString(",")
		Elements(cxt, "", cf.Elements()...)
	}
	cxt.WriteString(">>")
	cxt.EndBlock()
	return
}
