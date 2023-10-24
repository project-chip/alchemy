package ascii

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/matter"
	"github.com/hasty/matterfmt/parse"
)

func (s *Section) toFeatures() (features []*matter.Feature, err error) {
	var rows []*types.TableRow
	var headerRowIndex int
	var columnMap map[matter.TableColumn]int
	rows, headerRowIndex, columnMap, _, err = parseFirstTable(s)
	if err != nil {
		return
	}
	for i := headerRowIndex + 1; i < len(rows); i++ {
		row := rows[i]
		f := &matter.Feature{}
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
		var b string
		b, err = readRowValue(row, columnMap, matter.TableColumnBit)
		if err != nil {
			return
		}
		var bv uint64
		bv, err = parse.ID(b)
		if err != nil {
			return
		}
		f.Bit = int(bv)
		features = append(features, f)
	}
	return
}
