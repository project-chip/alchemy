package mle

import (
	"fmt"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
)

func extractText(elements asciidoc.Elements) string {
	var sb strings.Builder
	for _, el := range elements {
		switch v := el.(type) {
		case *asciidoc.String:
			sb.WriteString(v.Value)
		}
	}
	return strings.TrimSpace(sb.String())
}

func tableHasColumns(table *asciidoc.Table, requiredCols []string) bool {
	if len(table.Elements) == 0 {
		return false
	}

	r, ok := table.Elements[0].(*asciidoc.TableRow)
	if !ok {
		return false
	}

	headerSet := make(map[string]bool)
	for _, cell := range r.TableCells() {
		text := extractText(cell.Elements)
		headerSet[text] = true
	}

	for _, col := range requiredCols {
		if !headerSet[col] {
			return false
		}
	}

	return true
}

func getCellTextAtCol(row *asciidoc.TableRow, colIndex int) string {
	cells := row.TableCells()

	if colIndex < 0 || colIndex >= len(cells) {
		return ""
	}

	return extractText(cells[colIndex].Elements)
}

func getColumnIndex(table *asciidoc.Table, columnName string) (idx int, err error) {
	if len(table.Elements) == 0 {
		err = fmt.Errorf("table has no header row")
	}

	header, ok := table.Elements[0].(*asciidoc.TableRow)
	if !ok {
		err = fmt.Errorf("table has no header row")
	}

	for i, cell := range header.TableCells() {
		text := strings.TrimSpace(extractText(cell.Elements))
		if text == columnName {
			idx = i
			return
		}
	}

	err = fmt.Errorf("column '%s' not found in table", columnName)
	return
}

func findTableWithColumns(elements asciidoc.Elements, requiredCols []string) *asciidoc.Table {
	for _, el := range elements {
		if table, ok := el.(*asciidoc.Table); ok {
			if tableHasColumns(table, requiredCols) {
				return table
			}
		}

		if parent, ok := el.(asciidoc.ParentElement); ok {
			t := findTableWithColumns(parent.Children(), requiredCols)
			if t != nil {
				return t
			}
		}
	}
	return nil
}
