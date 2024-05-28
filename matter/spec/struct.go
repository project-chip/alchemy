package spec

import (
	"fmt"
	"strings"

	"github.com/hasty/alchemy/asciidoc"
	"github.com/hasty/alchemy/matter"
	mattertypes "github.com/hasty/alchemy/matter/types"
)

func (s *Section) toStruct(d *Doc, entityMap map[asciidoc.Attributable][]mattertypes.Entity) (ms *matter.Struct, err error) {
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
	ms.Fields, err = d.readFields(headerRowIndex, rows, columnMap, mattertypes.EntityTypeStruct)
	entityMap[s.Base] = append(entityMap[s.Base], ms)
	return
}