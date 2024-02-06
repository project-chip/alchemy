package disco

import (
	"fmt"
	"log/slog"

	"github.com/hasty/alchemy/matter"
)

func (b *Ball) organizeClassificationSection(cxt *discoContext, dp *docParse) (err error) {
	for _, classification := range dp.classification {
		classificationTable := classification.table
		if classificationTable.element == nil {
			return fmt.Errorf("no classification table found")
		}
		if classificationTable.columnMap == nil {
			slog.Debug("can't rearrange classification table without header row")
			return nil
		}

		if len(classificationTable.columnMap) < 3 {
			slog.Debug("can't rearrange classification table with so few matches")
			return nil
		}

		err = b.renameTableHeaderCells(classificationTable.rows, classificationTable.headerRow, classificationTable.columnMap, matter.ClassificationTableColumnNames)
		if err != nil {
			return fmt.Errorf("error renaming table header cells in section %s in %s: %w", classification.section.Name, dp.doc.Path, err)
		}

		var order []matter.TableColumn

		switch dp.docType {
		case matter.DocTypeCluster:
			order = matter.AppClusterClassificationTableColumnOrder[:]
		case matter.DocTypeDeviceType:
			order = matter.DeviceTypeClassificationTableColumnOrder[:]
		}

		if len(order) > 0 {
			b.reorderColumns(dp.doc, classification.section, classificationTable.rows, order, classificationTable.columnMap, classificationTable.extraColumns)
		}
	}
	return
}
