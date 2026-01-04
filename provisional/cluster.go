package provisional

import (
	"iter"
	"log/slog"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
)

func compareClusters(specs spec.SpecPullRequest, violations entityViolations) {
	var clusterStates []EntityState[*matter.Cluster]
	for cluster := range specs.HeadInProgress.Clusters {
		state := EntityState[*matter.Cluster]{HeadInProgress: cluster}
		state.Head = findExistingClusterInSpec(cluster, specs.Head)
		state.BaseInProgress = findExistingClusterInSpec(cluster, specs.BaseInProgress)
		state.Base = findExistingClusterInSpec(cluster, specs.Base)
		clusterStates = append(clusterStates, state)
	}

	for _, state := range clusterStates {
		presence := state.Presence()
		novelty, err := presence.Novelty()
		if err != nil {
			continue
		}
		if novelty.IsNew() {
			violationType := checkProvisionalityOfNewEntity(specs.HeadInProgress, state.HeadInProgress)
			if !novelty.IsIfDefd() {
				violationType |= spec.ViolationTypeNotIfDefd
			}
			violations.add(state.HeadInProgress, violationType)
		} else {
			compareClusterEntities(specs.HeadInProgress, violations, state)
		}
	}
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

func compareChildEntity[Parent ComparableEntity, Child ComparableEntity](s *spec.Specification, violations entityViolations, entity Child, parentState EntityState[Parent], childIterator func(parent Parent) iter.Seq[Child]) {
	entityState := getEntityState(entity, parentState, childIterator)
	presence := entityState.Presence()
	novelty, err := presence.Novelty()
	if err != nil {
		slog.Error("Invalid presence on entity", matter.LogEntity("entity", entity), slog.String("presence", entityState.Presence().String()), log.Path("source", entity))
		return
	}
	if !novelty.IsNew() {
		return
	}
	violationType := checkProvisionalityOfNewEntity(s, entity)

	if !novelty.IsIfDefd() {
		violationType |= spec.ViolationTypeNotIfDefd
	}
	violations.add(entity, violationType)
}

func compareClusterParentEntity[Parent ComparableEntity, Child ComparableEntity](s *spec.Specification, violations entityViolations, clusterState EntityState[*matter.Cluster], parentIterator func(c *matter.Cluster) iter.Seq[Parent], childIterator func(parent Parent) iter.Seq[Child]) {
	for e := range parentIterator(clusterState.HeadInProgress) {
		entityState := getEntityState(e, clusterState, parentIterator)
		novelty, err := entityState.Presence().Novelty()
		if err != nil {
			slog.Error("Invalid presence on entity", matter.LogEntity("entity", e), slog.String("presence", entityState.Presence().String()), log.Path("source", e))
			continue
		}
		if novelty.IsNew() {
			violationType := checkProvisionalityOfNewEntity(s, entityState.HeadInProgress)
			if !novelty.IsIfDefd() {
				violationType |= spec.ViolationTypeNotIfDefd
			}
			violations.add(e, violationType)
		} else { // We don't need to check the children of new entities
			for child := range childIterator(e) {
				compareChildEntity(s, violations, child, entityState, childIterator)
			}
		}
	}
}

func compareClusterEntities(spec *spec.Specification, violations entityViolations, clusterState EntityState[*matter.Cluster]) {

	for f := range clusterState.HeadInProgress.Features.FeatureBits() {
		compareChildEntity(spec, violations, f, clusterState, iterateFeatures)
	}
	compareClusterParentEntity(spec, violations, clusterState, iterateBitmaps, iterateBits)
	compareClusterParentEntity(spec, violations, clusterState, iterateEnums, iterateEnumValues)
	compareClusterParentEntity(spec, violations, clusterState, iterateStructs, iterateStructFields)

	for _, a := range clusterState.HeadInProgress.Attributes {
		compareChildEntity(spec, violations, a, clusterState, iterateAttributes)
	}

	compareClusterParentEntity(spec, violations, clusterState, iterateCommands, iterateCommandFields)
	compareClusterParentEntity(spec, violations, clusterState, iterateEvents, iterateEventFields)

}
