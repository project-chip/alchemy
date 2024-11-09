package disco

import (
	"fmt"
	"log/slog"

	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
)

func (b *Baller) organizeClassificationSection(cxt *discoContext) (err error) {
	for _, classification := range cxt.parsed.classification {
		classificationTable := classification.table
		if classificationTable == nil || classificationTable.Element == nil {
			return fmt.Errorf("no classification table found")
		}
		if classificationTable.ColumnMap == nil {
			slog.Debug("can't rearrange classification table without header row")
			return nil
		}

		if len(classificationTable.ColumnMap) < 3 {
			slog.Debug("can't rearrange classification table with so few matches")
			return nil
		}

		err = b.renameTableHeaderCells(cxt, classification.section, classificationTable, matter.Tables[matter.TableTypeClassification].ColumnRenames)
		if err != nil {
			return fmt.Errorf("error renaming table header cells in section %s in %s: %w", classification.section.Name, cxt.doc.Path, err)
		}

		var tableType matter.TableType

		switch cxt.parsed.docType {
		case matter.DocTypeCluster:
			tableType = matter.TableTypeAppClusterClassification
		case matter.DocTypeDeviceType:
			tableType = matter.TableTypeDeviceTypeClassification
		}

		if tableType != matter.TableTypeUnknown {
			err = b.reorderColumns(cxt, classification.section, classificationTable, tableType)
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

func getClassificationInfo(classificationTable *spec.TableInfo) (ci *classificationInfo) {
	ci = &classificationInfo{}
	hierarchyIndex, hasHierarchy := classificationTable.ColumnMap[matter.TableColumnHierarchy]
	for i, row := range classificationTable.Rows {
		if i == classificationTable.HeaderRowIndex {
			continue
		}
		if hasHierarchy {
			hierarchyCell := row.Cell(hierarchyIndex)
			vc, e := spec.RenderTableCell(hierarchyCell)
			if e != nil {
				continue
			}
			ci.hierarchy = vc
		}
	}
	return
}
