package testscript

import (
	"log/slog"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/types"
)

func isExcludedConformance(c conformance.Conformance) bool {
	if conformance.IsDeprecated(c) {
		return true
	}
	if conformance.IsDisallowed(c) {
		return true
	}
	if conformance.IsProvisional(c) {
		return true
	}
	return false
}

func findVariables(cluster *matter.Cluster) (variables map[types.Entity]struct{}) {
	variables = make(map[types.Entity]struct{})

	for _, a := range cluster.Attributes {
		findVariablesForConstraint(a.Constraint, variables)
	}
	return
}

func findVariablesForConstraint(c constraint.Constraint, variables map[types.Entity]struct{}) {
	switch c := c.(type) {
	case constraint.Set:
		for _, c := range c {
			findVariablesForConstraint(c, variables)
		}
	case *constraint.AllConstraint, *constraint.DescribedConstraint, *constraint.GenericConstraint:
		return
	case *constraint.ExactConstraint:
		if c.Value == nil {
			return
		}
		findVariablesForLimit(c.Value, variables)
	case *constraint.RangeConstraint:
		findVariablesForLimit(c.Minimum, variables)
		findVariablesForLimit(c.Maximum, variables)
	case *constraint.MaxConstraint:
		findVariablesForLimit(c.Maximum, variables)
	case *constraint.MinConstraint:
		findVariablesForLimit(c.Minimum, variables)
	default:
		slog.Warn("Unexpected constraint type for variables", log.Type("type", c))
	}
}

func findVariablesForLimit(l constraint.Limit, variables map[types.Entity]struct{}) {
	switch l := l.(type) {
	case *constraint.IdentifierLimit:
		if l.Entity != nil {
			variables[l.Entity] = struct{}{}
		}
	case *constraint.ReferenceLimit:
		if l.Entity != nil {
			variables[l.Entity] = struct{}{}
		}
	case *constraint.TagIdentifierLimit:
		if l.Entity != nil {
			variables[l.Entity] = struct{}{}
		}
	case *constraint.MathExpressionLimit:
		findVariablesForLimit(l.Left, variables)
		findVariablesForLimit(l.Right, variables)
	case *constraint.LogicalLimit:
		findVariablesForLimit(l.Left, variables)
		for _, r := range l.Right {
			findVariablesForLimit(r, variables)
		}
	case *constraint.IntLimit, *constraint.TemperatureLimit, *constraint.PercentLimit:
		return
	default:
		slog.Warn("Unexpected limit type for variables", log.Type("type", l))
	}
}
