package dm

import (
	"slices"
	"strconv"
	"strings"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/constraint"
)

func renderAttributes(cluster *matter.Cluster, c *etree.Element) (err error) {
	if len(cluster.Attributes) == 0 {
		return
	}
	as := make([]*matter.Field, len(cluster.Attributes))
	copy(as, cluster.Attributes)
	slices.SortStableFunc(as, func(a, b *matter.Field) int {
		return a.ID.Compare(b.ID)
	})
	attributes := c.CreateElement("attributes")
	for _, a := range as {
		if conformance.IsZigbee(a.Conformance) {
			continue
		}
		ax := attributes.CreateElement("attribute")
		ax.CreateAttr("id", a.ID.HexString())
		ax.CreateAttr("name", a.Name)
		err = renderDataType(a, ax)
		if err != nil {
			return
		}
		if !constraint.IsBlankLimit(a.Fallback) {
			err = renderConstraintLimit(ax, ax, a.Fallback, a.Type, "default", cluster)
			if err != nil {
				return
			}
		}
		err = renderAnonymousType(ax, a)
		if err != nil {
			return
		}
		renderAttributeAccess(ax, a.Access)
		renderQuality(ax, a.Quality)
		err = renderConformanceElement(a.Conformance, ax, nil)
		if err != nil {
			return
		}
		err = renderConstraint(a.Constraint, a.Type, ax, cluster)
		if err != nil {
			return
		}
	}
	return
}

func renderAnonymousType(ax *etree.Element, field *matter.Field) error {
	switch at := field.AnonymousType.(type) {
	case *matter.AnonymousEnum:
		return renderAnonymousEnum(ax, at)
	case *matter.AnonymousBitmap:
		return renderAnonymousBitmap(ax, at)
	default:
	}
	return nil
}

func renderAnonymousEnum(ax *etree.Element, an *matter.AnonymousEnum) (err error) {
	en := ax.CreateElement("enum")
	for index, v := range an.Values {
		err = renderEnumValue(en, index, v, an)
		if err != nil {
			return
		}
	}
	return
}

func renderAnonymousBitmap(ax *etree.Element, bm *matter.AnonymousBitmap) (err error) {
	en := ax.CreateElement("bitmap")
	size := bm.Size()
	for _, v := range bm.Bits {
		err = renderBit(en, v, size, bm)
		if err != nil {
			return
		}
	}

	return
}

func renderAttributeAccess(ax *etree.Element, a matter.Access) {
	if a.Read == matter.PrivilegeUnknown && a.Write == matter.PrivilegeUnknown && !a.IsTimed() && a.FabricSensitivity != matter.FabricSensitivitySensitive && a.FabricScoping != matter.FabricScopingScoped {
		return
	}
	acx := ax.CreateElement("access")
	if a.Read != matter.PrivilegeUnknown {
		acx.CreateAttr("read", "true")
	}
	if a.Write != matter.PrivilegeUnknown {
		if a.OptionalWrite {
			acx.CreateAttr("write", "optional")
		} else {
			acx.CreateAttr("write", "true")
		}
	}
	if a.Read != matter.PrivilegeUnknown {
		acx.CreateAttr("readPrivilege", strings.ToLower(matter.PrivilegeNamesShort[a.Read]))
	}
	if a.Write != matter.PrivilegeUnknown {
		acx.CreateAttr("writePrivilege", strings.ToLower(matter.PrivilegeNamesShort[a.Write]))
	}
	if a.IsTimed() {
		acx.CreateAttr("timed", "true")
	}
	if a.FabricScoping == matter.FabricScopingScoped {
		acx.CreateAttr("fabricScoped", "true")
	}
	if a.FabricSensitivity == matter.FabricSensitivitySensitive {
		acx.CreateAttr("fabricSensitive", "true")
	}
}

func renderQuality(parent *etree.Element, q matter.Quality) {
	changeOmitted := q.Has(matter.QualityChangedOmitted)
	nullable := q.Has(matter.QualityNullable)
	scene := q.Has(matter.QualityScene)
	fixed := q.Has(matter.QualityFixed)
	nonvolatile := q.Has(matter.QualityNonVolatile)
	reportable := q.Has(matter.QualityReportable)
	singleton := q.Has(matter.QualitySingleton)
	atomicWrite := q.Has(matter.QualityAtomicWrite)
	diagnostics := q.Has(matter.QualityDiagnostics)
	quieterReporting := q.Has(matter.QualityQuieterReporting)
	sourceAttribution := q.Has(matter.QualitySourceAttribution)
	largeMessage := q.Has(matter.QualityLargeMessage)
	if !changeOmitted && !nullable && !scene && !fixed && !nonvolatile && !reportable && !singleton && !atomicWrite && !diagnostics && !quieterReporting && !sourceAttribution && !largeMessage {
		return
	}
	qx := parent.CreateElement("quality")
	if changeOmitted {
		qx.CreateAttr("changeOmitted", strconv.FormatBool(changeOmitted))
	}
	if nullable {
		qx.CreateAttr("nullable", strconv.FormatBool(nullable))
	}
	if scene {
		qx.CreateAttr("scene", strconv.FormatBool(scene))
	}
	if fixed || nonvolatile {
		if fixed {
			qx.CreateAttr("persistence", "fixed")
		} else {
			qx.CreateAttr("persistence", "nonVolatile")
		}
	}
	if reportable {
		qx.CreateAttr("reportable", strconv.FormatBool(reportable))
	}
	if singleton {
		qx.CreateAttr("singleton", strconv.FormatBool(singleton))
	}
	if atomicWrite {
		qx.CreateAttr("atomicWrite", strconv.FormatBool(atomicWrite))
	}
	if diagnostics {
		qx.CreateAttr("diagnostics", strconv.FormatBool(diagnostics))
	}
	if quieterReporting {
		qx.CreateAttr("quieterReporting", strconv.FormatBool(quieterReporting))
	}
	if sourceAttribution {
		qx.CreateAttr("sourceAttribution", strconv.FormatBool(sourceAttribution))
	}
	if largeMessage {
		qx.CreateAttr("largeMessage", strconv.FormatBool(largeMessage))
	}
}
