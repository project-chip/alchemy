package spec

import (
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func (library *Library) toAttributes(spec *Specification, reader asciidoc.Reader, d *asciidoc.Document, section *asciidoc.Section, cluster *matter.Cluster) (attributes matter.FieldSet, err error) {
	var ti *TableInfo
	ti, err = parseFirstTable(reader, d, section)
	if err != nil {
		if err == ErrNoTableFound {
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

func validateAttributes(spec *Specification) {
	for c := range spec.Clusters {
		validateFields(spec, c, c.Attributes)
		for _, a := range c.Attributes {
			if a.Type == nil {
				continue
			}
			switch et := a.Type.Entity.(type) {
			case *matter.Struct:
				if et.FabricScoping == matter.FabricScopingScoped {
					slog.Error("Fabric-scoped structs may not be used as a singular attribute type, only as a list", matter.LogEntity("attribute", a))
					spec.addError(&FabricScopedStructNotAllowedError{Entity: et})
				}

			}
		}
	}
}
