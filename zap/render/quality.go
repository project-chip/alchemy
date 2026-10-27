package render

import (
	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/internal/xml"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func (cr *configuratorRenderer) requiresQuality(entityType types.EntityType, q matter.Quality) bool {
	allowed, ok := matter.AllowedQualities[entityType]
	if !ok {
		return false
	}
	if cr.generator.options.ExtendedQuality {
		return q.Any(allowed)
	}
	return q.Any(allowed & (matter.QualityChangedOmitted | matter.QualityLargeMessage | matter.QualityFixed))
}

func (cr *configuratorRenderer) setQuality(parent *etree.Element, entityType types.EntityType, q matter.Quality, alternatives ...string) {
	requiresQuality := cr.requiresQuality(entityType, q)
	if requiresQuality {
		for _, el := range parent.SelectElements("quality") {
			if requiresQuality {
				cr.setQualityAttributes(el, q)
				requiresQuality = false
			} else {
				parent.RemoveChild(el)
			}
		}
		if requiresQuality {
			el := etree.NewElement("quality")
			xml.AppendElement(parent, el, alternatives...)
			cr.setQualityAttributes(el, q)
		}
	} else {
		for _, el := range parent.SelectElements("quality") {
			parent.RemoveChild(el)
		}
	}

}

func (cr *configuratorRenderer) setQualityAttributes(element *etree.Element, q matter.Quality) {
	setQualityAttribute(element, "changeOmitted", q, matter.QualityChangedOmitted)
	setQualityAttribute(element, "largeMessage", q, matter.QualityLargeMessage)
	if cr.generator.options.ExtendedQuality {
		setQualityAttribute(element, "nullable", q, matter.QualityNullable)
		setQualityAttribute(element, "scene", q, matter.QualityScene)
		setQualityAttribute(element, "diagnostics", q, matter.QualityDiagnostics)
		setQualityAttribute(element, "singleton", q, matter.QualitySingleton)
		setQualityAttribute(element, "sourceAttribution", q, matter.QualitySourceAttribution)
		setQualityAttribute(element, "atomicWrite", q, matter.QualityAtomicWrite)
		setQualityAttribute(element, "quieterReporting", q, matter.QualityQuieterReporting)
		element.RemoveAttr("fixed")
	}
	switch {
	case q.Has(matter.QualityFixed):
		// ZAP's ZCL loader reads the persistence attribute of quality, but not a fixed attribute
		element.CreateAttr("persistence", "fixed")
	case cr.generator.options.ExtendedQuality && q.Has(matter.QualityNonVolatile):
		element.CreateAttr("persistence", "nonVolatile")
	default:
		element.RemoveAttr("persistence")
	}
}

func setQualityAttribute(element *etree.Element, name string, quality matter.Quality, desiredQuality matter.Quality) {
	if quality.Has(desiredQuality) {
		element.CreateAttr(name, "true")
	} else {
		element.RemoveAttr(name)
	}
}
