package disco

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
)

func (b *Ball) organizeFeaturesSection(cxt *discoContext, dp *docParse) (err error) {
	for _, features := range dp.features {
		featuresTable := features.table
		if featuresTable.Element == nil {
			return fmt.Errorf("no features section found")
		}

		if featuresTable.ColumnMap == nil {
			return fmt.Errorf("can't rearrange features table without header row in %s", dp.doc.Path)
		}

		err = b.renameTableHeaderCells(dp.doc, features.section, featuresTable, matter.Tables[matter.TableTypeFeatures].ColumnNames)
		if err != nil {
			return fmt.Errorf("error renaming table header cells in features table in %s: %w", dp.doc.Path, err)
		}

		err = b.reorderColumns(dp.doc, features.section, featuresTable, matter.TableTypeFeatures)
		if err != nil {
			return err
		}

		featureIndex, ok := featuresTable.ColumnMap[matter.TableColumnFeature]
		if ok {
			for i, row := range featuresTable.Rows {
				if i == featuresTable.HeaderRowIndex {
					continue
				}
				featureCell := row.Cell(featureIndex)
				vc, e := spec.RenderTableCell(featureCell)
				if e != nil {
					continue
				}
				vc = strings.TrimSpace(vc)
				if strings.Contains(vc, " ") {
					vc = matter.Case(vc)
					slog.Debug("fixing feature name", "name", vc)
					err = setCellString(featureCell, vc)
					if err != nil {
						return
					}
				}

			}
		}

		codeIndex, ok := featuresTable.ColumnIndex(matter.TableColumnCode)
		if ok {
			for i, row := range featuresTable.Rows {
				if i == featuresTable.HeaderRowIndex {
					continue
				}
				codeCell := row.Cell(codeIndex)
				vc, e := spec.RenderTableCell(codeCell)
				if e != nil {
					continue
				}
				vc = strings.TrimSpace(vc)
				uc := strings.ToUpper(vc)
				if uc != vc {
					slog.Debug("fixing feature code", "name", vc)
					err = setCellString(codeCell, uc)
					if err != nil {
						return
					}
				}
			}
		}
	}
	return
}
