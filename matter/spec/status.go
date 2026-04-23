package spec

import (
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func (library *Library) toStatusCodes(reader asciidoc.Reader, d *asciidoc.Document, s *asciidoc.Section, parent types.Entity) (e *matter.Enum, err error) {

	name := CanonicalName(text.TrimCaseInsensitiveSuffix(library.SectionName(s), " Type"))
	e = matter.NewEnum(s, parent)
	e.Name = "StatusCodeEnum"
	var dt *types.DataType
	dt, err = GetDataType(library, reader, d, s)
	if err != nil {
		return nil, newGenericParseError(s, "error parsing error code type: %v", err)
	}
	if dt == nil {
		dt = types.NewDataType(types.BaseDataTypeEnum8, types.DataTypeRankScalar)
		slog.Debug("Status code does not declare its derived data type; assuming enum8", log.Element("source", d.Path, s), slog.String("enum", name))
	} else if !dt.IsEnum() {
		return nil, newGenericParseError(s, "unknown status code data type: %s", dt.Name)
	}

	e.Type = dt

	e.Values, err = library.findEnumValues(reader, d, s, e)
	if err != nil {
		return
	}
	library.addEntity(s, e)
	e.Name = CanonicalName(e.Name)
	return
}
