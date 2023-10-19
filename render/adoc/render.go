package adoc

import (
	"context"
	"regexp"
	"strings"

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
	output := strings.TrimSpace(renderContext.String())
	return postProcess(output), nil
}

var eolWhitespacePattern = regexp.MustCompile(`(?m)[ ]+\n`)

func postProcess(s string) string {
	return eolWhitespacePattern.ReplaceAllString(s, "\n")
}
