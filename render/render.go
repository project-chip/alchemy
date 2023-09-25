package render

import (
	"context"
	"regexp"

	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/output"
)

func Render(cxt context.Context, doc *ascii.Doc) (string, error) {
	renderContext := output.NewContext(cxt, doc)
	err := RenderElements(renderContext, "", renderContext.Doc.Elements)
	if err != nil {
		return "", err
	}
	renderContext.WriteNewline()
	return postProcess(renderContext.String()), nil
}

var eolWhitespacePattern = regexp.MustCompile(`[ ]+\n`)

func postProcess(s string) string {
	return eolWhitespacePattern.ReplaceAllString(s, "\n")
}
