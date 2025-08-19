package parse

import (
	"context"
	"fmt"
	"strings"
	"unicode"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/render"
)

func parseTable(attributes any, els any) (table *asciidoc.Table, err error) {
	table = &asciidoc.Table{}

	if attributes != nil {
		as, ok := attributes.([]asciidoc.Attribute)
		if !ok {
			err = fmt.Errorf("could not cast table attributes to attribute list: %T", attributes)
			return
		}
		err = table.ReadAttributes(table, as...)
		if err != nil {
			return
		}
	}
	var elements []any
	switch els := els.(type) {
	case nil:
		return
	case []any:
		elements = els
	}
	table.ColumnCount, err = getColumnCount(table, elements)
	if err != nil {
		return
	}
	for _, el := range elements {
		if pos, ok := el.(asciidoc.HasPosition); ok {
			copyPosition(pos, table)
			break
		}
	}
	var rows []asciidoc.Element
	rows, err = parseTableRows(table, elements)
	if err != nil {
		return
	}
	table.Append(rows...)
	return
}

func parseTableRows(table *asciidoc.Table, elements []any) (rows asciidoc.Elements, err error) {
	cellIndex := 0
	var currentTableRow *asciidoc.TableRow
	colSkip := make(map[int]int)
	for _, el := range elements {
		switch el := el.(type) {
		case []*asciidoc.TableCell:
			for _, cell := range el {
				if currentTableRow == nil || cellIndex >= table.ColumnCount {
					currentTableRow = &asciidoc.TableRow{Parent: table}
					currentTableRow.SetPosition(cell.Position())
					currentTableRow.SetPath(cell.Path())
					rows = append(rows, currentTableRow)
					cellIndex = 0
				}
				cell.Parent = currentTableRow
				for cellIndex < table.ColumnCount {
					skip, ok := colSkip[cellIndex]
					if !ok || skip == 0 {
						break
					}
					currentTableRow.Append(&asciidoc.TableCell{Blank: true})
					colSkip[cellIndex] = skip - 1
					cellIndex++
				}
				if cellIndex >= table.ColumnCount {
					currentTableRow = &asciidoc.TableRow{}
					currentTableRow.SetPosition(cell.Position())
					currentTableRow.SetPath(cell.Path())
					rows = append(rows, currentTableRow)
					cellIndex = 0
				}
				currentTableRow.Append(cell)
				if cell.Format != nil {
					rowSpan := cell.Format.Span.Row.Value
					colSpan := cell.Format.Span.Column.Value
					if rowSpan > 1 {
						colSkip[cellIndex] = rowSpan - 1
					}
					cellIndex++
					if cellIndex >= table.ColumnCount {
						continue
					}
					if colSpan > 1 {
						for i := 0; i < colSpan-1; i++ {
							currentTableRow.Append(&asciidoc.TableCell{Blank: true})
							cellIndex++
						}
					}
				} else {
					cellIndex++
				}
			}
		case asciidoc.Element:
			rows = append(rows, el)
		case asciidoc.Elements:
			rows = append(rows, el...)
		default:
			err = fmt.Errorf("unexpected type in table elements: %T", el)
			return
		}
	}
	if currentTableRow == nil {
		return
	}
	for cellIndex < table.ColumnCount {
		currentTableRow.Append(&asciidoc.TableCell{Blank: true})
		cellIndex++
	}
	return
}

func getColumnCount(table *asciidoc.Table, els []any) (columnCount int, err error) {
	for _, a := range table.Attributes() {
		switch a := a.(type) {
		case *asciidoc.TableColumnsAttribute:
			for _, c := range a.Columns {
				if c.Multiplier.IsSet {
					columnCount += c.Multiplier.Value
				} else {
					columnCount++
				}
			}
			if columnCount > 0 {
				return
			}
		}
	}
	if len(els) == 0 {
		return
	}
	var foundTableCell bool

	// First, we look for an empty line after a row of table cells
	for _, el := range els[0 : len(els)-1] {
		if _, ok := el.(asciidoc.EmptyLine); ok && foundTableCell {
			break
		}
		if cells, ok := el.([]*asciidoc.TableCell); ok {
			if foundTableCell {
				break
			}
			for _, cell := range cells {
				if cell.Format != nil && cell.Format.Span.Column.IsSet {
					columnCount += cell.Format.Span.Column.Value
				} else {
					columnCount++
				}
			}
			foundTableCell = true
		}
	}
	// OK, no empty lines, so we just take the first full row
	columnCount = 0
	for _, el := range els {
		if cells, ok := el.([]*asciidoc.TableCell); ok {
			for _, cell := range cells {
				if cell.Format != nil && cell.Format.Span.Column.IsSet {
					columnCount += cell.Format.Span.Column.Value
				} else {
					columnCount++
				}
				children := cell.Children()
				if len(children) == 0 {
					continue
				}
				lastChild := children[len(children)-1]
				switch lastChild.(type) {
				case *asciidoc.NewLine, *asciidoc.EmptyLine:
					return
				}
			}
			return
		}
	}
	err = fmt.Errorf("unable to determine column count")
	return
}

