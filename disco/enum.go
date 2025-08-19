package disco

import (
	"fmt"
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func (b *Baller) organizeEnumSections(cxt *discoContext) (err error) {
	for _, es := range cxt.parsed.enums {
		err = b.organizeEnumSection(cxt, es)
		if err != nil {
			return
		}
	}
	return
}

func (b *Baller) organizeEnumSection(cxt *discoContext, es *subSection) (err error) {

	b.canonicalizeDataTypeSectionName(cxt, es.section, "Enum")

	enumTable := es.table
	if enumTable == nil || enumTable.Element == nil {
		slog.Warn("Could not organize enum section, as no table of enum values was found", log.Path("source", es.section))
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

	err = b.renameTableHeaderCells(cxt, es.section, enumTable, matter.Tables[matter.TableTypeEnum].ColumnRenames)
	if err != nil {
		return fmt.Errorf("error renaming table header cells in enum table in section %s in %s: %w", cxt.doc.SectionName(es.section), cxt.doc.Path, err)
	}

	err = b.addMissingColumns(cxt, es.section, enumTable, matter.Tables[matter.TableTypeEnum], types.EntityTypeEnumValue)
	if err != nil {
		return fmt.Errorf("error adding missing table columns in enum section %s in %s: %w", cxt.doc.SectionName(es.section), cxt.doc.Path, err)
	}

	err = enumTable.Rescan(cxt.doc, asciidoc.NewRawReader())
	if err != nil {
		return fmt.Errorf("error reordering columns in enum table in section %s in %s: %w", cxt.doc.SectionName(es.section), cxt.doc.Path, err)

	}
	enumTable = es.table
	b.removeMandatoryFallbacks(enumTable)

	err = b.reorderColumns(cxt, es.section, enumTable, matter.TableTypeEnum)
	if err != nil {
		return err
	}

	b.appendSubsectionTypes(cxt, es.section, enumTable.ColumnMap, enumTable.Rows)
	return
}
