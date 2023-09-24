package disco

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
)

func organizeCommandsSection(doc *ascii.Doc, section *ascii.Section) error {
	t := findFirstTable(section)
	if t == nil {
		return fmt.Errorf("no commands section found")
	}
	return organizeCommandsTable(doc, section, t)
}

func organizeCommandsTable(doc *ascii.Doc, section *ascii.Section, attributesTable *types.Table) error {

	setSectionTitle(section, matter.CommandsSectionName)

	rows := combineRows(attributesTable)

	headerRowIndex, columnMap, extraColumns := findColumns(rows)

	if columnMap == nil {
		return fmt.Errorf("can't rearrange commands table without header row")
	}

	if len(columnMap) < 2 {
		return fmt.Errorf("can't rearrange commands table with so few matches")
	}

	err := fixAccessCells(doc, rows, columnMap)
	if err != nil {
		return err
	}

	err = fixCommandDirection(doc, rows, columnMap)
	if err != nil {
		return err
	}

	err = renameTableHeaderCells(rows, headerRowIndex, columnMap, matter.CommandsTableColumnNames)
	if err != nil {
		return err
	}

	reorderColumns(doc, section, rows, matter.CommandsTableColumnOrder[:], columnMap, extraColumns)
	return nil
}

func fixCommandDirection(doc *ascii.Doc, rows []*types.TableRow, columnMap map[matter.TableColumn]int) (err error) {
	if len(rows) < 2 {
		return
	}
	accessIndex, ok := columnMap[matter.TableColumnDirection]
	if !ok {
		return
	}
	for _, row := range rows[1:] {
		cell := row.Cells[accessIndex]

		vc, e := getCellValue(cell)
		if e != nil {
			continue
		}
		err = setCellValue(cell, strings.ToLower(vc))
		if err != nil {
			return
		}
	}
	return
}
