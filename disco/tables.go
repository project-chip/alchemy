package disco

import (
	"fmt"
	"log/slog"
	"slices"

	"github.com/hasty/adoc/elements"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/internal/parse"
	"github.com/hasty/alchemy/matter"
)

func (b *Ball) ensureTableOptions(elements []any) {
	if !b.options.normalizeTableOptions {
		return
	}
	parse.Search(elements, func(t *elements.Table) bool {
		if t.Attributes == nil {
			t.Attributes = make(elements.Attributes)
		}
		var excludedKeys []string
		for k := range t.Attributes {
			if _, ok := matter.AllowedTableAttributes[k]; !ok {
				excludedKeys = append(excludedKeys, k)
			}
		}
		for _, k := range excludedKeys {
			delete(t.Attributes, k)
		}
		for k, v := range matter.AllowedTableAttributes {
			if v != nil {
				t.Attributes[k] = v
			}
		}
		return false
	})

}

func (b *Ball) addMissingColumns(doc *ascii.Doc, section *ascii.Section, table *elements.Table, rows []*elements.TableRow, tableTemplate matter.Table, overrides map[matter.TableColumn]string, headerRowIndex int, columnMap ascii.ColumnIndex, entityType matterelements.EntityType) (err error) {
	if !b.options.addMissingColumns {
		return
	}
	delete(table.AttributeList, "cols")
	var order []matter.TableColumn
	if len(tableTemplate.RequiredColumns) > 0 {
		order = tableTemplate.RequiredColumns
	} else {
		order = tableTemplate.ColumnOrder
	}
	for _, column := range order {
		if _, ok := columnMap[column]; !ok {
			_, err = b.appendColumn(rows, columnMap, headerRowIndex, column, overrides, entityType)
			if err != nil {
				return
			}
		}
	}
	return
}

func (b *Ball) appendColumn(rows []*elements.TableRow, columnMap ascii.ColumnIndex, headerRowIndex int, column matter.TableColumn, overrides map[matter.TableColumn]string, entityType matterelements.EntityType) (appendedIndex int, err error) {
	if len(rows) == 0 {
		appendedIndex = -1
		return
	}
	appendedIndex = len(rows[0].Cells)
	for i, row := range rows {
		cell := &elements.TableCell{}
		last := row.Cells[len(row.Cells)-1]
		if i == headerRowIndex {
			if last.Formatter != nil {
				cell.Formatter = &elements.TableCellFormat{
					ColumnSpan:          1,
					RowSpan:             1,
					HorizontalAlignment: last.Formatter.HorizontalAlignment,
					VerticalAlignment:   last.Formatter.VerticalAlignment,
					Style:               last.Formatter.Style,
					Content:             last.Formatter.Content,
				}
			}
			name, ok := matter.GetColumnName(column, overrides)
			if !ok {
				err = fmt.Errorf("unknown column name: %s", column.String())
				return
			}
			err = setCellString(cell, name)
			if err != nil {
				return
			}
		} else {
			cell.Blank = last.Blank
			if !cell.Blank {
				err = setCellString(cell, b.getDefaultColumnValue(entityType, column, row, columnMap))
			}
		}
		row.Cells = append(row.Cells, cell)
	}
	columnMap[column] = appendedIndex
	return
}

func (b *Ball) getDefaultColumnValue(entityType matterelements.EntityType, column matter.TableColumn, row *elements.TableRow, columnMap ascii.ColumnIndex) string {
	switch column {
	case matter.TableColumnConformance:
		switch entityType {
		case matterelements.EntityTypeBitmapValue, matterelements.EntityTypeEnumValue:
			return "M"
		}
	case matter.TableColumnName:
		switch entityType {
		case matterelements.EntityTypeBitmapValue, matterelements.EntityTypeEnumValue:
			val, _ := ascii.ReadRowValue(b.doc, row, columnMap, matter.TableColumnSummary)
			if val != "" {
				return matter.Case(val)
			}
		}
	case matter.TableColumnSummary:
		switch entityType {
		case matterelements.EntityTypeBitmapValue, matterelements.EntityTypeEnumValue:
			val, _ := ascii.ReadRowValue(b.doc, row, columnMap, matter.TableColumnName)
			if val != "" {
				return matter.Uncase(val)
			}
		}
	case matter.TableColumnAccess:
		switch entityType {
		case matterelements.EntityTypeEvent, matterelements.EntityTypeAttribute:
			return "R V"
		case matterelements.EntityTypeField:
			return ""
		}
	}
	slog.Debug("no default value for column", slog.String("column", column.String()), slog.String("entity", entityType.String()))
	return ""
}

