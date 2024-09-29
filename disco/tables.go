package disco

import (
	"fmt"
	"log/slog"
	"slices"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func (b *Ball) ensureTableOptions(els asciidoc.Set) {
	if !b.options.normalizeTableOptions {
		return
	}
	parse.Search(els, func(t *asciidoc.Table) parse.SearchShould {
		var excludedIndexes []int
		for i, attr := range t.Attributes() {
			na, ok := attr.(*asciidoc.NamedAttribute)
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
		return parse.SearchShouldContinue
	})

}

func (b *Ball) addMissingColumns(section *spec.Section, ti *spec.TableInfo, tableTemplate matter.Table, entityType types.EntityType) (err error) {
	if !b.options.addMissingColumns {
		return
	}
	if b.errata.IgnoreSection(section.Name, errata.DiscoPurposeTableAddMissingColumns) {
		return
	}
	ti.Element.DeleteAttribute(asciidoc.AttributeNameColumns)
	var order []matter.TableColumn
	if len(tableTemplate.RequiredColumns) > 0 {
		order = tableTemplate.RequiredColumns
	} else {
		order = tableTemplate.ColumnOrder
	}
	for _, column := range order {
		if _, ok := ti.ColumnMap[column]; !ok {
			_, err = b.appendColumn(ti, column, entityType)
			if err != nil {
				return
			}
		}
	}
	return
}

func (b *Ball) appendColumn(ti *spec.TableInfo, column matter.TableColumn, entityType types.EntityType) (appendedIndex int, err error) {
	rows := ti.Rows
	if len(rows) == 0 {
		appendedIndex = -1
		return
	}
	ti.Element.ColumnCount += 1
	var cols *asciidoc.TableColumnsAttribute
	var ok bool
	for _, a := range ti.Element.Attributes() {
		cols, ok = a.(*asciidoc.TableColumnsAttribute)
		if ok {
			break
		}
	}
	if cols != nil {
		cols.Columns = append(cols.Columns, asciidoc.NewTableColumn())
	}
	appendedIndex = len(rows[0].Set)
	for i, row := range rows {
		cell := &asciidoc.TableCell{}
		last := row.Cell(len(row.Set) - 1)
		if i == ti.HeaderRowIndex {
			if last.Format != nil {
				cell.Format = asciidoc.NewTableCellFormat()
				cell.Format.HorizontalAlign = last.Format.HorizontalAlign
				cell.Format.VerticalAlign = last.Format.VerticalAlign
				cell.Format.Style = last.Format.Style
			}
			name, ok := matter.TableColumnNames[column]
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
				err = setCellString(cell, b.getDefaultColumnValue(ti, row, column, entityType))
			}
		}
		row.Append(cell)
	}
	ti.ColumnMap[column] = appendedIndex
	return
}

func (b *Ball) getDefaultColumnValue(ti *spec.TableInfo, row *asciidoc.TableRow, column matter.TableColumn, entityType types.EntityType) string {
	switch column {
	case matter.TableColumnConformance:
		switch entityType {
		case types.EntityTypeBitmapValue, types.EntityTypeEnumValue:
			return "M"
		}
	case matter.TableColumnName:
		switch entityType {
		case types.EntityTypeBitmapValue, types.EntityTypeEnumValue:
			val, _ := ti.ReadValue(row, matter.TableColumnSummary, matter.TableColumnDescription)
			if val != "" {
				return matter.Case(val)
			}
		}
	case matter.TableColumnSummary:
		switch entityType {
		case types.EntityTypeBitmapValue, types.EntityTypeEnumValue:
			val, _ := ti.ReadValue(row, matter.TableColumnName)
			if val != "" {
				return matter.Uncase(val)
			}
		}
	case matter.TableColumnAccess:
		switch entityType {
		case types.EntityTypeEvent, types.EntityTypeAttribute:
			return "R V"
		case types.EntityTypeStructField:
			return ""
		}
	}
	slog.Debug("no default value for column", slog.String("column", column.String()), slog.String("entity", entityType.String()))
	return ""
}

func (b *Ball) reorderColumns(doc *spec.Doc, section *spec.Section, ti *spec.TableInfo, tableType matter.TableType) (err error) {
	if !b.options.reorderColumns {
		return
	}
	if b.errata.IgnoreSection(section.Name, errata.DiscoPurposeTableReorderColumns) {
		return
	}
	rows := ti.Rows
	columnMap := ti.ColumnMap
	extraColumns := ti.ExtraColumns
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
			tableCells := row.TableCells()
			cell := row.Cell(i)
			if cell.Blank {
				// OK, this cell is blank, likely because of a horizontal span
				k := i - 1
				for k >= 0 {
					left := tableCells[k]
					if left.Format != nil && left.Format.Span.Column.Value > 1 {
						left.Format.Span.Column = asciidoc.One(left.Format.Span.Column.Value - 1)
						break
					}
					k--
				}
			}
		}
	}
	for i, row := range rows {
		newCells := make(asciidoc.Set, index)
		tableCells := row.TableCells()
		if len(tableCells) != len(newColumnIndexes) {
			err = fmt.Errorf("cell length mismatch; row %d has %d cells, expected %d", i, len(tableCells), len(newColumnIndexes))
			return
		}
		for j, cell := range tableCells {
			newIndex := newColumnIndexes[j]
			if newIndex < 0 {
				continue
			}
			newCells[newIndex] = cell
		}
		newCells = slices.Clip(newCells)
		row.Set = newCells
	}
	var cols *asciidoc.TableColumnsAttribute
	for _, a := range ti.Element.Attributes() {
		cols, ok = a.(*asciidoc.TableColumnsAttribute)
		if ok {
			break
		}
	}
	if cols != nil {
		newColumns := make([]*asciidoc.TableColumn, index)
		for i, col := range cols.Columns {
			newIndex := newColumnIndexes[i]
			if newIndex < 0 {
				continue
			}
			newColumns[newIndex] = col
		}
		cols.Columns = newColumns
	}
	err = ti.Rescan(doc)
	return
}

