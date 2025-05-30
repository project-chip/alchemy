package spec

import (
	"fmt"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/parse"
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

func (doc *Doc) parseEntities(spec *Specification) error {
	pc := &parseContext{
		entitiesByElement: make(map[asciidoc.Attributable][]types.Entity),
	}
	for top := range parse.Skim[*Section](doc.Elements()) {
		err := AssignSectionTypes(doc, top)
		if err != nil {
			return err
		}

		err = top.toEntities(spec, doc, pc)
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
