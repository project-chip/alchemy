package disco

import (
	"fmt"
	"log/slog"

	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func (b *Ball) organizeStructSections(cxt *discoContext, dp *docParse) (err error) {
	for _, ss := range dp.structs {
		err = b.organizeStructSection(cxt, dp, ss)
		if err != nil {
			return
		}
	}
	return
}

func (b *Ball) organizeStructSection(cxt *discoContext, dp *docParse, ss *subSection) (err error) {
	if b.doc.Errata().IgnoreSection(ss.section.Name, errata.PurposeDataTypesStruct) {
		return
	}
	b.canonicalizeDataTypeSectionName(dp, ss.section, "Struct")

	fieldsTable := &ss.table
	if fieldsTable.element == nil {
		slog.Debug("no struct table found")
		return nil
	}
	if fieldsTable.columnMap == nil {
		slog.Debug("can't rearrange struct table without header row")
		return nil
	}

	if len(fieldsTable.columnMap) < 2 {
		slog.Debug("can't rearrange struct table with so few matches")
		return nil
	}

	err = b.fixAccessCells(dp, ss, types.EntityTypeStruct)
	if err != nil {
		return fmt.Errorf("error fixing access cells in struct table in %s: %w", dp.doc.Path, err)
	}

	err = fixConstraintCells(dp.doc, fieldsTable)
	if err != nil {
		return err
	}

	err = b.renameTableHeaderCells(dp.doc, ss.section, fieldsTable, nil)
	if err != nil {
		return fmt.Errorf("error renaming table header cells in struct table in section %s in %s: %w", ss.section.Name, dp.doc.Path, err)
	}

	err = b.addMissingColumns(ss.section, fieldsTable, matter.Tables[matter.TableTypeStruct], types.EntityTypeStructField)
	if err != nil {
		return fmt.Errorf("error adding missing table columns in struct table in section %s in %s: %w", ss.section.Name, dp.doc.Path, err)
	}

	err = b.reorderColumns(dp.doc, ss.section, fieldsTable, matter.TableTypeStruct)
	if err != nil {
		return err
	}

	b.appendSubsectionTypes(ss.section, fieldsTable.columnMap, fieldsTable.rows)

	err = b.linkIndexTables(cxt, ss)
	if err != nil {
		return err
	}
	return
}
