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

	//cellSpans, indent := calculateCellWidths(tbl)

	//headerCount := len(tbl.header.cells)

	renderAttributes(cxt, t, t.Attributes, false)

	var colOffsets []int
	colOffsets, err = calculateColumnOffsets(tbl)
	if err != nil {
		return
	}

	cxt.WriteString("|===")
	cxt.WriteNewline()
	fmt.Printf("Rendering header cells %d...\n", len(tbl.header.cells))
	err = renderRow(cxt, tbl.header.cells, colOffsets)
	if err != nil {
		return
	}
	for i, row := range tbl.rows {
		fmt.Printf("Rendering row %d...\n", i)
		err = renderRow(cxt, row.cells, colOffsets)
		if err != nil {
			return
		}
	}
	//cxt.WriteString(renderTableHeaders(tbl, cellSpans, headerCount, indent))
	//cxt.WriteNewline()
	//	renderTableRows(cxt, tbl, cellSpans, headerCount, indent)
	cxt.WriteString("|===\n")
	return
}

func renderRow(cxt *output.Context, cells []*tableCell, colOffsets []int) (err error) {
	var row strings.Builder
	//var colIndex int
	var index int

	fmt.Printf("rendering row (%d cells)\n", len(cells))
	for i, c := range cells {
		if c.blank {
			fmt.Printf("skipping blank cell %d: %s (%d)\n", i, c.value, c.width)
			continue
		}
		fmt.Printf("rendering cell %d: %s (%d)\n", i, c.value, c.width)
		offset := colOffsets[i]
		fmt.Printf("render cell %d; index %d, offset %d colIndex %d\n", i, index, offset, i)
		var indentLength int
		if index < offset {
			indentLength = offset - index
		}
		//fmt.Printf("multiline indent length %d\n", indentLength)
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
			fmt.Printf("completing row @ %d: %s (%d) colIndex: %d colSpan: %d\n", i, c.value, c.width, i, colSpan)
			writeCellValue(c, c.width, &row)
			break
		}
		//fmt.Printf("colSpan %d\n", colSpan)
		//colIndex += colSpan
		//fmt.Printf("colIndex %d\n", colIndex)
		offset = colOffsets[i]
		//fmt.Printf("contentLength %d - %d = %d\n", offset, index, offset-index)
		contentLength := offset - index
		if c.margin > 0 {
			contentLength -= (c.margin - 1)
		}
		//fmt.Printf("contentLength  - %d = %d\n", c.margin, contentLength)
		index += writeCellValue(c, contentLength, &row)
		//fmt.Printf("index %d\n", index)
		//index += c.margin
	}
	row.WriteRune('\n')
	cxt.WriteString(row.String())
	return
}

func shiftColumnOffsets(offsets []int, index int, amount int) {
	for i := index; i < len(offsets); i++ {
		if i < len(offsets)-1 {
			offset := offsets[i]
			nextOffset := offsets[i+1]
			if nextOffset-offset >= amount {
				break
			}
		}
		offsets[i] += amount
	}
}

func calculateColumnOffsets(tbl *table) (colOffsets []int, err error) {
	colCount := len(tbl.header.cells)
	/*for _, c := range tbl.header.cells {
		if c.formatter != nil {
			colCount += c.formatter.ColumnSpan
		} else {
			colCount++
		}
	}*/
	colOffsets = make([]int, colCount)
	fmt.Printf("colCount: %d from %d cells\n", colCount, len(tbl.header.cells))
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
		fmt.Printf("\n\noffsets after pass %d: ", i)
		for i, o := range colOffsets {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%d: %d", i, o)
		}
		fmt.Print("\n")
	}
	return
}

func shiftColumns(offsets []int, index int, amount int) {
	for i := index; i < len(offsets); i++ {
		offsets[i] += amount
	}
}

