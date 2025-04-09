package spec

import (
	"fmt"
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

type parseContext struct {
	entities          []types.Entity
	orderedEntities   []types.Entity
	globalObjects     []types.Entity
	entitiesByElement map[asciidoc.Attributable][]types.Entity
}

func (pc *parseContext) addRootEntity(e types.Entity, source asciidoc.Attributable) {
	pc.entities = append(pc.entities, e)
	pc.addEntity(e, source)
}

func (pc *parseContext) addEntity(e types.Entity, source asciidoc.Attributable) {
	pc.orderedEntities = append(pc.orderedEntities, e)
	pc.entitiesByElement[source] = append(pc.entitiesByElement[source], e)
}

func (doc *Doc) parseEntities() error {
	pc := &parseContext{
		entitiesByElement: make(map[asciidoc.Attributable][]types.Entity),
	}
	for _, top := range parse.Skim[*Section](doc.Elements()) {
		err := AssignSectionTypes(doc, top)
		if err != nil {
			return err
		}

		err = top.toEntities(doc, pc)
		if err != nil {
			return fmt.Errorf("failed converting doc %s to entities: %w", doc.Path, err)
		}

	}
	doc.entities = pc.entities
	doc.orderedEntities = pc.orderedEntities
	doc.entitiesBySection = pc.entitiesByElement
	doc.globalObjects = pc.globalObjects
	doc.entitiesParsed = true
	return nil
}

type entityFinder interface {
	setIdentity(entity types.Entity)
	findEntityByIdentifier(identifier string, source log.Source) types.Entity
	findEntityByReference(reference string, label string, source log.Source) types.Entity
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

type entityFilter struct {
	entityFinderCommon

	filterTypes []types.EntityType
}

func (ef *entityFilter) findEntityByIdentifier(identifier string, source log.Source) types.Entity {
	if ef.inner == nil {
		return nil
	}
	e := ef.inner.findEntityByIdentifier(identifier, source)
	if e == nil {
		return nil
	}
	for _, t := range ef.filterTypes {
		if t == e.EntityType() {
			return e
		}
	}
	return nil
}
