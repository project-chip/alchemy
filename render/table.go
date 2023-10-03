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
	footer *tableRow
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
	margin int // width of the formatter of the next cell
}

func renderTable(cxt *output.Context, t *types.Table) (err error) {

	tbl := &table{header: &tableRow{}, footer: &tableRow{}}

	err = renderTableSubElements(cxt, t, tbl)
	if err != nil {
		return
	}

	renderAttributes(cxt, t, t.Attributes, false)

	var colOffsets []int
	colOffsets, err = calculateColumnOffsets(tbl)
	if err != nil {
		return
	}

	cxt.WriteString("|===")
	cxt.WriteNewline()
	err = renderRow(cxt, tbl.header.cells, colOffsets)
	if err != nil {
		return
	}
	for _, row := range tbl.rows {
		err = renderRow(cxt, row.cells, colOffsets)
		if err != nil {
			return
		}
	}
	err = renderRow(cxt, tbl.footer.cells, colOffsets)
	if err != nil {
		return
	}
	cxt.WriteString("|===\n")
	return
}

func renderRow(cxt *output.Context, cells []*tableCell, colOffsets []int) (err error) {
	if len(cells) == 0 {
		return
	}
	var row strings.Builder
	var index int

	for i, c := range cells {
		if c.blank {
			continue
		}
		offset := colOffsets[i]
		var indentLength int
		if index < offset {
			indentLength = offset - index
		}
		var format string
		colSpan := 1
		if c.formatter != nil {
			format = c.formatter.Content
			colSpan = c.formatter.ColumnSpan
		}
		indent := fmt.Sprintf("%*s", indentLength, format)
		row.WriteString(indent)
		index += utf8.RuneCountInString(indent)
		row.WriteString("| ")
		index += 2
		if i == len(cells)-1 || i+colSpan > len(cells)-1 {
			writeCellValue(c, c.width, &row)
			break
		}
		offset = colOffsets[i]
		contentLength := offset - index
		if c.margin > 0 {
			contentLength -= (c.margin - 1)
		}
		index += writeCellValue(c, contentLength, &row)
	}
	row.WriteRune('\n')
	cxt.WriteString(row.String())
	return
}

func calculateColumnOffsets(tbl *table) (colOffsets []int, err error) {
	colCount := len(tbl.header.cells)
	colOffsets = make([]int, colCount)
	// We do two passes on the column widths; first to make sure everything is spaced out
	// enough for each cell, and then again to compact any unnecessary white space
	for i := 0; i < 2; i++ {
		err = calculateColumnOffsetsForRow(tbl.header.cells, colOffsets, i > 0)
		if err != nil {
			return
		}
		for _, r := range tbl.rows {
			err = calculateColumnOffsetsForRow(r.cells, colOffsets, i > 0)
			if err != nil {
				return
			}
		}
		err = calculateColumnOffsetsForRow(tbl.footer.cells, colOffsets, i > 0)
		if err != nil {
			return
		}
	}
	return
}

func calculateColumnOffsetsForRow(cells []*tableCell, colOffsets []int, expand bool) (err error) {
	if len(cells) > len(colOffsets) {
		return fmt.Errorf("more cells than offsets: %d > %d", len(cells), len(colOffsets))
	}
	index := 0
	for i, c := range cells {
		colSpan := 1
		if c.formatter != nil {
			if i == 0 {
				indent := utf8.RuneCountInString(c.formatter.Content)
				offset := colOffsets[0]
				if indent > offset {
					shiftColumns(colOffsets, 0, indent-offset)
					index += indent
				}
			}
			colSpan = c.formatter.ColumnSpan
		}
		width := getCellWidth(c)
		if width != 0 {
			// 2 spaces on either side
			width += 2
		} else if !c.blank {
			// Empty cells take up one space
			width = 1
		}

		for j := i + 1; j < len(cells); j++ {
			nextCell := cells[j]
			if nextCell.blank {
				continue
			}
			if nextCell.formatter != nil {
				c.margin = utf8.RuneCountInString(nextCell.formatter.Content)
				width += c.margin
				break
			}
		}

		c.width = width

		if i+colSpan > len(cells)-1 {
			continue
		}

		offset := colOffsets[i]
		nextOffset := colOffsets[i+colSpan]
		availableSpace := max(0, (nextOffset-offset)-1)
		if availableSpace <= width {
			if expand {
				shiftColumns(colOffsets, i+colSpan, (width - availableSpace))
			} else {
				colOffsets[i+colSpan] = offset + width + 1
			}
		}
	}
	return
}

func shiftColumns(offsets []int, index int, amount int) {
	for i := index; i < len(offsets); i++ {
		offsets[i] += amount
	}
}

func renderTableSubElements(cxt *output.Context, t *types.Table, tbl *table) (err error) {
	if t.Header != nil {
		for _, c := range t.Header.Cells {
			renderContext := output.NewContext(cxt, cxt.Doc)
			err = RenderElements(renderContext, "", c.Elements)
			if err != nil {
				return
			}
			tbl.header.cells = append(tbl.header.cells, &tableCell{value: strings.TrimSpace(renderContext.String()), format: c.Format, formatter: c.Formatter, blank: c.Blank})
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
			tr.cells = append(tr.cells, &tableCell{value: strings.TrimSpace(renderContext.String()), format: c.Format, formatter: c.Formatter, blank: c.Blank})
		}
		tbl.rows = append(tbl.rows, tr)
	}
	if t.Footer != nil {
		for _, c := range t.Footer.Cells {
			renderContext := output.NewContext(cxt, cxt.Doc)
			err = RenderElements(renderContext, "", c.Elements)
			if err != nil {
				return
			}
			tbl.footer.cells = append(tbl.footer.cells, &tableCell{value: strings.TrimSpace(renderContext.String()), format: c.Format, formatter: c.Formatter, blank: c.Blank})
		}
	}
	return
}

func getCellWidth(c *tableCell) int {
	lines := strings.Split(c.value, "\n")
	if len(lines) == 1 {
		return utf8.RuneCountInString(c.value)
	}
	return utf8.RuneCountInString(lines[len(lines)-1])
	var width int
	for _, line := range lines {
		width = max(width, utf8.RuneCountInString(strings.TrimSpace(line)))
	}
	return width
}

func writeCellValue(c *tableCell, width int, out *strings.Builder) (index int) {
	lines := strings.Split(c.value, "\n")
	if len(lines) == 1 {
		v := fmt.Sprintf("%-*s", width, c.value)
		out.WriteString(v)
		index = utf8.RuneCountInString(v)
		return
	}
	length := out.Len()
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if i > 0 && !strings.HasPrefix(line, "// ") {
			out.WriteString(strings.Repeat(" ", length))
		}
		v := fmt.Sprintf("%-*s", width, line)
		out.WriteString(v)
		if i < len(lines)-1 {
			out.WriteRune('\n')
		} else {
			index = utf8.RuneCountInString(v)
		}
	}
	return
}
