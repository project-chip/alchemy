package dump

import (
	"fmt"
	"strings"

	"github.com/hasty/adoc/elements"
	"github.com/hasty/alchemy/ascii"
)

func dumpTable(doc *ascii.Doc, tbl *elements.Table, indent int) {
	fmt.Print(strings.Repeat("\t", indent))
	fmt.Print("{head}:\n")
	dumpTableCells(doc, tbl.Header.Cells, indent+1)
	fmt.Print(strings.Repeat("\t", indent))
	fmt.Print("{body}:\n")
	for _, row := range tbl.Rows {
		dumpTableRow(doc, row, indent+1)
	}
}

func dumpTableRow(doc *ascii.Doc, row *elements.TableRow, indent int) {
	fmt.Print(strings.Repeat("\t", indent))
	fmt.Print("{row}:\n")
	dumpTableCells(doc, row.Cells, indent+1)
}

func dumpTableCells(doc *ascii.Doc, cells []*elements.TableCell, indent int) {
	for _, c := range cells {
		fmt.Print(strings.Repeat("\t", indent))
		if c.Blank {
			fmt.Print("{cellblank}:\n")
		} else {
			fmt.Print("{cell}:\n")
			if c.Formatter != nil {
				fmt.Print(strings.Repeat("\t", indent+1))
				fmt.Printf("{format: %s (cell %d row %d)}\n", c.Formatter.Content, c.Formatter.ColumnSpan, c.Formatter.RowSpan)
			}
			dumpElements(doc, c.Elements, indent+1)

		}
	}

}
