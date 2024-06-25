package spec

import (
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

func (s *Section) toAttributes(d *Doc, cluster *matter.Cluster, entityMap map[asciidoc.Attributable][]types.Entity) (attributes matter.FieldSet, err error) {
	var rows []*asciidoc.TableRow
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
				slog.Warn("error reading attribute data type", log.Element("path", d.Path, row), slog.String("name", attr.Name), slog.Any("error", err))
			}
			err = nil
		}
		attr.Constraint = d.getRowConstraint(row, columnMap, matter.TableColumnConstraint)
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
		attr.Access, _ = ParseAccess(a, types.EntityTypeAttribute)
		attributes = append(attributes, attr)
		attributeMap[attr.Name] = attr
	}
	err = s.mapFields(attributeMap, entityMap)
	return
}
