package testscript

import (
	"log/slog"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter/types"
)

type dependencyGraph struct {
	dependencies map[types.Entity]map[*dependencyNode]struct{}
	nodes        []*dependencyNode
}

func reorderDependencies(t *Test) {

	nodes := make([]*dependencyNode, 0, len(t.Steps))

	for _, a := range t.Steps {
		node := &dependencyNode{step: a, requires: make(map[*dependencyNode]struct{}), satisfies: make(map[*dependencyNode]struct{})}
		nodes = append(nodes, node)
	}

	satisfies := make(map[types.Entity]map[*dependencyNode]struct{})
	for _, n := range nodes {
		getSatisfies(n, satisfies)
	}

	dependsOn := make(map[types.Entity]map[*dependencyNode]struct{})
	for _, n := range nodes {
		getDepends(n, dependsOn)
	}

	for de, dn := range dependsOn {
		ds, ok := satisfies[de]
		if !ok {
			slog.Warn("missing satisfies for dependency")
		}
		for s := range ds {
			dn[s] = struct{}{}
		}
	}
}

type dependencyNode struct {
	step      *TestStep
	requires  map[*dependencyNode]struct{}
	satisfies map[*dependencyNode]struct{}
}

func getSatisfies(n *dependencyNode, s map[types.Entity]map[*dependencyNode]struct{}) {
	for _, a := range n.step.Actions {
		switch a := a.(type) {
		case *ReadAttribute:
			addDependency(n, s, a.Attribute)
		}
	}
}

func getDepends(n *dependencyNode, s map[types.Entity]map[*dependencyNode]struct{}) {
	for _, a := range n.step.Actions {
		switch a := a.(type) {
		case *CheckType:
			variables := make(map[types.Entity]struct{})
			findReferencedEntities(a.Field, variables)
			for e := range variables {
				addDependency(n, s, e)
			}
		case *CheckMinConstraint:
			addDependency(n, s, a.Field)
		case *CheckMaxConstraint:
			addDependency(n, s, a.Field)
		default:
			slog.Warn("Unexpected action type for dependency graph", log.Type("type", a))
		}
	}
}

func addDependency(n *dependencyNode, d map[types.Entity]map[*dependencyNode]struct{}, e types.Entity) {
	var ok bool
	var ds map[*dependencyNode]struct{}
	ds, ok = d[e]
	if !ok {
		ds = make(map[*dependencyNode]struct{})
		d[e] = ds
	}
	ds[n] = struct{}{}
}
