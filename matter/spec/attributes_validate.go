package spec

import (
	"log/slog"

	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func validateAttributes(spec *Specification, c *matter.Cluster) {
	validateFields(spec, c, c, c.Attributes, types.EntityTypeAttribute)
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
