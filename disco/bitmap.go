package disco

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/hasty/alchemy/matter"
)

func (b *Ball) organizeBitmapSections(cxt *discoContext, dp *docParse) (err error) {
	for _, bms := range dp.bitmaps {
		err = b.organizeBitmapSection(cxt, dp, bms)
		if err != nil {
			return
		}
	}
	return
}

func (b *Ball) organizeBitmapSection(cxt *discoContext, dp *docParse, bms *subSection) (err error) {
	name := strings.TrimSpace(bms.section.Name)
	if strings.HasSuffix(strings.ToLower(name), "bitmap") {
		setSectionTitle(bms.section, name+" Type")
	}
	bitsTable := bms.table
	if bitsTable.element == nil {
		return
	}
	if bitsTable.columnMap == nil {
		slog.Debug("can't rearrange bitmap table without header row")
		return nil
	}

	if len(bitsTable.columnMap) < 2 {
		slog.Debug("can't rearrange bitmap table with so few matches")
		return nil
	}

	err = b.renameTableHeaderCells(bitsTable.rows, bitsTable.headerRow, bitsTable.columnMap, nil)
	if err != nil {
		return fmt.Errorf("error renaming table header cells in section %s in %s: %w", bms.section.Name, dp.doc.Path, err)
	}

	b.addMissingColumns(dp.doc, bms.section, bitsTable.element, bitsTable.rows, matter.BitmapTableColumnOrder[:], nil, bitsTable.headerRow, bitsTable.columnMap)

	b.reorderColumns(dp.doc, bms.section, bitsTable.rows, matter.BitmapTableColumnOrder[:], bitsTable.columnMap, bitsTable.extraColumns)

	b.appendSubsectionTypes(bms.section, bitsTable.columnMap, bitsTable.rows)
	return
}

/*
func (b *Ball) organizeBitmapSection(doc *ascii.Doc, section *ascii.Section) error {
	bitsTable := ascii.FindFirstTable(section)
	if bitsTable == nil {
		return nil
	}
	name := strings.TrimSpace(section.Name)
	if strings.HasSuffix(strings.ToLower(name), "bitmap") {
		setSectionTitle(section, name+" Type")
	}
	return b.organizeBitmapTable(doc, section, bitsTable)
}

func (b *Ball) organizeBitmapTable(doc *ascii.Doc, section *ascii.Section, bitsTable *types.Table) error {
	rows := ascii.TableRows(bitsTable)

	headerRowIndex, columnMap, extraColumns, err := ascii.MapTableColumns(doc, rows)
	if err != nil {
		return fmt.Errorf("failed mapping table columns for bitmap table in section %s: %w", section.Name, err)
	}

	if columnMap == nil {
		slog.Debug("can't rearrange bitmap table without header row")
		return nil
	}

	if len(columnMap) < 2 {
		slog.Debug("can't rearrange bitmap table with so few matches")
		return nil
	}

	_, ok := columnMap[matter.TableColumnBit]
	if !ok {
		idIndex, ok := columnMap[matter.TableColumnID]
		if ok {
			delete(columnMap, matter.TableColumnID)
			columnMap[matter.TableColumnBit] = idIndex
		}
	}

	err = b.renameTableHeaderCells(rows, headerRowIndex, columnMap, nil)
	if err != nil {
		return fmt.Errorf("error renaming table header cells in section %s in %s: %w", section.Name, doc.Path, err)
	}

	b.addMissingColumns(doc, section, bitsTable, rows, matter.BitmapTableColumnOrder[:], nil, headerRowIndex, columnMap)

	b.reorderColumns(doc, section, rows, matter.BitmapTableColumnOrder[:], columnMap, extraColumns)

	b.appendSubsectionTypes(section, columnMap, rows)

	return nil
}
*/
