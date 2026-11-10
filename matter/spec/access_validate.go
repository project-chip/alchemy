package spec

import (
	"log/slog"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func validateAccess(spec *Specification, entity types.Entity, access matter.Access) {
	if access.IsFabricScoped() && !matter.IsFabricScopingAllowed(entity) {
		slog.Error("Fabric scoping is not allowed on this entity", matter.LogEntity("entity", entity), log.Path("source", entity))
		spec.addError(&FabricScopingNotAllowedError{Entity: entity})
	}
	if access.IsFabricSensitive() && !matter.IsFabricSensitivityAllowed(entity) {
		slog.Error("Fabric sensitivity is not allowed on this entity", matter.LogEntity("entity", entity), log.Path("source", entity))
		spec.addError(&FabricSensitivityNotAllowedError{Entity: entity})
	}
}
