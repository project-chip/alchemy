package generate

import (
	"log/slog"
	"regexp"
	"strings"

	"github.com/beevik/etree"
	"github.com/iancoleman/strcase"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/xml"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/zap"
)

func (cr *configuratorRenderer) generateAttributes(cle *etree.Element, cluster *matter.Cluster, attributes map[*matter.Field]struct{}, clusterPrefix string) (err error) {
	for _, ae := range cle.SelectElements("attribute") {
		ce := ae.SelectAttr("code")
		if ce == nil {
			slog.Warn("missing code attribute in cluster", slog.String("path", cr.configurator.OutPath), slog.String("clusterName", cluster.Name))
			continue
		}
		attributeID := matter.ParseNumber(ce.Value)
		if !attributeID.Valid() {
			slog.Warn("invalid code attribute value in cluster", slog.String("path", cr.configurator.OutPath), slog.String("clusterName", cluster.Name), slog.String("id", attributeID.Text()))
			continue
		}
		var attribute *matter.Field
		for a := range attributes {
			if !a.ID.Equals(attributeID) {
				continue
			}

			if conformance.IsZigbee(cluster, a.Conformance) || conformance.IsDisallowed(a.Conformance) {
				continue
			}

			if matter.NonGlobalIDInvalidForEntity(a.ID, types.EntityTypeAttribute) {
				continue
			}

			attribute = a
			delete(attributes, a)
		}
		if attribute == nil {
			slog.Warn("unrecognized attribute in cluster", slog.String("path", cr.configurator.OutPath), slog.String("clusterName", cluster.Name), slog.String("code", attributeID.Text()))
			cle.RemoveChild(ae)
			continue
		}
		delete(attributes, attribute)
		err = cr.populateAttribute(ae, attribute, cluster, clusterPrefix)
		if err != nil {
			return
		}
	}

	for a := range attributes {
		if conformance.IsZigbee(cluster, a.Conformance) {
			continue
		}
		if conformance.IsDeprecated(a.Conformance) {
			continue
		}
		if conformance.IsDisallowed(a.Conformance) {
			continue
		}
		if !a.ID.Valid() {
			continue
		}
		if matter.NonGlobalIDInvalidForEntity(a.ID, types.EntityTypeAttribute) {
			continue
		}

		ae := etree.NewElement("attribute")
		err = cr.populateAttribute(ae, a, cluster, clusterPrefix)
		if err != nil {
			return
		}
		xml.InsertElementByAttributeNumber(cle, ae, "code", a.ID, "globalAttribute", "server", "client", "domain")
	}
	return
}

