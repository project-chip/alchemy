package spec

import (
	"log/slog"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func (s *Section) toStatusCodes(d *Doc, pc *parseContext, parent types.Entity) (e *matter.Enum, err error) {

	name := CanonicalName(text.TrimCaseInsensitiveSuffix(s.Name, " Type"))
	e = matter.NewEnum(s.Base, parent)
	e.Name = "StatusCodeEnum"
	dt := s.GetDataType()
	if dt == nil {
		dt = types.NewDataType(types.BaseDataTypeEnum8, false)
		slog.Warn("Status code does not declare its derived data type; assuming enum8", log.Element("source", d.Path, s.Base), slog.String("enum", name))
	} else if !dt.IsEnum() {
		return nil, newGenericParseError(s.Base, "unknown status code data type: %s", dt.Name)
	}

	e.Type = dt

	e.Values, err = s.findEnumValues(e)
	if err != nil {
		return
	}

	pc.orderedEntities = append(pc.orderedEntities, e)
	pc.entitiesByElement[s.Base] = append(pc.entitiesByElement[s.Base], e)
	e.Name = CanonicalName(e.Name)
	return
}
