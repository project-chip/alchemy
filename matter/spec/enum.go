package spec

import (
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/suggest"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

func (library *Library) toEnum(reader asciidoc.Reader, d *asciidoc.Document, s *asciidoc.Section, parent types.Entity) (e *matter.Enum, err error) {

	name := CanonicalName(text.TrimCaseInsensitiveSuffix(library.SectionName(s), " Type"))
	e = matter.NewEnum(s, parent)
	e.Name = name
	var dt *types.DataType
	dt, err = GetDataType(library, reader, d, s)
	if err != nil {
		return nil, newGenericParseError(s, "error parsing enum data type: %v", err)
	}
	if dt == nil {
		dt = types.NewDataType(types.BaseDataTypeEnum8, false)
		slog.Warn("Enum does not declare its derived data type; assuming enum8", log.Element("source", d.Path, s), slog.String("enum", name))
	} else if !dt.IsEnum() {
		return nil, newGenericParseError(s, "unknown enum data type: \"%s\"", dt.Name)
	}

	e.Type = dt

	e.Values, err = library.findEnumValues(reader, d, s, e)
	if err != nil {
		return
	}

	if len(e.Values) == 0 {
		var subSectionValues matter.EnumValueSet
		for el := range reader.Iterate(s, reader.Children(s)) {
			switch el := el.(type) {
			case *asciidoc.Section:
				if strings.HasSuffix(library.SectionName(el), " Range") {
					var ssv matter.EnumValueSet
					ssv, err = library.findEnumValues(reader, d, el, e)
					if err != nil {
						continue
					}
					if len(ssv) > 0 {
						subSectionValues = append(subSectionValues, ssv...)
					}
				}
			}
		}
		e.Values = subSectionValues
	}
	library.addEntity(s, e)
	e.Name = CanonicalName(e.Name)
	return
}

func (library *Library) findEnumValues(reader asciidoc.Reader, doc *asciidoc.Document, s *asciidoc.Section, e *matter.Enum) (matter.EnumValueSet, error) {
	var tables []*asciidoc.Table
	parse.SkimFunc(reader, s, reader.Children(s), func(t *asciidoc.Table) bool {
		tables = append(tables, t)
		return false
	})
	if len(tables) == 0 {
		return nil, newGenericParseError(e, "no enum field tables found")
	}
	for _, t := range tables {
		ti, err := parseTable(reader, doc, s, t)
		if err != nil {
			return nil, err
		}
		var values matter.EnumValueSet
		for row := range ti.ContentRows() {
			ev := matter.NewEnumValue(row, e)
			ev.Name, err = ti.ReadValue(reader, row, matter.TableColumnName)
			if err != nil {
				return nil, err
			}
			ev.Name = matter.StripTypeSuffixes(ev.Name)
			if len(ev.Name) == 0 {
				ev.Name, err = ti.ReadValue(reader, row, matter.TableColumnSummary)
				if err != nil {
					return nil, err
				}
				ev.Name = matter.StripTypeSuffixes(ev.Name)
				if len(ev.Name) == 0 {
					slog.Debug("skipping enum with no name", slog.String("path", doc.Path.String()), slog.String("section", library.SectionName(s)))
					continue
				}
			}
			ev.Name = CanonicalName(ev.Name)
			ev.Summary, err = ti.ReadValue(reader, row, matter.TableColumnSummary, matter.TableColumnDescription)
			if err != nil {
				return nil, err
			}
			ev.Conformance = ti.ReadConformance(reader, row, matter.TableColumnConformance)
			if ev.Conformance == nil {

				ev.Conformance = conformance.Set{&conformance.Mandatory{}}
			}
			ev.Value, err = ti.ReadID(reader, row, matter.TableColumnValue)
			if err != nil {
				return nil, err
			}
			if !ev.Value.Valid() {
				ev.Value, err = ti.ReadID(reader, row, matter.TableColumnStatusCode)
				if err != nil {
					return nil, err
				}
			}
			values = append(values, ev)
		}
		validValues := make(matter.EnumValueSet, 0, len(values))
		for _, v := range values {
			if v.Value.Valid() {
				validValues = append(validValues, v)
			}
		}
		if len(validValues) > 0 {
			return validValues, nil
		}
	}
	return nil, nil
}

func (library *Library) toModeTags(reader asciidoc.Reader, d *asciidoc.Document, s *asciidoc.Section, parent types.Entity) (e *matter.Enum, err error) {
	var ti *TableInfo
	ti, err = parseFirstTable(reader, d, s)
	if err != nil {
		return nil, newGenericParseError(s, "failed reading mode tags: %w", err)
	}
	e = matter.NewEnum(s, parent)
	e.Name = "ModeTag"
	e.Type = types.NewDataType(types.BaseDataTypeEnum16, false)
	for row := range ti.ContentRows() {
		ev := matter.NewEnumValue(s, e)
		ev.Name, err = ti.ReadString(reader, row, matter.TableColumnName)
		if err != nil {
			return
		}
		ev.Value, err = ti.ReadID(reader, row, matter.TableColumnModeTagValue)
		if err != nil {
			return
		}
		e.Values = append(e.Values, ev)
	}
	return
}

type enumFinder struct {
	entityFinderCommon

	en *matter.Enum
}

func newEnumFinder(en *matter.Enum, inner entityFinder) *enumFinder {
	return &enumFinder{entityFinderCommon: entityFinderCommon{inner: inner}, en: en}
}

func (ef *enumFinder) findEntityByIdentifier(identifier string, source log.Source) types.Entity {
	for _, ev := range ef.en.Values {
		if ev.Name == identifier && ev != ef.identity {
			return ev
		}
	}
	if ef.inner != nil {
		return ef.inner.findEntityByIdentifier(identifier, source)
	}
	return nil
}

func (ef *enumFinder) suggestIdentifiers(identifier string, suggestions map[types.Entity]int) {
	suggest.PossibleEntities(identifier, suggestions, func(yield func(string, types.Entity) bool) {
		for _, ev := range ef.en.Values {

			if ev == ef.identity {
				continue
			}
			if !yield(ev.Name, ev) {
				return
			}

		}
	})
	if ef.inner != nil {
		ef.inner.suggestIdentifiers(identifier, suggestions)
	}
}

func validateEnums(spec *Specification) {
	for c := range spec.Clusters {
		for _, e := range c.Enums {
			validateEnum(spec, e)
		}
	}
	for _, en := range types.FilterSet[*matter.Enum](spec.GlobalObjects) {
		validateEnum(spec, en)
	}
}

func validateEnum(spec *Specification, en *matter.Enum) {
	idu := make(idUniqueness[*matter.EnumValue])
	nu := make(nameUniqueness[*matter.EnumValue])
	cv := make(conformanceValidation)
	for _, ev := range en.Values {
		if !ev.Value.Valid() {
			slog.Warn("Enum value has invalid ID", log.Path("source", ev), matter.LogEntity("parent", en), slog.String("valueName", en.Name))
			continue
		}
		cv.add(ev, ev.Conformance)
		idu.check(spec, ev.Value, ev)
		nu.check(spec, ev)
	}
	cv.check(spec)
}
