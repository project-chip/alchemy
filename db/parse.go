package db

import (
	"context"
	"fmt"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
)

type sectionInfo struct {
	id     int32
	values *dbRow

	parent *sectionInfo

	children map[string][]*sectionInfo
}

var errMissingTable = fmt.Errorf("no table found")

func appendSectionToRow(cxt context.Context, doc *spec.Doc, section *spec.Section, row *dbRow) error {
	t := spec.FindFirstTable(section)
	if t == nil {
		return fmt.Errorf("no table found")
	}
	ti, err := spec.ReadTable(doc, t)
	if err != nil {
		return fmt.Errorf("failed mapping table columns for section %s: %w", section.Name, err)
	}
	if len(ti.Rows) < 2 {
		return fmt.Errorf("not enough rows in table")
	}
	if len(ti.Rows) < ti.HeaderRowIndex+2 {
		return fmt.Errorf("not enough value rows in table")
	}
	if ti.ColumnMap == nil {
		return fmt.Errorf("can't read table without columns")
	}
	valueRow := ti.Rows[ti.HeaderRowIndex+1]
	err = readTableRow(valueRow, ti.ColumnMap, ti.ExtraColumns, row)
	if err != nil {
		return err
	}
	return nil
}

func (h *Host) readTableSection(cxt context.Context, doc *spec.Doc, parent *sectionInfo, section *spec.Section, name string) error {
	rows, err := readTable(cxt, doc, section)
	if err == errMissingTable {
		return nil
	}
	if err != nil {
		return err
	}
	if parent.children == nil {
		parent.children = make(map[string][]*sectionInfo, len(rows))
	}
	for _, r := range rows {
		ci := &sectionInfo{id: h.nextID(name), parent: parent, values: r}
		parent.children[name] = append(parent.children[name], ci)
	}
	return nil
}

func readTable(cxt context.Context, doc *spec.Doc, section *spec.Section) (rs []*dbRow, err error) {
	t := spec.FindFirstTable(section)
	if t == nil {
		return nil, errMissingTable
	}

	ti, err := spec.ReadTable(doc, t)
	if err != nil {
		return nil, err
	}
	if len(ti.Rows) < 2 {
		return nil, fmt.Errorf("not enough rows in table")
	}
	if len(ti.Rows) < ti.HeaderRowIndex+2 {
		return nil, fmt.Errorf("not enough value rows in table")
	}
	if ti.ColumnMap == nil {
		return nil, fmt.Errorf("can't read table without columns")
	}
	for _, valueRow := range ti.Rows[ti.HeaderRowIndex+1:] {
		row := &dbRow{}
		err = readTableRow(valueRow, ti.ColumnMap, ti.ExtraColumns, row)
		if err != nil {
			return
		}
		rs = append(rs, row)
	}

	return
}

func readTableRow(valueRow *asciidoc.TableRow, columnMap spec.ColumnIndex, extraColumns []spec.ExtraColumn, row *dbRow) error {
	if row.values == nil {
		row.values = make(map[matter.TableColumn]any)
	}

	for col, index := range columnMap {
		val, err := spec.RenderTableCell(valueRow.Cell(index))
		if err != nil {
			return err
		}
		row.values[col] = val
	}
	if len(extraColumns) > 0 {
		if row.extras == nil {
			row.extras = make(map[string]any)

		}
		for _, e := range extraColumns {
			val, err := spec.RenderTableCell(valueRow.Cell(e.Offset))
			if err != nil {
				return err
			}
			row.extras[e.Name] = val
		}
	}
	return nil
}
