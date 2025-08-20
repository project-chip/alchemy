package dump

import (
	"fmt"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
)

func dumpTable(reader Reader, tbl *asciidoc.Table, indent int) {
	fmt.Print(strings.Repeat("\t", indent))
	fmt.Print("{cells}:\n")
	for row := range reader.Iterate(tbl, tbl.Children()) {
		switch row := row.(type) {
		case *asciidoc.TableRow:
			dumpTableRow(reader, row, indent+1)
		default:
			els := asciidoc.Elements{row}
			dumpElements(reader, &els, els, indent+1)
		}
	}
}

func dumpTableRow(reader Reader, row *asciidoc.TableRow, indent int) {
	fmt.Print(strings.Repeat("\t", indent))
	fmt.Printf("{row%s}:\n", dumpPosition(row))
	dumpTableCells(reader, row.TableCells(), indent+1)
}

func dumpTableCells(reader Reader, cells []*asciidoc.TableCell, indent int) {
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
			dumpElements(reader, c, c.Children(), indent+1)

		}
	}

}
