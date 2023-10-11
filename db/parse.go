package db

import (
	"context"
	"fmt"

	asciitypes "github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
	"github.com/hasty/matterfmt/parse"
)

func readSimpleSection(cxt context.Context, section *ascii.Section, row *dbRow) error {
	t := parse.FindFirstTable(section)
	if t == nil {
		return fmt.Errorf("no table found")
	}
	rows := parse.TableRows(t)
	if len(rows) < 2 {
		return fmt.Errorf("not enough rows in table")
	}
	headerRowIndex, columnMap, extraColumns, err := parse.MapTableColumns(rows)
	if err != nil {
		return err
	}
	if len(rows) < headerRowIndex+2 {
		return fmt.Errorf("not enough value rows in table")
	}
	if columnMap == nil {
		return fmt.Errorf("can't read table without columns")
	}
	valueRow := rows[headerRowIndex+1]
	err = readTableRow(valueRow, columnMap, extraColumns, row)
	if err != nil {
		return err
	}
	return nil
}

func (h *Host) readTableSection(cxt context.Context, parent *sectionInfo, section *ascii.Section, name string) error {
	rows, err := readTable(cxt, section)
	if err != nil {
		return err
	}
	if parent.children == nil {
		parent.children = make(map[string][]*sectionInfo, len(rows))
	}
	for _, r := range rows {
		ci := &sectionInfo{id: h.nextId(name), parent: parent, values: r}
		parent.children[name] = append(parent.children[name], ci)
	}
	return nil
}

func readTable(cxt context.Context, section *ascii.Section) (rs []*dbRow, err error) {
	t := parse.FindFirstTable(section)
	if t == nil {
		return nil, fmt.Errorf("no table found")
	}
	rows := parse.TableRows(t)
	if len(rows) < 2 {
		return nil, fmt.Errorf("not enough rows in table")
	}
	headerRowIndex, columnMap, extraColumns, err := parse.MapTableColumns(rows)
	if err != nil {
		return nil, err
	}
	if len(rows) < headerRowIndex+2 {
		return nil, fmt.Errorf("not enough value rows in table")
	}
	if columnMap == nil {
		return nil, fmt.Errorf("can't read table without columns")
	}
	for _, valueRow := range rows[headerRowIndex+1:] {
		row := &dbRow{}
		err = readTableRow(valueRow, columnMap, extraColumns, row)
		if err != nil {
			return
		}
		rs = append(rs, row)
	}

	return
}

func readTableRow(valueRow *asciitypes.TableRow, columnMap map[matter.TableColumn]int, extraColumns []parse.ExtraColumn, row *dbRow) error {
	if row.values == nil {
		row.values = make(map[matter.TableColumn]interface{})
	}
	for col, index := range columnMap {
		val, err := parse.GetTableCellValue(valueRow.Cells[index])
		if err != nil {
			return err
		}
		row.values[col] = val
	}
	if len(extraColumns) > 0 {
		if row.extras == nil {
			row.extras = make(map[string]interface{})

		}
		for _, e := range extraColumns {
			val, err := parse.GetTableCellValue(valueRow.Cells[e.Offset])
			if err != nil {
				return err
			}
			row.extras[e.Name] = val
		}
	}
	return nil
}
