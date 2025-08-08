package spec

import (
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/internal/suggest"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

func (s *Section) toEnum(d *Doc, pc *parseContext, parent types.Entity) (e *matter.Enum, err error) {

	name := CanonicalName(text.TrimCaseInsensitiveSuffix(s.Name, " Type"))
	e = matter.NewEnum(s.Base, parent)
	e.Name = name
	dt := s.GetDataType()
	if dt == nil {
		dt = types.NewDataType(types.BaseDataTypeEnum8, false)
		slog.Warn("Enum does not declare its derived data type; assuming enum8", log.Element("source", d.Path, s.Base), slog.String("enum", name))
	} else if !dt.IsEnum() {
		return nil, newGenericParseError(s.Base, "unknown enum data type: \"%s\"", dt.Name)
	}

	e.Type = dt

	e.Values, err = s.findEnumValues(e)
	if err != nil {
		return
	}

	if len(e.Values) == 0 {
		var subSectionValues matter.EnumValueSet
		for _, el := range s.Children() {
			switch el := el.(type) {
			case *Section:
				if strings.HasSuffix(el.Name, " Range") {
					var ssv matter.EnumValueSet
					ssv, err = el.findEnumValues(e)
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
	pc.orderedEntities = append(pc.orderedEntities, e)
	pc.entitiesByElement[s.Base] = append(pc.entitiesByElement[s.Base], e)
	e.Name = CanonicalName(e.Name)
	return
}

func (s *Section) findEnumValues(e *matter.Enum) (matter.EnumValueSet, error) {
	var tables []*asciidoc.Table
	parse.SkimFunc(s.Children(), func(t *asciidoc.Table) bool {
		tables = append(tables, t)
		return false
	})
	if len(tables) == 0 {
		return nil, newGenericParseError(e, "no enum field tables found")
	}
	for _, t := range tables {
		ti, err := parseTable(s.Doc, s, t)
		if err != nil {
			return nil, err
		}
		var values matter.EnumValueSet
		for row := range ti.ContentRows() {
			ev := matter.NewEnumValue(row, e)
			ev.Name, err = ti.ReadValue(row, matter.TableColumnName)
			if err != nil {
				return nil, err
			}
			ev.Name = matter.StripTypeSuffixes(ev.Name)
			if len(ev.Name) == 0 {
				ev.Name, err = ti.ReadValue(row, matter.TableColumnSummary)
				if err != nil {
					return nil, err
				}
				ev.Name = matter.StripTypeSuffixes(ev.Name)
				if len(ev.Name) == 0 {
					slog.Debug("skipping enum with no name", slog.String("path", s.Doc.Path.String()), slog.String("section", s.Name))
					continue
				}
			}
			ev.Name = CanonicalName(ev.Name)
			ev.Summary, err = ti.ReadValue(row, matter.TableColumnSummary, matter.TableColumnDescription)
			if err != nil {
				return nil, err
			}
			ev.Conformance = ti.ReadConformance(row, matter.TableColumnConformance)
			if ev.Conformance == nil {

				ev.Conformance = conformance.Set{&conformance.Mandatory{}}
			}
			ev.Value, err = ti.ReadID(row, matter.TableColumnValue)
			if err != nil {
				return nil, err
			}
			if !ev.Value.Valid() {
				ev.Value, err = ti.ReadID(row, matter.TableColumnStatusCode)
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

func (s *Section) toModeTags(d *Doc, parent types.Entity) (e *matter.Enum, err error) {
	var ti *TableInfo
	ti, err = parseFirstTable(d, s)
	if err != nil {
		return nil, newGenericParseError(s.Base, "failed reading mode tags: %w", err)
	}
	e = matter.NewEnum(s.Base, parent)
	e.Name = "ModeTag"
	e.Type = types.NewDataType(types.BaseDataTypeEnum16, false)
	for row := range ti.ContentRows() {
		ev := matter.NewEnumValue(s.Base, e)
		ev.Name, err = ti.ReadString(row, matter.TableColumnName)
		if err != nil {
			return
		}
		ev.Value, err = ti.ReadID(row, matter.TableColumnModeTagValue)
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
	enumValues := make(map[uint64]*matter.EnumValue)
	for _, ev := range en.Values {
		if !ev.Value.Valid() {
			slog.Warn("Enum value has invalid ID", log.Path("source", ev), matter.LogEntity("parent", en), slog.String("valueName", en.Name))
			continue
		}
		valueId := ev.Value.Value()
		existing, ok := enumValues[valueId]
		if ok {
			slog.Error("Duplicate enum value", log.Path("source", ev), matter.LogEntity("parent", en), slog.String("enumValue", ev.Value.HexString()), slog.String("enumValueName", ev.Name), slog.String("previousEnumValueName", existing.Name))
			spec.addError(&DuplicateEntityIDError{Entity: ev, Previous: existing})
		} else {
			enumValues[valueId] = ev
		}
	}
}
