package regen

import (
	"slices"
	"strings"

	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
)

func enumsHelper(spec *spec.Specification) func(enums matter.EnumSet, options *raymond.Options) raymond.SafeString {
	return func(enums matter.EnumSet, options *raymond.Options) raymond.SafeString {
		sortedEnums := make(matter.EnumSet, len(enums))
		copy(sortedEnums, enums)
		slices.SortStableFunc(sortedEnums, func(a *matter.Enum, b *matter.Enum) int {
			return strings.Compare(a.Name, b.Name)
		})
		return enumerateEntitiesHelper(sortedEnums, spec, options)
	}
}
