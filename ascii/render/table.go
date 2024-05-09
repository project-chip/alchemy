package render

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/hasty/adoc/elements"
	"github.com/hasty/alchemy/internal/text"
)

type table struct {
	columnCount int
	rows        []*tableRow
}

type tableRow struct {
	index   int
	element elements.Element
	cells   []*tableCell
}

type tableCell struct {
	value     string
	formatter *elements.TableCellFormat
	blank     bool

	width  int // width of the actual content in this cell
	margin int // width of the formatter of the next cell
}

func renderTable(cxt *Context, t *elements.Table) (err error) {

	tbl := &table{columnCount: t.ColumnCount}

	err = renderTableSubElements(cxt, t, tbl)
	if err != nil {
		return
	}

	err = renderAttributes(cxt, t, t.Attributes(), false)
	if err != nil {
		return
	}

	colOffsets := calculateColumnOffsets(tbl)

	cxt.WriteString("|===")
	cxt.WriteNewline()
	for _, row := range tbl.rows {
		if row.element != nil {
			Elements(cxt, "", row.element)
		} else {
			err = renderRow(cxt, row.cells, colOffsets)
		}
		if err != nil {
			return err
		}
	}
	cxt.WriteString("|===\n")
	return
}

func renderRow(cxt *Context, cells []*tableCell, colOffsets []int) error {
	if len(cells) == 0 {
		return nil
	}
	//var row strings.Builder
	var index int

	for i, c := range cells {
		if c.blank {
			continue
		}
		if i >= len(colOffsets) {
			return fmt.Errorf("column offset out of bounds: %d vs %d", i, len(colOffsets))
		}
		offset := colOffsets[i]
		var indentLength int
		if index < offset {
			indentLength = offset - index
		}
		var format string
		colSpan := 1
		if c.formatter != nil {
			format = renderTableCellFormat(c.formatter)
			colSpan = c.formatter.Span.Column.Value
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
	cxt.WriteNewline()
	return nil
}

func calculateColumnOffsets(tbl *table) (colOffsets []int) {
	colCount := tbl.columnCount
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
	offsetMatrix := make([][]int, len(tbl.rows))

	for i, r := range tbl.rows {
		if r.element != nil {
			continue
		}
		offsetMatrix[i] = offsetsForRow(r.cells)
	}
	return offsetMatrix
}

func offsetsForRow(cells []*tableCell) (offsets []int) {
	offsets = make([]int, len(cells))
	for i, c := range cells {
		colSpan := 1
		var format string
		if c.formatter != nil {
			format = renderTableCellFormat(c.formatter)
			if i == 0 {
				indent := utf8.RuneCountInString(format)
				offsets[i] = indent
			}
			colSpan = c.formatter.Span.Column.Value
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
				format = renderTableCellFormat(nextCell.formatter)
				c.margin = utf8.RuneCountInString(format)
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

func renderTableSubElements(cxt *Context, t *elements.Table, tbl *table) (err error) {
	var rowCount = 0
	for _, row := range t.Elements() {
		tr := &tableRow{}
		switch row := row.(type) {
		case *elements.TableRow:
			tr.index = rowCount
			rowCount++
			for _, c := range row.TableCells {
				renderContext := NewContext(cxt, cxt.Doc)
				err = Elements(renderContext, "", c.Elements()...)
				if err != nil {
					return
				}
				tr.cells = append(tr.cells, &tableCell{value: text.TrimWhitespace(renderContext.String()), formatter: c.Format, blank: c.Blank})
			}
		default:
			tr.element = row
		}
		tbl.rows = append(tbl.rows, tr)
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

func writeCellValue(out *Context, c *tableCell, width int, indent int) (count int) {
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
		directivePrefix := strings.HasPrefix(line, "ifdef::") || strings.HasPrefix(line, "ifndef::") || strings.HasPrefix(line, "endif::")
		if directivePrefix {
			out.WriteNewline()
		}
		if i > 0 && !strings.HasPrefix(line, "// ") && !directivePrefix {
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

func renderTableCellFormat(format *elements.TableCellFormat) string {
	if format == nil {
		return ""
	}
	var s strings.Builder
	colSpan := format.Span.Column.Value
	rowSpan := format.Span.Row.Value
	if (format.Span.Column.IsSet && colSpan > 1) || (format.Span.Row.IsSet && rowSpan > 1) {
		if colSpan > 1 {
			s.WriteString(strconv.Itoa(colSpan))
		}
		if rowSpan > 1 {
			s.WriteRune('.')
			s.WriteString(strconv.Itoa(rowSpan))
		}
		s.WriteRune('+')
	}
	if format.Multiplier.IsSet {
		multipler := format.Multiplier.Value
		if multipler > 1 {
			s.WriteString(strconv.Itoa(multipler))
			s.WriteRune('*')

		}
	}
	if format.HorizontalAlign.IsSet {
		switch format.HorizontalAlign.Value {
		case elements.TableCellHorizontalAlignLeft:
			s.WriteRune('<')
		case elements.TableCellHorizontalAlignCenter:
			s.WriteRune('^')
		case elements.TableCellHorizontalAlignRight:
			s.WriteRune('>')
		}
	}
	if format.VerticalAlign.IsSet {
		switch format.VerticalAlign.Value {
		case elements.TableCellVerticalAlignTop:
			s.WriteString(".<")
		case elements.TableCellVerticalAlignMiddle:
			s.WriteString(".^")
		case elements.TableCellVerticalAlignBottom:
			s.WriteString(".>")
		}

	}
	if format.Style.IsSet {
		switch format.Style.Value {
		case elements.TableCellStyleAsciiDoc:
			s.WriteRune('a')
		case elements.TableCellStyleEmphasis:
			s.WriteRune('e')
		case elements.TableCellStyleHeader:
			s.WriteRune('h')
		case elements.TableCellStyleLiteral:
			s.WriteRune('l')
		case elements.TableCellStyleMonospace:
			s.WriteRune('m')
		case elements.TableCellStyleStrong:
			s.WriteRune('s')
		}
	}
	return s.String()
}
