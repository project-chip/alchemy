package disco

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/parse"
)

func ensureTableOptions(elements []interface{}) {
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

func addMissingColumns(doc *ascii.Doc, section *ascii.Section, rows []*types.TableRow, order []matter.TableColumn, nameMap map[matter.TableColumn]string, headerRowIndex int, columnMap map[matter.TableColumn]int) []*types.TableRow {
	newRows := make([]*types.TableRow, 0, len(rows))
	for _, column := range order {
		if _, ok := columnMap[column]; !ok {
			fmt.Printf("missing column %v\n", column)
			for i, row := range rows {
				cell := &types.TableCell{}
				if i == headerRowIndex {
					if headerRowIndex > 0 {
						cell.Format = "h"
					}
					err := setCellString(cell, nameMap[column])
					if err != nil {
						fmt.Printf("errored: %v\n", err)
					}
				} else {
					last := row.Cells[len(row.Cells)-1]
					cell.Blank = last.Blank
				}
				fmt.Printf("added column %v\n", cell)
				row.Cells = append(row.Cells, cell)
			}
		}
	}
	return newRows
}

func reorderColumns(doc *ascii.Doc, section *ascii.Section, rows []*types.TableRow, order []matter.TableColumn, columnMap map[matter.TableColumn]int, extraColumns []ascii.ExtraColumn) {
	for _, row := range rows {
		newCells := make([]*types.TableCell, len(columnMap)+len(extraColumns))
		var newOffset int
		for _, column := range order {
			if offset, ok := columnMap[column]; ok {
				newCells[newOffset] = row.Cells[offset]
				newOffset++
			}
		}
		for _, extra := range extraColumns {
			newCells[newOffset] = row.Cells[extra.Offset]
			newOffset++
		}
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
		cell.SetElements([]interface{}{p})
	} else {
		var ok bool
		p, ok = cell.Elements[0].(*types.Paragraph)
		if !ok {
			return fmt.Errorf("table cell does not have paragraph child")
		}
	}
	se, _ := types.NewStringElement(v)
	p.SetElements([]interface{}{se})
	return
}

func setCellValue(cell *types.TableCell, val []interface{}) (err error) {
	var p *types.Paragraph

	if len(cell.Elements) == 0 {
		p, err = types.NewParagraph(nil)
		if err != nil {
			return
		}
		cell.SetElements([]interface{}{p})
	} else {
		var ok bool
		p, ok = cell.Elements[0].(*types.Paragraph)
		if !ok {
			return fmt.Errorf("table cell does not have paragraph child")
		}
	}
	p.SetElements(val)
	return
}

func renameTableHeaderCells(rows []*types.TableRow, headerRowIndex int, columnMap map[matter.TableColumn]int, nameMap map[matter.TableColumn]string) (err error) {
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
			err = setCellString(cell, name)
			if err != nil {
				return
			}
		}
	}
	return
}
