package spec

import (
	"fmt"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func (s *Section) toTypeDef(d *Doc, entityMap map[asciidoc.Attributable][]types.Entity) (ms *matter.TypeDef, err error) {
	name := text.TrimCaseInsensitiveSuffix(s.Name, " Type")
	ms = matter.NewTypeDef(s.Base)
	ms.Name = name

	dt := s.GetDataType()
	if (dt == nil) || !dt.BaseType.IsSimple() {
		return nil, fmt.Errorf("unknown typedef data type: %s", dt.Name)
	}
	ms.Type = dt
	entityMap[s.Base] = append(entityMap[s.Base], ms)
	ms.Name = CanonicalName(ms.Name)
	return
}
