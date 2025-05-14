package provisional

import (
	"log/slog"

	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

type State uint8

const (
	StateUnknown State = iota
	StateExplicitlyProvisional
	StateAllDataTypeReferencesProvisional
	StateAllDataTypeReferencesNonProvisional
	StateSomeDataTypeReferencesProvisional
	StateAllClustersProvisional
	StateAllClustersNonProvisional
	StateSomeClustersProvisional
	StateUnreferenced
)

func (s State) String() string {
	switch s {
	case StateUnknown:
		return "unknown"
	case StateExplicitlyProvisional:
		return "explicitlyProvisional"
	case StateAllDataTypeReferencesProvisional:
		return "allDataTypeReferencesProvisional"
	case StateAllDataTypeReferencesNonProvisional:
		return "allDataTypeReferencesNonProvisional"
	case StateSomeDataTypeReferencesProvisional:
		return "someDataTypeReferencesProvisional"
	case StateAllClustersProvisional:
		return "allClustersProvisional"
	case StateAllClustersNonProvisional:
		return "allClustersNonProvisional"
	case StateSomeClustersProvisional:
		return "someClustersProvisional"
	case StateUnreferenced:
		return "unreferenced"
	default:
		return "invalid"
	}
}

func IsProvisional(spec *spec.Specification, entity types.Entity) bool {
	is := Check(spec, entity, entity)
	switch is {
	case StateAllClustersProvisional,
		StateAllDataTypeReferencesProvisional,
		StateExplicitlyProvisional:
		return true
	default:
		return false
	}
}

func Check(spec *spec.Specification, entity types.Entity, originalEntity types.Entity) State {
	if conformance.IsProvisional(matter.EntityConformance(entity)) {
		// This is explicitly marked provisional
		return StateExplicitlyProvisional
	}
	// There are a couple of test data types that should just be marked as non-provisional
	switch entity := entity.(type) {
	case *matter.Struct:
		if entity.Name == "TestGlobalStruct" {
			return StateAllDataTypeReferencesNonProvisional
		}
	case *matter.Enum:
		if entity.Name == "TestGlobalEnum" {
			return StateAllDataTypeReferencesNonProvisional
		}
	}
	switch entity := entity.(type) {
	case *matter.Cluster:
		b := conformance.IsProvisional(matter.EntityConformance(entity))
		if b {
			return StateAllClustersProvisional
		}
		return StateAllClustersNonProvisional
	case *matter.Features:
		b := Check(spec, entity.Parent(), originalEntity)
		return b

	case *matter.Enum, *matter.Bitmap, *matter.Struct:
		refs, ok := spec.DataTypeRefs.Get(entity)
		if !ok || refs.Size() == 0 {
			slog.Warn("Enum, bitmap or struct has no references; assuming provisional", matter.LogEntity("entity", entity))
			return StateUnreferenced
		}
		var hasProvisionalRef bool
		var hasNonProvisionalRef bool
		refs.Range(func(ref types.Entity, value struct{}) bool {
			switch Check(spec, ref, originalEntity) {
			case StateExplicitlyProvisional, StateAllClustersProvisional, StateAllDataTypeReferencesProvisional:
				hasProvisionalRef = true
			case StateSomeClustersProvisional, StateSomeDataTypeReferencesProvisional:
				hasNonProvisionalRef = true
				hasProvisionalRef = true
			case StateAllDataTypeReferencesNonProvisional, StateAllClustersNonProvisional:
				hasNonProvisionalRef = true
			case StateUnreferenced:

			default:
				slog.Warn("Unexpected provisional state", "state", Check(spec, ref, originalEntity))
			}
			return true
		})
		if hasProvisionalRef {
			if hasNonProvisionalRef {
				return StateSomeDataTypeReferencesProvisional
			}
			return StateAllDataTypeReferencesProvisional
		} else if hasNonProvisionalRef {
			return StateAllDataTypeReferencesNonProvisional
		}
		return StateUnreferenced
	case *matter.EnumValue, matter.Bit, *matter.Field:
		if conformance.IsProvisional(matter.EntityConformance(entity)) {
			// This is explicitly marked provisional
			return StateExplicitlyProvisional
		}
		b := Check(spec, entity.Parent(), originalEntity)
		return b
	case *matter.Command, *matter.Event:
		cluster := entity.Parent()
		if cluster != nil {
			clusterProvisional := Check(spec, cluster, originalEntity)
			switch clusterProvisional {
			case StateAllClustersProvisional, StateExplicitlyProvisional:
				return clusterProvisional
			case StateAllClustersNonProvisional:
				return clusterProvisional
			}
		}
		clusters, ok := spec.ClusterRefs.Get(entity)
		if !ok || clusters.Size() == 0 {
			slog.Warn("command or event has no clusters", matter.LogEntity("entity", entity))
			return StateUnreferenced
		}
		var hasProvisionalCluster bool
		var hasNonProvisionalCluster bool
		clusters.Range(func(cluster *matter.Cluster, value struct{}) bool {
			if !conformance.IsProvisional(cluster.Conformance) {
				hasNonProvisionalCluster = true
			} else {
				hasProvisionalCluster = true
			}
			return true
		})
		if hasProvisionalCluster {
			if !hasNonProvisionalCluster {
				return StateAllClustersProvisional
			}
			return StateSomeClustersProvisional
		}
		return StateAllClustersNonProvisional
	default:
		slog.Error("Unexpected entity type checking provisional status", matter.LogEntity("entity", entity))
	}
	return StateUnknown
}
