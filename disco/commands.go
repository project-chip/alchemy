package disco

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func (b *Baller) organizeCommandsSection(cxt *discoContext) (err error) {
	for _, commands := range cxt.parsed.commands {
		if commands.table == nil || commands.table.Element == nil {
			slog.Warn("Could not organize Commands section, as no table of commands was found", log.Path("source", commands.section))
			return
		}
		if commands.table.ColumnMap == nil {
			return fmt.Errorf("can't rearrange commands table without header row")
		}

		if len(commands.table.ColumnMap) < 2 {
			return fmt.Errorf("can't rearrange commands table with so few matches")
		}
		err = b.renameTableHeaderCells(cxt, commands.section, commands.table, matter.Tables[matter.TableTypeCommands].ColumnRenames)
		if err != nil {
			return fmt.Errorf("error renaming table header cells in section %s in %s: %w", commands.section.Name, cxt.doc.Path, err)
		}
		err = b.fixAccessCells(cxt, commands, types.EntityTypeCommand)
		if err != nil {
			return fmt.Errorf("error fixing access cells in commands table in %s: %w", cxt.doc.Path, err)
		}

		err = b.fixConformanceCells(cxt, commands, commands.table.Rows, commands.table.ColumnMap)
		if err != nil {
			return fmt.Errorf("error fixing conformance cells in commands table in %s: %w", cxt.doc.Path, err)
		}

		err = b.fixCommandDirection(cxt, commands.section, commands.table.Rows, commands.table.ColumnMap)
		if err != nil {
			return fmt.Errorf("error fixing command direction in commands table in %s: %w", cxt.doc.Path, err)
		}

		err = b.renameTableHeaderCells(cxt, commands.section, commands.table, nil)
		if err != nil {
			return fmt.Errorf("error table header cells in commands table in %s: %w", cxt.doc.Path, err)
		}

		err = b.linkIndexTables(cxt, commands)
		if err != nil {
			return err
		}

		for _, command := range commands.children {
			if command.table == nil || command.table.Element == nil {
				continue
			}
			err = b.renameTableHeaderCells(cxt, command.section, command.table, matter.Tables[matter.TableTypeCommandFields].ColumnRenames)
			if err != nil {
				return fmt.Errorf("error table header cells in commands table in %s: %w", cxt.doc.Path, err)
			}
			err = b.fixConstraintCells(cxt, command.section, command.table)
			if err != nil {
				return fmt.Errorf("error fixing command constraint cells in %s in %s: %w", command.section.Name, cxt.doc.Path, err)
			}
			err = b.fixConformanceCells(cxt, command, command.table.Rows, command.table.ColumnMap)
			if err != nil {
				return fmt.Errorf("error fixing command conformance cells in %s in %s: %w", command.section.Name, cxt.doc.Path, err)
			}
			b.appendSubsectionTypes(cxt, command.section, command.table.ColumnMap, command.table.Rows)
			b.removeMandatoryFallbacks(command.table)

			err = b.linkIndexTables(cxt, command)
			if err != nil {
				return err
			}
		}
	}
	return
}

func (b *Baller) fixCommandDirection(cxt *discoContext, section *asciidoc.Section, rows []*asciidoc.TableRow, columnMap spec.ColumnIndex) (err error) {
	if len(rows) < 2 {
		return
	}
	if cxt.errata.IgnoreSection(cxt.doc.SectionName(section), errata.DiscoPurposeDataTypeCommandFixDirection) {
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
		setCellString(cell, strings.ToLower(vc))
	}
	return
}
