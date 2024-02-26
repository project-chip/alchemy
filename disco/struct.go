package disco

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/hasty/alchemy/matter"
	mattertypes "github.com/hasty/alchemy/matter/types"
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
	name := strings.TrimSpace(ss.section.Name)
	if strings.HasSuffix(strings.ToLower(name), "struct") {
		setSectionTitle(ss.section, name+" Type")
	}
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

	err = b.fixAccessCells(dp.doc, fieldsTable, mattertypes.EntityTypeStruct)
	if err != nil {
		return fmt.Errorf("error fixing access cells in struct table in %s: %w", dp.doc.Path, err)
	}

	err = b.renameTableHeaderCells(fieldsTable.rows, fieldsTable.headerRow, fieldsTable.columnMap, nil)
	if err != nil {
		return fmt.Errorf("error renaming table header cells in struct table in section %s in %s: %w", ss.section.Name, dp.doc.Path, err)
	}

	err = b.addMissingColumns(dp.doc, ss.section, fieldsTable.element, fieldsTable.rows, matter.StructTableColumnOrder[:], nil, fieldsTable.headerRow, fieldsTable.columnMap)
	if err != nil {
		return fmt.Errorf("error adding missing table columns in struct table in section %s in %s: %w", ss.section.Name, dp.doc.Path, err)
	}

	b.reorderColumns(dp.doc, ss.section, fieldsTable.rows, matter.StructTableColumnOrder[:], fieldsTable.columnMap, fieldsTable.extraColumns)

	b.appendSubsectionTypes(ss.section, fieldsTable.columnMap, fieldsTable.rows)
	return
}
