package spec

import (
	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func (library *Library) toTypeDef(reader asciidoc.Reader, d *asciidoc.Document, s *asciidoc.Section, parent types.Entity) (ms *matter.TypeDef, err error) {
	name := text.TrimCaseInsensitiveSuffix(library.SectionName(s), " Type")
	ms = matter.NewTypeDef(s, parent)
	ms.Name = name

	var dt *types.DataType
	dt, err = GetDataType(library, reader, d, s)
	if err != nil {
		return nil, newGenericParseError(s, "error parsing typedef: %v", err)
	}
	if (dt == nil) || !dt.BaseType.IsSimple() {
		return nil, newGenericParseError(s, "unknown typedef data type: \"%s\"", dt.Name)
	}
	ms.Type = dt
	library.addEntity(s, ms)
	ms.Name = CanonicalName(ms.Name)
	return
}
