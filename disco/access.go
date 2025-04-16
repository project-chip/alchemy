package disco

import (
	"strings"

	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func (b *Baller) fixAccessCells(cxt *discoContext, subSection *subSection, entityType types.EntityType) (err error) {
	if !b.options.FormatAccess {
		return nil
	}
	if cxt.errata.IgnoreSection(subSection.section.Name, errata.DiscoPurposeTableAccess) {
		return nil
	}
	table := subSection.table
	if len(table.Rows) < 2 {
		return
	}
	accessIndex, ok := table.ColumnMap[matter.TableColumnAccess]
	if !ok {
		return
	}
	var directionIndex int
	directionIndex = -1
	if entityType == types.EntityTypeCommand {
		directionIndex, ok = table.ColumnMap[matter.TableColumnDirection]
		if !ok {
			directionIndex = -1
		}
	}
	for _, row := range table.Rows[1:] {
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
			c := getSubsectionCluster(cxt.parsed, subSection.section)
			if c != nil {
				ci := getClassificationInfo(c.classification.table)
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
			setCellString(accessCell, replacementAccess)

		}
	}
	return
}
