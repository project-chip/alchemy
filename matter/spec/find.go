package spec

import (
	"log/slog"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

type entityFinder interface {
	setIdentity(entity types.Entity)
	findEntityByIdentifier(identifier string, source log.Source) types.Entity
	findEntityByReference(reference string, label string, source log.Source) types.Entity
	suggestIdentifiers(identifier string, suggestions map[types.Entity]int)
}

func makeEntityFinder(entity types.Entity, inner entityFinder) entityFinder {
	switch entity := entity.(type) {
	case *matter.Bitmap:
		return newBitmapFinder(entity, inner)
	case *matter.Enum:
		return newEnumFinder(entity, inner)
	case *matter.Command:
		return newFieldFinder(entity.Fields, inner)
	case *matter.Event:
		return newFieldFinder(entity.Fields, inner)
	case *matter.Struct:
		return newFieldFinder(entity.Fields, inner)
	case *matter.Field:
		if entity.Type.Entity != nil {
			return makeEntityFinder(entity.Type.Entity, inner)
		}
		return inner
	case *matter.Namespace:
		return newTagFinder(entity, inner)
	case *matter.TypeDef: // There are no fields in a typedef
		return inner
	case nil:
		slog.Warn("Unable to make named entry finder for nil type")
		return inner
	default:
		slog.Warn("Unable to make named entry finder for type", log.Type("type", entity), matter.LogEntity("entity", entity))
		return inner
	}
}

type entityFinderCommon struct {
	identity types.Entity
	inner    entityFinder
}

func (bf *entityFinderCommon) setIdentity(entity types.Entity) {
	bf.identity = entity
}

func (bf *entityFinderCommon) findEntityByReference(reference string, label string, source log.Source) types.Entity {
	if bf.inner != nil {
		return bf.inner.findEntityByReference(reference, label, source)
	}
	return nil
}

type findEntityByID[T types.Entity] func(id *matter.Number) (T, bool)
type findEntityByName[T types.Entity] func(name string) (T, bool)
type findEntityByNameAndField[T types.Entity] func(name, field string) (T, bool)

func (s *Specification) findClusterByID(id *matter.Number) (cluster *matter.Cluster, ok bool) {
	cluster, ok = s.ClustersByID[id.Value()]
	return
}

func (s *Specification) findClusterByName(name string) (cluster *matter.Cluster, ok bool) {
	cluster, ok = s.ClustersByName[name]
	return
}

func (s *Specification) findDeviceTypeByID(id *matter.Number) (deviceType *matter.DeviceType, ok bool) {
	deviceType, ok = s.DeviceTypesByID[id.Value()]
	return
}

func (s *Specification) findDeviceTypeByName(name string) (deviceType *matter.DeviceType, ok bool) {
	deviceType, ok = s.DeviceTypesByName[name]
	return
}

func (s *Specification) findNamespaceByID(id *matter.Number) (*matter.Namespace, bool) {
	for _, ns := range s.Namespaces {
		if ns.ID.Equals(id) {
			return ns, true
		}
	}
	return nil, false
}

func (s *Specification) findNamespaceByName(name string) (*matter.Namespace, bool) {
	for _, ns := range s.Namespaces {
		if ns.Name == name {
			return ns, true
		}
	}
	return nil, false
}

func findTagByID(namespace *matter.Namespace) func(id *matter.Number) (*matter.SemanticTag, bool) {
	return func(id *matter.Number) (*matter.SemanticTag, bool) {
		if id.Valid() {
			for _, t := range namespace.SemanticTags {
				if t.ID.Equals(id) {
					return t, true
				}
			}
		}
		return nil, false
	}
}

func findTagByName(namespace *matter.Namespace) func(name string) (*matter.SemanticTag, bool) {
	return func(name string) (*matter.SemanticTag, bool) {
		if name != "" {
			for _, t := range namespace.SemanticTags {
				if t.Name == name {
					return t, true
				}
			}
		}
		return nil, false
	}
}

func findConditionByName(deviceType *matter.DeviceType) func(string) (*matter.Condition, bool) {
	return func(name string) (*matter.Condition, bool) {
		for _, condition := range deviceType.Conditions {
			if condition.Feature == name {
				return condition, true
			}
		}
		return nil, false
	}
}
