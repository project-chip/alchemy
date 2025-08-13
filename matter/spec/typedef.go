package spec

import (
	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func toTypeDef(d *Doc, s *asciidoc.Section, pc *parseContext, parent types.Entity) (ms *matter.TypeDef, err error) {
	name := text.TrimCaseInsensitiveSuffix(d.SectionName(s), " Type")
	ms = matter.NewTypeDef(s, parent)
	ms.Name = name

	dt := GetDataType(d, s)
	if (dt == nil) || !dt.BaseType.IsSimple() {
		return nil, newGenericParseError(s, "unknown typedef data type: \"%s\"", dt.Name)
	}
	ms.Type = dt
	pc.orderedEntities = append(pc.orderedEntities, ms)
	pc.entitiesByElement[s] = append(pc.entitiesByElement[s], ms)
	ms.Name = CanonicalName(ms.Name)
	return
}
