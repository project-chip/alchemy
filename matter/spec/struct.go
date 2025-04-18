package spec

import (
	"fmt"
	"strings"

	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func (s *Section) toStruct(d *Doc, pc *parseContext, parent types.Entity) (ms *matter.Struct, err error) {
	name := text.TrimCaseInsensitiveSuffix(s.Name, " Type")
	var ti *TableInfo
	ti, err = parseFirstTable(d, s)
	if err != nil {
		return nil, fmt.Errorf("failed reading struct %s: %w", name, err)
	}
	ms = matter.NewStruct(s.Base, parent)
	ms.Name = name

	if ti.HeaderRowIndex > 0 {
		firstRow := ti.Rows[0]
		tableCells := firstRow.TableCells()
		if len(tableCells) > 0 {
			cv, rowErr := RenderTableCell(tableCells[0])
			if rowErr == nil {
				cv = strings.ToLower(cv)
				if strings.Contains(cv, "fabric scoped") || strings.Contains(cv, "fabric-scoped") {
					ms.FabricScoping = matter.FabricScopingScoped
				}
			}
		}
	}
	var fieldMap map[string]*matter.Field
	ms.Fields, fieldMap, err = d.readFields(ti, types.EntityTypeStructField, ms)
	if err != nil {
		return
	}
	pc.orderedEntities = append(pc.orderedEntities, ms)
	pc.entitiesByElement[s.Base] = append(pc.entitiesByElement[s.Base], ms)
	err = s.mapFields(fieldMap, pc)
	if err != nil {
		return
	}
	ms.Name = CanonicalName(ms.Name)
	return
}
