package render

import (
	"context"

	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/output"
)

func Render(cxt context.Context, doc *ascii.Doc) string {
	renderContext := output.NewContext(cxt, doc)
	RenderElements(renderContext, "", renderContext.Doc.Elements)
	renderContext.WriteNewline()
	return renderContext.String()
}
