package disco

import (
	"strings"

	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/spec"
	mattertypes "github.com/hasty/alchemy/matter/types"
)

func (b *Ball) fixAccessCells(dp *docParse, subSection *subSection, entityType mattertypes.EntityType) (err error) {
	if !b.options.formatAccess {
		return nil
	}
	table := &subSection.table
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
		accessCell := row.Cell(accessIndex)
		vc, e := spec.RenderTableCell(accessCell)
		if e != nil {
			continue
		}
		var access matter.Access
		var parsed bool
		if len(strings.TrimSpace(vc)) > 0 {
			access, parsed = spec.ParseAccess(vc, entityType)
			if !parsed {
				continue
			}
		} else {
			c := getSubsectionCluster(dp, subSection.section)
			if c != nil {
				ci := getClassificationInfo(&c.classification.table)
				if ci.hierarchy != "Base" {
					continue
				}
			}
			access = matter.DefaultAccess(entityType)
		}
		if directionIndex >= 0 {
			directionCell := row.Cell(directionIndex)
			rc, e := spec.RenderTableCell(directionCell)
			if e != nil {
				continue
			}
			direction := spec.ParseCommandDirection(rc)
			if direction == matter.InterfaceClient {
				access.Invoke = matter.PrivilegeUnknown
			}
		}
		replacementAccess := spec.AccessToASCIIDocString(access, entityType)
		if vc != replacementAccess {
			err = setCellString(accessCell, replacementAccess)
			if err != nil {
				return
			}

		}
	}
	return
}
