package spec

import (
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/types"
)

func (sp *Builder) resolveConstraints() {
	specEntityFinder := newSpecEntityFinder(sp.Spec, nil, nil)
	for cluster := range sp.Spec.Clusters {
		specEntityFinder.cluster = cluster
		for _, a := range cluster.Attributes {
			sp.resolveFieldConstraints(cluster.Attributes, a, a.Type, specEntityFinder)
		}
		attributeFinder := newFieldFinder(cluster.Attributes, specEntityFinder)
		for _, s := range cluster.Structs {
			structFinder := newFieldFinder(s.Fields, attributeFinder)
			for _, f := range s.Fields {
				sp.resolveFieldConstraints(s.Fields, f, f.Type, structFinder)
			}
		}
		commandFinder := newCommandFinder(cluster.Commands, attributeFinder)
		for _, command := range cluster.Commands {
			for _, f := range command.Fields {
				sp.resolveFieldConstraints(command.Fields, f, f.Type, commandFinder)
			}
		}
		eventFinder := newEventFinder(cluster.Events, attributeFinder)
		for _, event := range cluster.Events {
			for _, f := range event.Fields {
				sp.resolveFieldConstraints(event.Fields, f, f.Type, eventFinder)
			}
		}
	}
	specEntityFinder.cluster = nil
	for o := range sp.Spec.GlobalObjects {
		switch o := o.(type) {
		case *matter.Struct:
			for _, f := range o.Fields {
				sp.resolveFieldConstraints(o.Fields, f, f.Type, specEntityFinder)
			}
		}
	}
	for _, deviceType := range sp.Spec.DeviceTypes {
		conditionFinder := newConditionFinder(deviceType, sp.Spec.BaseDeviceType, specEntityFinder)
		for _, er := range deviceType.ElementRequirements {
			sp.resolveElementRequirementConstraint(er, conditionFinder)
		}
		for _, tr := range deviceType.TagRequirements {
			sp.resolveFieldConstraintReferences(conditionFinder, tr, tr.Constraint)
		}
		for _, dtr := range deviceType.DeviceTypeRequirements {
			sp.resolveFieldConstraintReferences(conditionFinder, dtr, dtr.Constraint)
		}
		for _, er := range deviceType.ComposedDeviceTypeElementRequirements {
			sp.resolveElementRequirementConstraint(er.ElementRequirement, conditionFinder)
		}
	}
}

func (sp *Builder) resolveElementRequirementConstraint(er *matter.ElementRequirement, finder entityFinder) {
	switch e := er.Entity.(type) {
	case *matter.Field:
		if e.Type != nil && e.Type.Entity != nil {
			finder = makeEntityFinder(e.Type.Entity, finder)
		}
	}
	sp.resolveFieldConstraintReferences(finder, er, er.Constraint)
}

func (sp *Builder) resolveFieldConstraints(fieldSet matter.FieldSet, field *matter.Field, dataType *types.DataType, finder entityFinder) {
	var fieldFinder entityFinder = newFieldFinder(fieldSet, finder)
	if dataType != nil && dataType.Entity != nil {
		dataTypeFinder := makeEntityFinder(dataType.Entity, fieldFinder)
		if dataTypeFinder != nil {
			fieldFinder = dataTypeFinder
		}
	}
	sp.resolveFieldConstraintReferences(fieldFinder, field, field.Constraint)
	sp.resolveFieldConstraintLimit(fieldFinder, field, field.Fallback)
}

func (sp *Builder) resolveFieldConstraintReferences(finder entityFinder, source types.Entity, con constraint.Constraint) {
	for _, limit := range constraint.TraverseConstraintLimits(con) {
		sp.resolveFieldConstraintLimit(finder, source, limit)
	}
}

func (sp *Builder) resolveFieldConstraintLimit(finder entityFinder, source types.Entity, limit constraint.Limit) {
	switch limit := limit.(type) {
	case *constraint.IdentifierLimit:
		if limit.Entity == nil {
			limit.Entity = finder.findEntityByIdentifier(limit.ID, source)
		}
		if limit.Entity != nil && limit.Field != nil {
			f := makeEntityFinder(limit.Entity, finder)
			for l := range constraint.TraverseLimit(limit.Field) {
				sp.resolveFieldConstraintLimit(f, source, l)
			}
		}
	case *constraint.ReferenceLimit:
		if limit.Entity == nil {
			limit.Entity = finder.findEntityByReference(limit.Reference, limit.Label, source)
		}
		if limit.Entity != nil && limit.Field != nil {
			f := makeEntityFinder(limit.Entity, finder)
			for l := range constraint.TraverseLimit(limit.Field) {
				sp.resolveFieldConstraintLimit(f, source, l)
			}
		}
	}
}
