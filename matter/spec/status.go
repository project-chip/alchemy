package spec

import (
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func toStatusCodes(d *Doc, s *asciidoc.Section, pc *parseContext, parent types.Entity) (e *matter.Enum, err error) {

	name := CanonicalName(text.TrimCaseInsensitiveSuffix(d.SectionName(s), " Type"))
	e = matter.NewEnum(s, parent)
	e.Name = "StatusCodeEnum"
	dt := GetDataType(d, s)
	if dt == nil {
		dt = types.NewDataType(types.BaseDataTypeEnum8, false)
		slog.Warn("Status code does not declare its derived data type; assuming enum8", log.Element("source", d.Path, s), slog.String("enum", name))
	} else if !dt.IsEnum() {
		return nil, newGenericParseError(s, "unknown status code data type: %s", dt.Name)
	}

	e.Type = dt

	e.Values, err = findEnumValues(d, s, e)
	if err != nil {
		return
	}

	pc.orderedEntities = append(pc.orderedEntities, e)
	pc.entitiesByElement[s] = append(pc.entitiesByElement[s], e)
	e.Name = CanonicalName(e.Name)
	return
}
