package spec

import (
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
			sp.resolveFieldConformances(clusterFinder, cluster.Attributes, a, a.Type)
		}

		sp.resolveBitmapConformances(cluster.Bitmaps, clusterFinder)
		sp.resolveEnumConformances(cluster.Enums, clusterFinder)
		sp.resolveStructConformances(cluster.Structs, clusterFinder)
		sp.resolveEventConformances(cluster.Events, clusterFinder)
		sp.resolveCommandConformances(cluster.Commands, clusterFinder)
	}
	specEntityFinder.cluster = nil

	sp.resolveBitmapConformances(sp.Spec.GlobalObjects.OfType[*matter.Bitmap](), specEntityFinder)
	sp.resolveEnumConformances(sp.Spec.GlobalObjects.OfType[*matter.Enum](), specEntityFinder)
	sp.resolveStructConformances(sp.Spec.GlobalObjects.OfType[*matter.Struct](), specEntityFinder)
	sp.resolveEventConformances(sp.Spec.GlobalObjects.OfType[*matter.Event](), specEntityFinder)
	sp.resolveCommandConformances(sp.Spec.GlobalObjects.OfType[*matter.Command](), specEntityFinder)

	for _, deviceType := range sp.Spec.DeviceTypes {
		conditionFinder := newConditionFinder(deviceType, sp.Spec.BaseDeviceType, specEntityFinder)
		for _, cr := range deviceType.ClusterRequirements {
			var finder entityFinder = conditionFinder
			if cr.Cluster != nil {
				finder = newFeatureFinder(cr.Cluster.Features, conditionFinder)
			}
			sp.resolveEntityConformance(finder, cr, cr.Conformance)
		}
		for _, er := range deviceType.ElementRequirements {
			var finder entityFinder = conditionFinder
			if er.Cluster != nil {
				finder = newFeatureFinder(er.Cluster.Features, conditionFinder)
			}
			sp.resolveEntityConformance(finder, er, er.Conformance)
		}
		for _, er := range deviceType.ConditionRequirements {
			sp.resolveEntityConformance(conditionFinder, er, er.Conformance)
		}
		for _, tr := range deviceType.TagRequirements {
			sp.resolveEntityConformance(conditionFinder, tr, tr.Conformance)
		}
		for _, dtr := range deviceType.DeviceTypeRequirements {
			sp.resolveEntityConformance(conditionFinder, dtr, dtr.Conformance)
		}
		for _, dtcr := range deviceType.ComposedDeviceTypeClusterRequirements {
			var finder entityFinder = conditionFinder
			if dtcr.DeviceType != nil {
				finder = newConditionFinder(dtcr.DeviceType, sp.Spec.BaseDeviceType, finder)
			}
			if dtcr.Cluster != nil {
				finder = newFeatureFinder(dtcr.Cluster.Features, finder)
			}
			sp.resolveEntityConformance(finder, dtcr, dtcr.Conformance)
		}
		for _, dter := range deviceType.ComposedDeviceTypeElementRequirements {
			var finder entityFinder = conditionFinder
			if dter.DeviceType != nil {
				finder = newConditionFinder(dter.DeviceType, sp.Spec.BaseDeviceType, finder)
			}
			if dter.ElementRequirement.Cluster != nil {
				finder = newFeatureFinder(dter.ElementRequirement.Cluster.Features, finder)
			}
			sp.resolveEntityConformance(finder, dter.ElementRequirement, dter.ElementRequirement.Conformance)
		}
		for _, dttr := range deviceType.ComposedDeviceTagRequirements {
			var finder entityFinder = conditionFinder
			if dttr.DeviceType != nil {
				finder = newConditionFinder(dttr.DeviceType, sp.Spec.BaseDeviceType, finder)
			}
			sp.resolveEntityConformance(finder, dttr.TagRequirement, dttr.TagRequirement.Conformance)
		}
	}
}

func (sp *Builder) resolveFeatureConformances(cluster *matter.Cluster, finder *featureFinder) {
	if cluster.Features == nil {
		return
	}
	for feature := range finder.features.FeatureBits() {
		sp.resolveEntityConformance(finder, feature, feature.Conformance())
	}
}

func (sp *Builder) resolveFieldConformances(finder entityFinder, fieldSet matter.FieldSet, field *matter.Field, dataType *types.DataType) {

	fieldFinder := newFieldFinder(fieldSet, finder)

	sp.resolveEntityConformance(fieldFinder, field, field.Conformance)
}

func (sp *Builder) resolveBitmapConformances(bitmaps matter.BitmapSet, finder entityFinder) {
	for _, bm := range bitmaps {
		bitmapValueFinder := makeEntityFinder(bm, finder)
		for _, bmv := range bm.Bits {
			bitmapValueFinder.setIdentity(bmv)
			sp.resolveEntityConformance(bitmapValueFinder, bmv, bmv.Conformance())
		}
	}
}

func (sp *Builder) resolveEnumConformances(enums matter.EnumSet, finder entityFinder) {
	for _, e := range enums {
		enumValueFinder := makeEntityFinder(e, finder)
		for _, ev := range e.Values {
			enumValueFinder.setIdentity(ev)
			sp.resolveEntityConformance(enumValueFinder, ev, ev.Conformance)
		}
	}
}

func (sp *Builder) resolveStructConformances(structs matter.StructSet, finder entityFinder) {
	for _, s := range structs {
		for _, f := range s.Fields {
			sp.resolveFieldConformances(finder, s.Fields, f, f.Type)
		}
	}
}

