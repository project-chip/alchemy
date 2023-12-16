package dm

import (
	"slices"
	"strconv"
	"strings"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
)

func renderAttributes(cluster *matter.Cluster, c *etree.Element) (err error) {
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
		renderAttributeAccess(ax, a.Access)
		renderQuality(ax, a.Quality, matter.QualityAll^(matter.QualitySingleton))
		err = renderConformanceString(cluster, a.Conformance, ax)
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

func renderAttributeAccess(ax *etree.Element, a matter.Access) {
	if a.Read == matter.PrivilegeUnknown && a.Write == matter.PrivilegeUnknown && !a.Timed {
		return
	}
	acx := ax.CreateElement("access")
	if a.Read != matter.PrivilegeUnknown {
		acx.CreateAttr("read", "true")
	}
	if a.Write != matter.PrivilegeUnknown {
		if a.Write == matter.PrivilegeOperate {
			acx.CreateAttr("write", "true")
		} else {
			acx.CreateAttr("write", "optional")
		}
	}
	if a.Read != matter.PrivilegeUnknown {
		acx.CreateAttr("readPrivilege", strings.ToLower(matter.PrivilegeNamesShort[a.Read]))
	}
	if a.Write != matter.PrivilegeUnknown {
		acx.CreateAttr("writePrivilege", strings.ToLower(matter.PrivilegeNamesShort[a.Write]))
	}
	if a.Timed {
		acx.CreateAttr("timed", "true")
	}
}

func renderQuality(parent *etree.Element, q matter.Quality, mask matter.Quality) {
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
	if mask&matter.QualityChangedOmitted == matter.QualityChangedOmitted {
		qx.CreateAttr("changeOmitted", strconv.FormatBool(changeOmitted))
	}
	if mask&matter.QualityNullable == matter.QualityNullable {
		qx.CreateAttr("nullable", strconv.FormatBool(nullable))
	}
	if mask&matter.QualityScene == matter.QualityScene {
		qx.CreateAttr("scene", strconv.FormatBool(scene))
	}
	if mask&matter.QualityFixed == matter.QualityFixed || mask&matter.QualityNonVolatile == matter.QualityNonVolatile {
		if fixed {
			qx.CreateAttr("persistence", "fixed")
		} else if nonvolatile {
			qx.CreateAttr("persistence", "nonVolatile")
		} else {
			qx.CreateAttr("persistence", "volatile")
		}
	}
	if mask&matter.QualityReportable == matter.QualityReportable {
		qx.CreateAttr("reportable", strconv.FormatBool(reportable))
	}
	if mask&matter.QualitySingleton == matter.QualitySingleton {
		qx.CreateAttr("singleton", strconv.FormatBool(singleton))
	}
}
