package spec

import (
	"log/slog"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/constraint"
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

type referenceFailure struct {
	source    types.Entity
	reference constraint.Limit
	finder    entityFinder
}
