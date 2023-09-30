package render

import (
	"fmt"
	"strings"
	"unicode/utf8"

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
	blank     bool

	width  int // width of the actual content in this cell
	indent int // width of the formatter of this cell
	margin int // width of the formatter of the next cell
}

type colSpan struct {
	//indent int
	column int
	span   int
	width  int
}

func renderTable(cxt *output.Context, t *types.Table) (err error) {

	tbl := &table{header: &tableRow{}}

	err = renderTableSubElements(cxt, t, tbl)
	if err != nil {
		return
	}

	cellSpans, indent := calculateCellWidths(tbl)

	headerCount := len(tbl.header.cells)

	renderAttributes(cxt, t, t.Attributes, false)

	cxt.WriteString("|===")
	cxt.WriteNewline()
	cxt.WriteString(renderTableHeaders(tbl, cellSpans, headerCount, indent))
	cxt.WriteNewline()
	renderTableRows(cxt, tbl, cellSpans, headerCount, indent)
	cxt.WriteString("|===\n")
	return
}

func renderTableRows(cxt *output.Context, tbl *table, colSpans map[int]*colSpan, headerCount int, indent int) {
	for _, tr := range tbl.rows {
		var row strings.Builder
		for i, c := range tr.cells {
			var ind = 0
			if i == 0 {
				ind = indent
			}
			var width int
			rowSpan, ok := colSpans[i]
			if ok {
				width = rowSpan.width
			}
			var format string
			if c.formatter != nil {
				format = c.formatter.Content
				ind = max(ind, utf8.RuneCountInString(format))
			}

			row.WriteString(fmt.Sprintf("%*s", ind, format))

			if c.blank {
				row.WriteRune(' ')
			} else {
				row.WriteRune('|')
			}

			if c.margin > 0 {
				width -= (c.margin)
			}

			if ok && i+1 != len(tr.cells) {
				row.WriteRune(' ')
				writeCellValue(c, width, &row)
			} else {
				if utf8.RuneCountInString(c.value) > 0 {
					row.WriteRune(' ')
					writeCellValue(c, 0, &row)
				}
			}
			/*if i+1 != len(tr.cells) {
				row.WriteRune(' ')
			}*/

		}
		row.WriteRune('\n')
		cxt.WriteString(row.String())
	}

}

func renderTableHeaders(tbl *table, colSpans map[int]*colSpan, headerCount int, indent int) string {
	var out strings.Builder
	for i, c := range tbl.header.cells {

		var ind int
		if i == 0 {
			ind = indent
		}
		var format string
		if c.formatter != nil {
			format = c.formatter.Content
			ind = max(ind, utf8.RuneCountInString(format))
		}

		out.WriteString(fmt.Sprintf("%*s", ind, format))

		if c.blank {
			out.WriteString("  ")
		} else {
			out.WriteString("| ")
		}

		if cs, ok := colSpans[i]; ok && i+1 != headerCount {
			w := cs.width
			if c.margin > 0 {
				w -= c.margin
			} else if i+1 == headerCount {
				w -= 1
			}
			writeCellValue(c, w, &out)
		} else {
			out.WriteString(c.value)
		}
		/*if i+1 != headerCount {
			out.WriteRune(' ')
		}*/

	}
	return out.String()
}

func calculateCellWidths(tbl *table) (colSpans map[int]*colSpan, indent int) {

	colSpans = make(map[int]*colSpan)

	columnIndex := 0
	for i, c := range tbl.header.cells {
		c.width = getCellWidth(c)
		if c.formatter != nil {
			c.indent = utf8.RuneCountInString(c.formatter.Content)
		}
		if i < len(tbl.header.cells)-1 {
			nextHeader := tbl.header.cells[i+1]
			if nextHeader.formatter != nil {
				c.margin = utf8.RuneCountInString(nextHeader.formatter.Content)
			}
		}
		if i == 0 {
			indent = max(indent, c.indent)
		}
		thw := &colSpan{column: columnIndex, width: c.width + 1, span: 1}
		if c.margin > 0 {
			thw.width += c.margin
		}
		colSpans[columnIndex] = thw
		columnIndex++
	}

	for _, tr := range tbl.rows {

		for i, c := range tr.cells[0 : len(tr.cells)-1] {

			cs, rowOK := colSpans[i]
			if !rowOK {
				cs = &colSpan{}
				colSpans[i] = cs
			}

			c.width = getCellWidth(c)
			if c.formatter != nil {
				c.indent = utf8.RuneCountInString(c.formatter.Content)
			}

			nextCell := tr.cells[i+1]

			if nextCell.formatter != nil {
				c.margin = utf8.RuneCountInString(nextCell.formatter.Content)
			}
			if i == 0 {
				indent = max(indent, c.indent)
			}
			width := c.width + 1
			if c.margin > 0 {
				width += c.margin
			}
			cs.width = max(cs.width, width)
		}
	}
	return
}

func renderTableSubElements(cxt *output.Context, t *types.Table, tbl *table) (err error) {
	if t.Header != nil {
		for _, c := range t.Header.Cells {
			renderContext := output.NewContext(cxt, cxt.Doc)
			err = RenderElements(renderContext, "", c.Elements)
			if err != nil {
				return
			}
			tbl.header.cells = append(tbl.header.cells, &tableCell{value: renderContext.String(), format: c.Format, formatter: c.Formatter, blank: c.Blank})
		}
	}

	for _, row := range t.Rows {
		tr := &tableRow{}
		for _, c := range row.Cells {
			renderContext := output.NewContext(cxt, cxt.Doc)
			err = RenderElements(renderContext, "", c.Elements)
			if err != nil {
				return
			}
			tr.cells = append(tr.cells, &tableCell{value: renderContext.String(), format: c.Format, formatter: c.Formatter, blank: c.Blank})
		}
		tbl.rows = append(tbl.rows, tr)
	}
	return
}

func getCellWidth(c *tableCell) int {
	lines := strings.Split(c.value, "\n")
	if len(lines) == 1 {
		return utf8.RuneCountInString(c.value)
	}
	var width int
	for _, line := range lines {
		width = max(width, utf8.RuneCountInString(strings.TrimSpace(line)))
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
