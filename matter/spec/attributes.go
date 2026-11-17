package spec

import (
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func (library *Library) toAttributes(spec *Specification, reader asciidoc.Reader, d *asciidoc.Document, section *asciidoc.Section, cluster *matter.Cluster) (attributes matter.FieldSet, err error) {
	var ti *TableInfo
	ti, err = parseFirstTable(reader, d, section)
	if err != nil {
		if err == ErrNoTableFound {
			if !isBaseOrDerivedCluster(cluster) {
				slog.Warn("no table found for attributes", slog.String("name", cluster.Name), log.Element("source", d.Path, section))
			}
			err = nil
		}
		return
	}
	var attributeMap map[string]*matter.Field
	attributes, attributeMap, err = library.readFields(spec, reader, ti, types.EntityTypeAttribute, cluster)
	if err != nil {
		return
	}
	err = library.mapFields(reader, d, section, attributeMap)
	return
}
