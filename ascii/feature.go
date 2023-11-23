package ascii

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/matter"
)

func (s *Section) toFeatures(d *Doc) (features []*matter.Feature, err error) {
	var rows []*types.TableRow
	var headerRowIndex int
	var columnMap ColumnIndex
	rows, headerRowIndex, columnMap, _, err = parseFirstTable(s)
	if err != nil {
		return nil, fmt.Errorf("failed reading features: %w", err)

	}
	for i := headerRowIndex + 1; i < len(rows); i++ {
		row := rows[i]
		f := &matter.Feature{}
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
		f.Code, err = readRowValue(row, columnMap, matter.TableColumnCode)
		if err != nil {
			return
		}
		f.Summary, err = readRowValue(row, columnMap, matter.TableColumnSummary)
		if err != nil {
			return
		}
		f.Conformance = d.getRowConformance(row, columnMap, matter.TableColumnConformance)
		features = append(features, f)
	}
	return
}
