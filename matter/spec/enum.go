package spec

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

func (s *Section) toEnum(d *Doc, entityMap map[asciidoc.Attributable][]types.Entity) (e *matter.Enum, err error) {

	name := specName(text.TrimCaseInsensitiveSuffix(s.Name, " Type"))
	e = &matter.Enum{
		Name: name,
	}
	dt := s.GetDataType()
	if dt == nil {
		dt = types.ParseDataType("enum8", false)
	}

	if !dt.IsEnum() {
		return nil, fmt.Errorf("unknown enum data type: %s", dt.Name)
	}

	e.Type = dt

	e.Values, err = s.findEnumValues()
	if err != nil {
		slog.Warn("error finding enum values", log.Element("path", d.Path, s.Base), slog.Any("err", err))
		return
	}

	if len(e.Values) > 0 {
		return
	}

	var subSectionValues matter.EnumValueSet
	for _, el := range s.Elements() {
		switch el := el.(type) {
		case *Section:
			if strings.HasSuffix(el.Name, " Range") {
				var ssv matter.EnumValueSet
				ssv, err = el.findEnumValues()
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
	entityMap[s.Base] = append(entityMap[s.Base], e)
	return
}

func (s *Section) findEnumValues() (matter.EnumValueSet, error) {
	var tables []*asciidoc.Table
	parse.SkimFunc(s.Elements(), func(t *asciidoc.Table) bool {
		tables = append(tables, t)
		return false
	})
	if len(tables) == 0 {
		return nil, fmt.Errorf("no enum field tables found")
	}
	for _, t := range tables {
		var rows []*asciidoc.TableRow
		var headerRowIndex int
		var columnMap ColumnIndex
		var err error
		rows, headerRowIndex, columnMap, _, err = parseTable(s.Doc, s, t)
		if err != nil {
			return nil, err
		}
		var values matter.EnumValueSet
		for i := headerRowIndex + 1; i < len(rows); i++ {
			row := rows[i]
			ev := &matter.EnumValue{}
			ev.Name, err = ReadRowValue(s.Doc, row, columnMap, matter.TableColumnName)
			if err != nil {
				return nil, err
			}
			ev.Name = matter.StripTypeSuffixes(ev.Name)
			if len(ev.Name) == 0 {
				ev.Name, err = ReadRowValue(s.Doc, row, columnMap, matter.TableColumnSummary)
				if err != nil {
					return nil, err
				}
				ev.Name = matter.StripTypeSuffixes(ev.Name)
				if len(ev.Name) == 0 {
					slog.Debug("skipping enum with no name", slog.String("path", s.Doc.Path), slog.String("section", s.Name))
					continue
				}
			}
			ev.Summary, err = ReadRowValue(s.Doc, row, columnMap, matter.TableColumnSummary, matter.TableColumnDescription)
			if err != nil {
				return nil, err
			}
			ev.Conformance = s.Doc.getRowConformance(row, columnMap, matter.TableColumnConformance)
			if ev.Conformance == nil {

				ev.Conformance = conformance.Set{&conformance.Mandatory{}}
			}
			ev.Value, err = readRowID(row, columnMap, matter.TableColumnValue)
			if err != nil {
				return nil, err
			}
			if !ev.Value.Valid() {
				ev.Value, err = readRowID(row, columnMap, matter.TableColumnStatusCode)
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

func (s *Section) toModeTags(d *Doc) (e *matter.Enum, err error) {
	var rows []*asciidoc.TableRow
	var headerRowIndex int
	var columnMap ColumnIndex
	rows, headerRowIndex, columnMap, _, err = parseFirstTable(d, s)
	if err != nil {
		return nil, fmt.Errorf("failed reading mode tags: %w", err)
	}
	e = &matter.Enum{
		Name: "ModeTag",
		Type: types.ParseDataType("enum16", false),
	}

	e.Type = types.ParseDataType("enum16", false)

	for i := headerRowIndex + 1; i < len(rows); i++ {
		row := rows[i]
		ev := &matter.EnumValue{}
		ev.Name, err = readRowASCIIDocString(row, columnMap, matter.TableColumnName)
		if err != nil {
			return
		}
		ev.Value, err = readRowID(row, columnMap, matter.TableColumnModeTagValue)
		if err != nil {
			return
		}
		e.Values = append(e.Values, ev)
	}
	return
}
