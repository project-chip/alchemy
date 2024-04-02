package disco

import (
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/matter"
	mattertypes "github.com/hasty/alchemy/matter/types"
)

func (b *Ball) fixAccessCells(doc *ascii.Doc, table *tableInfo, entityType mattertypes.EntityType) (err error) {
	if !b.options.formatAccess {
		return nil
	}
	if len(table.rows) < 2 {
		return
	}
	accessIndex, ok := table.columnMap[matter.TableColumnAccess]
	if !ok {
		return
	}
	for _, row := range table.rows[1:] {
		cell := row.Cells[accessIndex]
		vc, e := ascii.RenderTableCell(cell)
		if e != nil {
			continue
		}
		access := ascii.ParseAccess(vc, entityType)
		replacementAccess := ascii.AccessToASCIIDocString(access, entityType)
		if vc != replacementAccess {
			err = setCellString(cell, replacementAccess)
			if err != nil {
				return
			}

		}
	}
	return
}
