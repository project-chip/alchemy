package disco

import (
	"fmt"
	"log/slog"
	"slices"

	"github.com/hasty/adoc/elements"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/internal/parse"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/types"
)

func (b *Ball) ensureTableOptions(els elements.Set) {
	if !b.options.normalizeTableOptions {
		return
	}
	parse.Search(els, func(t *elements.Table) bool {
		var excludedIndexes []int
		for i, attr := range t.Attributes() {
			na, ok := attr.(*elements.NamedAttribute)
			if !ok {
				continue
			}
			if _, ok := matter.AllowedTableAttributes[na.Name]; !ok {
				excludedIndexes = append(excludedIndexes, i)
			}
		}
		for i := len(excludedIndexes) - 1; i >= 0; i-- {
			t.AttributeList = slices.Delete(t.AttributeList, i, i)
		}
		for k, v := range matter.AllowedTableAttributes {
			if v != nil {
				t.SetAttribute(k, v)
			}
		}
		return false
	})

}

func (b *Ball) addMissingColumns(doc *ascii.Doc, section *ascii.Section, table *elements.Table, rows []*elements.TableRow, tableTemplate matter.Table, overrides map[matter.TableColumn]string, headerRowIndex int, columnMap ascii.ColumnIndex, entityType types.EntityType) (err error) {
	if !b.options.addMissingColumns {
		return
	}
	table.DeleteAttribute(elements.AttributeNameColumns)
	var order []matter.TableColumn
	if len(tableTemplate.RequiredColumns) > 0 {
		order = tableTemplate.RequiredColumns
	} else {
		order = tableTemplate.ColumnOrder
	}
	for _, column := range order {
		if _, ok := columnMap[column]; !ok {
			_, err = b.appendColumn(table, columnMap, headerRowIndex, column, overrides, entityType)
			if err != nil {
				return
			}
		}
	}
	return
}

func (b *Ball) appendColumn(table *elements.Table, columnMap ascii.ColumnIndex, headerRowIndex int, column matter.TableColumn, overrides map[matter.TableColumn]string, entityType types.EntityType) (appendedIndex int, err error) {
	rows := table.TableRows()
	if len(rows) == 0 {
		appendedIndex = -1
		return
	}
	table.ColumnCount += 1
	var cols *elements.TableColumnsAttribute
	var ok bool
	for _, a := range table.Attributes() {
		cols, ok = a.(*elements.TableColumnsAttribute)
		if ok {
			break
		}
	}
	if cols != nil {
		cols.Columns = append(cols.Columns, elements.NewTableColumn())
	}
	appendedIndex = len(rows[0].TableCells)
	for i, row := range rows {
		cell := &elements.TableCell{}
		last := row.TableCells[len(row.TableCells)-1]
		if i == headerRowIndex {
			if last.Format != nil {
				cell.Format = elements.NewTableCellFormat()
				cell.Format.HorizontalAlign = last.Format.HorizontalAlign
				cell.Format.VerticalAlign = last.Format.VerticalAlign
				cell.Format.Style = last.Format.Style
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
		row.TableCells = append(row.TableCells, cell)
	}
	columnMap[column] = appendedIndex
	return
}

func (b *Ball) getDefaultColumnValue(entityType types.EntityType, column matter.TableColumn, row *elements.TableRow, columnMap ascii.ColumnIndex) string {
	switch column {
	case matter.TableColumnConformance:
		switch entityType {
		case types.EntityTypeBitmapValue, types.EntityTypeEnumValue:
			return "M"
		}
	case matter.TableColumnName:
		switch entityType {
		case types.EntityTypeBitmapValue, types.EntityTypeEnumValue:
			val, _ := ascii.ReadRowValue(b.doc, row, columnMap, matter.TableColumnSummary)
			if val != "" {
				return matter.Case(val)
			}
		}
	case matter.TableColumnSummary:
		switch entityType {
		case types.EntityTypeBitmapValue, types.EntityTypeEnumValue:
			val, _ := ascii.ReadRowValue(b.doc, row, columnMap, matter.TableColumnName)
			if val != "" {
				return matter.Uncase(val)
			}
		}
	case matter.TableColumnAccess:
		switch entityType {
		case types.EntityTypeEvent, types.EntityTypeAttribute:
			return "R V"
		case types.EntityTypeField:
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
			cell := row.TableCells[i]
			if cell.Blank {
				// OK, this cell is blank, likely because of a horizontal span
				k := i - 1
				for k >= 0 {
					left := row.TableCells[k]
					if left.Format != nil && left.Format.Span.Column.Value > 1 {
						left.Format.Span.Column = elements.One(left.Format.Span.Column.Value - 1)
						break
					}
					k--
				}
			}
		}
	}
	for i, row := range rows {
		newCells := make([]*elements.TableCell, index)
		if len(row.TableCells) != len(newColumnIndexes) {
			err = fmt.Errorf("cell length mismatch; row %d has %d cells, expected %d", i, len(row.TableCells), len(newColumnIndexes))
			return
		}
		for j, cell := range row.TableCells {
			newIndex := newColumnIndexes[j]
			if newIndex < 0 {
				continue
			}
			newCells[newIndex] = cell
		}
		newCells = slices.Clip(newCells)
		row.TableCells = newCells
	}
	var cols *elements.TableColumnsAttribute
	for _, a := range ti.element.Attributes() {
		cols, ok = a.(*elements.TableColumnsAttribute)
		if ok {
			break
		}
	}
	if cols != nil {
		newColumns := make([]*elements.TableColumn, index)
		for i, col := range cols.Columns {
			newIndex := newColumnIndexes[i]
			if newIndex < 0 {
				continue
			}
			newColumns[newIndex] = col
		}
		cols.Columns = newColumns
	}
	ti.headerRow, ti.columnMap, ti.extraColumns, err = ascii.MapTableColumns(doc, ti.rows)
	return
}

func setCellString(cell *elements.TableCell, v string) (err error) {
	se := elements.NewString(v)
	err = cell.SetElements(elements.Set{se})
	return
}

func setCellValue(cell *elements.TableCell, val elements.Set) (err error) {
	return cell.SetElements(val)
}

func copyCells(rows []*elements.TableRow, headerRowIndex int, fromIndex int, toIndex int, transformer func(s string) string) (err error) {
	for i, row := range rows {
		if i == headerRowIndex {
			continue
		}
		var value string
		value, err = ascii.RenderTableCell(row.TableCells[fromIndex])
		if err != nil {
			return
		}
		if transformer != nil {
			value = transformer(value)
		}
		err = setCellString(row.TableCells[toIndex], value)
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
	for i, cell := range headerRow.TableCells {
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
