package disco

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func (b *Ball) organizeEnumSections(cxt *discoContext, dp *docParse) (err error) {
	for _, es := range dp.enums {
		err = b.organizeEnumSection(cxt, dp, es)
		if err != nil {
			return
		}
	}
	return
}

func (b *Ball) organizeEnumSection(cxt *discoContext, dp *docParse, es *subSection) (err error) {
	name := strings.TrimSpace(es.section.Name)
	if strings.HasSuffix(strings.ToLower(name), "enum") {
		setSectionTitle(es.section, name+" Type")
	}
	enumTable := es.table
	if enumTable.element == nil {
		return
	}
	if enumTable.columnMap == nil {
		slog.Debug("can't rearrange enum table without header row")
		return nil
	}

	if len(enumTable.columnMap) < 2 {
		slog.Debug("can't rearrange enum table with so few matches")
		return nil
	}

	err = b.renameTableHeaderCells(dp.doc, &enumTable, matter.Tables[matter.TableTypeEnum].ColumnNames)
	if err != nil {
		return fmt.Errorf("error renaming table header cells in enum table in section %s in %s: %w", es.section.Name, dp.doc.Path, err)
	}

	err = b.addMissingColumns(&enumTable, matter.Tables[matter.TableTypeEnum], nil, types.EntityTypeEnumValue)
	if err != nil {
		return fmt.Errorf("error adding missing table columns in enum section %s in %s: %w", es.section.Name, dp.doc.Path, err)
	}

	es.table.headerRow, es.table.columnMap, es.table.extraColumns, err = spec.MapTableColumns(dp.doc, enumTable.rows)
	if err != nil {
		return fmt.Errorf("error reordering columns in enum table in section %s in %s: %w", es.section.Name, dp.doc.Path, err)

	}
	enumTable = es.table

	err = b.reorderColumns(dp.doc, es.section, &enumTable, matter.TableTypeEnum)
	if err != nil {
		return err
	}

	b.appendSubsectionTypes(es.section, enumTable.columnMap, enumTable.rows)
	return
}
