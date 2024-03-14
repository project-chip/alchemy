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

	err = b.renameTableHeaderCells(bitsTable.rows, bitsTable.headerRow, bitsTable.columnMap, nil)
	if err != nil {
		return fmt.Errorf("error renaming table header cells in section %s in %s: %w", bms.section.Name, dp.doc.Path, err)
	}

	b.addMissingColumns(dp.doc, bms.section, bitsTable.element, bitsTable.rows, matter.BitmapTableColumnOrder[:], nil, bitsTable.headerRow, bitsTable.columnMap, types.EntityTypeBitmapValue)

	b.reorderColumns(dp.doc, bms.section, bitsTable.rows, matter.BitmapTableColumnOrder[:], bitsTable.columnMap, bitsTable.extraColumns)

	b.appendSubsectionTypes(bms.section, bitsTable.columnMap, bitsTable.rows)
	return
}
