package disco

import (
	"github.com/hasty/alchemy/asciidoc"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
	"github.com/hasty/alchemy/matter/spec"
)

func fixConformanceCells(doc *spec.Doc, rows []*asciidoc.TableRow, columnMap spec.ColumnIndex) (err error) {
	if len(rows) < 2 {
		return
	}
	conformanceIndex, ok := columnMap[matter.TableColumnConformance]
	if !ok {
		return
	}
	for _, row := range rows[1:] {
		cell := row.Cell(conformanceIndex)
		vc, e := spec.RenderTableCell(cell)
		if e != nil {
			continue
		}

		conf := conformance.ParseConformance(vc)

		cs := conf.ASCIIDocString()

		if cs != vc {
			err = setCellString(cell, cs)
			if err != nil {
				return
			}
		}

	}
	return
}
