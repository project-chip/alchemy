package disco

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
)

func (b *Baller) organizeFeaturesSection(cxt *discoContext) (err error) {
	for _, features := range cxt.parsed.features {
		featuresTable := features.table
		if featuresTable == nil || featuresTable.Element == nil {
			slog.Warn("Could not organize Features section, as no table of features was found", log.Path("source", features.section.Base))
			return
		}

		if featuresTable.ColumnMap == nil {
			return fmt.Errorf("can't rearrange features table without header row in %s", cxt.doc.Path)
		}

		err = b.renameTableHeaderCells(cxt, features.section, featuresTable, matter.Tables[matter.TableTypeFeatures].ColumnRenames)
		if err != nil {
			return fmt.Errorf("error renaming table header cells in features table in %s: %w", cxt.doc.Path, err)
		}

		err = b.reorderColumns(cxt, features.section, featuresTable, matter.TableTypeFeatures)
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
					setCellString(featureCell, vc)
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
					setCellString(codeCell, uc)
				}
			}
		}
	}
	return
}
