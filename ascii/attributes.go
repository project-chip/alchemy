package ascii

import (
	"log/slog"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/matter"
	mattertypes "github.com/hasty/alchemy/matter/types"
	"github.com/hasty/alchemy/parse"
)

func (s *Section) toAttributes(d *Doc, entityMap map[types.WithAttributes][]mattertypes.Entity) (attributes matter.FieldSet, err error) {
	var rows []*types.TableRow
	var headerRowIndex int
	var columnMap ColumnIndex
	rows, headerRowIndex, columnMap, _, err = parseFirstTable(d, s)
	if err != nil {
		if err == NoTableFound {
			err = nil
		}
		return
	}
	attributeMap := make(map[string]*matter.Field)
	for i := headerRowIndex + 1; i < len(rows); i++ {
		row := rows[i]
		attr := matter.NewAttribute()
		attr.ID, err = readRowID(row, columnMap, matter.TableColumnID)
		if err != nil {
			return
		}
		attr.Name, err = readRowValue(d, row, columnMap, matter.TableColumnName)
		if err != nil {
			return
		}
		attr.Name = StripTypeSuffixes(attr.Name)
		attr.Type = d.ReadRowDataType(row, columnMap, matter.TableColumnType)
		attr.Constraint = d.getRowConstraint(row, columnMap, matter.TableColumnConstraint, attr.Type)
		if err != nil {
			return
		}
		var q string
		q, err = readRowAsciiDocString(row, columnMap, matter.TableColumnQuality)
		if err != nil {
			return
		}
		attr.Quality = matter.ParseQuality(q)
		attr.Default, err = readRowAsciiDocString(row, columnMap, matter.TableColumnDefault)
		if err != nil {
			return
		}
		attr.Conformance = d.getRowConformance(row, columnMap, matter.TableColumnConformance)
		var a string
		a, err = readRowAsciiDocString(row, columnMap, matter.TableColumnAccess)
		if err != nil {
			return
		}
		attr.Access = ParseAccess(a, mattertypes.EntityTypeAttribute)
		attributes = append(attributes, attr)
		attributeMap[attr.Name] = attr
	}
	for _, s := range parse.Skim[*Section](s.Elements) {
		switch s.SecType {
		case matter.SectionAttribute:

			name := strings.TrimSuffix(s.Name, " Attribute")
			a, ok := attributeMap[name]
			if !ok {
				slog.Debug("unknown attribute", "attribute", name)
				continue
			}

			entityMap[s.Base] = append(entityMap[s.Base], a)
		}
	}
	return
}
