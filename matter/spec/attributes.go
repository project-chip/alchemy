package spec

import (
	"log/slog"

	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func (s *Section) toAttributes(spec *Specification, d *Doc, cluster *matter.Cluster, pc *parseContext) (attributes matter.FieldSet, err error) {
	var ti *TableInfo
	ti, err = parseFirstTable(d, s)
	if err != nil {
		if err == ErrNoTableFound {
			err = nil
		}
		return
	}
	var attributeMap map[string]*matter.Field
	attributes, attributeMap, err = d.readFields(spec, ti, types.EntityTypeAttribute, cluster)
	if err != nil {
		return
	}
	err = s.mapFields(attributeMap, pc)
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
