package disco

import (
	"fmt"
	"log/slog"
	"slices"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/internal/parse"
	"github.com/hasty/alchemy/matter"
	mattertypes "github.com/hasty/alchemy/matter/types"
)

func (b *Ball) ensureTableOptions(elements []interface{}) {
	if !b.options.normalizeTableOptions {
		return
	}
	parse.Search(elements, func(t *types.Table) bool {
		if t.Attributes == nil {
			t.Attributes = make(types.Attributes)
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

func (b *Ball) addMissingColumns(doc *ascii.Doc, section *ascii.Section, table *types.Table, rows []*types.TableRow, order []matter.TableColumn, overrides map[matter.TableColumn]string, headerRowIndex int, columnMap ascii.ColumnIndex, entityType mattertypes.EntityType) (err error) {
	if !b.options.addMissingColumns {
		return
	}
	delete(table.Attributes, "cols")
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

func (b *Ball) appendColumn(rows []*types.TableRow, columnMap ascii.ColumnIndex, headerRowIndex int, column matter.TableColumn, overrides map[matter.TableColumn]string, entityType mattertypes.EntityType) (appendedIndex int, err error) {
	if len(rows) == 0 {
		appendedIndex = -1
		return
	}
	appendedIndex = len(rows[0].Cells)
	for i, row := range rows {
		cell := &types.TableCell{}
		if i == headerRowIndex {
			if headerRowIndex > 0 {
				cell.Format = "h"
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
			last := row.Cells[len(row.Cells)-1]
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

func (b *Ball) getDefaultColumnValue(entityType mattertypes.EntityType, column matter.TableColumn, row *types.TableRow, columnMap ascii.ColumnIndex) string {
	switch column {
	case matter.TableColumnConformance:
		switch entityType {
		case mattertypes.EntityTypeBitmapValue, mattertypes.EntityTypeEnumValue:
			return "M"
		}
	case matter.TableColumnName:
		switch entityType {
		case mattertypes.EntityTypeBitmapValue, mattertypes.EntityTypeEnumValue:
			val, _ := ascii.ReadRowValue(b.doc, row, columnMap, matter.TableColumnSummary)
			if val != "" {
				return matter.Case(val)
			}
		}
	case matter.TableColumnSummary:
		switch entityType {
		case mattertypes.EntityTypeBitmapValue, mattertypes.EntityTypeEnumValue:
			val, _ := ascii.ReadRowValue(b.doc, row, columnMap, matter.TableColumnName)
			if val != "" {
				return matter.Uncase(val)
			}
		}
	case matter.TableColumnAccess:
		switch entityType {
		case mattertypes.EntityTypeEvent, mattertypes.EntityTypeAttribute:
			return "R V"
		case mattertypes.EntityTypeField:
			return ""
		}
	}
	slog.Debug("no default value for column", slog.String("column", column.String()), slog.String("entity", entityType.String()))
	return ""
}

func (b *Ball) reorderColumns(doc *ascii.Doc, section *ascii.Section, rows []*types.TableRow, order []matter.TableColumn, columnMap ascii.ColumnIndex, extraColumns []ascii.ExtraColumn) {
	if !b.options.reorderColumns {
		return
	}
	if len(rows) == 0 {
		return
	}
	newColumnIndexes := make([]int, len(columnMap)+len(extraColumns))
	for i := range newColumnIndexes {
		newColumnIndexes[i] = -1
	}
	for i, column := range order {
		if offset, ok := columnMap[column]; ok {
			// column exists in existing order
			newColumnIndexes[offset] = i
		}
	}
	for i, extra := range extraColumns {
		newColumnIndexes[extra.Offset] = len(order) + i
	}
	for _, row := range rows {
		newCells := make([]*types.TableCell, 0, len(order)+len(extraColumns))
		for _, column := range order {
			if offset, ok := columnMap[column]; ok {
				newCells = append(newCells, row.Cells[offset])
			}
		}
		for _, extra := range extraColumns {
			newCells = append(newCells, row.Cells[extra.Offset])
		}
		slices.Clip(newCells)
		row.Cells = newCells
	}
}

func setCellString(cell *types.TableCell, v string) (err error) {
	var p *types.Paragraph

	if len(cell.Elements) == 0 {
		p, err = types.NewParagraph(nil)
		if err != nil {
			return
		}
		err = cell.SetElements([]interface{}{p})
		if err != nil {
			return
		}
	} else {
		var ok bool
		p, ok = cell.Elements[0].(*types.Paragraph)
		if !ok {
			return fmt.Errorf("table cell does not have paragraph child")
		}
	}
	se, _ := types.NewStringElement(v)
	err = p.SetElements([]interface{}{se})
	return
}

func setCellValue(cell *types.TableCell, val []interface{}) (err error) {
	var p *types.Paragraph

	if len(cell.Elements) == 0 {
		p, err = types.NewParagraph(nil)
		if err != nil {
			return
		}
		err = cell.SetElements([]interface{}{p})
		if err != nil {
			return
		}
	} else {
		var ok bool
		p, ok = cell.Elements[0].(*types.Paragraph)
		if !ok {
			return fmt.Errorf("table cell does not have paragraph child")
		}
	}
	err = p.SetElements(val)
	return
}

func copyCells(rows []*types.TableRow, headerRowIndex int, fromIndex int, toIndex int, transformer func(s string) string) (err error) {
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

func (b *Ball) renameTableHeaderCells(rows []*types.TableRow, headerRowIndex int, columnMap ascii.ColumnIndex, overrides map[matter.TableColumn]string) (err error) {
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
