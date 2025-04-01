package spec

import (
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

func (sp *Builder) resolveConformances(spec *Specification) {
	for cluster := range spec.Clusters {
		sp.resolveFeatureConformances(spec, cluster)
		for _, a := range cluster.Attributes {
			sp.resolveFieldConformances(spec, cluster, cluster.Attributes, a, a.Type)
		}
		for _, s := range cluster.Structs {
			for _, f := range s.Fields {
				sp.resolveFieldConformances(spec, cluster, s.Fields, f, f.Type)
			}
		}
		sp.resolveBitmapConformances(spec, cluster)
		sp.resolveEnumConformances(spec, cluster)
		sp.resolveEventConformances(spec, cluster)
		sp.resolveCommandConformances(spec, cluster)
	}
	for _, s := range spec.structIndex {
		for _, f := range s.Fields {
			sp.resolveFieldConformances(spec, nil, s.Fields, f, f.Type)
		}
	}
	for _, deviceType := range spec.DeviceTypes {
		conditionFinder := func(identifier string) types.Entity {
			for _, con := range deviceType.Conditions {
				if strings.EqualFold(con.Feature, identifier) {
					return con
				}
			}
			if spec.BaseDeviceType != nil {
				for _, con := range spec.BaseDeviceType.Conditions {
					if strings.EqualFold(con.Feature, identifier) {
						return con
					}
				}
			}
			return nil
		}
		for _, cr := range deviceType.ClusterRequirements {
			resolveEntityConformanceReferences(spec, cr.Cluster, conditionFinder, cr, nil, cr.Conformance)
		}
		for _, er := range deviceType.ElementRequirements {
			resolveEntityConformanceReferences(spec, er.Cluster, conditionFinder, er, nil, er.Conformance)
		}
		for _, dtr := range deviceType.DeviceTypeRequirements {
			resolveEntityConformanceReferences(spec, nil, conditionFinder, dtr, nil, dtr.Conformance)
		}
	}

}

func (sp *Builder) resolveFeatureConformances(spec *Specification, cluster *matter.Cluster) {
	if cluster.Features == nil {
		return
	}
	featureFinder := func(identifier string) types.Entity {
		for f := range cluster.Features.FeatureBits() {
			if f.Code == identifier {
				return f
			}
		}
		return nil
	}
	for feature := range cluster.Features.FeatureBits() {
		resolveEntityConformanceReferences(spec, cluster, featureFinder, feature, feature, feature.Conformance())
	}
}

func (sp *Builder) resolveFieldConformances(spec *Specification, cluster *matter.Cluster, fieldSet matter.FieldSet, field *matter.Field, dataType *types.DataType) {

	fieldFinder := func(identifier string) types.Entity {
		for _, of := range fieldSet {
			if strings.EqualFold(of.Name, identifier) {
				return of
			}
		}
		return nil
	}
	var fieldDataTypeEntity types.Entity
	if dataType != nil {
		fieldDataTypeEntity = dataType.Entity
	}
	resolveEntityConformanceReferences(spec, cluster, fieldFinder, field, fieldDataTypeEntity, field.Conformance)
}

func (sp *Builder) resolveBitmapConformances(spec *Specification, cluster *matter.Cluster) {
	for _, bm := range cluster.Bitmaps {
		bitmapValueFinder := makeEntityFinder(bm)
		for _, bmv := range bm.Bits {
			resolveEntityConformanceReferences(spec, cluster, bitmapValueFinder, bmv, bm, bmv.Conformance())
		}
	}
}

func (sp *Builder) resolveEnumConformances(spec *Specification, cluster *matter.Cluster) {
	for _, e := range cluster.Enums {
		enumValueFinder := makeEntityFinder(e)
		for _, ev := range e.Values {
			resolveEntityConformanceReferences(spec, cluster, enumValueFinder, ev, e, ev.Conformance)
		}
	}
}

func (sp *Builder) resolveCommandConformances(spec *Specification, cluster *matter.Cluster) {
	commandFinder := func(identifier string) types.Entity {
		for _, c := range cluster.Commands {
			if strings.EqualFold(c.Name, identifier) {
				return c
			}
		}
		return nil
	}
	for _, command := range cluster.Commands {
		resolveEntityConformanceReferences(spec, cluster, commandFinder, command, command, command.Conformance)
		for _, field := range command.Fields {
			sp.resolveFieldConformances(spec, cluster, command.Fields, field, field.Type)
		}
	}
}

func (sp *Builder) resolveEventConformances(spec *Specification, cluster *matter.Cluster) {
	eventFinder := func(identifier string) types.Entity {
		for _, e := range cluster.Events {
			if strings.EqualFold(e.Name, identifier) {
				return e
			}
		}
		return nil
	}
	for _, event := range cluster.Events {
		resolveEntityConformanceReferences(spec, cluster, eventFinder, event, event, event.Conformance)
		for _, field := range event.Fields {
			sp.resolveFieldConformances(spec, cluster, event.Fields, field, field.Type)
		}
	}
}

func resolveEntityConformanceReferences(spec *Specification, cluster *matter.Cluster, finder findNamedEntity, source log.Source, entity types.Entity, con conformance.Conformance) {
	switch con := con.(type) {
	case *conformance.Mandatory:
		resolveEntityConformanceExpressionReferences(spec, cluster, finder, source, entity, con.Expression)
	case *conformance.Optional:
		resolveEntityConformanceExpressionReferences(spec, cluster, finder, source, entity, con.Expression)
	case conformance.Set:
		for _, c := range con {
			resolveEntityConformanceReferences(spec, cluster, finder, source, entity, c)
		}
	case *conformance.Disallowed, *conformance.Provisional, *conformance.Described, *conformance.Deprecated:
	default:
		slog.Warn("Unexpected field conformance type", log.Type("type", con))
	}
}

func resolveEntityConformanceExpressionReferences(spec *Specification, cluster *matter.Cluster, finder findNamedEntity, source log.Source, entity types.Entity, exp conformance.Expression) {
	switch exp := exp.(type) {
	case *conformance.ReferenceExpression:
		if exp.Entity == nil {
			exp.Entity = getCustomDataTypeFromReference(spec, cluster, exp.Reference, exp.Label)
			if exp.Entity == nil {
				slog.Warn("failed to resolve conformance expression reference", "ref", exp.Reference, log.Path("path", source), matter.LogEntity("entity", exp.Entity))
			}

		}
		if exp.Field != nil && exp.Entity != nil {
			resolveEntityConformanceValueReferences(spec, cluster, makeEntityFinder(exp.Entity), source, exp.Entity, exp.Field)
		}
	case *conformance.IdentifierExpression:
		if exp.Entity == nil {
			exp.Entity = findEntityForEntityIdentifier(spec, cluster, finder, source, entity, exp.ID)
			if exp.Entity == nil {
				slog.Warn("failed to resolve conformance expression identifier", "ref", exp.ID, log.Path("path", source), matter.LogEntity("entity", exp.Entity))
			}
		}
		if exp.Field != nil && exp.Entity != nil {
			resolveEntityConformanceValueReferences(spec, cluster, makeEntityFinder(exp.Entity), source, exp.Entity, exp.Field)
		}
	case *conformance.EqualityExpression:
		resolveEntityConformanceExpressionReferences(spec, cluster, finder, source, entity, exp.Left)
		resolveEntityConformanceExpressionReferences(spec, cluster, finder, source, entity, exp.Right)
	case *conformance.LogicalExpression:
		resolveEntityConformanceExpressionReferences(spec, cluster, finder, source, entity, exp.Left)
		for _, re := range exp.Right {
			resolveEntityConformanceExpressionReferences(spec, cluster, finder, source, entity, re)
		}
	case *conformance.FeatureExpression:
		if exp.Entity == nil {
			exp.Entity = findEntityForEntityIdentifier(spec, cluster, finder, source, entity, exp.Feature)
		}
	case *conformance.ComparisonExpression:
		resolveEntityConformanceValueReferences(spec, cluster, finder, source, entity, exp.Left)
		resolveEntityConformanceValueReferences(spec, cluster, finder, source, entity, exp.Right)

	}
}

func resolveEntityConformanceValueReferences(spec *Specification, cluster *matter.Cluster, finder findNamedEntity, source log.Source, entity types.Entity, cv conformance.ComparisonValue) {
	switch cv := cv.(type) {
	case *conformance.IdentifierValue:
		if cv.Entity == nil {
			cv.Entity = findEntityForEntityIdentifier(spec, cluster, finder, source, entity, cv.ID)
			if cv.Entity != nil && cv.Field != nil {
				resolveEntityConformanceValueReferences(spec, cluster, finder, source, cv.Entity, cv.Field)
			}
		}
	case *conformance.ReferenceValue:
		if cv.Entity == nil {
			cv.Entity = getCustomDataTypeFromReference(spec, cluster, cv.Reference, cv.Label)
			if cv.Entity == nil {
				slog.Warn("failed to resolve conformance value reference", "ref", cv.Reference, log.Path("path", source), slog.Any("entity", cv.Entity))
			}
			if cv.Entity != nil && cv.Field != nil {
				resolveEntityConformanceValueReferences(spec, cluster, finder, source, cv.Entity, cv.Field)
			}
		}
	case *conformance.MathOperation:
		resolveEntityConformanceValueReferences(spec, cluster, finder, source, entity, cv.Left)
		resolveEntityConformanceValueReferences(spec, cluster, finder, source, entity, cv.Right)
	}
}
