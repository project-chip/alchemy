package zcl

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/beevik/etree"
	"github.com/hasty/matterfmt/matter"
	"github.com/hasty/matterfmt/parse"
	"github.com/iancoleman/strcase"
)

func renderAttributes(cluster *matter.Cluster, cx *etree.Element, clusterPrefix string, errata *errata) {
	if len(cluster.Attributes) > 0 {
		cx.CreateComment("Attributes")
	}
	for _, a := range cluster.Attributes {
		if a.Conformance == "Zigbee" {
			continue
		}
		mandatory := (a.Conformance == "M")
		readAccess, writeAccess, _, _, _ := matter.ParseAccessValues(a.Access)

		id, err := parse.ID(a.ID)
		if err != nil {
			continue
		}

		if !mandatory && len(a.Conformance) > 0 && a.Conformance != "O" {
			cx.CreateComment(fmt.Sprintf("Conformance feature %s - for now optional", a.Conformance))
		}
		attr := cx.CreateElement("attribute")
		attr.CreateAttr("side", "server")
		attr.CreateAttr("code", fmt.Sprintf("%#04x", id))
		define := getDefine(a.Name, clusterPrefix, errata)
		attr.CreateAttr("define", define)
		writeDataType(attr, a.Type)
		if a.Quality.Has(matter.QualityNullable) {
			attr.CreateAttr("isNullable", "true")
		} else {
			attr.CreateAttr("isNullable", "false")
		}
		if a.Quality.Has(matter.QualityReportable) {
			attr.CreateAttr("reportable", "true")
		}
		if a.Default != "" {
			def, err := parse.ID(a.Default)
			if err == nil {
				attr.CreateAttr("default", strconv.Itoa(int(def)))
			}
		}
		if a.Quality.Has(matter.QualityFixed) || readAccess == "V" && writeAccess == "" || errata.suppressAttributePermissions {
			if writeAccess != "" {
				attr.CreateAttr("writeable", "true")
			} else {
				attr.CreateAttr("writeable", "false")

			}
			attr.SetText(a.Name)
		} else {
			attr.CreateElement("description").SetText(a.Name)
			if readAccess != "" {
				ax := attr.CreateElement("access")
				ax.CreateAttr("op", "read")
				ax.CreateAttr("privilege", renderAccess(readAccess))
			}
			if writeAccess != "" {
				ax := attr.CreateElement("access")
				ax.CreateAttr("op", "write")
				ax.CreateAttr("privilege", renderAccess(writeAccess))
				attr.CreateAttr("writeable", "true")
			} else {
				attr.CreateAttr("writeable", "false")
			}
		}
		if a.Conformance != "M" {
			attr.CreateAttr("optional", "true")
		} else {
			attr.CreateAttr("optional", "false")
		}

	}
}

func renderAccess(a string) string {
	switch a {
	case "V":
		return "view"
	case "M":
		return "manage"
	case "A":
		return "administer"
	case "O":
		return "operate"
	default:
		return a
	}
}

func getDefine(name string, prefix string, errata *errata) string {
	define := strcase.ToScreamingDelimited(name, '_', "", true)
	if !strings.HasPrefix(define, prefix) {
		define = prefix + define
	}
	if errata.defineOverrides != nil {
		if override, ok := errata.defineOverrides[define]; ok {
			return override
		}
	}
	return define
}
