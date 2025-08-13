package spec

import (
	"strings"

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
