package spec

import (
	"fmt"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func (s *Section) toStruct(d *Doc, entityMap map[asciidoc.Attributable][]types.Entity) (ms *matter.Struct, err error) {
	name := text.TrimCaseInsensitiveSuffix(s.Name, " Type")
	var ti *TableInfo
	ti, err = parseFirstTable(d, s)
	if err != nil {
		return nil, fmt.Errorf("failed reading struct %s: %w", name, err)
	}
	ms = matter.NewStruct(s.Base)
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
	ms.Fields, err = d.readFields(ti, types.EntityTypeStructField)
	if err != nil {
		return
	}
	entityMap[s.Base] = append(entityMap[s.Base], ms)
	fieldMap := make(map[string]*matter.Field, len(ms.Fields))
	for _, f := range ms.Fields {
		fieldMap[f.Name] = f
	}
	err = s.mapFields(fieldMap, entityMap)
	if err != nil {
		return
	}
	ms.Name = CanonicalName(ms.Name)
	return
}
