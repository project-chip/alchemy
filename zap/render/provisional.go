package render

import (
	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/provisional"
)

func (cr *configuratorRenderer) isProvisional(entity types.Entity) bool {
	is := provisional.Check(cr.generator.spec, entity, entity)
	switch is {
	case provisional.StateAllClustersProvisional,
		provisional.StateAllDataTypeReferencesProvisional,
		provisional.StateExplicitlyProvisional,
		provisional.StateUnreferenced:
		return true
	default:
		return false
	}
}

func (cr *configuratorRenderer) setProvisional(e *etree.Element, entity types.Entity) {
	switch provisional.Policy(cr.generator.options.ProvisionalPolicy) {
	case provisional.PolicyStrict, provisional.PolicyLoose:
		if cr.isProvisional(entity) {
			e.CreateAttr("apiMaturity", "provisional")
		} else {
			e.RemoveAttr("apiMaturity")
		}
	case provisional.PolicyNone:
		switch entity := entity.(type) {
		case *matter.Cluster:
			if conformance.IsProvisional(entity.Conformance) {
				e.CreateAttr("apiMaturity", "provisional")
			} else {
				e.RemoveAttr("apiMaturity")
			}
		}
	}
}

func (cr *configuratorRenderer) isProvisionalViolation(entity types.Entity) bool {
	switch provisional.Policy(cr.generator.options.ProvisionalPolicy) {
	case provisional.PolicyStrict:
		switch provisional.Check(cr.generator.spec, entity, entity) {
		case provisional.StateAllClustersNonProvisional,
			provisional.StateSomeClustersProvisional,
			provisional.StateAllDataTypeReferencesNonProvisional,
			provisional.StateSomeDataTypeReferencesProvisional,
			provisional.StateUnreferenced:
			return true
		case provisional.StateAllDataTypeReferencesProvisional, provisional.StateAllClustersProvisional, provisional.StateExplicitlyProvisional:
			return false
		}
	}
	return false
}
