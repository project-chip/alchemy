package ascii

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
	mattertypes "github.com/hasty/alchemy/matter/types"
)

func (s *Section) toFeatures(d *Doc) (features *matter.Bitmap, err error) {
	var rows []*types.TableRow
	var headerRowIndex int
	var columnMap ColumnIndex
	rows, headerRowIndex, columnMap, _, err = parseFirstTable(d, s)
	if err != nil {
		return nil, fmt.Errorf("failed reading features: %w", err)

	}
	features = &matter.Bitmap{
		Name: "Feature",
		Type: mattertypes.NewDataType("map32", false),
	}
	for i := headerRowIndex + 1; i < len(rows); i++ {
		row := rows[i]
		f := &matter.BitmapBit{}
		f.Bit, err = readRowValue(row, columnMap, matter.TableColumnBit)
		if err != nil {
			return
		}
		if len(f.Bit) == 0 {
			f.Bit, err = readRowValue(row, columnMap, matter.TableColumnID)
			if err != nil {
				return
			}
		}

		f.Name, err = readRowValue(row, columnMap, matter.TableColumnFeature)
		if err != nil {
			return
		}
		if len(f.Name) == 0 {
			f.Name, err = readRowValue(row, columnMap, matter.TableColumnName)
			if err != nil {
				return
			}
		}
		f.Code, err = readRowValue(row, columnMap, matter.TableColumnCode)
		if err != nil {
			return
		}
		f.Summary, err = readRowValue(row, columnMap, matter.TableColumnSummary, matter.TableColumnDescription)
		if err != nil {
			return
		}
		f.Conformance = d.getRowConformance(row, columnMap, matter.TableColumnConformance)
		if f.Conformance == nil {
			f.Conformance = conformance.Set{&conformance.Optional{}}
		}
		features.Bits = append(features.Bits, f)
	}
	return
}
