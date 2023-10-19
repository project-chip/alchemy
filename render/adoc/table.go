package adoc

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

	colOffsets := calculateColumnOffsets(tbl)

	cxt.WriteString("|===")
	cxt.WriteNewline()
	renderRow(cxt, tbl.header.cells, colOffsets)
	for _, row := range tbl.rows {
		renderRow(cxt, row.cells, colOffsets)
	}
	renderRow(cxt, tbl.footer.cells, colOffsets)
	cxt.WriteString("|===\n")
	return
}

func renderRow(cxt *output.Context, cells []*tableCell, colOffsets []int) {
	if len(cells) == 0 {
		return
	}
	//var row strings.Builder
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
		cxt.WriteString(indent)
		index += utf8.RuneCountInString(indent)
		cxt.WriteString("| ")
		index += 2
		if i == len(cells)-1 || i+colSpan > len(cells)-1 {
			writeCellValue(cxt, c, c.width, index)
			break
		}
		contentLength := c.width
		if c.margin > 0 {
			contentLength -= (c.margin + 1)
		}
		wb := writeCellValue(cxt, c, contentLength, index)
		index += wb

	}
	cxt.WriteRune('\n')
}

func calculateColumnOffsets(tbl *table) (colOffsets []int) {
	colCount := len(tbl.header.cells)
	colOffsets = make([]int, colCount)

	offsetMatrix := offsetMatrixForTable(tbl)

	for colIndex := 0; colIndex < colCount; colIndex++ {
		var longestOffset int
		var longestOffsetRowIndex int
		for rowIndex := 0; rowIndex < len(offsetMatrix); rowIndex++ {
			os := offsetMatrix[rowIndex]
			if len(os) < colIndex+1 { // Empty header or footer
				continue
			}
			o := os[colIndex]
			if o > longestOffset {
				longestOffset = o
				longestOffsetRowIndex = rowIndex
			}
		}
		colOffsets[colIndex] = longestOffset
		for rowIndex := 0; rowIndex < len(offsetMatrix); rowIndex++ {
			if rowIndex == longestOffsetRowIndex {
				continue
			}
			offsets := offsetMatrix[rowIndex]
			if len(offsets) < colIndex+1 {
				continue
			}
			offset := offsets[colIndex]
			if offset == -1 { // Skip empty cells
				continue
			}
			shift := longestOffset - offset
			offsets[colIndex] = longestOffset
			if colIndex+2 > len(offsets) { // Last cell in row
				continue
			}
			// Shift all subsequent cells by the amount we shifted this cell
			for k := colIndex + 1; k < len(offsets); k++ {
				nextOffset := offsets[k]
				if nextOffset >= 0 { // ignore empty cells
					nextOffset += shift
					offsets[k] = nextOffset
				}
			}
		}
	}
	return
}

func offsetMatrixForTable(tbl *table) [][]int {
	offsetMatrix := make([][]int, len(tbl.rows)+2)

	offsetMatrix[0] = offsetsForRow(tbl.header.cells)

	for i, r := range tbl.rows {
		offsetMatrix[i+1] = offsetsForRow(r.cells)
	}
	offsetMatrix[len(offsetMatrix)-1] = offsetsForRow(tbl.footer.cells)
	return offsetMatrix
}

func offsetsForRow(cells []*tableCell) (offsets []int) {
	offsets = make([]int, len(cells))
	for i, c := range cells {
		colSpan := 1
		if c.formatter != nil {
			if i == 0 {
				indent := utf8.RuneCountInString(c.formatter.Content)
				offsets[i] = indent
			}
			colSpan = c.formatter.ColumnSpan
		}
		width := getCellWidth(c)
		c.width = width
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
			}
			break
		}

		if c.blank {
			offsets[i] = -1 // Blank cells don't have offsets
			continue
		}
		if i+colSpan > len(cells)-1 {
			continue
		}
		offset := offsets[i]
		offsets[i+colSpan] = offset + width + 1
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
		return utf8.RuneCountInString(strings.TrimSpace(c.value))
	}
	return utf8.RuneCountInString(strings.TrimSpace(lines[len(lines)-1]))
}

func writeCellValue(out *output.Context, c *tableCell, width int, indent int) (count int) {
	lines := strings.Split(c.value, "\n")
	if len(lines) == 1 {
		v := fmt.Sprintf("%-*s", width, c.value)
		out.WriteString(v)
		count = utf8.RuneCountInString(v)
		return
	}
	//length := out.Len()
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if i > 0 && !strings.HasPrefix(line, "// ") && !strings.HasPrefix(line, "ifdef::") && !strings.HasPrefix(line, "ifndef::") && !strings.HasPrefix(line, "endif::") {
			out.WriteString(strings.Repeat(" ", indent))
		}
		v := fmt.Sprintf("%-*s", width, line)
		out.WriteString(v)
		if i < len(lines)-1 {
			out.WriteRune('\n')
		} else {
			count = utf8.RuneCountInString(v)
		}
	}
	return
}
