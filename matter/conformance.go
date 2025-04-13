package matter

import (
	"log/slog"

	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

func EntityConformance(entity types.Entity) conformance.Set {
	switch entity := entity.(type) {
	case Bit:
		return entity.Conformance()
	case *EnumValue:
		return entity.Conformance
	case *Field:
		return entity.Conformance
	case *Command:
		return entity.Conformance
	case *Event:
		return entity.Conformance
	case *Cluster:
		return entity.Conformance
	case *DeviceTypeRequirement:
		return entity.Conformance
	case *ClusterRequirement:
		return entity.Conformance
	case *ElementRequirement:
		return entity.Conformance
	case *Struct, *Enum, *Bitmap:
		return nil
	case nil:
		slog.Warn("Unexpected nil entity fetching conformance")
	default:
		slog.Warn("Unexpected entity fetching conformance", LogEntity("entity", entity))
	}
	return nil
}
