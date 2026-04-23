package render

import "github.com/project-chip/alchemy/asciidoc"

func renderFileInclude(cxt Target, el *asciidoc.FileInclude) (err error) {
	cxt.StartBlock()
	cxt.WriteString("include::")
	err = Elements(cxt, "", el.Children()...)
	if err != nil {
		return
	}
	attributes := el.Attributes()
	if len(attributes) == 0 {
		cxt.WriteString("[]\n")
	} else {
		err = renderAttributes(cxt, el.Attributes(), attributeRenderTypeInline)
		if err != nil {
			return
		}
		cxt.WriteRune('\n')
	}
	cxt.EndBlock()
	return
}

func renderCounter(cxt Target, el *asciidoc.Counter) {
	cxt.StartBlock()
	cxt.WriteString("{counter")
	if !el.Display.Visible() {
		cxt.WriteRune('2')
	}
	cxt.WriteRune(':')
	cxt.WriteString(el.Name)
	if len(el.InitialValue) > 0 {
		cxt.WriteRune(':')
		cxt.WriteString(el.InitialValue)
	}
	cxt.WriteString("}")
	cxt.EndBlock()
}
