package render

import "github.com/project-chip/alchemy/asciidoc"

func renderFileInclude(cxt *Context, el *asciidoc.FileInclude) {
	cxt.WriteString("include::")
	Elements(cxt, "", el.Elements()...)
	attributes := el.Attributes()
	if len(attributes) == 0 {
		cxt.WriteString("[]\n")
	} else {
		renderAttributes(cxt, el.Attributes(), true)
		cxt.WriteRune('\n')
	}
}

func renderCounter(cxt *Context, el *asciidoc.Counter) {
	cxt.WriteString("{counter")
	if !el.Display {
		cxt.WriteRune('2')
	}
	cxt.WriteRune(':')
	cxt.WriteString(el.Name)
	if len(el.InitialValue) > 0 {
		cxt.WriteRune(':')
		cxt.WriteString(el.InitialValue)
	}
	cxt.WriteString("}")
}
