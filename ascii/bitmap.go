package ascii

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
	mattertypes "github.com/hasty/alchemy/matter/types"
	"github.com/iancoleman/strcase"
)

func (s *Section) toBitmap(d *Doc) (e *matter.Bitmap, err error) {
	name := strings.TrimSuffix(s.Name, " Type")

	dt := s.GetDataType()
	if dt == nil {
		dt = mattertypes.NewDataType("map8", false)
	}

	if !dt.IsMap() {
		return nil, fmt.Errorf("unknown bitmap data type: %s", dt.Name)
	}

	e = &matter.Bitmap{
		Name: name,
		Type: dt,
	}
	var rows []*types.TableRow
	var headerRowIndex int
	var columnMap ColumnIndex
	rows, headerRowIndex, columnMap, _, err = parseFirstTable(d, s)

	if err != nil {
		if err == NoTableFound {
			slog.Warn("no table found for bitmap", slog.String("name", e.Name))
			return e, nil
		}
		return nil, fmt.Errorf("failed reading bitmap %s: %w", name, err)
	}

	for i := headerRowIndex + 1; i < len(rows); i++ {
		row := rows[i]
		bv := &matter.BitmapBit{}
		bv.Name, err = readRowValue(row, columnMap, matter.TableColumnName)
		if err != nil {
			return
		}
		bv.Summary, err = readRowValue(row, columnMap, matter.TableColumnSummary, matter.TableColumnDescription)
		if err != nil {
			return
		}
		bv.Conformance = d.getRowConformance(row, columnMap, matter.TableColumnConformance)
		if bv.Conformance == nil {
			bv.Conformance = conformance.Set{&conformance.Mandatory{}}
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
