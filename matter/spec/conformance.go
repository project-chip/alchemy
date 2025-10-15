package spec

import (
	"log/slog"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/suggest"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

func (sp *Builder) resolveConformances() {
	specEntityFinder := newSpecEntityFinder(sp.Spec, nil, nil)
	for cluster := range sp.Spec.Clusters {
		specEntityFinder.cluster = cluster
		featureFinder := newFeatureFinder(cluster.Features, specEntityFinder)
		sp.resolveFeatureConformances(cluster, featureFinder)
		clusterFinder := newClusterEntityFinder(cluster, specEntityFinder)
		for _, a := range cluster.Attributes {
			sp.resolveFieldConformances(cluster, clusterFinder, cluster.Attributes, a, a.Type)
		}
		for _, s := range cluster.Structs {
			for _, f := range s.Fields {
				sp.resolveFieldConformances(cluster, clusterFinder, s.Fields, f, f.Type)
			}
		}
		sp.resolveBitmapConformances(cluster, clusterFinder)
		sp.resolveEnumConformances(cluster, clusterFinder)
		sp.resolveEventConformances(cluster, clusterFinder)
		sp.resolveCommandConformances(cluster, clusterFinder)
	}
	specEntityFinder.cluster = nil
	for o := range sp.Spec.GlobalObjects {
		switch o := o.(type) {
		case *matter.Struct:
			for _, f := range o.Fields {
				sp.resolveFieldConformances(nil, specEntityFinder, o.Fields, f, f.Type)
			}
		}
	}
	for _, deviceType := range sp.Spec.DeviceTypes {
		conditionFinder := newConditionFinder(deviceType, sp.Spec.BaseDeviceType, specEntityFinder)
		for _, cr := range deviceType.ClusterRequirements {
			var finder entityFinder = conditionFinder
			if cr.Cluster != nil {
				finder = newFeatureFinder(cr.Cluster.Features, conditionFinder)
			}
			sp.resolveEntityConformanceReferences(cr.Cluster, finder, cr, cr.Conformance)
		}
		for _, er := range deviceType.ElementRequirements {
			var finder entityFinder = conditionFinder
			if er.Cluster != nil {
				finder = newFeatureFinder(er.Cluster.Features, conditionFinder)
			}
			sp.resolveEntityConformanceReferences(er.Cluster, finder, er, er.Conformance)
		}
		for _, dtr := range deviceType.DeviceTypeRequirements {
			sp.resolveEntityConformanceReferences(nil, conditionFinder, dtr, dtr.Conformance)
		}
	}
}

func (sp *Builder) noteConformanceResolutionFailures(spec *Specification) {
	for exp, failure := range sp.conformanceFailures {
		switch exp := exp.(type) {
		case *conformance.ReferenceExpression:
			if exp.Entity == nil {
				if comparableEntity, ok := failure.source.(types.ComparableEntity); ok {
					slog.Error("Failed to resolve conformance expression reference", "ref", exp.Reference, log.Path("source", comparableEntity))
					spec.addError(&UnknownConformanceReferenceError{Entity: comparableEntity, Reference: exp.Reference})
				} else {
					slog.Error("Conformance failure source is not a ComparableEntity", log.Type("type", failure.source), log.Path("source", failure.source))
				}
			}
		case *conformance.IdentifierExpression:
			if exp.Entity == nil {
				if comparableEntity, ok := failure.source.(types.ComparableEntity); ok {
					slog.Error("Failed to resolve conformance expression identifier", "ref", exp.ID, log.Path("source", comparableEntity))
					spec.addError(&UnknownConformanceIdentifierError{Entity: comparableEntity, Identifier: exp.ID})
					suggestions := make(map[types.Entity]int)
					failure.finder.suggestIdentifiers(exp.ID, suggestions)
					suggest.ListPossibilities(exp.ID, suggestions)
				} else {
					slog.Error("Conformance failure source is not a ComparableEntity", log.Type("type", failure.source), log.Path("source", failure.source))
				}
			}
		case *conformance.IdentifierValue:
			if exp.Entity == nil {
				if comparableEntity, ok := failure.source.(types.ComparableEntity); ok {
					slog.Error("failed to resolve conformance value identifier", "id", exp.ID, log.Path("source", comparableEntity))
					spec.addError(&UnknownConformanceIdentifierError{Entity: comparableEntity, Identifier: exp.ID})
					suggestions := make(map[types.Entity]int)
					failure.finder.suggestIdentifiers(exp.ID, suggestions)
					suggest.ListPossibilities(exp.ID, suggestions)
				} else {
					slog.Error("Conformance failure source is not a ComparableEntity", log.Type("type", failure.source), log.Path("source", failure.source))
				}
			}
		case *conformance.ReferenceValue:
			if exp.Entity == nil {
				if comparableEntity, ok := failure.source.(types.ComparableEntity); ok {
					slog.Error("failed to resolve conformance value reference", "ref", exp.Reference, log.Path("source", comparableEntity))
					spec.addError(&UnknownConformanceReferenceError{Entity: comparableEntity, Reference: exp.Reference})
				} else {
					slog.Error("Conformance failure source is not a ComparableEntity", log.Type("type", failure.source), log.Path("source", failure.source))
				}
			}
		default:
			slog.Warn("Unexpected failed conformance entity", log.Type("type", exp), log.Path("source", failure.source))
		}
	}
}

func (sp *Builder) resolveFeatureConformances(cluster *matter.Cluster, finder *featureFinder) {
	if finder.features == nil {
		return
	}
	for feature := range finder.features.FeatureBits() {
		sp.resolveEntityConformanceReferences(cluster, finder, feature, feature.Conformance())
	}
}

func (sp *Builder) resolveFieldConformances(cluster *matter.Cluster, finder entityFinder, fieldSet matter.FieldSet, field *matter.Field, dataType *types.DataType) {

	fieldFinder := newFieldFinder(fieldSet, finder)

	sp.resolveEntityConformanceReferences(cluster, fieldFinder, field, field.Conformance)
}

func (sp *Builder) resolveBitmapConformances(cluster *matter.Cluster, finder entityFinder) {
	for _, bm := range cluster.Bitmaps {
		bitmapValueFinder := makeEntityFinder(bm, finder)
		for _, bmv := range bm.Bits {
			bitmapValueFinder.setIdentity(bmv)
			sp.resolveEntityConformanceReferences(cluster, bitmapValueFinder, bm, bmv.Conformance())
		}
	}
}

func (sp *Builder) resolveEnumConformances(cluster *matter.Cluster, finder entityFinder) {
	for _, e := range cluster.Enums {
		enumValueFinder := makeEntityFinder(e, finder)
		for _, ev := range e.Values {
			enumValueFinder.setIdentity(ev)
			sp.resolveEntityConformanceReferences(cluster, enumValueFinder, e, ev.Conformance)
		}
	}
}

func (sp *Builder) resolveCommandConformances(cluster *matter.Cluster, finder entityFinder) {
	commandFinder := newCommandFinder(cluster.Commands, finder)

	for _, command := range cluster.Commands {
		commandFinder.setIdentity(command)
		sp.resolveEntityConformanceReferences(cluster, commandFinder, command, command.Conformance)
		for _, field := range command.Fields {
			fieldFinder := newFieldFinder(command.Fields, finder)
			fieldFinder.setIdentity(field)
			sp.resolveFieldConformances(cluster, fieldFinder, command.Fields, field, field.Type)
		}
	}
}

func (sp *Builder) resolveEventConformances(cluster *matter.Cluster, finder entityFinder) {

	eventFinder := newEventFinder(cluster.Events, finder)

	for _, event := range cluster.Events {
		eventFinder.setIdentity(event)
		sp.resolveEntityConformanceReferences(cluster, eventFinder, event, event.Conformance)
		for _, field := range event.Fields {
			fieldFinder := newFieldFinder(event.Fields, finder)
			fieldFinder.setIdentity(field)
			sp.resolveFieldConformances(cluster, fieldFinder, event.Fields, field, field.Type)
		}
	}
}

func (sp *Builder) resolveEntityConformanceReferences(cluster *matter.Cluster, finder entityFinder, source types.Entity, con conformance.Conformance) {
	switch con := con.(type) {
	case *conformance.Mandatory:
		sp.resolveEntityConformanceExpressionReferences(cluster, finder, source, con.Expression)
	case *conformance.Optional:
		sp.resolveEntityConformanceExpressionReferences(cluster, finder, source, con.Expression)
	case conformance.Set:
		for _, c := range con {
			sp.resolveEntityConformanceReferences(cluster, finder, source, c)
		}
	case *conformance.Disallowed, *conformance.Provisional, *conformance.Described, *conformance.Deprecated:
	case *conformance.Generic:
		if !conformance.IsBlank(con) {
			slog.Warn("Can not resolve entities on generic conformance", slog.String("conformance", con.RawText()), log.Path("source", source))

		}
	default:
		slog.Warn("Unexpected field conformance type", log.Type("type", con))
	}
}

func (sp *Builder) resolveEntityConformanceExpressionReferences(cluster *matter.Cluster, finder entityFinder, source types.Entity, exp conformance.Expression) (resolved types.Entity, failed bool) {
	switch exp := exp.(type) {
	case *conformance.ReferenceExpression:
		if exp.Entity == nil {
			exp.Entity = finder.findEntityByReference(exp.Reference, exp.Label, source)
			if exp.Entity == nil {
				//slog.Warn("failed to resolve conformance expression reference", "ref", exp.Reference, log.Path("source", source))
				sp.conformanceFailures[exp] = referenceFailure{source: source, finder: finder}
				failed = true
				return
			}
			resolved = exp.Entity
		} else {
			resolved = exp.Entity
		}
		if exp.Field != nil {
			fieldResolved, fieldFailed := sp.resolveEntityConformanceValueReferences(cluster, makeEntityFinder(exp.Entity, finder), source, exp.Field)
			if !fieldFailed && fieldResolved != nil {
				resolved = fieldResolved
				failed = fieldFailed
			}
		}
	case *conformance.IdentifierExpression:
		if exp.Entity == nil {
			exp.Entity = finder.findEntityByIdentifier(exp.ID, source)
			if exp.Entity == nil {
				//slog.Warn("failed to resolve conformance expression identifier", "ref", exp.ID, log.Path("source", source))
				sp.conformanceFailures[exp] = referenceFailure{source: source, finder: finder}
				failed = true
				return
			}
			resolved = exp.Entity
		} else {
			resolved = exp.Entity
		}
		if exp.Field != nil {
			fieldResolved, fieldFailed := sp.resolveEntityConformanceValueReferences(cluster, makeEntityFinder(exp.Entity, finder), source, exp.Field)
			if !fieldFailed && fieldResolved != nil {
				resolved = fieldResolved
				failed = fieldFailed
			}
		}
	case *conformance.EqualityExpression:
		leftResolved, leftFailed := sp.resolveEntityConformanceExpressionReferences(cluster, finder, source, exp.Left)
		rightResolved, rightFailed := sp.resolveEntityConformanceExpressionReferences(cluster, finder, source, exp.Right)
		if leftFailed && rightResolved != nil {
			sp.resolveEntityConformanceExpressionReferences(cluster, makeEntityFinder(rightResolved, finder), source, exp.Left)
		}
		if rightFailed && leftResolved != nil {
			sp.resolveEntityConformanceExpressionReferences(cluster, makeEntityFinder(leftResolved, finder), source, exp.Right)
		}
	case *conformance.LogicalExpression:
		sp.resolveEntityConformanceExpressionReferences(cluster, finder, source, exp.Left)
		for _, re := range exp.Right {
			sp.resolveEntityConformanceExpressionReferences(cluster, finder, source, re)
		}
	case *conformance.ComparisonExpression:
		leftResolved, leftFailed := sp.resolveEntityConformanceValueReferences(cluster, finder, source, exp.Left)
		rightResolved, rightFailed := sp.resolveEntityConformanceValueReferences(cluster, finder, source, exp.Right)
		if leftFailed && rightResolved != nil {
			sp.resolveEntityConformanceValueReferences(cluster, makeEntityFinder(rightResolved, finder), source, exp.Left)
		}
		if rightFailed && leftResolved != nil {
			sp.resolveEntityConformanceValueReferences(cluster, makeEntityFinder(leftResolved, finder), source, exp.Right)
		}
	}
	return
}

func (sp *Builder) resolveEntityConformanceValueReferences(cluster *matter.Cluster, finder entityFinder, source types.Entity, cv conformance.ComparisonValue) (resolved types.Entity, failed bool) {
	switch cv := cv.(type) {
	case *conformance.IdentifierValue:
		if cv.Entity == nil {
			cv.Entity = finder.findEntityByIdentifier(cv.ID, source)
			if cv.Entity == nil {
				sp.conformanceFailures[cv] = referenceFailure{source: source, finder: finder}
				failed = true
				return
			}
			resolved = cv.Entity
			if cv.Entity != nil && cv.Field != nil {
				resolved, failed = sp.resolveEntityConformanceValueReferences(cluster, finder, source, cv.Field)
			}
		} else {
			resolved = cv.Entity
		}
	case *conformance.ReferenceValue:
		if cv.Entity == nil {
			cv.Entity = finder.findEntityByReference(cv.Reference, cv.Label, source)
			if cv.Entity == nil {
				sp.conformanceFailures[cv] = referenceFailure{source: source, finder: finder}

				//slog.Warn("failed to resolve conformance value reference", "ref", cv.Reference, log.Path("source", source), slog.Any("entity", cv.Entity))
				failed = true
				return
			}
			resolved = cv.Entity
			if cv.Entity != nil && cv.Field != nil {
				resolved, failed = sp.resolveEntityConformanceValueReferences(cluster, makeEntityFinder(cv.Entity, finder), source, cv.Field)
			}
		} else {
			resolved = cv.Entity
		}
	case *conformance.MathOperation:
		sp.resolveEntityConformanceValueReferences(cluster, finder, source, cv.Left)
		sp.resolveEntityConformanceValueReferences(cluster, finder, source, cv.Right)
	}
	return
}
