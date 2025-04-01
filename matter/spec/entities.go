package spec

import (
	"fmt"
	"log/slog"
	"strings"

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

type findNamedEntity func(name string) types.Entity
type filterNamedEntity func(e types.Entity) bool

func makeEntityFinder(entity types.Entity) findNamedEntity {
	switch entity := entity.(type) {

	case *matter.Bitmap:
		return func(identifier string) types.Entity {
			for _, bmv := range entity.Bits {
				if strings.EqualFold(bmv.Name(), identifier) {
					return bmv
				}
			}
			return nil
		}
	case *matter.Enum:
		return func(identifier string) types.Entity {
			for _, ev := range entity.Values {
				if strings.EqualFold(ev.Name, identifier) {
					return ev
				}
			}
			return nil
		}
	case *matter.Struct:
		return func(identifier string) types.Entity {
			for _, ev := range entity.Fields {
				if strings.EqualFold(ev.Name, identifier) {
					return ev
				}
			}
			return nil
		}
	case *matter.Command:
		return func(identifier string) types.Entity {
			for _, ev := range entity.Fields {
				if strings.EqualFold(ev.Name, identifier) {
					return ev
				}
			}
			return nil
		}
	case *matter.Event:
		return func(identifier string) types.Entity {
			for _, ev := range entity.Fields {
				if strings.EqualFold(ev.Name, identifier) {
					return ev
				}
			}
			return nil
		}
	case *matter.Field:
		return makeEntityFinder(entity.Type.Entity)

	default:
		slog.Warn("Unable to make named entry finder for type", log.Type("type", entity))
		return nil
	}
}

func findEntityForEntityIdentifier(spec *Specification, cluster *matter.Cluster, finder findNamedEntity, source log.Source, entity types.Entity, identifier string) (e types.Entity) {

	if spec.BaseDeviceType != nil {
		for _, c := range spec.BaseDeviceType.Conditions {
			if strings.EqualFold(c.Feature, identifier) {
				e = c
				return
			}
		}
	}
	if finder != nil {
		e = finder(identifier)
		if e != nil {
			return
		}
	}

	e = getCustomDataTypeFromIdentifier(spec, cluster, source, identifier)
	if e != nil {
		return
	}

	field, isField := entity.(*matter.Field)
	if isField {
		entity = field.Type.Entity
	}
	if entity != nil {
		var fieldSet matter.FieldSet
		switch entity := entity.(type) {
		case *matter.Struct:
			fieldSet = entity.Fields
		case *matter.Command:
			fieldSet = entity.Fields
		case *matter.Event:
			fieldSet = entity.Fields
		case *matter.Enum:
			for _, v := range entity.Values {
				if strings.EqualFold(v.Name, identifier) {
					e = v
					return
				}
			}

		case *matter.Bitmap:
			for _, v := range entity.Bits {
				if strings.EqualFold(v.Name(), identifier) {
					e = v
					return
				}
			}
		case *matter.Features:
			for f := range entity.FeatureBits() {
				if f.Code == identifier {
					e = f
					return
				}
			}
		default:
			slog.Warn("referenced identifier field has a type without fields", log.Path("source", source), slog.String("identifier", identifier), log.Type("type", entity))
		}
		if len(fieldSet) > 0 {
			childField := fieldSet.Get(identifier)
			if childField != nil {
				e = childField
				return
			}
		}
	}
	if cluster != nil {
		if cluster.Features != nil {
			for fb := range cluster.Features.FeatureBits() {
				if fb.Code == identifier { // Require case to match
					e = fb
					return
				}
			}
		}
		for _, a := range cluster.Attributes {
			if strings.EqualFold(a.Name, identifier) {
				e = a
				return
			}
		}
	}
	if spec.BaseDeviceType != nil {
		for _, con := range spec.BaseDeviceType.Conditions {
			if strings.EqualFold(identifier, con.Feature) {
				e = con
				return
			}
		}
	}
	slog.Error("Unable to find matching entity for identifier", log.Path("source", source), slog.String("identifier", identifier))
	return
}
