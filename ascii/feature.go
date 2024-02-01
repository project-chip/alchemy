package ascii

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
	mattertypes "github.com/hasty/alchemy/matter/types"
)

func (s *Section) toFeatures(d *Doc) (features *matter.Features, err error) {
	var rows []*types.TableRow
	var headerRowIndex int
	var columnMap ColumnIndex
	rows, headerRowIndex, columnMap, _, err = parseFirstTable(d, s)
	if err != nil {
		return nil, fmt.Errorf("failed reading features: %w", err)

	}
	features = &matter.Features{
		Bitmap: matter.Bitmap{
			Name: "Feature",
			Type: mattertypes.NewDataType("map32", false),
		},
	}
	for i := headerRowIndex + 1; i < len(rows); i++ {
		row := rows[i]
		var bit, code, name, summary string
		var conf conformance.Set
		bit, err = readRowValue(row, columnMap, matter.TableColumnBit)
		if err != nil {
			return
		}
		if len(bit) == 0 {
			bit, err = readRowValue(row, columnMap, matter.TableColumnID)
			if err != nil {
				return
			}
		}

		name, err = readRowValue(row, columnMap, matter.TableColumnFeature)
		if err != nil {
			return
		}
		if len(name) == 0 {
			name, err = readRowValue(row, columnMap, matter.TableColumnName)
			if err != nil {
				return
			}
		}
		code, err = readRowValue(row, columnMap, matter.TableColumnCode)
		if err != nil {
			return
		}
		summary, err = readRowValue(row, columnMap, matter.TableColumnSummary, matter.TableColumnDescription)
		if err != nil {
			return
		}
		conf = d.getRowConformance(row, columnMap, matter.TableColumnConformance)
		if conf == nil {
			conf = conformance.Set{&conformance.Optional{}}
		}
		features.Bits = append(features.Bits, matter.NewFeature(bit, code, name, summary, conf))
	}
	return
}
