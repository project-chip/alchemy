package ascii

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/matter"
)

func (s *Section) toFeatures() (features []*matter.Feature, err error) {
	var rows []*types.TableRow
	var headerRowIndex int
	var columnMap map[matter.TableColumn]int
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

		f.Name, err = readRowValue(row, columnMap, matter.TableColumnFeature)
		if err != nil {
			return
		}
		f.Code, err = readRowValue(row, columnMap, matter.TableColumnCode)
		if err != nil {
			return
		}
		f.Description, err = readRowValue(row, columnMap, matter.TableColumnDescription)
		if err != nil {
			return
		}
		f.Conformance, err = readRowValue(row, columnMap, matter.TableColumnConformance)
		if err != nil {
			return
		}
		features = append(features, f)
	}
	return
}