func newTableCell(format any) *asciidoc.TableCell {
	if format, ok := format.(*asciidoc.TableCellFormat); ok {
		return asciidoc.NewTableCell(format)
	}
	return asciidoc.NewTableCell(nil)
}

func reparseTables(els asciidoc.Elements) (err error) {

	for _, el := range els {
		switch el := el.(type) {
		case *asciidoc.Table:
			err = ReparseTable(el, el.Children())
			if err != nil {
				return
			}
		}
	}
	return
}

func ReparseTable(table *asciidoc.Table, elements asciidoc.Elements) (err error) {
	columnStyles := make([]asciidoc.TableCellStyle, table.ColumnCount)
	for _, a := range table.Attributes() {
		switch a := a.(type) {
		case *asciidoc.TableColumnsAttribute:
			for i, c := range a.Columns {
				if c.Style.IsSet {
					columnStyles[i] = c.Style.Value
				}
			}

		}
	}
	for _, e := range elements {
		switch e := e.(type) {
		case *asciidoc.TableRow:
			for i, c := range e.TableCells() {
				if c.Blank {
					continue
				}
				var style asciidoc.TableCellStyle
				if c.Format != nil && c.Format.Style.IsSet {
					style = c.Format.Style.Value
				} else {
					style = columnStyles[i]
				}
				switch style {
				case asciidoc.TableCellStyleAsciiDoc:
					err = parseBlockCell(c)
				case asciidoc.TableCellStyleLiteral: // Leave the strings alone for a literal cell
				default:
					var tcels asciidoc.Elements
					tcels, err = trimCell(c)
					if err != nil {
						return
					}
					c.SetChildren(tcels)
				}
				if err != nil {
					return
				}
			}
		}
	}
	return
}

func parseBlockCell(tc *asciidoc.TableCell) error {
	out := render.NewUnwrappedTarget(context.Background())
	err := render.Elements(out, "", tc.Children()...)
	if err != nil {
		return err
	}
	val := out.String()
	if val == "" {
		return nil
	}

	line, col, offset := tc.Position()
	col++
	vals, err := Parse(tc.Path(), []byte(val), initialPosition(line, col, offset))
	if err != nil {
		return err
	}
	els, ok := vals.(asciidoc.Elements)
	if !ok {
		return fmt.Errorf("unexpected type for table cell set: %T", vals)
	}
	if len(els) > 0 {
		if _, ok := els[0].(*asciidoc.EmptyLine); ok {
			// If the first character is a new line, the parser will interpret that as an empty line, since it thinks there's nothing before it
			els[0] = &asciidoc.NewLine{}
		}
	}
	els, err = coalesce(els)
	if err != nil {
		return err
	}
	tc.SetChildren(els)
	return nil
}

func trimCell(tc *asciidoc.TableCell) (els asciidoc.Elements, err error) {
	els = tc.Children()
	leftIndex := 0
	rightIndex := len(els) - 1
	switch len(els) {
	case 0:
	case 1:
		switch e := els[0].(type) {
		case *asciidoc.String:
			e.Value = strings.TrimSpace(e.Value)
		}
	default:
		switch e := els[leftIndex].(type) {
		case *asciidoc.String:
			e.Value = strings.TrimLeftFunc(e.Value, unicode.IsSpace)
			if len(e.Value) == 0 {
				leftIndex = 1
			}
		}
		switch e := els[rightIndex].(type) {
		case *asciidoc.String:
			e.Value = strings.TrimRightFunc(e.Value, unicode.IsSpace)
			if len(e.Value) == 0 {
				rightIndex -= 1
			}
		}
	}
	if leftIndex == 0 && rightIndex == len(els)-1 {
		return
	}

	els = els[leftIndex : rightIndex+1]

	return
}
