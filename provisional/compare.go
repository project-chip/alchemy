package provisional

import (
	"iter"
	"log/slog"

	"github.com/project-chip/alchemy/internal"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

type entityViolations map[types.Entity]spec.ViolationType

func (ev entityViolations) add(entity types.Entity, violationType spec.ViolationType) {
	if violationType != spec.ViolationTypeNone {
		ev[entity] = violationType
	}
}

type EntityState[T types.Entity] struct {
	HeadInProgress T
	Head           T
	BaseInProgress T
	Base           T
}

func (es EntityState[T]) Presence() (p Presence) {
	if !internal.IsNil(es.Base) {
		p |= PresenceBase
	}
	if !internal.IsNil(es.BaseInProgress) {
		p |= PresenceBaseInProgress
	}
	if !internal.IsNil(es.Head) {
		p |= PresenceHead
	}
	if !internal.IsNil(es.HeadInProgress) {
		p |= PresenceHeadInProgress
	}
	return
}

func compare(specs spec.SpecPullRequest) (violationsByPath map[string][]spec.Violation) {
	violationsByPath = make(map[string][]spec.Violation)
	violations := make(entityViolations)
	compareClusters(specs, violations)
	compareGlobals(specs, violations)

	for entity, violationType := range violations {
		if violationType == spec.ViolationTypeNone {
			continue
		}
		parent := entity.Parent()
		for parent != nil {
			if conformance.IsProvisional(matter.EntityConformance(parent)) {
				break
			}
			parent = parent.Parent()
		}

		// When an ancestor is provisional, we don't need to report non-provisional violations for this entity
		if parent != nil {
			violationType &= ^spec.ViolationTypeNonProvisional
			if violationType == spec.ViolationTypeNone {
				continue
			}
		}

		v := spec.Violation{Entity: entity, Type: violationType}
		v.Path, v.Line = entity.Origin()
		violationsByPath[v.Path] = append(violationsByPath[v.Path], v)
	}
	return
}

func getEntityState[T ComparableEntity, Parent types.Entity](e T, parentState EntityState[Parent], iterator func(p Parent) iter.Seq[T]) EntityState[T] {
	state := EntityState[T]{HeadInProgress: e}
	if !internal.IsNil(parentState.Head) {
		state.Head = findExistingEntity(e, iterator(parentState.Head))
	}
	if !internal.IsNil(parentState.BaseInProgress) {
		state.BaseInProgress = findExistingEntity(e, iterator(parentState.BaseInProgress))
	}
	if !internal.IsNil(parentState.Base) {
		state.Base = findExistingEntity(e, iterator(parentState.Base))
	}
	return state
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

func checkProvisionalityOfNewEntity(s *spec.Specification, e types.Entity) (violationType spec.ViolationType) {
	switch e.(type) {
	case *matter.Cluster,
		*matter.DeviceType,
		*matter.Feature,
		*matter.Command,
		*matter.Event,
		*matter.Field,
		*matter.EnumValue,
		matter.Bit:
		// All these are types that can be marked explicitly provisional
		if !conformance.IsProvisional(matter.EntityConformance(e)) {
			violationType = spec.ViolationTypeNonProvisional
		}
	case *matter.Bitmap, *matter.Enum, *matter.Struct:
		// These types can't be marked provsional, so we'll rely on ifdefs
	default:
		slog.Error("Unexpected provisionality entity", matter.LogEntity("entity", e))
	}

	return
}
