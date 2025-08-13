package dump

import (
	"fmt"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/matter/spec"
)

func dumpTable(doc *spec.Doc, tbl *asciidoc.Table, indent int) {
	fmt.Print(strings.Repeat("\t", indent))
	fmt.Print("{cells}:\n")
	for row := range doc.Reader().Iterate(tbl, tbl.Children()) {
		switch row := row.(type) {
		case *asciidoc.TableRow:
			dumpTableRow(doc, row, indent+1)
		default:
			els := asciidoc.Elements{row}
			dumpElements(doc, &els, els, indent+1)
		}
	}
}

func dumpTableRow(doc *spec.Doc, row *asciidoc.TableRow, indent int) {
	fmt.Print(strings.Repeat("\t", indent))
	fmt.Printf("{row%s}:\n", dumpPosition(row))
	dumpTableCells(doc, row.TableCells(), indent+1)
}

func dumpTableCells(doc *spec.Doc, cells []*asciidoc.TableCell, indent int) {
	for _, c := range cells {
		fmt.Print(strings.Repeat("\t", indent))
		if c.Blank {
			fmt.Print("{cellblank}:\n")
		} else {
			fmt.Printf("{cell%s}:\n", dumpPosition(c))
			if c.Format != nil {
				fmt.Print(strings.Repeat("\t", indent+1))
				fmt.Printf("{format: %v (cell %d row %d)}\n", c.Format, c.Format.Span.Column.Value, c.Format.Span.Row.Value)
			}
			dumpElements(doc, c, c.Children(), indent+1)

		}
	}

}
