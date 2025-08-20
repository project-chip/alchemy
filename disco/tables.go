package disco

import (
	"fmt"
	"log/slog"
	"slices"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func (b *Baller) ensureTableOptions(doc *asciidoc.Document, root asciidoc.ParentElement) {
	if !b.options.NormalizeTableOptions {
		return
	}
	parse.Search(asciidoc.RawReader, root, root.Children(), func(t *asciidoc.Table, parent asciidoc.Parent, index int) parse.SearchShould {
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

func (b *Baller) addMissingColumns(cxt *discoContext, section *asciidoc.Section, ti *spec.TableInfo, tableTemplate matter.Table, entityType types.EntityType) (err error) {
	if !b.options.AddMissingColumns {
		return
	}
	if cxt.errata.IgnoreSection(cxt.doc.SectionName(section), errata.DiscoPurposeTableAddMissingColumns) {
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
	err = ti.Rescan(cxt.doc, asciidoc.RawReader)
	return
}

func (b *Baller) appendColumn(ti *spec.TableInfo, column matter.TableColumn, entityType types.EntityType) (appendedIndex int, err error) {
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
	appendedIndex = len(rows[0].Elements)
	for i, row := range rows {
		cell := &asciidoc.TableCell{}
		last := row.Cell(len(row.Elements) - 1)
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
			setCellString(cell, name)
		} else {
			cell.Blank = last.Blank
			if !cell.Blank {
				setCellString(cell, b.getDefaultColumnValue(ti, row, column, entityType))
			}
		}
		row.Append(cell)
	}
	ti.ColumnMap[column] = appendedIndex
	return
}

func (b *Baller) getDefaultColumnValue(ti *spec.TableInfo, row *asciidoc.TableRow, column matter.TableColumn, entityType types.EntityType) string {
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

func (b *Baller) reorderColumns(cxt *discoContext, section *asciidoc.Section, ti *spec.TableInfo, tableType matter.TableType) (err error) {
	if !b.options.ReorderColumns {
		return
	}
	if cxt.errata.IgnoreSection(cxt.doc.SectionName(section), errata.DiscoPurposeTableReorderColumns) {
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
	var recognizedColumnIndexes []int
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
		recognizedColumnIndexes = append(recognizedColumnIndexes, i)
	}
	slices.Sort(recognizedColumnIndexes)
	for _, extraColumnIndex := range recognizedColumnIndexes {
		newColumnIndexes[extraColumnIndex] = index
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
		newCells := make(asciidoc.Elements, index)
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
		row.Elements = newCells
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
	err = ti.Rescan(cxt.doc, asciidoc.RawReader)
	return
}

func setCellString(cell *asciidoc.TableCell, v string) {
	se := asciidoc.NewString(v)
	cell.SetChildren(asciidoc.Elements{se})
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
		setCellString(tableCells[toIndex], value)
	}
	return
}

func (b *Baller) renameTableHeaderCells(cxt *discoContext, section *asciidoc.Section, table *spec.TableInfo, overrides map[matter.TableColumn]matter.TableColumn) (err error) {
	if !b.options.RenameTableHeaders {
		return
	}
	if cxt.errata.IgnoreSection(cxt.doc.SectionName(section), errata.DiscoPurposeTableRenameHeaders) {
		return
	}
	if table.HeaderRowIndex == -1 {
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
				slog.Warn("Can not rename column; column with same name already exists", slog.String("from", tc.String()), slog.String("to", overrideColumn.String()), slog.Int("existingColumnIndex", existingIndex), log.Element("source", cxt.doc.Path, table.Element))
				continue
			}
			tc = overrideColumn
		}
		name, ok := matter.GetColumnName(tc, overrides)
		if ok {
			setCellString(cell, name)
		}
	}
	err = table.Rescan(cxt.doc, asciidoc.RawReader)
	return
}
