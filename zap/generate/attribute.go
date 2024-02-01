package generate

import (
	"log/slog"
	"regexp"
	"strings"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
	"github.com/hasty/alchemy/zap"
	"github.com/iancoleman/strcase"
)

func generateAttributes(configurator *zap.Configurator, cle *etree.Element, cluster *matter.Cluster, attributes map[*matter.Field]struct{}, clusterPrefix string, errata *zap.Errata) (err error) {

	for _, ae := range cle.SelectElements("attribute") {
		ce := ae.SelectAttr("code")
		if ce == nil {
			slog.Warn("missing code attribute in cluster", slog.String("path", configurator.Doc.Path), slog.String("clusterName", cluster.Name))
			continue
		}
		attributeId := matter.ParseNumber(ce.Value)
		if !attributeId.Valid() {
			slog.Warn("invalid code attribute value in cluster", slog.String("path", configurator.Doc.Path), slog.String("clusterName", cluster.Name), slog.String("id", attributeId.Text()))
			continue
		}
		var attribute *matter.Field
		for a := range attributes {
			if !a.ID.Equals(attributeId) {
				continue
			}

			if conformance.IsZigbee(cluster, a.Conformance) {
				continue
			}

			attribute = a
			delete(attributes, a)
		}
		if attribute == nil {
			slog.Warn("unrecognized code value in cluster", slog.String("path", configurator.Doc.Path), slog.String("clusterName", cluster.Name), slog.String("code", attributeId.Text()))
			cle.RemoveChild(ae)
			continue
		}
		delete(attributes, attribute)
		populateAttribute(ae, attribute, cluster, clusterPrefix, errata)
	}
	for a := range attributes {
		if conformance.IsZigbee(cluster, a.Conformance) {
			continue
		}
		if !a.ID.Valid() {
			continue
		}
		ae := etree.NewElement("attribute")
		populateAttribute(ae, a, cluster, clusterPrefix, errata)
		insertElementByName(cle, ae, "code", "globalAttribute", "server", "client", "domain")
	}
	return
}

func populateAttribute(ae *etree.Element, attribute *matter.Field, cluster *matter.Cluster, clusterPrefix string, errata *zap.Errata) (err error) {
	ae.CreateAttr("code", attribute.ID.HexString())
	ae.CreateAttr("side", "server")
	define := getDefine(attribute.Name, clusterPrefix, errata)
	ae.CreateAttr("define", define)
	writeAttributeDataType(ae, cluster.Attributes, attribute)
	if attribute.Quality.Has(matter.QualityNullable) {
		ae.CreateAttr("isNullable", "true")
	} else {
		ae.RemoveAttr("isNullable")
	}
	if attribute.Quality.Has(matter.QualityReportable) {
		ae.CreateAttr("reportable", "true")
	} else {
		ae.RemoveAttr("reportable")
	}
	renderConstraint(ae, cluster.Attributes, attribute)
	if attribute.Default != "" {
		defaultValue := zap.GetDefaultValue(&matter.ConstraintContext{Field: attribute, Fields: cluster.Attributes})
		if defaultValue.Defined() && !defaultValue.IsNull() {
			ae.CreateAttr("default", defaultValue.ZapString(attribute.Type))
		} else {
			ae.RemoveAttr("default")
		}
	} else {
		ae.RemoveAttr("default")
	}
	if attribute.Quality.Has(matter.QualityFixed) || ((attribute.Access.Read == matter.PrivilegeUnknown || attribute.Access.Read == matter.PrivilegeView) && (attribute.Access.Write == matter.PrivilegeUnknown || attribute.Access.Write == matter.PrivilegeOperate)) || errata.SuppressAttributePermissions {
		if attribute.Access.Write != matter.PrivilegeUnknown {
			ae.CreateAttr("writable", "true")
		} else {
			ae.RemoveAttr("writable")
		}
		ae.SetText(attribute.Name)
	} else {
		setOrCreateSimpleElement(ae, "description", attribute.Name)
		needsRead := attribute.Access.Read != matter.PrivilegeUnknown && attribute.Access.Read != matter.PrivilegeView
		var needsWrite bool
		if attribute.Access.Write != matter.PrivilegeUnknown {
			needsWrite = attribute.Access.Write != matter.PrivilegeOperate
			ae.CreateAttr("writable", "true")
		} else {
			ae.CreateAttr("writable", "false")
		}
		accessElements := ae.SelectElements("access")
		for _, ax := range accessElements {
			if needsRead {
				ax.CreateAttr("op", "read")
				ax.CreateAttr("privilege", renderPrivilege(attribute.Access.Read))
				ax.RemoveAttr("role")
				needsRead = false
			} else if needsWrite {
				ax.CreateAttr("op", "write")
				ax.CreateAttr("privilege", renderPrivilege(attribute.Access.Write))
				ax.RemoveAttr("role")
				needsWrite = false
			} else {
				ae.RemoveChild(ax)
			}
		}
		if needsRead {
			ax := ae.CreateElement("access")
			ax.CreateAttr("op", "read")
			ax.CreateAttr("privilege", renderPrivilege(attribute.Access.Read))
		}
		if needsWrite {
			ax := ae.CreateElement("access")
			ax.CreateAttr("op", "write")
			ax.CreateAttr("privilege", renderPrivilege(attribute.Access.Write))
		}
	}
	if !conformance.IsMandatory(attribute.Conformance) {
		ae.CreateAttr("optional", "true")
	} else {
		ae.RemoveAttr("optional")
	}
	return
}

func writeAttributeDataType(x *etree.Element, fs matter.FieldSet, f *matter.Field) {
	if f.Type == nil {
		return
	}
	dts := zap.FieldToZapDataType(fs, f)
	if f.Type.IsArray() {
		x.CreateAttr("type", "array")
		x.CreateAttr("entryType", dts)
	} else {
		x.CreateAttr("type", dts)
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

func getDefine(name string, prefix string, errata *zap.Errata) string {
	define := strcase.ToScreamingDelimited(cleanAcronyms(name), '_', "", true)
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

var acronymPattern = regexp.MustCompile(`[A-Z]{2,}`)

func cleanAcronyms(s string) string {
	s2 := acronymPattern.ReplaceAllStringFunc(s, func(match string) string {
		return string(match[0]) + strings.ToLower(string(match[1:]))
	})
	return s2
}
