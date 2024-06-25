package render

import "github.com/project-chip/alchemy/asciidoc"

func renderFormattedText(cxt *Context, el asciidoc.BlockElement, wrapper string) (err error) {
	err = renderAttributes(cxt, el.Attributes(), true)
	if err != nil {
		return
	}
	cxt.WriteString(wrapper)
	err = Elements(cxt, "", el.Elements()...)
	cxt.WriteString(wrapper)
	return
}

func renderSpecialCharacter(cxt *Context, s asciidoc.SpecialCharacter) error {
	cxt.WriteString(s.Character)
	return nil
}
