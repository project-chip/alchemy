package render

import "github.com/project-chip/alchemy/asciidoc"

func renderDelimitedLines(cxt Target, el asciidoc.HasLines, delimiter asciidoc.Delimiter) (err error) {
	cxt.FlushWrap()
	if ae, ok := el.(asciidoc.Attributable); ok {
		err = renderAttributes(cxt, ae.Attributes(), attributeRenderTypeBlock)
		if err != nil {
			return
		}
	}
	cxt.DisableWrap()
	renderDelimiter(cxt, delimiter)
	for _, l := range el.Lines() {
		cxt.WriteString(l)
		cxt.WriteRune('\n')
	}
	renderDelimiter(cxt, delimiter)
	cxt.EnableWrap()
	return
}

func renderDelimitedElements(cxt Target, el asciidoc.ParentElement, delimiter asciidoc.Delimiter) (err error) {
	cxt.FlushWrap()
	if ae, ok := el.(asciidoc.Attributable); ok {
		err = renderAttributes(cxt, ae.Attributes(), attributeRenderTypeBlock)
		if err != nil {
			return
		}
	}
	cxt.DisableWrap()
	renderDelimiter(cxt, delimiter)
	cxt.EnableWrap()
	err = Elements(cxt, "", el.Children()...)
	if err != nil {
		return
	}
	cxt.DisableWrap()
	renderDelimiter(cxt, delimiter)
	cxt.EnableWrap()
	return
}

func renderFencedBlock(cxt Target, el *asciidoc.FencedBlock) (err error) {
	cxt.FlushWrap()
	err = renderAttributes(cxt, el.Attributes(), attributeRenderTypeBlock)
	if err != nil {
		return
	}
	cxt.DisableWrap()
	renderDelimiter(cxt, el.Delimiter.Delimiter)
	if len(el.Delimiter.Language) > 0 {
		err = Elements(cxt, "", el.Delimiter.Language...)
	}
	if err != nil {
		return
	}
	cxt.EnableWrap()
	err = Elements(cxt, "", el.Children()...)
	if err != nil {
		return
	}
	renderDelimiter(cxt, el.Delimiter.Delimiter)
	cxt.EnableWrap()
	return
}
