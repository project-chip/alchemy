package parse

import (
	"fmt"

	"github.com/project-chip/alchemy/asciidoc"
)

func parseTable(attributes any, els any) (table *asciidoc.Table, err error) {
	/*fmt.Fprintf(os.Stderr, "parseTable: %T\n", els)
	for _, el := range els.([]any) {
		fmt.Fprintf(os.Stderr, "\t%T\n", el)
		switch el := el.(type) {
		case []any:
			for _, e := range el {
				fmt.Fprintf(os.Stderr, "\t\t%T\n", e)
			}
		}

	}*/
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
	cellIndex := 0
	var currentTableRow *asciidoc.TableRow
	colSkip := make(map[int]int)
	var copiedPosition bool
	for _, el := range elements {
		if !copiedPosition {
			if pos, ok := el.(asciidoc.HasPosition); ok {
				copyPosition(pos, table)
				copiedPosition = true
			}
		}
		switch el := el.(type) {
		case []*asciidoc.TableCell:
			for _, cell := range el {
				if currentTableRow == nil || cellIndex >= table.ColumnCount {
					currentTableRow = &asciidoc.TableRow{Parent: table}
					currentTableRow.SetPosition(cell.Position())
					currentTableRow.SetPath(cell.Path())
					table.Append(currentTableRow)
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
					table.Append(currentTableRow)
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
			table.Append(el)
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
				children := cell.Elements()
				if len(children) == 0 {
					continue
				}
				lastChild := children[len(children)-1]
				switch lastChild.(type) {
				case *asciidoc.NewLine, asciidoc.EmptyLine:
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

func reparseTables(els asciidoc.Set) (err error) {

	for _, el := range els {
		switch el := el.(type) {
		case *asciidoc.Table:
			err = reparseTable(el)
			if err != nil {
				return
			}
		}
	}
	return
}

func reparseTable(table *asciidoc.Table) (err error) {
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
	for _, e := range table.Elements() {
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
					err = parseInlineCell(c)
				}
				if err != nil {
					return
				}
			}
		}
	}
	return
}

func parseInlineCell(tc *asciidoc.TableCell) error {
	val, err := renderPreParsedDoc(tc.Elements())
	if err != nil {
		return err
	}
	if val == "" {
		return nil
	}
	line, col, offset := tc.Position()
	col++
	vals, err := Parse("", []byte(val), Entrypoint("TableCellInlineContent"), initialPosition(line, col, offset))
	if err != nil {
		return err
	}
	els, ok := vals.(asciidoc.Set)
	if !ok {
		return fmt.Errorf("unexpected type for table cell set: %T", vals)
	}
	if len(els) > 0 {
		if _, ok := els[0].(asciidoc.EmptyLine); ok {
			// If the first character is a new line, the parser will interpret that as an empty line, since it thinks there's nothing before it
			els[0] = &asciidoc.NewLine{}
		}
	}

	els, err = coalesce(els)
	if err != nil {
		return err
	}

	tc.SetElements(els)
	return nil
}

func parseBlockCell(tc *asciidoc.TableCell) error {
	val, err := renderPreParsedDoc(tc.Elements())
	if err != nil {
		return err
	}
	if val == "" {
		return nil
	}
	line, col, offset := tc.Position()
	col++
	vals, err := Parse("", []byte(val), initialPosition(line, col, offset))
	if err != nil {
		return err
	}
	els, ok := vals.(asciidoc.Set)
	if !ok {
		return fmt.Errorf("unexpected type for table cell set: %T", vals)
	}
	if len(els) > 0 {
		if _, ok := els[0].(asciidoc.EmptyLine); ok {
			// If the first character is a new line, the parser will interpret that as an empty line, since it thinks there's nothing before it
			els[0] = &asciidoc.NewLine{}
		}
	}
	els, err = coalesce(els)
	if err != nil {
		return err
	}
	tc.SetElements(els)
	return nil
}