func (sp *Builder) resolveCommandConformances(commands []*matter.Command, finder entityFinder) {
	commandFinder := newCommandFinder(commands, finder)

	for _, command := range commands {
		commandFinder.setIdentity(command)
		sp.resolveEntityConformance(commandFinder, command, command.Conformance)
		for _, field := range command.Fields {
			fieldFinder := newFieldFinder(command.Fields, commandFinder)
			fieldFinder.setIdentity(field)
			sp.resolveFieldConformances(fieldFinder, command.Fields, field, field.Type)
		}
	}
}

func (sp *Builder) resolveEventConformances(events matter.EventSet, finder entityFinder) {

	eventFinder := newEventFinder(events, finder)

	for _, event := range events {
		eventFinder.setIdentity(event)
		for _, exp := range conformance.TraverseExpressions(event.Conformance) {
			sp.resolveEntityConformanceExpressionReferences(eventFinder, event, exp)
		}
		for _, field := range event.Fields {
			fieldFinder := newFieldFinder(event.Fields, finder)
			fieldFinder.setIdentity(field)
			sp.resolveFieldConformances(fieldFinder, event.Fields, field, field.Type)
		}
	}
}

func (sp *Builder) resolveEntityConformance(finder entityFinder, source types.Entity, con conformance.Conformance) {
	for _, exp := range conformance.TraverseExpressions(con) {
		sp.resolveEntityConformanceExpressionReferences(finder, source, exp)
	}
}

func (sp *Builder) resolveEntityConformanceExpressionReferences(finder entityFinder, source types.Entity, exp conformance.Expression) (resolved types.Entity, failed bool) {
	switch exp := exp.(type) {
	case *conformance.ReferenceExpression:
		if exp.Entity == nil {
			exp.Entity = finder.findEntityByReference(exp.Reference, exp.Label, source)
			if exp.Entity == nil {
				failed = true
				return
			}
			resolved = exp.Entity
		} else {
			resolved = exp.Entity
		}
		if exp.Field != nil {
			fieldResolved, fieldFailed := sp.resolveEntityConformanceValueReferences(makeEntityFinder(exp.Entity, finder), source, exp.Field)
			if !fieldFailed && fieldResolved != nil {
				resolved = fieldResolved
				failed = fieldFailed
			}
		}
	case *conformance.IdentifierExpression:
		if exp.Entity == nil {
			exp.Entity = finder.findEntityByIdentifier(exp.ID, source)
			if exp.Entity == nil {
				failed = true
				return
			}
			resolved = exp.Entity
		} else {
			resolved = exp.Entity
		}
		if exp.Field != nil {
			fieldResolved, fieldFailed := sp.resolveEntityConformanceValueReferences(makeEntityFinder(exp.Entity, finder), source, exp.Field)
			if !fieldFailed && fieldResolved != nil {
				resolved = fieldResolved
				failed = fieldFailed
			}
		}
	case *conformance.EqualityExpression:
		leftResolved, leftFailed := sp.resolveEntityConformanceExpressionReferences(finder, source, exp.Left)
		rightResolved, rightFailed := sp.resolveEntityConformanceExpressionReferences(finder, source, exp.Right)
		if leftFailed && rightResolved != nil {
			sp.resolveEntityConformanceExpressionReferences(makeEntityFinder(rightResolved, finder), source, exp.Left)
		}
		if rightFailed && leftResolved != nil {
			sp.resolveEntityConformanceExpressionReferences(makeEntityFinder(leftResolved, finder), source, exp.Right)
		}
	case *conformance.LogicalExpression:
		sp.resolveEntityConformanceExpressionReferences(finder, source, exp.Left)
		for _, re := range exp.Right {
			sp.resolveEntityConformanceExpressionReferences(finder, source, re)
		}
	case *conformance.ComparisonExpression:
		leftResolved, leftFailed := sp.resolveEntityConformanceValueReferences(finder, source, exp.Left)
		rightResolved, rightFailed := sp.resolveEntityConformanceValueReferences(finder, source, exp.Right)
		if leftFailed && rightResolved != nil {
			sp.resolveEntityConformanceValueReferences(makeEntityFinder(rightResolved, finder), source, exp.Left)
		}
		if rightFailed && leftResolved != nil {
			sp.resolveEntityConformanceValueReferences(makeEntityFinder(leftResolved, finder), source, exp.Right)
		}
	}
	return
}

func (sp *Builder) resolveEntityConformanceValueReferences(finder entityFinder, source types.Entity, cv conformance.ComparisonValue) (resolved types.Entity, failed bool) {
	for v := range conformance.TraverseComparisonValue(cv) {
		switch v := v.(type) {
		case *conformance.IdentifierValue:
			if v.Entity == nil {
				v.Entity = finder.findEntityByIdentifier(v.ID, source)
				if v.Entity == nil {
					failed = true
					continue
				}
				resolved = v.Entity
				if v.Entity != nil && v.Field != nil {
					resolved, failed = sp.resolveEntityConformanceValueReferences(finder, source, v.Field)
				}
			} else {
				resolved = v.Entity
			}
		case *conformance.ReferenceValue:
			if v.Entity == nil {
				v.Entity = finder.findEntityByReference(v.Reference, v.Label, source)
				if v.Entity == nil {
					failed = true
					continue
				}
				resolved = v.Entity
				if v.Entity != nil && v.Field != nil {
					resolved, failed = sp.resolveEntityConformanceValueReferences(makeEntityFinder(v.Entity, finder), source, v.Field)
				}
			} else {
				resolved = v.Entity
			}
		}
	}

	return
}
