package provisional

import (
	"iter"

	"github.com/project-chip/alchemy/internal"
	"github.com/project-chip/alchemy/matter/types"
)

func compareGlobals(specs specs, violations entityViolations) {
	compareGlobalEntities(specs, iterateBits, violations)
	compareGlobalEntities(specs, iterateEnumValues, violations)
	compareGlobalEntities(specs, iterateStructFields, violations)
	compareGlobalEntities(specs, iterateCommandFields, violations)
	compareGlobalEntities(specs, iterateEventFields, violations)
}

func compareGlobalEntities[Parent ComparableEntity, Child ComparableEntity](specs specs, iterator func(p Parent) iter.Seq[Child], violations entityViolations) {

	baseEntities := types.FilterSet[Parent](specs.Base.GlobalObjects)
	baseInProgressEntities := types.FilterSet[Parent](specs.BaseInProgress.GlobalObjects)
	headEntities := types.FilterSet[Parent](specs.Head.GlobalObjects)
	headInProgressEntities := types.FilterSet[Parent](specs.HeadInProgress.GlobalObjects)

	for _, e := range headInProgressEntities {
		state := EntityState[Parent]{HeadInProgress: e}
		state.Head = findExistingEntity(e, internal.Iterate(headEntities))
		state.BaseInProgress = findExistingEntity(e, internal.Iterate(baseInProgressEntities))
		state.Base = findExistingEntity(e, internal.Iterate(baseEntities))

		presence := state.Presence()
		novelty, err := presence.Novelty()
		if err != nil {
			continue
		}
		if novelty.IsNew() {
			violationType := checkProvisionality(specs.HeadInProgress, state.HeadInProgress)
			if !novelty.IsIfDefd() {
				violationType |= ViolationTypeNotIfDefd
			}
			violations.add(state.HeadInProgress, violationType)
		} else {
			for c := range iterator(e) {
				compareChildEntity(specs.HeadInProgress, violations, c, state, iterator)
			}
		}

	}
}
