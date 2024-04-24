package disco

import (
	"fmt"
	"log/slog"

	"github.com/hasty/alchemy/ascii"
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

		err = b.renameTableHeaderCells(classificationTable.rows, classificationTable.headerRow, classificationTable.columnMap, matter.Tables[matter.TableTypeClassification].ColumnNames)
		if err != nil {
			return fmt.Errorf("error renaming table header cells in section %s in %s: %w", classification.section.Name, dp.doc.Path, err)
		}

		var tableType matter.TableType

		switch dp.docType {
		case matter.DocTypeCluster:
			tableType = matter.TableTypeAppClusterClassification
		case matter.DocTypeDeviceType:
			tableType = matter.TableTypeDeviceTypeClassification
		}

		if tableType != matter.TableTypeUnknown {
			err = b.reorderColumns(dp.doc, classification.section, &classificationTable, tableType)
			if err != nil {
				return err
			}
		}
	}
	return
}

type classificationInfo struct {
	hierarchy string
}

func getClassificationInfo(classificationTable *tableInfo) (ci *classificationInfo) {
	ci = &classificationInfo{}
	hierarchyIndex, hasHierarchy := classificationTable.columnMap[matter.TableColumnHierarchy]
	for i, row := range classificationTable.rows {
		if i == classificationTable.headerRow {
			continue
		}
		if hasHierarchy {
			hierarchyCell := row.Cells[hierarchyIndex]
			vc, e := ascii.RenderTableCell(hierarchyCell)
			if e != nil {
				continue
			}
			ci.hierarchy = vc
		}
	}
	return
}
