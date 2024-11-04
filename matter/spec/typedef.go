package spec

import (
	"fmt"

	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
)

func (s *Section) toTypeDef(d *Doc, pc *parseContext) (ms *matter.TypeDef, err error) {
	name := text.TrimCaseInsensitiveSuffix(s.Name, " Type")
	ms = matter.NewTypeDef(s.Base)
	ms.Name = name

	dt := s.GetDataType()
	if (dt == nil) || !dt.BaseType.IsSimple() {
		return nil, fmt.Errorf("unknown typedef data type: %s", dt.Name)
	}
	ms.Type = dt
	pc.orderedEntities = append(pc.orderedEntities, ms)
	pc.entitiesByElement[s.Base] = append(pc.entitiesByElement[s.Base], ms)
	ms.Name = CanonicalName(ms.Name)
	return
}
