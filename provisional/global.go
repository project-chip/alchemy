package provisional

import (
	"iter"

	"github.com/project-chip/alchemy/internal"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func compareGlobals(specs spec.SpecPullRequest, violations entityViolations) {
	compareGlobalEntities(specs, iterateBits, violations)
	compareGlobalEntities(specs, iterateEnumValues, violations)
	compareGlobalEntities(specs, iterateStructFields, violations)
	compareGlobalEntities(specs, iterateCommandFields, violations)
	compareGlobalEntities(specs, iterateEventFields, violations)
}

func compareGlobalEntities[Parent ComparableEntity, Child ComparableEntity](pullRequest spec.SpecPullRequest, iterator func(p Parent) iter.Seq[Child], violations entityViolations) {

	baseEntities := types.FilterSet[Parent](pullRequest.Base.GlobalObjects)
	baseInProgressEntities := types.FilterSet[Parent](pullRequest.BaseInProgress.GlobalObjects)
	headEntities := types.FilterSet[Parent](pullRequest.Head.GlobalObjects)
	headInProgressEntities := types.FilterSet[Parent](pullRequest.HeadInProgress.GlobalObjects)

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
			violationType := checkProvisionalityOfNewEntity(pullRequest.HeadInProgress, state.HeadInProgress)
			if !novelty.IsIfDefd() {
				violationType |= spec.ViolationTypeNotIfDefd
			}
			violations.add(state.HeadInProgress, violationType)
		} else {
			for c := range iterator(e) {
				compareChildEntity(pullRequest.HeadInProgress, violations, c, state, iterator)
			}
		}

	}
}
