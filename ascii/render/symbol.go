package render

import (
	"github.com/hasty/adoc/elements"
)

func renderFormattedText(cxt *Context, el elements.BlockElement, wrapper string) (err error) {
	err = renderAttributes(cxt, el, el.Attributes(), true)
	if err != nil {
		return
	}
	cxt.WriteString(wrapper)
	err = Elements(cxt, "", el.Elements()...)
	cxt.WriteString(wrapper)
	return
}

func renderSpecialCharacter(cxt *Context, s elements.SpecialCharacter) error {
	cxt.WriteString(s.Character)
	return nil
}
