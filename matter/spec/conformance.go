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
			resolveFieldConformanceReferences(spec, cr.Cluster, conditionFinder, cr, nil, cr.Conformance)
		}
		for _, er := range deviceType.ElementRequirements {
			resolveFieldConformanceReferences(spec, er.Cluster, conditionFinder, er, nil, er.Conformance)
		}
		for _, dtr := range deviceType.DeviceTypeRequirements {
			resolveFieldConformanceReferences(spec, nil, conditionFinder, dtr, nil, dtr.Conformance)
		}
	}

}

func (sp *Builder) resolveFeatureConformances(spec *Specification, cluster *matter.Cluster) {
	if cluster.Features == nil {
		return
	}
	featureFinder := func(identifier string) types.Entity {
		for f := range cluster.Features.FeatureBits() {
			if strings.EqualFold(f.Code, identifier) {
				return f
			}
		}
		return nil
	}
	for feature := range cluster.Features.FeatureBits() {
		resolveFieldConformanceReferences(spec, cluster, featureFinder, feature, feature, feature.Conformance())
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
	resolveFieldConformanceReferences(spec, cluster, fieldFinder, field, fieldDataTypeEntity, field.Conformance)
}

func (sp *Builder) resolveBitmapConformances(spec *Specification, cluster *matter.Cluster) {
	for _, bm := range cluster.Bitmaps {
		bitmapValueFinder := makeFinder(bm)
		for _, bmv := range bm.Bits {
			resolveFieldConformanceReferences(spec, cluster, bitmapValueFinder, bmv, bm, bmv.Conformance())
		}
	}
}

func (sp *Builder) resolveEnumConformances(spec *Specification, cluster *matter.Cluster) {
	for _, e := range cluster.Enums {
		enumValueFinder := makeFinder(e)
		for _, ev := range e.Values {
			resolveFieldConformanceReferences(spec, cluster, enumValueFinder, ev, e, ev.Conformance)
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
		resolveFieldConformanceReferences(spec, cluster, commandFinder, command, command, command.Conformance)
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
		resolveFieldConformanceReferences(spec, cluster, eventFinder, event, event, event.Conformance)
		for _, field := range event.Fields {
			sp.resolveFieldConformances(spec, cluster, event.Fields, field, field.Type)
		}
	}
}

func resolveFieldConformanceReferences(spec *Specification, cluster *matter.Cluster, finder findNamedEntity, source log.Source, entity types.Entity, con conformance.Conformance) {
	switch con := con.(type) {
	case *conformance.Mandatory:
		resolveFieldConformanceExpressionReferences(spec, cluster, finder, source, entity, con.Expression)
	case *conformance.Optional:
		resolveFieldConformanceExpressionReferences(spec, cluster, finder, source, entity, con.Expression)
	case conformance.Set:
		for _, c := range con {
			resolveFieldConformanceReferences(spec, cluster, finder, source, entity, c)
		}
	case *conformance.Disallowed, *conformance.Provisional, *conformance.Described, *conformance.Deprecated:
	default:
		slog.Warn("Unexpected field conformance type", log.Type("type", con))
	}
}

func resolveFieldConformanceExpressionReferences(spec *Specification, cluster *matter.Cluster, finder findNamedEntity, source log.Source, entity types.Entity, exp conformance.Expression) {
	switch exp := exp.(type) {
	case *conformance.ReferenceExpression:
		if exp.Entity == nil {
			exp.Entity = getCustomDataTypeFromReference(spec, cluster, exp.Reference, exp.Label)
			if exp.Entity == nil {
				slog.Warn("failed to resolve conformance expression reference", "ref", exp.Reference, log.Path("path", source), slog.Any("entity", exp.Entity))
			}

		}
		if exp.Field != nil && exp.Entity != nil {
			resolveFieldConformanceValueReferences(spec, cluster, makeFinder(exp.Entity), source, exp.Entity, exp.Field)
		}
	case *conformance.IdentifierExpression:
		if exp.Entity == nil {
			exp.Entity = findEntityForFieldIdentifier(spec, cluster, finder, source, entity, exp.ID)
			if exp.Entity == nil {
				slog.Warn("failed to resolve conformance expression identifier", "ref", exp.ID, log.Path("path", source), slog.Any("entity", exp.Entity))
			}
		}
		if exp.Field != nil && exp.Entity != nil {
			resolveFieldConformanceValueReferences(spec, cluster, makeFinder(exp.Entity), source, exp.Entity, exp.Field)
		}
	case *conformance.EqualityExpression:
		resolveFieldConformanceExpressionReferences(spec, cluster, finder, source, entity, exp.Left)
		resolveFieldConformanceExpressionReferences(spec, cluster, finder, source, entity, exp.Right)
	case *conformance.LogicalExpression:
		resolveFieldConformanceExpressionReferences(spec, cluster, finder, source, entity, exp.Left)
		for _, re := range exp.Right {
			resolveFieldConformanceExpressionReferences(spec, cluster, finder, source, entity, re)
		}
	case *conformance.FeatureExpression:
		if exp.Entity == nil {
			exp.Entity = findEntityForFieldIdentifier(spec, cluster, finder, source, entity, exp.Feature)
		}
	case *conformance.ComparisonExpression:
		resolveFieldConformanceValueReferences(spec, cluster, finder, source, entity, exp.Left)
		resolveFieldConformanceValueReferences(spec, cluster, finder, source, entity, exp.Right)

	}
}

func resolveFieldConformanceValueReferences(spec *Specification, cluster *matter.Cluster, finder findNamedEntity, source log.Source, entity types.Entity, cv conformance.ComparisonValue) {
	switch cv := cv.(type) {
	case *conformance.IdentifierValue:
		if cv.Entity == nil {
			cv.Entity = findEntityForFieldIdentifier(spec, cluster, finder, source, entity, cv.ID)
			if cv.Entity != nil && cv.Field != nil {
				resolveFieldConformanceValueReferences(spec, cluster, finder, source, cv.Entity, cv.Field)
			}
		}
	case *conformance.ReferenceValue:
		if cv.Entity == nil {
			cv.Entity = getCustomDataTypeFromReference(spec, cluster, cv.Reference, cv.Label)
			if cv.Entity == nil {
				slog.Warn("failed to resolve conformance value reference", "ref", cv.Reference, log.Path("path", source), slog.Any("entity", cv.Entity))
			}
			if cv.Entity != nil && cv.Field != nil {
				resolveFieldConformanceValueReferences(spec, cluster, finder, source, cv.Entity, cv.Field)
			}
		}
	case *conformance.MathOperation:
		resolveFieldConformanceValueReferences(spec, cluster, finder, source, entity, cv.Left)
		resolveFieldConformanceValueReferences(spec, cluster, finder, source, entity, cv.Right)
	}
}
