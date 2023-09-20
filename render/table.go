package render

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/output"
)

type table struct {
	header *tableRow
	rows   []*tableRow
}

type tableRow struct {
	cells []*tableCell
}

type tableCell struct {
	value     string
	format    string
	formatter *types.TableCellFormat
}

type tableHeaderSpan struct {
	column int
	span   int
	width  int
}

func renderTable(cxt *output.Context, t *types.Table) {

	tbl := &table{header: &tableRow{}}

	renderTableSubElements(cxt, t, tbl)

	headerSpans, rowWidths := calculateCellWidths(tbl)

	headerCount := len(tbl.header.cells)

	renderAttributes(cxt, t, t.Attributes)

	cxt.WriteString("|===")
	cxt.WriteNewline()
	cxt.WriteString(renderTableHeaders(tbl, headerSpans, headerCount))
	cxt.WriteNewline()
	renderTableRows(cxt, tbl, rowWidths, headerCount)
	cxt.WriteString("|===\n")
}

func renderTableRows(cxt *output.Context, tbl *table, rowWidths map[int]int, headerCount int) {
	for _, tr := range tbl.rows {
		var row strings.Builder
		for i, c := range tr.cells {
			if c.formatter != nil {
				row.WriteString(c.formatter.Content)
			}
			row.WriteRune('|')
			width, ok := rowWidths[i]

			if i < len(tr.cells)-1 {
				nextCell := tr.cells[i+1]
				if nextCell.formatter != nil {
					width -= len(nextCell.formatter.Content)
				}
			}

			if ok && i+1 != len(tr.cells) {
				row.WriteRune(' ')
				writeCellValue(c, width, &row)
			} else {
				if len(c.value) > 0 {
					row.WriteRune(' ')
					writeCellValue(c, 0, &row)
				}
			}
			if i+1 != len(tr.cells) {
				row.WriteRune(' ')
			}

		}
		row.WriteRune('\n')
		cxt.WriteString(row.String())
	}

}

func renderTableHeaders(tbl *table, headerSpans map[int]*tableHeaderSpan, headerCount int) string {
	var out strings.Builder
	for i, c := range tbl.header.cells {
		if c.formatter != nil {
			out.WriteString(c.formatter.Content)
		}
		out.WriteString("| ")

		hs, ok := headerSpans[i]
		if ok && i+1 != headerCount {
			w := hs.width
			if i < len(tbl.header.cells)-1 {
				next := tbl.header.cells[i+1]
				if next.formatter != nil {
					w -= len(next.formatter.Content)
				}
			}
			writeCellValue(c, w, &out)
		} else {
			out.WriteString(c.value)

		}
		if i+1 != headerCount {
			out.WriteRune(' ')
		}

	}
	return out.String()
}

func calculateCellWidths(tbl *table) (headerSpans map[int]*tableHeaderSpan, rowWidths map[int]int) {
	headerSpans = make(map[int]*tableHeaderSpan)
	rowWidths = make(map[int]int)

	columnIndex := 0
	for i, c := range tbl.header.cells {
		thw := &tableHeaderSpan{column: columnIndex, width: getCellWidth(c), span: 1}
		headerSpans[columnIndex] = thw
		if c.formatter != nil && c.formatter.ColumnSpan > 0 {
			thw.span = c.formatter.ColumnSpan
			columnIndex += thw.span - 1
		}
		if i < len(tbl.header.cells)-1 {
			nextHeader := tbl.header.cells[i+1]
			if nextHeader.formatter != nil {
				thw.width += len(nextHeader.formatter.Content)
			}
		}
		columnIndex++
	}

	for _, tr := range tbl.rows {
		var currentHeader *tableHeaderSpan
		var spanWidth int
		for i, c := range tr.cells[0 : len(tr.cells)-1] {

			if currentHeader == nil {
				currentHeader = headerSpans[i]
			} else {
				if hw, ok := headerSpans[i]; ok {
					if currentHeader != hw {
						currentHeader.width = max(currentHeader.width, spanWidth)
						spanWidth = 0
						currentHeader = hw
					}
				}
			}
			l := getCellWidth(c)

			nextCell := tr.cells[i+1]

			if nextCell.formatter != nil {
				l += len(nextCell.formatter.Content)
			}
			rowWidth := rowWidths[i]

			if currentHeader.span == 1 {
				if currentHeader.width < max(l, rowWidth) {
					currentHeader.width = max(l, rowWidth)
				} else {
					l = currentHeader.width
				}
			} else {
				spanWidth += (l + 1)
			}
			if rowWidth < l {
				rowWidths[i] = l
			}
		}
	}
	return
}

func renderTableSubElements(cxt *output.Context, t *types.Table, tbl *table) {
	for _, c := range t.Header.Cells {
		renderContext := output.NewContext(cxt, cxt.Doc)
		RenderElements(renderContext, "", c.Elements)
		tbl.header.cells = append(tbl.header.cells, &tableCell{value: renderContext.String(), format: c.Format, formatter: c.Formatter})
	}

	for _, row := range t.Rows {
		tr := &tableRow{}
		for _, c := range row.Cells {
			renderContext := output.NewContext(cxt, cxt.Doc)
			RenderElements(renderContext, "", c.Elements)
			tr.cells = append(tr.cells, &tableCell{value: renderContext.String(), format: c.Format, formatter: c.Formatter})
		}
		tbl.rows = append(tbl.rows, tr)
	}
}

func getCellWidth(c *tableCell) int {
	lines := strings.Split(c.value, "\n")
	if len(lines) == 1 {
		return len(c.value)
	}
	var width int
	for _, line := range lines {

		width = max(width, len(strings.TrimSpace(line)))
	}
	return width
}

func writeCellValue(c *tableCell, width int, out *strings.Builder) {
	lines := strings.Split(c.value, "\n")
	if len(lines) == 1 {
		out.WriteString(fmt.Sprintf("%-*s", width, c.value))
		return
	}
	length := out.Len()
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if i > 0 {
			out.WriteString(strings.Repeat(" ", length))
		}
		out.WriteString(fmt.Sprintf("%-*s", width, line))
		if i < len(lines)-1 {
			out.WriteRune('\n')
		}
	}
}
