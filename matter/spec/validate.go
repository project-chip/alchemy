package spec

import (
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

func Validate(spec *Specification) {
	validateClusters(spec)
	validateCommands(spec, nil, spec.GlobalObjects.OfType[*matter.Command]())
	validateEvents(spec, nil, spec.GlobalObjects.OfType[*matter.Event]())
	validateBitmaps(spec, spec.GlobalObjects.OfType[*matter.Bitmap]())
	validateEnums(spec, spec.GlobalObjects.OfType[*matter.Enum]())
	validateStructs(spec, nil, spec.GlobalObjects.OfType[*matter.Struct]())
	validateDeviceTypes(spec)
	validateNamespaces(spec)
}

func stripNonAlphabeticalCharacters(s string) string {
	var result strings.Builder
	for i := range len(s) {
		b := s[i]
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') {
			result.WriteByte(b)
		}
	}
	return result.String()
}

func checkEntityReference[EntityT types.Entity, SourceT types.Entity](spec *Specification, er matter.EntityReference, entity EntityT, source SourceT) {
	entityID := matter.EntityID(entity)
	if entityID.Valid() && er.ID.Valid() && !entityID.Equals(er.ID) {
		slog.Error("Mismatch between entity ID and reference entity ID", slog.String("entityId", entityID.HexString()), slog.String("referenceEntityId", er.ID.HexString()), log.Path("source", source))
		spec.addError(&ReferenceIDMismatchError[EntityT, SourceT]{Entity: entity, ID: er.ID, Source: source})
	}
	refName := stripNonAlphabeticalCharacters(er.Name)
	entityName := stripNonAlphabeticalCharacters(matter.EntityName(entity))
	if refName != "" && !strings.EqualFold(refName, entityName) {
		slog.Error("Mismatch between entity name and reference entity name", slog.String("entityName", entityName), slog.String("referenceEntityName", refName), slog.String("entityId", entityID.HexString()), log.Path("source", source))
		spec.addError(&ReferenceNameMismatchError[EntityT, SourceT]{Entity: entity, Name: er.Name, Source: source})
	}
}

type idUniqueness[T types.Entity] map[uint64]T

func (iu idUniqueness[T]) check(spec *Specification, id *matter.Number, entity T) {
	if id.Valid() {
		if previous, ok := iu[id.Value()]; ok {
			spec.addError(&DuplicateEntityIDError{Entity: entity, Previous: previous})
			return
		}
		iu[id.Value()] = entity
	}
}

type nameUniqueness[T types.Entity] map[string]T

func (nu nameUniqueness[T]) check(spec *Specification, entity T) {
	name := matter.EntityName(entity)
	if name == "" || text.HasCaseInsensitivePrefix(name, "reserved") {
		return
	}
	name = strings.ToLower(name)
	if previous, ok := nu[name]; ok {
		spec.addError(&DuplicateEntityNameError{Entity: entity, Previous: previous})
		return
	}

	nu[name] = entity

}

type conformanceValidation map[types.Entity][]conformance.Set

func (cv conformanceValidation) add(entity types.Entity, set conformance.Set) {
	if len(set) > 0 {
		cv[entity] = append(cv[entity], set)
	}
}

func (cv *conformanceValidation) check(spec *Specification) {

	type choiceLimit struct {
		entity types.Entity
		limit  conformance.ChoiceLimit
	}

	choiceCount := make(map[string][]types.Entity)
	choiceLimits := make(map[string]choiceLimit)
	for entity, sets := range *cv {
		for _, cs := range sets {
			for _, c := range cs {
				switch c := c.(type) {
				case *conformance.Optional:
					if c.Choice == nil || c.Choice.Set == "" {
						continue
					}
					choiceCount[c.Choice.Set] = append(choiceCount[c.Choice.Set], entity)
					existingLimit, existing := choiceLimits[c.Choice.Set]
					if !existing {

						choiceLimits[c.Choice.Set] = choiceLimit{
							entity: entity,
							limit:  c.Choice.Limit,
						}
					} else {
						if (c.Choice.Limit != nil && existingLimit.limit == nil) ||
							(c.Choice.Limit == nil && existingLimit.limit != nil) ||
							(c.Choice.Limit != nil && existingLimit.limit != nil && !existingLimit.limit.Equal(c.Choice.Limit)) {
							spec.addError(&ConformanceChoiceMismatchError{
								Set:                 c.Choice.Set,
								Source:              entity,
								Entity:              entity,
								ChoiceLimit:         c.Choice.Limit,
								Previous:            existingLimit.entity,
								PreviousChoiceLimit: existingLimit.limit,
							})
						}

					}
				}
			}
		}
	}
	for set, entities := range choiceCount {
		if len(entities) == 1 {
			spec.addError(&ConformanceChoiceOrphanError{Set: set, Source: entities[0]})
		}
	}
}
