package render

import "github.com/project-chip/alchemy/asciidoc"

func renderEscape(cxt Target, el *asciidoc.AlchemyEscape) (err error) {
	cxt.FlushWrap()
	cxt.DisableWrap()
	cxt.WriteString(":alchemy-escape: ")
	err = Elements(cxt, "", el.Elements()...)
	if err != nil {
		return
	}
	cxt.WriteRune('\n')
	for _, l := range el.Lines() {
		cxt.WriteString(l)
		cxt.WriteRune('\n')
	}
	cxt.WriteString(":!alchemy-escape:\n")
	cxt.EnableWrap()
	return
}
