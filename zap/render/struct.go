package render

import (
	"github.com/beevik/etree"
	"github.com/hasty/alchemy/matter"
)

func renderStructs(structs []*matter.Struct, clusterIDs []string, cx *etree.Element) {
	for _, v := range structs {
		en := cx.CreateElement("struct")
		en.CreateAttr("name", v.Name)
		en.CreateAttr("apiMaturity", "provisional")
		if v.FabricScoped {
			en.CreateAttr("isFabricScoped", "true")
		}
		for _, cid := range clusterIDs {
			en.CreateElement("cluster").CreateAttr("code", cid)
		}
		for _, f := range v.Fields {
			if f.Conformance == "Zigbee" {
				continue
			}
			fx := en.CreateElement("item")
			fx.CreateAttr("fieldId", f.ID.IntString())
			fx.CreateAttr("name", f.Name)
			writeDataType(fx, f.Type)
			renderConstraint(v.Fields, f, fx)
			if f.Quality.Has(matter.QualityNullable) {
				fx.CreateAttr("isNullable", "true")
			}
			if f.Conformance != "M" {
				fx.CreateAttr("optional", "true")
			}
		}

	}
}
