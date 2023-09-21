package parse

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/output"
	"github.com/hasty/matterfmt/render"
)

func Dump(cxt *output.Context) {
	dumpElements(cxt, cxt.Doc.Base.Elements, 0)
}

func dumpElements(cxt *output.Context, elements []interface{}, indent int) {
	for _, e := range elements {
		fmt.Print(strings.Repeat(" ", indent*5))
		switch el := e.(type) {
		case *types.BlankLine:
			fmt.Print("{blank}\n")
		case *types.DelimitedBlock:
			fmt.Print("{delim}: ")
			switch el.Kind {
			case "comment":
				renderContext := output.NewContext(cxt, cxt.Doc)
				render.RenderElements(renderContext, "", []interface{}{el})
				fmt.Print(snippet(renderContext.String()))
			}
			fmt.Print("\n")
		case *types.AttributeDeclaration:
			fmt.Print("{attrib}: ")
			fmt.Print(snippet(el.RawText()))
			fmt.Print("\n")
		case *types.Paragraph:

			fmt.Print("{para}: ")
			fmt.Print("\n")
			dumpElements(cxt, el.Elements, indent+1)
		case *types.Section:
			fmt.Printf("{sec %d}: ", el.Level)
			fmt.Printf("\t{title:}\n")
			dumpElements(cxt, el.Title, indent+2)
			fmt.Printf("\t{body:}\n")
			dumpElements(cxt, el.Elements, indent+2)
		case *types.StringElement:
			fmt.Print("{str}: ", snippet(el.Content))
			fmt.Print("\n")
		case *types.Table:
			fmt.Print("{tab}:\n")
			dumpTable(cxt, el, indent+1)
		case *types.List:
			fmt.Print("{list}:\n")
			dumpElements(cxt, el.GetElements(), indent+1)
		case *types.OrderedListElement:
			fmt.Print("{ole}:\n")
			dumpElements(cxt, el.GetElements(), indent+1)
		case *types.UnorderedListElement:
			fmt.Print("{uole}:\n")
			dumpElements(cxt, el.GetElements(), indent+1)
		case *types.InternalCrossReference:
			fmt.Print("{xref}\n")
		case *types.SpecialCharacter:
			fmt.Printf("{sc: %s}\n", el.Name)
		case *types.Symbol:
			fmt.Printf("{sym: %s}\n", el.Name)
		case *types.InlineLink:
			fmt.Printf("{link: %s %s}\n", el.Location.Scheme, el.Location.Path.(string))
		default:
			fmt.Printf("unknown element type: %T\n", el)
		}
	}
}

func dumpTable(cxt *output.Context, tbl *types.Table, indent int) {
	fmt.Print(strings.Repeat(" ", indent*5))
	fmt.Print("{head}:\n")
	dumpTableCells(cxt, tbl.Header.Cells, indent+1)
	fmt.Print(strings.Repeat(" ", indent*5))
	fmt.Print("{body}:\n")
	for _, row := range tbl.Rows {
		dumpTableRow(cxt, row, indent+1)
	}
}

func dumpTableCells(cxt *output.Context, cells []*types.TableCell, indent int) {
	for _, c := range cells {
		fmt.Print(strings.Repeat(" ", indent*5))
		fmt.Print("{cell}:\n")
		dumpElements(cxt, c.Elements, indent+1)
	}

}

func dumpTableRow(cxt *output.Context, row *types.TableRow, indent int) {
	fmt.Print(strings.Repeat(" ", indent*5))
	fmt.Print("{cell}:\n")
	dumpTableCells(cxt, row.Cells, indent+1)
}

func snippet(str string) string {
	v := []rune(str)
	if 25 < len(v) {
		str = string(v[:25])
	}
	return strings.ReplaceAll(str, "\n", "")
}
