package disco

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/matter"
	mattertypes "github.com/hasty/alchemy/matter/types"
)

func (b *Ball) fixAccessCells(doc *ascii.Doc, rows []*types.TableRow, columnMap ascii.ColumnIndex, entityType mattertypes.EntityType) (err error) {
	if !b.options.formatAccess {
		return nil
	}
	if len(rows) < 2 {
		return
	}
	accessIndex, ok := columnMap[matter.TableColumnAccess]
	if !ok {
		return
	}
	for _, row := range rows[1:] {
		cell := row.Cells[accessIndex]
		vc, e := ascii.RenderTableCell(cell)
		if e != nil {
			continue
		}
		access := ascii.ParseAccess(vc, entityType)
		replacementAccess := ascii.AccessToAsciiString(access, entityType)
		if vc != replacementAccess {
			err = setCellString(cell, replacementAccess)
			if err != nil {
				return
			}

		}
	}
	return
}
