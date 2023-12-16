package render

import (
	"slices"
	"strings"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/conformance"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/zap"
)

func (r *renderer) renderStructs(structs map[*matter.Struct]struct{}, cx *etree.Element) {

	ss := make([]*matter.Struct, 0, len(structs))
	for s := range structs {
		ss = append(ss, s)
	}

	slices.SortFunc(ss, func(a, b *matter.Struct) int {
		return strings.Compare(a.Name, b.Name)
	})

	for _, s := range ss {
		en := cx.CreateElement("struct")
		en.CreateAttr("name", zap.CleanName(s.Name))
		if s.FabricScoped {
			en.CreateAttr("isFabricScoped", "true")
		}
		r.renderClusterCodes(en, s)
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
