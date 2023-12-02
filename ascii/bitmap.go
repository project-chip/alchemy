package ascii

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/conformance"
	"github.com/hasty/alchemy/matter"
	"github.com/iancoleman/strcase"
)

func (s *Section) toBitmap(d *Doc) (e *matter.Bitmap, err error) {
	var rows []*types.TableRow
	var headerRowIndex int
	var columnMap ColumnIndex
	rows, headerRowIndex, columnMap, _, err = parseFirstTable(s)
	if err != nil {
		return nil, fmt.Errorf("failed reading bitmap: %w", err)
	}
	name := strings.TrimSuffix(s.Name, " Type")

	e = &matter.Bitmap{
		Name: name,
	}

	dts := s.GetDataType()
	if dts == "" {
		dts = "map8"
	}

	dt := matter.NewDataType(dts, false)
	if !dt.IsMap() {
		return nil, fmt.Errorf("unknown bitmap data type: %s", dts)
	}

	e.Type = dt

	for i := headerRowIndex + 1; i < len(rows); i++ {
		row := rows[i]
		bv := &matter.Bit{}
		bv.Name, err = readRowValue(row, columnMap, matter.TableColumnName)
		if err != nil {
			return
		}
		bv.Summary, err = readRowValue(row, columnMap, matter.TableColumnSummary)
		if err != nil {
			return
		}
		bv.Conformance = d.getRowConformance(row, columnMap, matter.TableColumnConformance)
		if bv.Conformance == nil {
			bv.Conformance = &conformance.MandatoryConformance{}
		}
		bv.Bit, err = readRowValue(row, columnMap, matter.TableColumnBit)
		if err != nil {
			return
		}
		if len(bv.Bit) == 0 {
			bv.Bit, err = readRowValue(row, columnMap, matter.TableColumnValue)
			if err != nil {
				return
			}
		}
		if len(bv.Name) == 0 && len(bv.Summary) > 0 {
			bv.Name = strcase.ToCamel(bv.Summary)
		}
		e.Bits = append(e.Bits, bv)
	}
	return
}
