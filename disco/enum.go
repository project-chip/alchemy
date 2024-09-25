package disco

import (
	"fmt"
	"log/slog"

	"github.com/project-chip/alchemy/matter"
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

	b.canonicalizeDataTypeSectionName(dp, es.section, "Enum")

	enumTable := es.table
	if enumTable.Element == nil {
		return
	}
	if enumTable.ColumnMap == nil {
		slog.Debug("can't rearrange enum table without header row")
		return nil
	}

	if len(enumTable.ColumnMap) < 2 {
		slog.Debug("can't rearrange enum table with so few matches")
		return nil
	}

	err = b.renameTableHeaderCells(dp.doc, es.section, enumTable, matter.Tables[matter.TableTypeEnum].ColumnNames)
	if err != nil {
		return fmt.Errorf("error renaming table header cells in enum table in section %s in %s: %w", es.section.Name, dp.doc.Path, err)
	}

	err = b.addMissingColumns(es.section, enumTable, matter.Tables[matter.TableTypeEnum], types.EntityTypeEnumValue)
	if err != nil {
		return fmt.Errorf("error adding missing table columns in enum section %s in %s: %w", es.section.Name, dp.doc.Path, err)
	}

	err = enumTable.Rescan(dp.doc)
	if err != nil {
		return fmt.Errorf("error reordering columns in enum table in section %s in %s: %w", es.section.Name, dp.doc.Path, err)

	}
	enumTable = es.table

	err = b.reorderColumns(dp.doc, es.section, enumTable, matter.TableTypeEnum)
	if err != nil {
		return err
	}

	b.appendSubsectionTypes(es.section, enumTable.ColumnMap, enumTable.Rows)
	return
}
