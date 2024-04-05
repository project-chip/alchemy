package ascii

import (
	"log/slog"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/internal/parse"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
	mattertypes "github.com/hasty/alchemy/matter/types"
)

func (s *Section) toAttributes(d *Doc, cluster *matter.Cluster, entityMap map[types.WithAttributes][]mattertypes.Entity) (attributes matter.FieldSet, err error) {
	var rows []*types.TableRow
	var headerRowIndex int
	var columnMap ColumnIndex
	rows, headerRowIndex, columnMap, _, err = parseFirstTable(d, s)
	if err != nil {
		if err == ErrNoTableFound {
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
		attr.Name, err = ReadRowValue(d, row, columnMap, matter.TableColumnName)
		if err != nil {
			return
		}
		attr.Name = matter.StripTypeSuffixes(attr.Name)
		attr.Conformance = d.getRowConformance(row, columnMap, matter.TableColumnConformance)
		attr.Type, err = d.ReadRowDataType(row, columnMap, matter.TableColumnType)
		if err != nil {
			if cluster.Hierarchy == "Base" && !conformance.IsDeprecated(attr.Conformance) && !conformance.IsDisallowed(attr.Conformance) {
				// Clusters inheriting from other clusters don't supply type information, nor do attributes that are deprecated or disallowed
				slog.Warn("error reading attribute data type", slog.String("path", s.Doc.Path), slog.String("name", attr.Name), slog.Any("error", err))
			}
			err = nil
		}
		attr.Constraint = d.getRowConstraint(row, columnMap, matter.TableColumnConstraint, attr.Type)
		if err != nil {
			return
		}
		var q string
		q, err = readRowASCIIDocString(row, columnMap, matter.TableColumnQuality)
		if err != nil {
			return
		}
		attr.Quality = matter.ParseQuality(q)
		attr.Default, err = readRowASCIIDocString(row, columnMap, matter.TableColumnDefault)
		if err != nil {
			return
		}
		var a string
		a, err = readRowASCIIDocString(row, columnMap, matter.TableColumnAccess)
		if err != nil {
			return
		}
		attr.Access, _ = ParseAccess(a, mattertypes.EntityTypeAttribute)
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
