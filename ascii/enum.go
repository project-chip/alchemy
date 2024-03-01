package ascii

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
	mattertypes "github.com/hasty/alchemy/matter/types"
)

func (s *Section) toEnum(d *Doc, entityMap map[types.WithAttributes][]mattertypes.Entity) (e *matter.Enum, err error) {
	var rows []*types.TableRow
	var headerRowIndex int
	var columnMap ColumnIndex
	rows, headerRowIndex, columnMap, _, err = parseFirstTable(d, s)
	if err != nil {
		return nil, fmt.Errorf("failed reading enum: %w", err)
	}
	name := strings.TrimSuffix(s.Name, " Type")
	e = &matter.Enum{
		Name: name,
	}
	dt := s.GetDataType()
	if dt == nil {
		dt = mattertypes.NewDataType("enum8", false)
	}

	if !dt.IsEnum() {
		return nil, fmt.Errorf("unknown enum data type: %s", dt.Name)
	}

	e.Type = dt

	for i := headerRowIndex + 1; i < len(rows); i++ {
		row := rows[i]
		ev := &matter.EnumValue{}
		ev.Name, err = readRowValue(d, row, columnMap, matter.TableColumnName)
		if err != nil {
			return
		}
		ev.Name = StripTypeSuffixes(ev.Name)
		ev.Summary, err = readRowValue(d, row, columnMap, matter.TableColumnSummary, matter.TableColumnDescription)
		if err != nil {
			return
		}
		ev.Conformance = d.getRowConformance(row, columnMap, matter.TableColumnConformance)
		if ev.Conformance == nil {
			// Missing conformance should be treated as mandatory
			ev.Conformance = conformance.Set{&conformance.Mandatory{}}
		}
		ev.Value, err = readRowID(row, columnMap, matter.TableColumnValue)
		if err != nil {
			return
		}
		e.Values = append(e.Values, ev)
	}
	entityMap[s.Base] = append(entityMap[s.Base], e)
	return
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
		Type: mattertypes.NewDataType("enum16", false),
	}

	e.Type = mattertypes.NewDataType("enum16", false)

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
