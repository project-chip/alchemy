package disco

import (
	"strings"

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
	var directionIndex int
	directionIndex = -1
	if entityType == mattertypes.EntityTypeCommand {
		directionIndex, ok = table.columnMap[matter.TableColumnDirection]
		if !ok {
			directionIndex = -1
		}
	}
	for _, row := range table.rows[1:] {
		accessCell := row.Cells[accessIndex]
		vc, e := ascii.RenderTableCell(accessCell)
		if e != nil {
			continue
		}
		var access matter.Access
		var parsed bool
		if len(strings.TrimSpace(vc)) > 0 {
			access, parsed = ascii.ParseAccess(vc, entityType)
			if !parsed {
				continue
			}
		} else {
			access = matter.DefaultAccess(entityType)
		}
		if directionIndex >= 0 {
			directionCell := row.Cells[directionIndex]
			rc, e := ascii.RenderTableCell(directionCell)
			if e != nil {
				continue
			}
			direction := ascii.ParseCommandDirection(rc)
			if direction == matter.InterfaceClient {
				access.Invoke = matter.PrivilegeUnknown
			}
		}
		replacementAccess := ascii.AccessToASCIIDocString(access, entityType)
		if vc != replacementAccess {
			err = setCellString(accessCell, replacementAccess)
			if err != nil {
				return
			}

		}
	}
	return
}
