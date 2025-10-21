package provisional

import (
	"iter"
	"log/slog"
	"reflect"

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
	if !isNil(es.Base) {
		p |= PresenceBase
	}
	if !isNil(es.BaseInProgress) {
		p |= PresenceBaseInProgress
	}
	if !isNil(es.Head) {
		p |= PresenceHead
	}
	if !isNil(es.HeadInProgress) {
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

		v := spec.Violation{Entity: entity, Type: violationType}
		v.Path, v.Line = entity.Origin()
		violationsByPath[v.Path] = append(violationsByPath[v.Path], v)
	}
	return
}

func getEntityState[T ComparableEntity, Parent types.Entity](e T, parentState EntityState[Parent], iterator func(p Parent) iter.Seq[T]) EntityState[T] {
	state := EntityState[T]{HeadInProgress: e}
	if !isNil(parentState.Head) {
		state.Head = findExistingEntity(e, iterator(parentState.Head))
	}
	if !isNil(parentState.BaseInProgress) {
		state.BaseInProgress = findExistingEntity(e, iterator(parentState.BaseInProgress))
	}
	if !isNil(parentState.Base) {
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

func isNil[T any](val T) bool {
	rValue := reflect.ValueOf(val)
	switch rValue.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func, reflect.Interface:
		return rValue.IsNil()
	default:
		return false
	}
}

func checkProvisionality(s *spec.Specification, e types.Entity) (violationType spec.ViolationType) {
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