func calculateColumnOffsetsForRow(cells []*tableCell, colOffsets []int, expand bool) (err error) {
	//colIndex := 0
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
					//					colOffsets[0] = indent - offset
					//shiftColumnOffsets(colOffsets, 0, indent-offset)
				}
			}
			colSpan = c.formatter.ColumnSpan
		}
		width := getCellWidth(c)
		if width != 0 {
			// 2 spaces on either side
			width += 2
		} else if !c.blank {
			// 2 spaces on either side
			//width += 2
			width = 1
		}

		for j := i + 1; j < len(cells); j++ {
			nextCell := cells[j]
			if nextCell.blank {
				continue
			}
			if nextCell.formatter != nil {
				c.margin = utf8.RuneCountInString(nextCell.formatter.Content)
				fmt.Printf("cell %d: next cell adding width: %d; margin %d\n", i, width, c.margin)
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
		fmt.Printf("cell %d colIndex %d colSpan %d width %d availableSpace %d (%d - %d) %s\n", i, i, colSpan, width, availableSpace, nextOffset, offset, c.value)
		if availableSpace <= width {
			if expand {
				shiftColumns(colOffsets, i+colSpan, (width - availableSpace))
			} else {
				colOffsets[i+colSpan] = offset + width + 1
			}
		}
		//colIndex += colSpan
		fmt.Printf("offsets: ")
		for i, o := range colOffsets {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%d: %d", i, o)
		}
		fmt.Print("\n")
	}
	fmt.Printf("offsets: ")
	for i, o := range colOffsets {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("%d: %d", i, o)
	}
	fmt.Print("\n")
	/*for i, c := range cells {
	colSpan := 1
	if c.formatter != nil {
		if colIndex == 0 {
			indent := utf8.RuneCountInString(c.formatter.Content)
			offset := colOffsets[0]
			if indent > offset {
				colOffsets[0] = indent - offset
				//shiftColumnOffsets(colOffsets, 0, indent-offset)
			}
		}
		colSpan = c.formatter.ColumnSpan
	}

	width := getCellWidth(c)
	if width != 0 {
		// 2 spaces on either side
		width += 2
	} else if !c.blank {
		// 2 spaces on either side
		//width += 2
		width = 1
	}

	for j := i + 1; j < len(cells); j++ {
		nextCell := cells[j]
		if nextCell.blank {
			continue
		}
		if nextCell.formatter != nil {
			c.margin = utf8.RuneCountInString(nextCell.formatter.Content)
			fmt.Printf("cell %d: next cell adding width: %d; margin %d\n", i, width, c.margin)
			width += c.margin
			break
		}
	}

	c.width = width

	if colIndex+colSpan > len(colOffsets)-1 {
		fmt.Printf("cell past offsets %d: width: %d; %s\n", i, width, c.value)
		return
	}

	offset := colOffsets[colIndex]
	nextOffset := colOffsets[colIndex+colSpan]
	availableSpace := max(0, (nextOffset-offset)-1)
	fmt.Printf("cell %d colIndex %d colSpan %d width %d availableSpace %d (%d - %d) %s\n", i, colIndex, colSpan, width, availableSpace, nextOffset, offset, c.value)
	if availableSpace <= width {
		colOffsets[colIndex+colSpan] = offset + width + 1
		shift := (offset + width + 1) - nextOffset
		for j := colIndex + colSpan + 1; j < len(colOffsets); j++ {
			colOffsets[j] += shift
		}
	}
	/*if offset == nextOffset {
		fmt.Printf("cell %d: unexpanded offset %d @ index %d (%d+%d); width: %d; %s\n", i, nextOffset, colIndex+colSpan, colIndex, colSpan, width, c.value)
		shiftColumnOffsets(colOffsets, colIndex+colSpan, width+1)
	} else {
		availableSpace := (nextOffset - offset)
		fmt.Printf("cell %d colIndex %d colSpan %d width %d availableSpace %d (%d - %d) %s\n", i, colIndex, colSpan, width, availableSpace, nextOffset, offset, c.value)
		if availableSpace <= width {
			shiftColumnOffsets(colOffsets, colIndex+colSpan, (width-availableSpace)+1)
		}
	}*/
	/*availableSpace := max(0, (colOffsets[colIndex+colSpan]-colOffsets[colIndex])-1)
		fmt.Printf("cells %d colIndex %d colSpan %d width %d availableSpace %d (%d - %d - 1)\n", len(cells), colIndex, colSpan, width, availableSpace, colOffsets[colIndex+colSpan], colOffsets[colIndex])
		if availableSpace < width {
			shiftColumnOffsets(colOffsets, colIndex+colSpan, (width-availableSpace)+1)
		}
		fmt.Printf("\tavailableSpace %d (%d - %d)\n", (colOffsets[colIndex+colSpan]-colOffsets[colIndex])-1, colOffsets[colIndex+colSpan], colOffsets[colIndex])* /
		colIndex += colSpan
		fmt.Printf("offsets: ")
		for i, o := range colOffsets {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%d: %d", i, o)
		}
		fmt.Print("\n")
	}*/
	return
}

func renderTableSubElements(cxt *output.Context, t *types.Table, tbl *table) (err error) {
	if t.Header != nil {
		fmt.Printf("rendering headers %d\n", len(t.Header.Cells))
		for i, c := range t.Header.Cells {
			renderContext := output.NewContext(cxt, cxt.Doc)
			err = RenderElements(renderContext, "", c.Elements)
			if err != nil {
				fmt.Printf("rendering header %d err: %v\n", i, err)
				return
			}
			fmt.Printf("rendering header %d: %s\n", i, renderContext.String())
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
	return
}

func getCellWidth(c *tableCell) int {
	lines := strings.Split(c.value, "\n")
	if len(lines) == 1 {
		return utf8.RuneCountInString(c.value)
	}
	return utf8.RuneCountInString(strings.TrimSpace(lines[len(lines)-1]))
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
