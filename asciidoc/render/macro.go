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
	renderCounterText(cxt, el)
	cxt.EndBlock()
}

func renderCounterText(tb TextBuffer, el *asciidoc.Counter) {
	tb.WriteString("{counter")
	if !el.Display.Visible() {
		tb.WriteRune('2')
	}
	tb.WriteRune(':')
	tb.WriteString(el.Name)
	if len(el.InitialValue) > 0 {
		tb.WriteRune(':')
		tb.WriteString(el.InitialValue)
	}
	tb.WriteString("}")
}
