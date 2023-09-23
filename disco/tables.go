package disco

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
)

func findFirstTable(section *ascii.Section) *types.Table {
	var table *types.Table
	find(section.Elements, func(t *types.Table) bool {
		table = t
		return true
	})

	return table
}

func ensureTableOptions(elements []interface{}) {
	find(elements, func(t *types.Table) bool {
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

func combineRows(t *types.Table) (rows []*types.TableRow) {
	rows = make([]*types.TableRow, 0, len(t.Rows)+2)
	if t.Header != nil {
		rows = append(rows, t.Header)
	}
	rows = append(rows, t.Rows...)
	if t.Footer != nil {
		rows = append(rows, t.Footer)
	}
	return
}

func reorderColumns(doc *ascii.Doc, section *ascii.Section, rows []*types.TableRow, order []matter.TableColumn, columnMap map[matter.TableColumn]int, extraColumns []int) []*types.TableRow {
	newRows := make([]*types.TableRow, 0, len(rows))
	for _, row := range rows {
		newRow := &types.TableRow{Cells: make([]*types.TableCell, len(columnMap)+len(extraColumns))}
		var newOffset int
		for _, column := range order {
			if offset, ok := columnMap[column]; ok {
				newRow.Cells[newOffset] = row.Cells[offset]
				newOffset++
			}
		}
		for _, extra := range extraColumns {
			newRow.Cells[newOffset] = row.Cells[extra]
			newOffset++
		}
		newRows = append(newRows, newRow)
	}
	return newRows
}

func getCellValue(cell *types.TableCell) string {
	p, ok := cell.Elements[0].(*types.Paragraph)
	if !ok || len(p.Elements) == 0 {
		return ""
	}
	switch v := p.Elements[0].(type) {
	case *types.StringElement:
		return v.Content
	case string:
		return v
	}
	return ""
}

func setCellValue(cell *types.TableCell, v string) {
	p, ok := cell.Elements[0].(*types.Paragraph)
	if !ok || len(p.Elements) == 0 {
		return
	}
	switch val := p.Elements[0].(type) {
	case *types.StringElement:
		val.Content = v
	case string:
		p.Elements[0] = v
	default:
		fmt.Printf("not a string %T!\n", p.Elements[0])

	}
}

func getTableColumn(cell *types.TableCell) matter.TableColumn {
	switch strings.ToLower(getCellValue(cell)) {
	case "id", "identifier":
		return matter.TableColumnID
	case "name":
		return matter.TableColumnName
	case "type":
		return matter.TableColumnType
	case "constraint":
		return matter.TableColumnConstraint
	case "quality":
		return matter.TableColumnQuality
	case "default":
		return matter.TableColumnDefault
	case "access":
		return matter.TableColumnAccess
	case "conformance":
		return matter.TableColumnConformance
	case "hierarchy":
		return matter.TableColumnHierarchy
	case "role":
		return matter.TableColumnRole
	case "context":
		return matter.TableColumnContext
	case "pics code", "pics":
		return matter.TableColumnPICS
	case "scope":
		return matter.TableColumnScope
	case "value":
		return matter.TableColumnValue
	case "bit":
		return matter.TableColumnBit
	}
	return matter.TableColumnUnknown
}

func findColumns(rows []*types.TableRow) (int, map[matter.TableColumn]int, []int) {
	var columnMap map[matter.TableColumn]int
	var extraColumns []int
	var cellCount = -1
	var headerRow = -1
	for i, row := range rows {
		for _, cell := range row.Cells {
			if cell.Formatter != nil {
				if cell.Formatter.ColumnSpan > 0 || cell.Formatter.RowSpan > 0 {
					fmt.Println("can't rearrange attributes table with row or column spanning")
					return -1, nil, nil
				}
			}
		}
		if cellCount == -1 {
			cellCount = len(row.Cells)
		} else if cellCount != len(row.Cells) {
			fmt.Println("can't rearrange attributes table with unequal cell counts between rows")
			return -1, nil, nil
		}
		if columnMap == nil {
			var spares []int
			for j, cell := range row.Cells {
				attributeColumn := getTableColumn(cell)
				if attributeColumn != matter.TableColumnUnknown {
					if columnMap == nil {
						headerRow = i
						columnMap = make(map[matter.TableColumn]int)
					}
					if _, ok := columnMap[attributeColumn]; ok {
						fmt.Println("can't rearrange attributes table duplicate columns")
						return -1, nil, nil
					}
					columnMap[attributeColumn] = j
				} else {
					spares = append(spares, j)
				}
			}
			if columnMap != nil {
				extraColumns = spares
			}
		}
	}
	return headerRow, columnMap, extraColumns
}

func renameTableHeaderCells(rows []*types.TableRow, headerRowIndex int, columnMap map[matter.TableColumn]int, nameMap map[matter.TableColumn]string) {
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
		name, ok := nameMap[tc]
		if ok {
			setCellValue(cell, name)
		}
	}
}
