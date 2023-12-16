package ascii

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/matter"
)

func (s *Section) toStruct(d *Doc) (ms *matter.Struct, err error) {
	var rows []*types.TableRow
	var headerRowIndex int
	var columnMap ColumnIndex
	rows, headerRowIndex, columnMap, _, err = parseFirstTable(d, s)
	if err != nil {
		return nil, fmt.Errorf("failed reading struct: %w", err)
	}
	name := strings.TrimSuffix(s.Name, " Type")
	ms = &matter.Struct{
		Name: name,
	}

	if headerRowIndex > 0 {
		firstRow := rows[0]
		if len(firstRow.Cells) > 0 {
			cv, rowErr := GetTableCellValue(rows[0].Cells[0])
			if rowErr == nil {
				cv = strings.ToLower(cv)
				if strings.Contains(cv, "fabric scoped") || strings.Contains(cv, "fabric-scoped") {
					ms.FabricScoped = true
				}
			}
		}
	}
	ms.Fields, err = d.readFields(headerRowIndex, rows, columnMap)
	return
}
