package disco

import (
	"fmt"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func (b *Ball) organizeCommandsSection(cxt *discoContext, dp *docParse) (err error) {
	for _, commands := range dp.commands {
		if commands.table.element == nil {
			err = fmt.Errorf("no commands table found")
			return
		}
		if commands.table.columnMap == nil {
			return fmt.Errorf("can't rearrange commands table without header row")
		}

		if len(commands.table.columnMap) < 2 {
			return fmt.Errorf("can't rearrange commands table with so few matches")
		}
		err = b.fixAccessCells(dp, commands, types.EntityTypeCommand)
		if err != nil {
			return fmt.Errorf("error fixing access cells in commands table in %s: %w", dp.doc.Path, err)
		}

		err = fixConformanceCells(dp, commands.table.rows, commands.table.columnMap)
		if err != nil {
			return fmt.Errorf("error fixing conformance cells in commands table in %s: %w", dp.doc.Path, err)
		}

		err = b.fixCommandDirection(dp.doc, commands.table.rows, commands.table.columnMap)
		if err != nil {
			return fmt.Errorf("error fixing command direction in commands table in %s: %w", dp.doc.Path, err)
		}

		err = b.renameTableHeaderCells(dp.doc, commands.section, &commands.table, nil)
		if err != nil {
			return fmt.Errorf("error table header cells in commands table in %s: %w", dp.doc.Path, err)
		}

		err = b.linkIndexTables(cxt, commands)
		if err != nil {
			return err
		}

		for _, command := range commands.children {
			if command.table.element == nil {
				continue
			}
			err = fixConstraintCells(dp.doc, &command.table)
			if err != nil {
				return fmt.Errorf("error fixing command constraint cells in %s in %s: %w", command.section.Name, dp.doc.Path, err)
			}
			err = fixConformanceCells(dp, command.table.rows, command.table.columnMap)
			if err != nil {
				return fmt.Errorf("error fixing command conformance cells in %s in %s: %w", command.section.Name, dp.doc.Path, err)
			}
			b.appendSubsectionTypes(command.section, command.table.columnMap, command.table.rows)
			err = b.linkIndexTables(cxt, command)
			if err != nil {
				return err
			}
		}
	}
	return
}

func (b *Ball) fixCommandDirection(doc *spec.Doc, rows []*asciidoc.TableRow, columnMap spec.ColumnIndex) (err error) {
	if len(rows) < 2 {
		return
	}
	accessIndex, ok := columnMap[matter.TableColumnDirection]
	if !ok {
		return
	}
	for _, row := range rows[1:] {
		cell := row.Cell(accessIndex)

		vc, e := spec.RenderTableCell(cell)
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
