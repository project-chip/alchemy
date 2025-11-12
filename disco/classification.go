package disco

import (
	"fmt"
	"log/slog"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
)

func (b *Baller) organizeClassificationSection(cxt *discoContext) (err error) {
	for _, classification := range cxt.parsed.classification {
		classificationTable := classification.table
		if classificationTable == nil || classificationTable.Element == nil {
			slog.Warn("Could not organize classification section, as no table was found", log.Path("source", classification.section))
			return nil
		}
		if classificationTable.ColumnMap == nil {
			slog.Debug("can't rearrange classification table without header row")
			return nil
		}

		if len(classificationTable.ColumnMap) < 3 {
			slog.Debug("can't rearrange classification table with so few matches")
			return nil
		}

		var tableType matter.TableType

		switch cxt.parsed.docType {
		case matter.DocTypeCluster:
			tableType = matter.TableTypeAppClusterClassification
		case matter.DocTypeDeviceType:
			tableType = matter.TableTypeDeviceTypeClassification
		}

		err = b.renameTableHeaderCells(cxt, classification.section, classificationTable, matter.Tables[tableType].ColumnRenames)
		if err != nil {
			return fmt.Errorf("error renaming table header cells in section %s in %s: %w", cxt.library.SectionName(classification.section), cxt.doc.Path, err)
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

func getClassificationInfo(cxt *discoContext, classificationTable *spec.TableInfo) (ci *classificationInfo) {
	ci = &classificationInfo{}
	hierarchyIndex, hasHierarchy := classificationTable.ColumnMap[matter.TableColumnHierarchy]
	for i, row := range classificationTable.Rows {
		if i == classificationTable.HeaderRowIndex {
			continue
		}
		if hasHierarchy {
			hierarchyCell := row.Cell(hierarchyIndex)
			vc, e := spec.RenderTableCell(cxt.library, hierarchyCell)
			if e != nil {
				continue
			}
			ci.hierarchy = vc
		}
	}
	return
}
