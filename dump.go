package main

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func dump(d *doc, doc *types.Document) {
	dumpElements(d, doc, doc.Elements, 0)
}

func dumpElements(d *doc, doc *types.Document, elements []interface{}, indent int) {
	for _, e := range elements {
		fmt.Print(strings.Repeat(" ", indent*5))
		switch el := e.(type) {
		case *types.BlankLine:
			fmt.Print("{blank}\n")
		case *types.DelimitedBlock:
			fmt.Print("{delim}: ")
			switch el.Kind {
			case "comment":
				var value output
				d.renderComment(el, &value)
				fmt.Print(snippet(value.String()))
			}
			fmt.Print("\n")
		case *types.AttributeDeclaration:
			fmt.Print("{attrib}: ")
			fmt.Print(snippet(el.RawText()))
			fmt.Print("\n")
		case *types.Paragraph:

			fmt.Print("{para}: ")
			fmt.Print("\n")
			dumpElements(d, doc, el.Elements, indent+1)
		case *types.Section:
			fmt.Printf("{sec %d}: ", el.Level)
			fmt.Print(getSectionTitle(el))
			fmt.Print("\n")
			dumpElements(d, doc, el.Elements, indent+1)
		case *types.StringElement:
			fmt.Print("{str}: ", snippet(el.Content))
			fmt.Print("\n")
		case *types.Table:
			fmt.Print("{tab}:\n")
			dumpTable(d, doc, el, indent+1)
		case *types.List:
			fmt.Print("{list}:\n")
			dumpElements(d, doc, el.GetElements(), indent+1)
		case *types.OrderedListElement:
			fmt.Print("{ole}:\n")
			dumpElements(d, doc, el.GetElements(), indent+1)
		case *types.UnorderedListElement:
			fmt.Print("{uole}:\n")
			dumpElements(d, doc, el.GetElements(), indent+1)
		case *types.InternalCrossReference:
			fmt.Print("{xref}\n")
		case *types.SpecialCharacter:
			fmt.Printf("{sc: %s}\n", el.Name)
		case *types.Symbol:
			fmt.Printf("{sym: %s}\n", el.Name)
		default:
			fmt.Printf("unknown element type: %T\n", el)
		}
	}
}

func dumpTable(d *doc, doc *types.Document, tbl *types.Table, indent int) {
	fmt.Print(strings.Repeat(" ", indent*5))
	fmt.Print("{head}:\n")
	dumpTableCells(d, doc, tbl.Header.Cells, indent+1)
	fmt.Print(strings.Repeat(" ", indent*5))
	fmt.Print("{body}:\n")
	for _, row := range tbl.Rows {
		dumpTableRow(d, doc, row, indent+1)
	}
}

func dumpTableCells(d *doc, doc *types.Document, cells []*types.TableCell, indent int) {
	for _, c := range cells {
		fmt.Print(strings.Repeat(" ", indent*5))
		fmt.Print("{cell}:\n")
		dumpElements(d, doc, c.Elements, indent+1)
	}

}

func dumpTableRow(d *doc, doc *types.Document, row *types.TableRow, indent int) {
	fmt.Print(strings.Repeat(" ", indent*5))
	fmt.Print("{cell}:\n")
	dumpTableCells(d, doc, row.Cells, indent+1)
}

func snippet(str string) string {
	v := []rune(str)
	if 25 < len(v) {
		str = string(v[:25])
	}
	return strings.ReplaceAll(str, "\n", "")
}
