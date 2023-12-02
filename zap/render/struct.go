package render

import (
	"github.com/beevik/etree"
	"github.com/hasty/alchemy/conformance"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/zap"
)

func renderStructs(structs map[*matter.Struct]struct{}, clusterIDs []string, cx *etree.Element) {
	for s := range structs {
		en := cx.CreateElement("struct")
		en.CreateAttr("name", zap.CleanName(s.Name))
		if s.FabricScoped {
			en.CreateAttr("isFabricScoped", "true")
		}
		for _, cid := range clusterIDs {
			en.CreateElement("cluster").CreateAttr("code", cid)
		}
		for _, f := range s.Fields {
			if conformance.IsZigbee(s.Fields, f.Conformance) {
				continue
			}
			fx := en.CreateElement("item")
			fx.CreateAttr("fieldId", f.ID.IntString())
			fx.CreateAttr("name", zap.CleanName(f.Name))
			writeDataType(fx, s.Fields, f)
			renderConstraint(s.Fields, f, fx)
			if f.Quality.Has(matter.QualityNullable) {
				fx.CreateAttr("isNullable", "true")
			}
			if !conformance.IsMandatory(f.Conformance) {
				fx.CreateAttr("optional", "true")
			}
			defaultValue := zap.GetDefaultValue(&matter.ConstraintContext{Field: f, Fields: s.Fields})
			if defaultValue.Defined() {
				fx.CreateAttr("default", defaultValue.ZapString(f.Type))
			}
			if f.Access.FabricSensitive {
				fx.CreateAttr("isFabricSensitive", "true")
			}
		}

	}
}