func (cr *configuratorRenderer) populateAttribute(ae *etree.Element, attribute *matter.Field, cluster *matter.Cluster, clusterPrefix string) (err error) {
	cr.elementMap[ae] = attribute
	patchNumberAttribute(ae, attribute.ID, "code")
	ae.CreateAttr("side", "server")
	define := getDefine(attribute.Name, clusterPrefix, cr.configurator.Errata)
	xml.SetNonexistentAttr(ae, "define", define)
	cr.writeAttributeDataType(ae, cluster.Attributes, attribute)
	if attribute.Quality.Has(matter.QualityNullable) {
		ae.CreateAttr("isNullable", "true")
	} else {
		ae.RemoveAttr("isNullable")
	}
	// This is a deprecated quality, so remove it if it exists
	ae.RemoveAttr("reportable")
	if attribute.Quality.Has(matter.QualityAtomicWrite) {
		ae.CreateAttr("mustUseAtomicWrite", "true")
	} else {
		ae.RemoveAttr("mustUseAtomicWrite")
	}
	renderConstraint(ae, cluster.Attributes, attribute)
	setFieldFallback(ae, attribute, cluster.Attributes)
	needsRead := attribute.Access.Read != matter.PrivilegeUnknown && attribute.Access.Read != matter.PrivilegeView
	var needsWrite bool
	if attribute.Access.Write != matter.PrivilegeUnknown {
		needsWrite = attribute.Access.Write != matter.PrivilegeOperate
		ae.CreateAttr("writable", "true")
	} else {
		ae.RemoveAttr("writable")
	}
	if attribute.Access.IsTimed() {
		ae.CreateAttr("mustUseTimedWrite", "true")
	} else {
		ae.RemoveAttr("mustUseTimedWrite")
	}
	if !conformance.IsMandatory(attribute.Conformance) {
		ae.CreateAttr("optional", "true")
	} else {
		ae.RemoveAttr("optional")
	}
	requiresConformance := cr.generator.generateConformanceXML && !conformance.IsBlank(attribute.Conformance) && !conformance.IsMandatory(attribute.Conformance)
	requiresPermissions := !cr.configurator.Errata.SuppressAttributePermissions && (needsRead || needsWrite)
	requiresQuality := attribute.Quality.Any(matter.QualityChangedOmitted)
	if !requiresPermissions && !requiresQuality && !requiresConformance {
		ae.Child = nil
		ae.SetText(attribute.Name)
	} else {
		ae.SetText("")
		de, exists := xml.CreateSimpleElementIfNotExists(ae, "description", attribute.Name)
		if exists {
			de.SetText(attribute.Name)
		}
		if requiresPermissions {
			accessElements := ae.SelectElements("access")
			for _, ax := range accessElements {
				if needsRead {
					cr.setAccessAttributes(ax, "read", attribute.Access.Read)
					needsRead = false
				} else if needsWrite {
					cr.setAccessAttributes(ax, "write", attribute.Access.Write)
					needsWrite = false
				} else {
					ae.RemoveChild(ax)
				}
			}
			if needsRead {
				ax := etree.NewElement("access")
				cr.setAccessAttributes(ax, "read", attribute.Access.Read)
				xml.AppendElement(ae, ax, "description")
			}
			if needsWrite {
				ax := etree.NewElement("access")
				cr.setAccessAttributes(ax, "write", attribute.Access.Write)
				xml.AppendElement(ae, ax, "access", "description")
			}
		}

		if requiresQuality {
			for _, el := range ae.SelectElements("quality") {
				if requiresQuality {
					cr.setQualityAttributes(el, attribute.Quality)
					requiresQuality = false
				} else {
					ae.RemoveChild(el)
				}
			}
			if requiresQuality {
				el := etree.NewElement("quality")
				xml.AppendElement(ae, el, "access", "description")
				cr.setQualityAttributes(el, attribute.Quality)
			}
		} else {
			for _, el := range ae.SelectElements("quality") {
				ae.RemoveChild(el)
			}
		}
		if cr.generator.generateConformanceXML {
			err = renderConformance(cr.generator.spec, attribute, cluster, attribute.Conformance, ae, "quality", "access", "description")
			if err != nil {
				return err
			}
		} else {
			removeConformance(ae)
		}
	}

	return
}

func setFieldFallback(e *etree.Element, field *matter.Field, fieldSet matter.FieldSet) {
	if !constraint.IsGenericLimit(field.Fallback) && !constraint.IsBlankLimit(field.Fallback) {
		fallbackValue := zap.GetFallbackValue(&matter.ConstraintContext{Field: field, Fields: fieldSet})
		patchDataExtremeAttribute(e, "default", &fallbackValue, field)
	} else {
		e.RemoveAttr("default")
	}
}

func (cr *configuratorRenderer) writeAttributeDataType(x *etree.Element, fs matter.FieldSet, f *matter.Field) {
	if f.Type == nil {
		return
	}
	dts := zap.FieldToZapDataType(fs, f)
	dts = cr.configurator.Errata.TypeName(dts)
	if f.Type.IsArray() {
		x.CreateAttr("type", "array")
		x.CreateAttr("entryType", dts)
	} else {
		x.CreateAttr("type", dts)
		x.RemoveAttr("entryType")
	}
}

func getDefine(name string, prefix string, errata *errata.ZAP) string {
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