func (b *Ball) reorderColumns(doc *ascii.Doc, section *ascii.Section, ti *tableInfo, tableType matter.TableType) (err error) {
	if !b.options.reorderColumns {
		return
	}
	rows := ti.rows
	columnMap := ti.columnMap
	extraColumns := ti.extraColumns
	tableTemplate, ok := matter.Tables[tableType]
	if !ok {
		err = fmt.Errorf("missing table template for table type: %v", tableType)
		return
	}
	order := tableTemplate.ColumnOrder
	if len(rows) == 0 {
		return
	}
	newColumnIndexes := make([]int, len(columnMap)+len(extraColumns))
	// In this little slice, -1 means "we don't know what to do with this column yet" and -2 means "we're excluding this column"
	for i := range newColumnIndexes {
		newColumnIndexes[i] = -1
	}
	var index int
	for _, column := range order {
		if offset, ok := columnMap[column]; ok {
			// column exists in existing order
			newColumnIndexes[offset] = index
			index++
		}
	}
	// Recognized columns, but not in the column order
	for tc, i := range columnMap {
		if newColumnIndexes[i] >= 0 {
			continue
		}
		var banned bool
		for _, bc := range tableTemplate.BannedColumns {
			if tc == bc {
				banned = true
				break
			}
		}
		if banned {
			newColumnIndexes[i] = -2
			continue
		}
		newColumnIndexes[i] = index
		index++

	}
	// Unrecognized columns come last
	for i, v := range newColumnIndexes {
		if v == -1 {
			newColumnIndexes[i] = index
			index++
		}
	}
	for i, v := range newColumnIndexes {
		if v != -2 {
			continue
		}
		for _, row := range rows {
			cell := row.Cells[i]
			if cell.Blank {
				// OK, this cell is blank, likely because of a horizontal span
				k := i - 1
				for k >= 0 {
					left := row.Cells[k]
					if left.Formatter != nil && left.Formatter.ColumnSpan > 1 {
						left.Formatter.ColumnSpan = left.Formatter.ColumnSpan - 1
						break
					}
					k--
				}
			}
		}
	}
	for i, row := range rows {
		newCells := make([]*elements.TableCell, index)
		if len(row.Cells) != len(newColumnIndexes) {
			err = fmt.Errorf("cell length mismatch; row %d has %d cells, expected %d", i, len(row.Cells), len(newColumnIndexes))
			return
		}
		for j, cell := range row.Cells {
			newIndex := newColumnIndexes[j]
			if newIndex < 0 {
				continue
			}
			newCells[newIndex] = cell
		}
		newCells = slices.Clip(newCells)
		row.Cells = newCells
	}
	ti.headerRow, ti.columnMap, ti.extraColumns, err = ascii.MapTableColumns(doc, ti.rows)
	return
}

func setCellString(cell *elements.TableCell, v string) (err error) {
	var p *elements.Paragraph

	if len(cell.Elements) == 0 {
		p, err = elements.NewParagraph(nil)
		if err != nil {
			return
		}
		err = cell.SetElements([]any{p})
		if err != nil {
			return
		}
	} else {
		var ok bool
		p, ok = cell.Elements[0].(*elements.Paragraph)
		if !ok {
			return fmt.Errorf("table cell does not have paragraph child")
		}
	}
	se, _ := elements.NewStringElement(v)
	err = p.SetElements([]any{se})
	return
}

func setCellValue(cell *elements.TableCell, val []any) (err error) {
	var p *elements.Paragraph

	if len(cell.Elements) == 0 {
		p, err = elements.NewParagraph(nil)
		if err != nil {
			return
		}
		err = cell.SetElements([]any{p})
		if err != nil {
			return
		}
	} else {
		var ok bool
		p, ok = cell.Elements[0].(*elements.Paragraph)
		if !ok {
			return fmt.Errorf("table cell does not have paragraph child")
		}
	}
	err = p.SetElements(val)
	return
}

func copyCells(rows []*elements.TableRow, headerRowIndex int, fromIndex int, toIndex int, transformer func(s string) string) (err error) {
	for i, row := range rows {
		if i == headerRowIndex {
			continue
		}
		var value string
		value, err = ascii.RenderTableCell(row.Cells[fromIndex])
		if err != nil {
			return
		}
		if transformer != nil {
			value = transformer(value)
		}
		err = setCellString(row.Cells[toIndex], value)
		if err != nil {
			return
		}
	}
	return
}

func (b *Ball) renameTableHeaderCells(rows []*elements.TableRow, headerRowIndex int, columnMap ascii.ColumnIndex, overrides map[matter.TableColumn]string) (err error) {
	if !b.options.renameTableHeaders {
		return
	}
	headerRow := rows[headerRowIndex]
	reverseMap := make(map[int]matter.TableColumn)
	for k, v := range columnMap {
		reverseMap[v] = k
	}
	for i, cell := range headerRow.Cells {
		tc, ok := reverseMap[i]
		if !ok {
			continue
		}
		name, ok := matter.GetColumnName(tc, overrides)
		if ok {
			err = setCellString(cell, name)
			if err != nil {
				return
			}
		}
	}
	return
}
