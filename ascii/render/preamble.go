package render

func renderPreamble(cxt *Context, p *elements.Preamble) error {
	return Elements(cxt, "", p.Elements)
}
