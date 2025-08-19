package spec

import (
	"strings"

	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func Validate(spec *Specification) {
	validateAttributes(spec)
	validateCommands(spec)
	validateEvents(spec)
	validateBitmaps(spec)
	validateEnums(spec)
	validateStructs(spec)
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
