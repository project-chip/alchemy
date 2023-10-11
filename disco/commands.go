package disco

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
	"github.com/hasty/matterfmt/parse"
)

func (b *Ball) organizeCommandsSection(cxt *discoContext, doc *ascii.Doc, commands *ascii.Section) error {
	t := parse.FindFirstTable(commands)
	if t == nil {
		return fmt.Errorf("no commands table found")
	}
	return b.organizeCommandsTable(cxt, doc, commands, t)
}

func (b *Ball) organizeCommandsTable(cxt *discoContext, doc *ascii.Doc, commands *ascii.Section, commandsTable *types.Table) error {

	setSectionTitle(commands, matter.CommandsSectionName)

	rows := parse.TableRows(commandsTable)

	headerRowIndex, columnMap, extraColumns, err := parse.MapTableColumns(rows)
	if err != nil {
		return err
	}

	if columnMap == nil {
		return fmt.Errorf("can't rearrange commands table without header row")
	}

	if len(columnMap) < 2 {
		return fmt.Errorf("can't rearrange commands table with so few matches")
	}

	err = b.fixAccessCells(doc, rows, columnMap)
	if err != nil {
		return err
	}

	err = b.fixCommandDirection(doc, rows, columnMap)
	if err != nil {
		return err
	}

	err = renameTableHeaderCells(rows, headerRowIndex, columnMap, matter.CommandsTableColumnNames)
	if err != nil {
		return err
	}

	err = organizeCommands(cxt, commands, commandsTable, columnMap)
	if err != nil {
		return err
	}

	reorderColumns(doc, commands, rows, matter.CommandsTableColumnOrder[:], columnMap, extraColumns)
	return nil
}

func organizeCommands(cxt *discoContext, commands *ascii.Section, commandsTable *types.Table, columnMap map[matter.TableColumn]int) error {
	nameIndex, ok := columnMap[matter.TableColumnName]
	if !ok {
		return nil
	}
	commandNames := make(map[string]struct{}, len(commandsTable.Rows))
	for _, row := range commandsTable.Rows {
		commandName, err := parse.GetTableCellValue(row.Cells[nameIndex])
		if err != nil {
			slog.Warn("could not get cell value for command", "err", err)
			continue
		}
		commandNames[commandName] = struct{}{}
	}
	subSections := ascii.FindAll[*ascii.Section](commands.Elements)
	for _, ss := range subSections {
		name := strings.TrimSuffix(ss.Name, " Command")
		if _, ok := commandNames[name]; !ok {
			continue
		}
		t := parse.FindFirstTable(ss)
		if t == nil {
			continue
		}
		rows := parse.TableRows(t)

		_, columnMap, _, err := parse.MapTableColumns(rows)
		if err != nil {
			return err
		}
		err = fixConstraintCells(rows, columnMap)
		if err != nil {
			return err
		}
		err = getPotentialDataTypes(cxt, ss, rows, columnMap)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *Ball) fixCommandDirection(doc *ascii.Doc, rows []*types.TableRow, columnMap map[matter.TableColumn]int) (err error) {
	if len(rows) < 2 {
		return
	}
	accessIndex, ok := columnMap[matter.TableColumnDirection]
	if !ok {
		return
	}
	for _, row := range rows[1:] {
		cell := row.Cells[accessIndex]

		vc, e := parse.GetTableCellValue(cell)
		if e != nil {
			continue
		}
		err = setCellString(cell, strings.ToLower(vc))
		if err != nil {
			return
		}
	}
	return
}
