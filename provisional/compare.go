package provisional

import (
	"iter"
	"log/slog"
	"reflect"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

type EntityState[T types.Entity] struct {
	HeadInProgress T
	Head           T
	BaseInProgress T
	Base           T
}

func compare(specs specs) (violations map[string][]Violation) {
	violations = make(map[string][]Violation)
	compareClusters(specs, violations)
	compareGlobals(specs, violations)
	return
}

func findExistingClusterInSpec(needle *matter.Cluster, haystack *spec.Specification) (existingCluster *matter.Cluster) {
	var ok bool
	if needle.ID.Valid() {
		existingCluster, ok = haystack.ClustersByID[needle.ID.Value()]
	}
	if !ok {
		existingCluster = haystack.ClustersByName[needle.Name]
	}
	return
}

type ComparableEntity interface {
	types.Entity
	Equals(types.Entity) bool
}

func findExistingEntity[T ComparableEntity](needle ComparableEntity, haystack iter.Seq[T]) (existing T) {
	for hay := range haystack {
		if hay.Equals(needle) {
			return hay
		}
	}
	return
}

func isNil[T any](val T) bool {
	rValue := reflect.ValueOf(val)
	switch rValue.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func, reflect.Interface:
		return rValue.IsNil()
	default:
		return false
	}
}

func compareStates[T types.Entity](spec *spec.Specification, violations map[string][]Violation, state EntityState[T]) {
	//slog.Info(state.HeadInProgress.EntityType().String(), "name", matter.EntityName(state.HeadInProgress), "inHead", state.Head, "inBaseIP", state.BaseInProgress, "inBase", state.Base)
	if !isNil(state.Head) {
		//	slog.Info("state.Head is not nil")
		// This entity is not wrapped in in-progress in HEAD
		if !isNil(state.BaseInProgress) {
			//		slog.Info("state.BaseInProgress is not nil")
			if !isNil(state.Base) {
				//			slog.Info("state.Base is not nil")
				// This entity already exists in the main spec, no in-progress ifdefs
				return
			} else {
				//			slog.Info("state.Base is nil")
				// This entity had in-progress around it, but that's being removed in this PR
				return
			}
		} else {
			//		slog.Info("state.BaseInProgress is nil")
			// This is a new entity that's not wrapped in in-progress!
			checkProvisionality(spec, violations, state.HeadInProgress, false)
			/*			switch Check(spec, state.HeadInProgress, state.HeadInProgress, false) {
						case StateAllClustersProvisional,
							StateAllDataTypeReferencesProvisional,
							StateExplicitlyProvisional:
						default:
						}*/
			/*if IsProvisional(spec, state.HeadInProgress) {
				slog.Error("New entity added without ifdef!", matter.LogEntity("entity", state.Head))
			} else {
				slog.Error("Non-provisional entity added without ifdef!", matter.LogEntity("entity", state.Head))
			}*/
		}
	} else {
		//slog.Info("state.Head is nil")
		// This entity is wrapped in in-progress in HEAD
		if !isNil(state.BaseInProgress) {
			//	slog.Info("state.BaseInProgress is not nil")
			if !isNil(state.Base) {
				//		slog.Info("state.Base is not nil")
				// This PR adds in-progress to this entity

			} else {
				//		slog.Info("state.Base is  nil")
				// This entity is wrapped in in-progress in the base ref
			}
		} else {
			//	slog.Info("state.BaseInProgress is nil")
			// This is a new entity that's wrapped in in-progress; check if it's provisional
			checkProvisionality(spec, violations, state.HeadInProgress, true)
			/*			if IsProvisional(spec, state.HeadInProgress) {
							//		slog.Info("entity is provisional")
							// We're cool; the whole cluster is provisional, so everything included within it will also be provisional
							return
						} else {
							// TODO: Problem! Non-provisional cluster being added
							slog.Error("Non-provisional entity added!", matter.LogEntity("entity", state.HeadInProgress))
						}*/
		}
	}
}

func checkProvisionality(spec *spec.Specification, violations map[string][]Violation, e types.Entity, ifDefd bool) {
	var violationType ViolationType
	state := Check(spec, e, e)
	switch e.(type) {
	case *matter.Cluster,
		*matter.Feature,
		*matter.Command,
		*matter.Field:
		switch state {
		case StateAllClustersProvisional,
			StateExplicitlyProvisional:
		case StateAllClustersNonProvisional, StateSomeClustersProvisional:
			violationType |= ViolationTypeNonProvisional
		default:
			slog.Warn("Unexpected provisional state for entity", matter.LogEntity("entity", e), slog.String("state", state.String()))
		}
	case *matter.Bitmap, *matter.Enum, *matter.Struct:
		switch state {
		case StateAllClustersProvisional,
			StateAllDataTypeReferencesProvisional:
		case StateAllClustersNonProvisional,
			StateSomeClustersProvisional,
			StateSomeDataTypeReferencesProvisional,
			StateAllDataTypeReferencesNonProvisional:
			violationType |= ViolationTypeNonProvisional
		case StateUnreferenced:
			if !ifDefd {
				violationType |= ViolationTypeNotIfDefd
			}
		default:
			slog.Warn("Unexpected provisional state for entity", matter.LogEntity("entity", e), slog.String("state", state.String()))

		}
	case *matter.EnumValue, matter.Bit:
		switch state {
		case StateAllClustersProvisional,
			StateAllDataTypeReferencesProvisional,
			StateExplicitlyProvisional:
		case StateAllClustersNonProvisional,
			StateSomeClustersProvisional,
			StateSomeDataTypeReferencesProvisional,
			StateAllDataTypeReferencesNonProvisional:
			violationType |= ViolationTypeNonProvisional
		case StateUnreferenced:
			if !ifDefd {
				violationType |= ViolationTypeNotIfDefd
			}
		default:
			slog.Warn("Unexpected provisional state for entity", matter.LogEntity("entity", e), slog.String("state", state.String()))
		}
	default:
		slog.Error("Unexpected provisionality entity", matter.LogEntity("entity", e))

	}
	if violationType != ViolationTypeNone {
		if !ifDefd {
			violationType |= ViolationTypeNotIfDefd
		}
		v := Violation{Entity: e, Type: violationType}
		source, ok := e.(log.Source)
		if ok {
			v.Path, v.Line = source.Origin()
		}
		violations[v.Path] = append(violations[v.Path], v)
	}
}

/*
	StateExplicitlyProvisional
	StateAllDataTypeReferencesProvisional
	StateAllDataTypeReferencesNonProvisional
	StateSomeDataTypeReferencesProvisional
	StateAllClustersProvisional
	StateAllClustersNonProvisional
	StateSomeClustersProvisional
	StateUnreferenced
*/
