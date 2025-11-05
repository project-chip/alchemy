package regen

import (
	"slices"
	"strings"

	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
)

func clusterAttributesHelper(spec *spec.Specification, commonAttributes matter.FieldSet) func(cluster matter.Cluster, options *raymond.Options) raymond.SafeString {
	return func(cluster matter.Cluster, options *raymond.Options) raymond.SafeString {
		cas := make(matter.FieldSet, 0, len(commonAttributes))
		for _, ca := range commonAttributes {
			ca = ca.Clone()
			ca.SetParent(&cluster)
			cas = append(cas, ca)
		}
		attributes := filterFields(cluster.Attributes, cas)
		slices.SortStableFunc(attributes, func(a *matter.Field, b *matter.Field) int {
			return a.ID.Compare(b.ID)
		})
		return enumerateEntitiesHelper(attributes, spec, options)
	}
}

func clusterStructsHelper(spec *spec.Specification) func(s matter.StructSet, options *raymond.Options) raymond.SafeString {
	return func(s matter.StructSet, options *raymond.Options) raymond.SafeString {
		structs := make(matter.StructSet, len(s))
		copy(structs, s)
		slices.SortStableFunc(structs, func(a *matter.Struct, b *matter.Struct) int {
			if a.FabricScoping != b.FabricScoping {
				if a.FabricScoping == matter.FabricScopingScoped {
					return 1
				}
				if b.FabricScoping == matter.FabricScopingScoped {
					return -1
				}
			}
			return strings.Compare(a.Name, b.Name)
		})
		return enumerateEntitiesHelper(structs, spec, options)

	}
}

func clusterEventsHelper(spec *spec.Specification) func(cluster matter.Cluster, options *raymond.Options) raymond.SafeString {
	return func(cluster matter.Cluster, options *raymond.Options) raymond.SafeString {
		events := make(matter.EventSet, len(cluster.Events))
		copy(events, cluster.Events)
		return enumerateEntitiesHelper(events, spec, options)
	}
}
