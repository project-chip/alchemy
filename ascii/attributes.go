package ascii

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/matter"
)

func (s *Section) toAttributes(d *Doc) (attributes []*matter.Field, err error) {
	var rows []*types.TableRow
	var headerRowIndex int
	var columnMap map[matter.TableColumn]int
	rows, headerRowIndex, columnMap, _, err = parseFirstTable(s)
	if err != nil {
		if err == NoTableFound {
			err = nil
		}
		return
	}
	for i := headerRowIndex + 1; i < len(rows); i++ {
		row := rows[i]
		attr := &matter.Field{}
		attr.ID, err = readRowID(row, columnMap, matter.TableColumnID)
		if err != nil {
			return
		}
		attr.Name, err = readRowValue(row, columnMap, matter.TableColumnName)
		if err != nil {
			return
		}
		attr.Type = d.getRowDataType(row, columnMap, matter.TableColumnType)
		attr.Constraint = d.getRowConstraint(row, columnMap, matter.TableColumnConstraint, attr.Type)
		if err != nil {
			return
		}
		var q string
		q, err = readRowValue(row, columnMap, matter.TableColumnQuality)
		if err != nil {
			return
		}
		attr.Quality = matter.ParseQuality(q)
		attr.Default, err = readRowValue(row, columnMap, matter.TableColumnDefault)
		if err != nil {
			return
		}
		attr.Conformance, err = readRowValue(row, columnMap, matter.TableColumnConformance)
		if err != nil {
			return
		}
		var a string
		a, err = readRowValue(row, columnMap, matter.TableColumnAccess)
		if err != nil {
			return
		}
		attr.Access = ParseAccess(a)
		attributes = append(attributes, attr)
	}
	return
}
