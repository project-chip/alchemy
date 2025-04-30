package provisional

import (
	"log/slog"
	"runtime/debug"

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
	slog.Info("CHECKING isProvisional", matter.LogEntity("entity", entity))
	is := Check(spec, entity, entity)
	slog.Info("CHECKED isProvisional", matter.LogEntity("entity", entity), "isProvisional", is.String())
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
	slog.Info("checking isProvisional", matter.LogEntity("entity", entity))
	if entity != nil {
		slog.Info("checking isProvisional", matter.LogEntity("parent", entity.Parent()))
	}
	if entity == nil {
		debug.PrintStack()
	}
	if conformance.IsProvisional(matter.EntityConformance(entity)) {
		// This is explicitly marked provisional
		slog.Info("explicitly isProvisional", matter.LogEntity("entity", entity))
		return StateExplicitlyProvisional
	}
	switch entity := entity.(type) {
	case *matter.Features:
		b := Check(spec, entity.Parent(), originalEntity)
		slog.Info("features parent isProvisional", matter.LogEntity("entity", entity), "isProvisional", b)
		return b
	case *matter.Command, *matter.Event:
		clusterProvisional := Check(spec, entity.Parent(), originalEntity)
		switch clusterProvisional {
		case StateAllClustersProvisional, StateExplicitlyProvisional:
			slog.Info("command or event parent cluster is provisional", matter.LogEntity("entity", entity))
			return clusterProvisional
		case StateAllClustersNonProvisional:
			slog.Info("command or event parent cluster is not provisional", matter.LogEntity("entity", entity))
			return clusterProvisional
		}
		clusters, ok := spec.ClusterRefs.Get(entity)
		if !ok || len(clusters) == 0 {
			slog.Info("command or event has no clusters", matter.LogEntity("entity", entity))
			return StateUnreferenced
		}
		var hasProvisionalCluster bool
		var hasNonProvisionalCluster bool
		for cluster := range clusters {
			slog.Info("checking associated cluster", matter.LogEntity("entity", entity), "cluster", cluster.Name, "isProvisional", conformance.IsProvisional(cluster.Conformance))
			if !conformance.IsProvisional(cluster.Conformance) {
				slog.Info("entity has non-provisional cluster reference", matter.LogEntity("entity", entity), "cluster", cluster.Name)
				hasNonProvisionalCluster = true

			} else {
				slog.Info("entity has provisional cluster reference", matter.LogEntity("entity", entity), "cluster", cluster.Name)
				hasProvisionalCluster = true
			}
		}
		slog.Info("command or event checked non-provisional clusters", matter.LogEntity("entity", entity), "hasNonProvisionalCluster", hasNonProvisionalCluster)
		if hasProvisionalCluster {
			if !hasNonProvisionalCluster {
				return StateAllClustersProvisional
			}
			return StateSomeClustersProvisional
		}
		return StateAllClustersNonProvisional
	case *matter.Enum, *matter.Bitmap, *matter.Struct:
		refs, ok := spec.DataTypeRefs.Get(entity)
		if !ok || len(refs) == 0 {
			slog.Info("enum, bitmap or struct has no clusters", matter.LogEntity("entity", entity))
			return StateUnreferenced
		}
		var hasProvisionalRef bool
		var hasNonProvisionalRef bool
		for ref := range refs {
			slog.Info("checking reference", matter.LogEntity("entity", entity), matter.LogEntity("ref", ref), "isProvisional", Check(spec, ref, originalEntity))
			switch Check(spec, ref, originalEntity) {
			case StateExplicitlyProvisional, StateAllClustersProvisional, StateAllDataTypeReferencesProvisional:
				hasProvisionalRef = true
			case StateSomeClustersProvisional, StateSomeDataTypeReferencesProvisional:
				slog.Info("entity has some non-provisional reference", matter.LogEntity("entity", entity), matter.LogEntity("ref", ref))
				hasNonProvisionalRef = true
				hasProvisionalRef = true
			case StateAllDataTypeReferencesNonProvisional, StateAllClustersNonProvisional:
				slog.Info("entity has non-provisional reference", matter.LogEntity("entity", entity), matter.LogEntity("ref", ref))
				hasNonProvisionalRef = true
			case StateUnreferenced:

			default:
				slog.Warn("Unexpected provisional state", "state", Check(spec, ref, originalEntity))
			}

		}
		if hasProvisionalRef {
			if hasNonProvisionalRef {
				return StateSomeDataTypeReferencesProvisional
			}
			return StateAllDataTypeReferencesProvisional
		} else if hasNonProvisionalRef {
			return StateAllDataTypeReferencesNonProvisional
		}
		return StateUnreferenced
	case *matter.EnumValue:
		if conformance.IsProvisional(matter.EntityConformance(entity)) {
			// This is explicitly marked provisional
			slog.Info("isProvisional explicitly", matter.LogEntity("entity", entity))
			return StateExplicitlyProvisional
		}
		b := Check(spec, entity.Parent(), originalEntity)
		slog.Info("enum value isProvisional", matter.LogEntity("entity", entity), "isProvisional", b)
		return b
	case matter.Bit:
		if conformance.IsProvisional(matter.EntityConformance(entity)) {
			// This is explicitly marked provisional
			slog.Info("isProvisional", matter.LogEntity("entity", entity))
			return StateExplicitlyProvisional
		}
		b := Check(spec, entity.Parent(), originalEntity)
		slog.Info("bit isProvisional", matter.LogEntity("entity", entity), "isProvisional", b)
		return b
	case *matter.Field:
		if conformance.IsProvisional(matter.EntityConformance(entity)) {
			// This is explicitly marked provisional
			slog.Info("isProvisional", matter.LogEntity("entity", entity))
			return StateExplicitlyProvisional
		}
		b := Check(spec, entity.Parent(), originalEntity)
		slog.Info("field parent isProvisional", matter.LogEntity("entity", entity), "isProvisional", b)
		return b
	case *matter.Cluster:
		b := conformance.IsProvisional(matter.EntityConformance(entity))
		slog.Info("cluster isProvisional", matter.LogEntity("entity", entity), "isProvisional", b)
		if b {
			return StateAllClustersProvisional
		}
		return StateAllClustersNonProvisional
	default:
		slog.Error("Unexpected entity type checking provisional status", matter.LogEntity("entity", entity))
	}
	return StateUnknown
}
