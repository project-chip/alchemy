package disco

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/types"
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

	err = b.renameTableHeaderCells(b.doc, &bitsTable, nil)
	if err != nil {
		return fmt.Errorf("error renaming table header cells in section %s in %s: %w", bms.section.Name, dp.doc.Path, err)
	}

	err = b.addMissingColumns(dp.doc, bms.section, bitsTable.element, bitsTable.rows, matter.Tables[matter.TableTypeBitmap], nil, bitsTable.headerRow, bitsTable.columnMap, types.EntityTypeBitmapValue)
	if err != nil {
		return fmt.Errorf("error adding missing table columns in bitmap section %s in %s: %w", bms.section.Name, dp.doc.Path, err)
	}

	err = b.reorderColumns(dp.doc, bms.section, &bitsTable, matter.TableTypeBitmap)
	if err != nil {
		return err
	}

	b.appendSubsectionTypes(bms.section, bitsTable.columnMap, bitsTable.rows)
	return
}
