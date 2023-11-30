package dm

import (
	"strconv"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/matter"
)

func renderQuality(parent *etree.Element, a *matter.Field) {
	changeOmitted := a.Quality.Has(matter.QualityChangedOmitted)
	nullable := a.Quality.Has(matter.QualityNullable)
	scene := a.Quality.Has(matter.QualityScene)
	fixed := a.Quality.Has(matter.QualityFixed)
	nonvolatile := a.Quality.Has(matter.QualityNonVolatile)
	reportable := a.Quality.Has(matter.QualityReportable)
	if !changeOmitted && !nullable && !scene && !fixed && !nonvolatile && !reportable {
		return
	}
	qx := parent.CreateElement("quality")
	qx.CreateAttr("changeOmitted", strconv.FormatBool(changeOmitted))
	qx.CreateAttr("nullable", strconv.FormatBool(nullable))
	qx.CreateAttr("scene", strconv.FormatBool(scene))
	if fixed {
		qx.CreateAttr("persistence", "fixed")
	} else if nonvolatile {
		qx.CreateAttr("persistence", "nonVolatile")
	} else {
		qx.CreateAttr("persistence", "volatile")
	}
	qx.CreateAttr("reportable", strconv.FormatBool(reportable))
}
