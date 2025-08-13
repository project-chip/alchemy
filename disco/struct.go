package disco

import (
	"fmt"
	"log/slog"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func (b *Baller) organizeStructSections(cxt *discoContext) (err error) {
	for _, ss := range cxt.parsed.structs {
		err = b.organizeStructSection(cxt, ss)
		if err != nil {
			return
		}
	}
	return
}

func (b *Baller) organizeStructSection(cxt *discoContext, ss *subSection) (err error) {

	b.canonicalizeDataTypeSectionName(cxt, ss.section, "Struct")

	fieldsTable := ss.table
	if fieldsTable == nil || fieldsTable.Element == nil {
		slog.Warn("Could not organize struct section, as no table of struct fields was found", log.Path("source", ss.section))
		return nil
	}
	if fieldsTable.ColumnMap == nil {
		slog.Debug("can't rearrange struct table without header row")
		return nil
	}

	if len(fieldsTable.ColumnMap) < 2 {
		slog.Debug("can't rearrange struct table with so few matches")
		return nil
	}

	err = b.renameTableHeaderCells(cxt, ss.section, fieldsTable, nil)
	if err != nil {
		return fmt.Errorf("error renaming table header cells in section %s in %s: %w", ss.section.Name, cxt.doc.Path, err)
	}

	err = b.fixAccessCells(cxt, ss, types.EntityTypeStruct)
	if err != nil {
		return fmt.Errorf("error fixing access cells in struct table in %s: %w", cxt.doc.Path, err)
	}

	err = b.fixConstraintCells(cxt, ss.section, fieldsTable)
	if err != nil {
		return err
	}

	err = b.renameTableHeaderCells(cxt, ss.section, fieldsTable, matter.Tables[matter.TableTypeStruct].ColumnRenames)
	if err != nil {
		return fmt.Errorf("error renaming table header cells in struct table in section %s in %s: %w", ss.section.Name, cxt.doc.Path, err)
	}

	err = b.addMissingColumns(cxt, ss.section, fieldsTable, matter.Tables[matter.TableTypeStruct], types.EntityTypeStructField)
	if err != nil {
		return fmt.Errorf("error adding missing table columns in struct table in section %s in %s: %w", ss.section.Name, cxt.doc.Path, err)
	}

	err = b.reorderColumns(cxt, ss.section, fieldsTable, matter.TableTypeStruct)
	if err != nil {
		return err
	}

	b.appendSubsectionTypes(cxt, ss.section, fieldsTable.ColumnMap, fieldsTable.Rows)
	b.removeMandatoryFallbacks(fieldsTable)

	err = b.linkIndexTables(cxt, ss)
	if err != nil {
		return err
	}
	return
}
