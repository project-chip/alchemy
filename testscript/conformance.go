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

func findReferencedEntities(entity types.Entity, variables map[types.Entity]struct{}) {
	switch entity := entity.(type) {
	case *matter.Cluster:
		for _, a := range entity.Attributes {
			findVariablesForConstraint(a.Constraint, variables)
		}
		for _, s := range entity.Structs {
			findReferencedEntities(s, variables)
		}
	case *matter.Struct:
		for _, f := range entity.Fields {
			findReferencedEntities(f, variables)
		}
	case *matter.Command:
		for _, f := range entity.Fields {
			findReferencedEntities(f, variables)
		}
	case *matter.Event:
		for _, f := range entity.Fields {
			findReferencedEntities(f, variables)
		}
	case *matter.Field:
		findVariablesForConstraint(entity.Constraint, variables)
	default:
		slog.Warn("Unexpected entity finding referenced entities", log.Type("type", entity))
	}
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
	case *constraint.ListConstraint:
		findVariablesForConstraint(c.Constraint, variables)
		findVariablesForConstraint(c.EntryConstraint, variables)
	case nil:
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
	case *constraint.CharacterLimit:
		findVariablesForLimit(l.ByteCount, variables)
		findVariablesForLimit(l.CodepointCount, variables)
	case *constraint.IntLimit,
		*constraint.HexLimit,
		*constraint.TemperatureLimit,
		*constraint.PercentLimit,
		*constraint.ManufacturerLimit,
		*constraint.ExpLimit:
		return
	default:
		slog.Warn("Unexpected limit type for variables", log.Type("type", l))
	}
}
