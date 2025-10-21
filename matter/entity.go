package matter

import (
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/matter/types"
)

type entity struct {
	parent types.Entity
	source asciidoc.Element
}

func (e entity) Parent() types.Entity {
	return e.parent
}

func (e *entity) SetParent(parent types.Entity) {
	e.parent = parent
}

func (e entity) Source() asciidoc.Element {
	return e.source
}

func (e entity) Origin() (path string, line int) {
	switch s := e.source.(type) {
	case Source:
		return s.Origin()
	default:
		return "", -1
	}
}

func (e entity) Cluster() *Cluster {
	return findCluster(e.parent)
}

func EntityName(e types.Entity) string {
	switch entity := e.(type) {
	case *Cluster:
		return entity.Name
	case *Struct:
		return entity.Name
	case *Field:
		return entity.Name
	case *Event:
		return entity.Name
	case *Command:
		return entity.Name
	case *Bitmap:
		return entity.Name
	case *Enum:
		return entity.Name
	case *TypeDef:
		return entity.Name
	case *Namespace:
		return entity.Name
	case *SemanticTag:
		return entity.Name
	case *Constant:
		return entity.Name
	case *Feature:
		return entity.Name()
	case *Condition:
		return entity.Feature
	case *DeviceType:
		return entity.Name
	case *ClusterGroup:
		return entity.Name
	case Bit:
		return entity.Name()
	case *EnumValue:
		return entity.Name
	default:
		slog.Error("Unknown entity type for name", LogEntity("entity", e))
		return ""
	}
}

func EntityID(e types.Entity) *Number {
	switch entity := e.(type) {
	case *Cluster:
		return entity.ID
	case *Field:
		return entity.ID
	case *Event:
		return entity.ID
	case *Command:
		return entity.ID
	case *Namespace:
		return entity.ID
	case *SemanticTag:
		return entity.ID
	case *Feature:
		return ParseNumber(entity.Bit())
	case *Condition:
		return entity.ID
	case *DeviceType:
		return entity.ID
	case Bit:
		return ParseNumber(entity.Bit())
	default:
		slog.Error("Unknown entity type for id", LogEntity("entity", e))
		return nil
	}
}
