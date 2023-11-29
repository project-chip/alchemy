package render

import (
	"strings"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/conformance"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/zap"
	"github.com/iancoleman/strcase"
)

func renderAttributes(cluster *matter.Cluster, cx *etree.Element, clusterPrefix string, errata *Errata) {
	if len(cluster.Attributes) > 0 {
		cx.CreateComment("Attributes")
	}
	for _, a := range cluster.Attributes {
		if conformance.IsZigbee(a.Conformance) || conformance.IsDeprecated(a.Conformance) {
			continue
		}

		if !a.ID.Valid() {
			continue
		}

		/*if !mandatory && len(a.Conformance) > 0 && a.Conformance != "O" {
			cx.CreateComment(fmt.Sprintf("Conformance feature %s - for now optional", a.Conformance))
		}*/
		attr := cx.CreateElement("attribute")
		attr.CreateAttr("code", a.ID.HexString())
		attr.CreateAttr("side", "server")
		define := GetDefine(a.Name, clusterPrefix, errata)
		attr.CreateAttr("define", define)
		writeAttributeDataType(attr, cluster.Attributes, a)
		if a.Quality.Has(matter.QualityNullable) {
			attr.CreateAttr("isNullable", "true")
		}
		if a.Quality.Has(matter.QualityReportable) {
			attr.CreateAttr("reportable", "true")
		}
		if a.Default != "" {
			defaultValue := zap.GetDefaultValue(&matter.ConstraintContext{Field: a, Fields: cluster.Attributes})
			if defaultValue.Defined() {
				attr.CreateAttr("default", defaultValue.ZapString(a.Type))
			}
		}
		renderConstraint(cluster.Attributes, a, attr)
		renderAttributeAccess(a, errata, attr)
		if !conformance.IsMandatory(a.Conformance) {
			attr.CreateAttr("optional", "true")
		} else {
			attr.CreateAttr("optional", "false")
		}

	}
}

func renderConstraint(fs matter.FieldSet, f *matter.Field, attr *etree.Element) {

	from, to := zap.GetMinMax(&matter.ConstraintContext{Fields: fs, Field: f})

	if f.Type != nil && f.Type.IsString() {
		if to.Defined() {
			attr.CreateAttr("length", to.ZapString(f.Type))
		}
		if from.Defined() {
			attr.CreateAttr("minLength", from.ZapString(f.Type))
		}
	} else {
		if from.Defined() {
			attr.CreateAttr("min", from.ZapString(f.Type))
		}
		if to.Defined() {
			attr.CreateAttr("max", to.ZapString(f.Type))
		}
	}
}

func renderAttributeAccess(a *matter.Field, errata *Errata, attr *etree.Element) {
	if a.Quality.Has(matter.QualityFixed) || (a.Access.Read == matter.PrivilegeUnknown || a.Access.Read == matter.PrivilegeView) && a.Access.Write == matter.PrivilegeUnknown || errata.SuppressAttributePermissions {
		if a.Access.Write != matter.PrivilegeUnknown {
			attr.CreateAttr("writable", "true")
		} else {
			attr.CreateAttr("writable", "false")
		}
		attr.SetText(a.Name)
	} else {
		if a.Access.Read != matter.PrivilegeUnknown {
			ax := attr.CreateElement("access")
			ax.CreateAttr("op", "read")
			ax.CreateAttr("privilege", renderPrivilege(a.Access.Read))
		}
		if a.Access.Write != matter.PrivilegeUnknown {
			ax := attr.CreateElement("access")
			ax.CreateAttr("op", "write")
			ax.CreateAttr("privilege", renderPrivilege(a.Access.Write))
			attr.CreateAttr("writable", "true")
		} else {
			attr.CreateAttr("writable", "false")
		}
		attr.CreateElement("description").SetText(a.Name)
	}
}

func renderPrivilege(a matter.Privilege) string {
	switch a {
	case matter.PrivilegeView:
		return "view"
	case matter.PrivilegeManage:
		return "manage"
	case matter.PrivilegeAdminister:
		return "administer"
	case matter.PrivilegeOperate:
		return "operate"
	default:
		return ""
	}
}

func GetDefine(name string, prefix string, errata *Errata) string {
	define := strcase.ToScreamingDelimited(name, '_', "", true)
	if !strings.HasPrefix(define, prefix) {
		define = prefix + define
	}
	if errata.DefineOverrides != nil {
		if override, ok := errata.DefineOverrides[define]; ok {
			return override
		}
	}
	return define
}
