package ascii

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/matter"
)

func (s *Section) toEnum(d *Doc) (e *matter.Enum, err error) {
	var rows []*types.TableRow
	var headerRowIndex int
	var columnMap ColumnIndex
	rows, headerRowIndex, columnMap, _, err = parseFirstTable(s)
	if err != nil {
		return nil, fmt.Errorf("failed reading enum: %w", err)
	}
	name := strings.TrimSuffix(s.Name, " Type")
	e = &matter.Enum{
		Name: name,
	}
	dts := s.GetDataType()
	if dts == "" {
		dts = "enum8"
	}

	dt := matter.NewDataType(dts, false)
	if !dt.IsEnum() {
		return nil, fmt.Errorf("unknown enum data type: %s", dts)
	}

	e.Type = dt

	for i := headerRowIndex + 1; i < len(rows); i++ {
		row := rows[i]
		ev := &matter.EnumValue{}
		ev.Name, err = readRowValue(row, columnMap, matter.TableColumnName)
		if err != nil {
			return
		}
		ev.Summary, err = readRowValue(row, columnMap, matter.TableColumnSummary)
		if err != nil {
			return
		}
		if ev.Summary == "" {
			ev.Summary, err = readRowValue(row, columnMap, matter.TableColumnDescription)
			if err != nil {
				return
			}
		}
		ev.Conformance = d.getRowConformance(row, columnMap, matter.TableColumnConformance)
		ev.Value, err = readRowValue(row, columnMap, matter.TableColumnValue)
		if err != nil {
			return
		}
		e.Values = append(e.Values, ev)
	}
	return
}
