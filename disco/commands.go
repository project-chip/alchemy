package disco

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/matter"
	mattertypes "github.com/hasty/alchemy/matter/types"
	"github.com/hasty/alchemy/parse"
)

func (b *Ball) organizeCommandsSection(cxt *discoContext, doc *ascii.Doc, commands *ascii.Section) error {
	t := ascii.FindFirstTable(commands)
	if t == nil {
		return fmt.Errorf("no commands table found")
	}
	return b.organizeCommandsTable(cxt, doc, commands, t)
}

func (b *Ball) organizeCommandsTable(cxt *discoContext, doc *ascii.Doc, commands *ascii.Section, commandsTable *types.Table) error {

	setSectionTitle(commands, matter.CommandsSectionName)

	rows := ascii.TableRows(commandsTable)

	headerRowIndex, columnMap, extraColumns, err := ascii.MapTableColumns(doc, rows)
	if err != nil {
		return fmt.Errorf("failed mapping table columns for commands table in section %s: %w", commands.Name, err)
	}

	if columnMap == nil {
		return fmt.Errorf("can't rearrange commands table without header row")
	}

	if len(columnMap) < 2 {
		return fmt.Errorf("can't rearrange commands table with so few matches")
	}

	err = b.fixAccessCells(doc, rows, columnMap, mattertypes.EntityTypeCommand)
	if err != nil {
		return fmt.Errorf("error fixing access cells in commands table in %s: %w", doc.Path, err)
	}

	err = b.fixCommandDirection(doc, rows, columnMap)
	if err != nil {
		return fmt.Errorf("error fixing command direction in commands table in %s: %w", doc.Path, err)
	}

	err = b.renameTableHeaderCells(rows, headerRowIndex, columnMap, nil)
	if err != nil {
		return fmt.Errorf("error table header cells in commands table in %s: %w", doc.Path, err)
	}

	err = b.organizeCommands(cxt, doc, commands, commandsTable, columnMap)
	if err != nil {
		return fmt.Errorf("error organizing commands in %s: %w", doc.Path, err)
	}

	b.reorderColumns(doc, commands, rows, matter.CommandsTableColumnOrder[:], columnMap, extraColumns)
	return nil
}

func (b *Ball) organizeCommands(cxt *discoContext, doc *ascii.Doc, commands *ascii.Section, commandsTable *types.Table, columnMap ascii.ColumnIndex) error {
	nameIndex, ok := columnMap[matter.TableColumnName]
	if !ok {
		return nil
	}
	commandNames := make(map[string]struct{}, len(commandsTable.Rows))
	for _, row := range commandsTable.Rows {
		commandName, err := ascii.GetTableCellValue(row.Cells[nameIndex])
		if err != nil {
			slog.Debug("could not get cell value for command", "err", err)
			continue
		}
		commandNames[commandName] = struct{}{}
	}
	subSections := parse.FindAll[*ascii.Section](commands.Elements)
	for _, ss := range subSections {
		name := strings.TrimSuffix(ss.Name, " Command")
		if _, ok := commandNames[name]; !ok {
			continue
		}
		t := ascii.FindFirstTable(ss)
		if t == nil {
			continue
		}
		rows := ascii.TableRows(t)

		_, columnMap, _, err := ascii.MapTableColumns(doc, rows)
		if err != nil {
			return fmt.Errorf("failed mapping table columns for fields table in section %s: %w", ss.Name, err)
		}
		err = fixConstraintCells(doc, rows, columnMap)
		if err != nil {
			return fmt.Errorf("error fixing command constraint cells in %s command in %s: %w", name, doc.Path, err)
		}
		err = b.getPotentialDataTypes(cxt, ss, rows, columnMap)
		if err != nil {
			return fmt.Errorf("error getting potential data types in %s command in %s: %w", name, doc.Path, err)
		}
		b.appendSubsectionTypes(ss, columnMap, rows)
	}

	return nil
}

func (b *Ball) fixCommandDirection(doc *ascii.Doc, rows []*types.TableRow, columnMap ascii.ColumnIndex) (err error) {
	if len(rows) < 2 {
		return
	}
	accessIndex, ok := columnMap[matter.TableColumnDirection]
	if !ok {
		return
	}
	for _, row := range rows[1:] {
		cell := row.Cells[accessIndex]

		vc, e := ascii.GetTableCellValue(cell)
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
