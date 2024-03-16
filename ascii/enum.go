package ascii

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/internal/parse"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
	mattertypes "github.com/hasty/alchemy/matter/types"
)

func (s *Section) toEnum(d *Doc, entityMap map[types.WithAttributes][]mattertypes.Entity) (e *matter.Enum, err error) {

	name := strings.TrimSuffix(s.Name, " Type")
	e = &matter.Enum{
		Name: name,
	}
	dt := s.GetDataType()
	if dt == nil {
		dt = mattertypes.ParseDataType("enum8", false)
	}

	if !dt.IsEnum() {
		return nil, fmt.Errorf("unknown enum data type: %s", dt.Name)
	}

	e.Type = dt

	e.Values, err = s.findEnumValues()
	if err != nil {
		return
	}

	if len(e.Values) > 0 {
		return
	}

	var subSectionValues matter.EnumValueSet
	for _, el := range s.Elements {
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
	return
}

func (s *Section) findEnumValues() (matter.EnumValueSet, error) {
	var tables []*types.Table
	parse.SkimFunc(s.Elements, func(t *types.Table) bool {
		tables = append(tables, t)
		return false
	})
	if len(tables) == 0 {
		return nil, fmt.Errorf("no enum field tables found")
	}
	for _, t := range tables {
		var rows []*types.TableRow
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
			ev.Name, err = readRowValue(s.Doc, row, columnMap, matter.TableColumnName)
			if err != nil {
				return nil, err
			}
			ev.Name = StripTypeSuffixes(ev.Name)
			ev.Summary, err = readRowValue(s.Doc, row, columnMap, matter.TableColumnSummary, matter.TableColumnDescription)
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
			values = append(values, ev)
		}
		valid := true
		for _, v := range values {
			if !v.Value.Valid() {
				valid = false
				break
			}
		}
		if valid {
			return values, nil
		}
	}
	return nil, nil
}

func (s *Section) toModeTags(d *Doc) (e *matter.Enum, err error) {
	var rows []*types.TableRow
	var headerRowIndex int
	var columnMap ColumnIndex
	rows, headerRowIndex, columnMap, _, err = parseFirstTable(d, s)
	if err != nil {
		return nil, fmt.Errorf("failed reading mode tags: %w", err)
	}
	e = &matter.Enum{
		Name: "ModeTag",
		Type: mattertypes.ParseDataType("enum16", false),
	}

	e.Type = mattertypes.ParseDataType("enum16", false)

	for i := headerRowIndex + 1; i < len(rows); i++ {
		row := rows[i]
		ev := &matter.EnumValue{}
		ev.Name, err = readRowAsciiDocString(row, columnMap, matter.TableColumnName)
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