func setCellString(cell *asciidoc.TableCell, v string) (err error) {
	se := asciidoc.NewString(v)
	err = cell.SetElements(asciidoc.Set{se})
	return
}

func setCellValue(cell *asciidoc.TableCell, val asciidoc.Set) (err error) {
	return cell.SetElements(val)
}

func copyCells(rows []*asciidoc.TableRow, headerRowIndex int, fromIndex int, toIndex int, transformer func(s string) string) (err error) {
	for i, row := range rows {
		if i == headerRowIndex {
			continue
		}
		tableCells := row.TableCells()
		var value string
		value, err = spec.RenderTableCell(tableCells[fromIndex])
		if err != nil {
			return
		}
		if transformer != nil {
			value = transformer(value)
		}
		err = setCellString(tableCells[toIndex], value)
		if err != nil {
			return
		}
	}
	return
}

func (b *Ball) renameTableHeaderCells(doc *spec.Doc, section *spec.Section, table *spec.TableInfo, overrides map[matter.TableColumn]matter.TableColumn) (err error) {
	if !b.options.renameTableHeaders {
		return
	}
	if b.errata.IgnoreSection(section.Name, errata.DiscoPurposeTableRenameHeaders) {
		return
	}
	headerRow := table.Rows[table.HeaderRowIndex]
	reverseMap := make(map[int]matter.TableColumn)
	for k, v := range table.ColumnMap {
		reverseMap[v] = k
	}
	for i, cell := range headerRow.TableCells() {
		tc, ok := reverseMap[i]
		if !ok {
			continue
		}
		overrideColumn, hasOverride := overrides[tc]
		if hasOverride {
			existingIndex, overrideAlreadyExists := table.ColumnMap[overrideColumn]
			if overrideAlreadyExists && existingIndex != i {
				slog.Warn("Can not rename column; column with same name already exists", slog.String("from", tc.String()), slog.String("to", overrideColumn.String()), slog.Int("existingColumnIndex", existingIndex))
				continue
			}
			tc = overrideColumn
		}
		name, ok := matter.GetColumnName(tc, overrides)
		if ok {
			err = setCellString(cell, name)
			if err != nil {
				return
			}
		}
	}
	err = table.Rescan(doc)
	return
}
