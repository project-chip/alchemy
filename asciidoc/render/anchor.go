package render

import "github.com/project-chip/alchemy/asciidoc"

func renderAnchor(cxt Target, el *asciidoc.Anchor) (err error) {
	cxt.FlushWrap()
	cxt.StartBlock()
	cxt.WriteString("[[")
	err = Elements(cxt, "", el.ID...)
	if err != nil {
		return
	}
	anchorElements := el.Children()
	if len(anchorElements) > 0 {
		cxt.WriteString(",")
		err = Elements(cxt, "", anchorElements...)
	}
	cxt.WriteString("]]")
	cxt.EndBlock()
	return
}
