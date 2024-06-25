package spec

import (
	"fmt"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func (s *Section) toStruct(d *Doc, entityMap map[asciidoc.Attributable][]types.Entity) (ms *matter.Struct, err error) {
	name := strings.TrimSuffix(s.Name, " Type")
	var rows []*asciidoc.TableRow
	var headerRowIndex int
	var columnMap ColumnIndex
	rows, headerRowIndex, columnMap, _, err = parseFirstTable(d, s)
	if err != nil {
		return nil, fmt.Errorf("failed reading struct %s: %w", name, err)
	}
	ms = &matter.Struct{
		Name: name,
	}

	if headerRowIndex > 0 {
		firstRow := rows[0]
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
	ms.Fields, err = d.readFields(headerRowIndex, rows, columnMap, types.EntityTypeStruct)
	entityMap[s.Base] = append(entityMap[s.Base], ms)
	return
}
