package generate

import (
	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/matter"
)

func requiresQuality(q matter.Quality) bool {
	return q.Any(matter.QualityChangedOmitted | matter.QualityLargeMessage)
}

func (cr *configuratorRenderer) setQualityAttributes(element *etree.Element, q matter.Quality) {
	setQualityAttribute(element, "changeOmitted", q, matter.QualityChangedOmitted)
	setQualityAttribute(element, "largeMessage", q, matter.QualityLargeMessage)
	// To be enabled when ZAP supports it
	/*
		setQualityAttribute(element, "nullable", q, matter.QualityNullable)
		setQualityAttribute(element, "scene", q, matter.QualityScene)
		setQualityAttribute(element, "fixed", q, matter.QualityFixed)
		setQualityAttribute(element, "diagnostics", q, matter.QualityDiagnostics)
		setQualityAttribute(element, "singleton", q, matter.QualitySingleton)
		setQualityAttribute(element, "sourceAttribution", q, matter.QualitySourceAttribution)
		setQualityAttribute(element, "atomicWrite", q, matter.QualityAtomicWrite)
		setQualityAttribute(element, "quieterReporting", q, matter.QualityQuieterReporting)
	*/
}

func setQualityAttribute(element *etree.Element, name string, quality matter.Quality, desiredQuality matter.Quality) {
	if quality.Has(desiredQuality) {
		element.CreateAttr(name, "true")
	} else {
		element.RemoveAttr(name)
	}
}
