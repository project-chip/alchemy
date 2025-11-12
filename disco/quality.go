package disco

import (
	"strings"

	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
)

func (b *Baller) fixQualityCells(cxt *discoContext, subSection *subSection) (err error) {
	if !b.options.FormatQuality {
		return nil
	}
	if cxt.errata.IgnoreSection(cxt.library.SectionName(subSection.section), errata.DiscoPurposeTableQuality) {
		return nil
	}
	table := subSection.table
	if len(table.Rows) < 2 {
		return
	}
	qualityIndex, ok := table.ColumnMap[matter.TableColumnQuality]
	if !ok {
		return
	}

	for _, row := range table.Rows[1:] {
		qualityCell := row.Cell(qualityIndex)
		vc, e := spec.RenderTableCell(cxt.library, qualityCell)
		if e != nil {
			continue
		}
		vc = strings.TrimSpace(vc)
		if len(vc) == 0 {
			continue
		}
		quality := matter.ParseQuality(vc)
		replacementQuality := quality.String()
		if vc != replacementQuality {
			setCellString(qualityCell, replacementQuality)
		}
	}
	return
}
