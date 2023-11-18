package render

import (
	"github.com/beevik/etree"
	"github.com/hasty/alchemy/matter"
)

func renderStructs(structs []*matter.Struct, clusterIDs []string, cx *etree.Element) {
	for _, s := range structs {
		en := cx.CreateElement("struct")
		en.CreateAttr("name", s.Name)
		en.CreateAttr("apiMaturity", "provisional")
		if s.FabricScoped {
			en.CreateAttr("isFabricScoped", "true")
		}
		for _, cid := range clusterIDs {
			en.CreateElement("cluster").CreateAttr("code", cid)
		}
		for _, f := range s.Fields {
			if f.Conformance == "Zigbee" {
				continue
			}
			fx := en.CreateElement("item")
			fx.CreateAttr("fieldId", f.ID.IntString())
			fx.CreateAttr("name", f.Name)
			writeDataType(fx, s.Fields, f)
			renderConstraint(s.Fields, f, fx)
			if f.Quality.Has(matter.QualityNullable) {
				fx.CreateAttr("isNullable", "true")
			}
			if f.Conformance != "M" {
				fx.CreateAttr("optional", "true")
			}
		}

	}
}
