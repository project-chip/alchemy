package dump

import (
	"fmt"
	"strings"

	"github.com/hasty/adoc/asciidoc"
	"github.com/hasty/alchemy/ascii"
)

func dumpTable(doc *ascii.Doc, tbl *asciidoc.Table, indent int) {
	fmt.Print(strings.Repeat("\t", indent))
	fmt.Print("{cells}:\n")
	for _, row := range tbl.Elements() {
		switch row := row.(type) {
		case *asciidoc.TableRow:
			dumpTableRow(doc, row, indent+1)
		default:
			dumpElements(doc, asciidoc.Set{row}, indent+1)
		}
	}
}

func dumpTableRow(doc *ascii.Doc, row *asciidoc.TableRow, indent int) {
	fmt.Print(strings.Repeat("\t", indent))
	fmt.Print("{row}:\n")
	dumpTableCells(doc, row.TableCells(), indent+1)
}

func dumpTableCells(doc *ascii.Doc, cells []*asciidoc.TableCell, indent int) {
	for _, c := range cells {
		fmt.Print(strings.Repeat("\t", indent))
		if c.Blank {
			fmt.Print("{cellblank}:\n")
		} else {
			fmt.Print("{cell}:\n")
			if c.Format != nil {
				fmt.Print(strings.Repeat("\t", indent+1))
				fmt.Printf("{format: %v (cell %d row %d)}\n", c.Format, c.Format.Span.Column.Value, c.Format.Span.Row.Value)
			}
			dumpElements(doc, c.Elements(), indent+1)

		}
	}

}
