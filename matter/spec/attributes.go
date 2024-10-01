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
	var ti *TableInfo
	ti, err = parseFirstTable(d, s)
	if err != nil {
		if err == ErrNoTableFound {
			err = nil
		}
		return
	}
	attributeMap := make(map[string]*matter.Field)
	for row := range ti.Body() {
		attr := matter.NewAttribute(row)
		attr.ID, err = ti.ReadID(row, matter.TableColumnID)
		if err != nil {
			return
		}
		attr.Name, err = ti.ReadValue(row, matter.TableColumnName)
		if err != nil {
			return
		}
		attr.Name = matter.StripTypeSuffixes(attr.Name)
		attr.Conformance = ti.ReadConformance(row, matter.TableColumnConformance)
		attr.Type, err = ti.ReadDataType(row, matter.TableColumnType)
		if err != nil {
			if cluster.Hierarchy == "Base" && !conformance.IsDeprecated(attr.Conformance) && !conformance.IsDisallowed(attr.Conformance) {
				// Clusters inheriting from other clusters don't supply type information, nor do attributes that are deprecated or disallowed
				slog.Warn("error reading attribute data type", log.Element("path", d.Path, row), slog.String("name", attr.Name), slog.Any("error", err))
			}
			err = nil
		}
		attr.Constraint = ti.ReadConstraint(row, matter.TableColumnConstraint)
		if err != nil {
			return
		}
		var q string
		q, err = ti.ReadString(row, matter.TableColumnQuality)
		if err != nil {
			return
		}
		attr.Quality = parseQuality(q, types.EntityTypeAttribute, d, row)
		attr.Default, err = ti.ReadString(row, matter.TableColumnDefault)
		if err != nil {
			return
		}
		var a string
		a, err = ti.ReadString(row, matter.TableColumnAccess)
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
