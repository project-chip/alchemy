package render

import (
	"context"
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/output"
)

func Render(cxt context.Context, doc *ascii.Doc) string {
	renderContext := output.NewContext(cxt, doc)
	RenderElements(renderContext, "", renderContext.Doc.Root.Elements)
	renderContext.WriteNewline()
	return renderContext.String()
}

func renderSection(cxt *output.Context, s *types.Section) {
	renderAttributes(cxt, s, s.Attributes)
	cxt.WriteString(strings.Repeat("=", s.Level+1))
	cxt.WriteRune(' ')
	cxt.WriteString(ascii.GetSectionTitle(s))
	cxt.WriteRune('\n')
}

func RenderElements(cxt *output.Context, prefix string, elements []interface{}) {
	var previous interface{}
	for _, e := range elements {
		if section, ok := e.(*ascii.Section); ok {
			renderSection(cxt, section.Base)
			RenderElements(cxt, "", section.Elements)
			continue
		}
		if el, ok := e.(*ascii.Element); ok {
			e = el.Base
		}
		switch el := e.(type) {
		case *types.DelimitedBlock:
			renderDelimitedBlock(cxt, el)
		case *types.Paragraph:
			cxt.WriteString(prefix)
			renderParagraph(cxt, el, &previous)
			continue // skip setting previous to the paragraph itself
		case *types.Table:
			renderTable(cxt, el)
		case *types.BlankLine:
			if _, ok := previous.(*types.StringElement); ok {
				cxt.WriteRune('\n')
			}
			cxt.WriteRune('\n')
		case *types.InternalCrossReference:
			renderInternalCrossReference(cxt, el)
		case *types.List:
			renderList(cxt, el)
		case *types.AttributeDeclaration:
			cxt.WriteString(el.RawText())
		case *types.StringElement:
			text, _ := el.RawText()
			cxt.WriteString(text)
		case *types.SinglelineComment:
			cxt.WriteString("//")
			cxt.WriteString(el.Content)
			cxt.WriteNewline()
		case *types.ImageBlock:
			text := el.Location.Scheme
			cxt.WriteString(text)
			switch p := el.Location.Path.(type) {
			case string:
				cxt.WriteString(p)
			default:
				fmt.Printf("unknown image location path type: %T\n", p)
			}
			renderAttributes(cxt, el, el.Attributes)
		default:
			fmt.Printf("unknown element type: %T\n", el)
		}
		previous = e
	}
}
