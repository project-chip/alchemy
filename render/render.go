package render

import (
	"context"
	"regexp"

	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/output"
)

func Render(cxt context.Context, doc *ascii.Doc) string {
	renderContext := output.NewContext(cxt, doc)
	RenderElements(renderContext, "", renderContext.Doc.Elements)
	renderContext.WriteNewline()
	return postProcess(renderContext.String())
}

var eolWhitespacePattern = regexp.MustCompile(`[ ]+\n`)

func postProcess(s string) string {
	return eolWhitespacePattern.ReplaceAllString(s, "\n")
}
