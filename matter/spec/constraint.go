package spec

import (
	"log/slog"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/types"
)

func (sp *Builder) resolveConstraints() {
	specEntityFinder := newSpecEntityFinder(sp.Spec, nil, nil)
	for cluster := range sp.Spec.Clusters {
		specEntityFinder.cluster = cluster
		for _, a := range cluster.Attributes {
			sp.resolveFieldConstraints(cluster, cluster.Attributes, a, a.Type, specEntityFinder)
		}
		attributeFinder := newFieldFinder(cluster.Attributes, specEntityFinder)
		for _, s := range cluster.Structs {
			structFinder := newFieldFinder(s.Fields, attributeFinder)
			for _, f := range s.Fields {
				sp.resolveFieldConstraints(cluster, s.Fields, f, f.Type, structFinder)
			}
		}
		commandFinder := newCommandFinder(cluster.Commands, attributeFinder)
		for _, command := range cluster.Commands {
			for _, f := range command.Fields {
				sp.resolveFieldConstraints(cluster, command.Fields, f, f.Type, commandFinder)
			}
		}
		eventFinder := newEventFinder(cluster.Events, attributeFinder)
		for _, event := range cluster.Events {
			for _, f := range event.Fields {
				sp.resolveFieldConstraints(cluster, event.Fields, f, f.Type, eventFinder)
			}
		}
	}
	specEntityFinder.cluster = nil
	for _, s := range sp.Spec.structIndex {
		if s.Parent() == nil {
			for _, f := range s.Fields {
				sp.resolveFieldConstraints(nil, s.Fields, f, f.Type, specEntityFinder)
			}
		}
	}
	for _, deviceType := range sp.Spec.DeviceTypes {
		conditionFinder := newConditionFinder(deviceType, sp.Spec.BaseDeviceType, specEntityFinder)
		for _, er := range deviceType.ElementRequirements {
			sp.resolveFieldConstraintReferences(er.Cluster, conditionFinder, er, er.Constraint)
		}
		for _, dtr := range deviceType.DeviceTypeRequirements {
			sp.resolveFieldConstraintReferences(nil, conditionFinder, dtr, dtr.Constraint)
		}
	}
}

func (sp *Builder) noteConstraintResolutionFailures() {
	for exp, source := range sp.constraintFailures {
		switch exp := exp.(type) {
		case *constraint.IdentifierLimit:
			if exp.Entity == nil {
				slog.Error("Failed to resolve constraint identifier", "ref", exp.ID, log.Path("path", source))
			}
		case *constraint.ReferenceLimit:
			if exp.Entity == nil {
				slog.Error("Failed to resolve constraint reference", "ref", exp.Reference, log.Path("path", source))
			}
		default:
			slog.Warn("Unexpected failed constraint entity", log.Type("type", exp), log.Path("path", source))
		}
	}
}

func (sp *Builder) resolveFieldConstraints(cluster *matter.Cluster, fieldSet matter.FieldSet, field *matter.Field, dataType *types.DataType, finder entityFinder) {
	var fieldFinder entityFinder = newFieldFinder(fieldSet, finder)
	if dataType != nil && dataType.Entity != nil {
		dataTypeFinder := makeEntityFinder(dataType.Entity, fieldFinder)
		if dataTypeFinder != nil {
			fieldFinder = dataTypeFinder
		}
	}
	sp.resolveFieldConstraintReferences(cluster, fieldFinder, field, field.Constraint)
	sp.resolveFieldConstraintLimit(cluster, fieldFinder, field, field.Fallback)
}

func (sp *Builder) resolveFieldConstraintReferences(cluster *matter.Cluster, finder entityFinder, source log.Source, con constraint.Constraint) {
	switch con := con.(type) {
	case *constraint.ExactConstraint:
		sp.resolveFieldConstraintLimit(cluster, finder, source, con.Value)
	case *constraint.ListConstraint:
		sp.resolveFieldConstraintReferences(cluster, finder, source, con.Constraint)
		sp.resolveFieldConstraintReferences(cluster, finder, source, con.EntryConstraint)
	case *constraint.MaxConstraint:
		sp.resolveFieldConstraintLimit(cluster, finder, source, con.Maximum)
	case *constraint.MinConstraint:
		sp.resolveFieldConstraintLimit(cluster, finder, source, con.Minimum)
	case *constraint.RangeConstraint:
		sp.resolveFieldConstraintLimit(cluster, finder, source, con.Minimum)
		sp.resolveFieldConstraintLimit(cluster, finder, source, con.Maximum)
	case constraint.Set:
		for _, c := range con {
			sp.resolveFieldConstraintReferences(cluster, finder, source, c)
		}
	}
}

func (sp *Builder) resolveFieldConstraintLimit(cluster *matter.Cluster, finder entityFinder, source log.Source, l constraint.Limit) {
	switch l := l.(type) {
	case *constraint.CharacterLimit:
		sp.resolveFieldConstraintLimit(cluster, finder, source, l.ByteCount)
		sp.resolveFieldConstraintLimit(cluster, finder, source, l.CodepointCount)
	case *constraint.LengthLimit:
		sp.resolveFieldConstraintLimit(cluster, finder, source, l.Reference)
	case *constraint.IdentifierLimit:
		if l.Entity == nil {
			l.Entity = finder.findEntityByIdentifier(l.ID, source)
			if l.Entity == nil {
				sp.constraintFailures[l] = source
				slog.Error("failed to resolve constraint identifier", "ref", l.ID, log.Path("path", source))
			}
		}
		if l.Entity != nil && l.Field != nil {
			sp.resolveFieldConstraintLimit(cluster, makeEntityFinder(l.Entity, finder), source, l.Field)
		}
	case *constraint.MathExpressionLimit:
		sp.resolveFieldConstraintLimit(cluster, finder, source, l.Left)
		sp.resolveFieldConstraintLimit(cluster, finder, source, l.Right)
	case *constraint.ReferenceLimit:
		if l.Entity == nil {
			l.Entity = finder.findEntityByReference(l.Reference, l.Label, source)
			if l.Entity == nil {
				slog.Error("failed to resolve constraint reference", "ref", l.Reference, log.Path("path", source))
				sp.constraintFailures[l] = source
			}
		}
		if l.Entity != nil && l.Field != nil {
			sp.resolveFieldConstraintLimit(cluster, makeEntityFinder(l.Entity, finder), source, l.Field)
		}
	case *constraint.MinOfLimit:
		for _, l := range l.Minimums {
			sp.resolveFieldConstraintLimit(cluster, finder, source, l)
		}
	case *constraint.MaxOfLimit:
		for _, l := range l.Maximums {
			sp.resolveFieldConstraintLimit(cluster, finder, source, l)
		}
	case nil, *constraint.ManufacturerLimit,
		*constraint.IntLimit,
		*constraint.NullLimit,
		*constraint.StatusCodeLimit,
		*constraint.EmptyLimit,
		*constraint.StringLimit,
		*constraint.GenericLimit,
		*constraint.BooleanLimit,
		*constraint.TemperatureLimit,
		*constraint.UnspecifiedLimit,
		*constraint.ExpLimit,
		*constraint.HexLimit,
		*constraint.PercentLimit: // None of these limits have references to be resolved
	default:
		slog.Warn("Unexpected field constraint limit type", log.Type("type", l))
	}
}
