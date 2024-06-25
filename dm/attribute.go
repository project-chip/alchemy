package dm

import (
	"slices"
	"strconv"
	"strings"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
)

func renderAttributes(doc *spec.Doc, cluster *matter.Cluster, c *etree.Element) (err error) {
	if len(cluster.Attributes) == 0 {
		return
	}
	as := make([]*matter.Field, len(cluster.Attributes))
	copy(as, cluster.Attributes)
	slices.SortFunc(as, func(a, b *matter.Field) int {
		return a.ID.Compare(b.ID)
	})
	attributes := c.CreateElement("attributes")
	for _, a := range as {
		if conformance.IsZigbee(cluster.Attributes, a.Conformance) {
			continue
		}
		ax := attributes.CreateElement("attribute")
		ax.CreateAttr("id", a.ID.HexString())
		ax.CreateAttr("name", a.Name)
		renderDataType(a, ax)
		if len(a.Default) > 0 {
			ax.CreateAttr("default", a.Default)
		}
		err = renderAnonymousType(doc, cluster, ax, a)
		if err != nil {
			return
		}
		renderAttributeAccess(ax, a.Access)
		renderQuality(ax, a.Quality)
		err = renderConformanceString(doc, cluster, a.Conformance, ax)
		if err != nil {
			return
		}
		err = renderConstraint(a.Constraint, a.Type, ax)
		if err != nil {
			return
		}
		renderDefault(cluster.Attributes, a, ax)
	}
	return
}

func renderAnonymousType(doc *spec.Doc, cluster *matter.Cluster, ax *etree.Element, field *matter.Field) error {
	switch at := field.AnonymousType.(type) {
	case *matter.AnonymousEnum:
		return renderAnonymousEnum(doc, cluster, ax, at)
	case *matter.AnonymousBitmap:
		return renderAnonymousBitmap(doc, cluster, ax, at)
	default:
	}
	return nil
}

func renderAnonymousEnum(doc *spec.Doc, cluster *matter.Cluster, ax *etree.Element, an *matter.AnonymousEnum) (err error) {
	en := ax.CreateElement("enum")
	for index, v := range an.Values {
		err = renderEnumValue(doc, cluster, en, index, v)
		if err != nil {
			return
		}
	}
	return
}

func renderAnonymousBitmap(doc *spec.Doc, cluster *matter.Cluster, ax *etree.Element, bm *matter.AnonymousBitmap) (err error) {
	en := ax.CreateElement("bitmap")
	size := bm.Size()
	for _, v := range bm.Bits {
		err = renderBit(doc, cluster, en, v, size)
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
	if !changeOmitted && !nullable && !scene && !fixed && !nonvolatile && !reportable && !singleton {
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
		} else if nonvolatile {
			qx.CreateAttr("persistence", "nonVolatile")
		} else {
			qx.CreateAttr("persistence", "volatile")
		}
	}
	if reportable {
		qx.CreateAttr("reportable", strconv.FormatBool(reportable))
	}
	if singleton {
		qx.CreateAttr("singleton", strconv.FormatBool(singleton))
	}
}
