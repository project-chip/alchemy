package render

import (
	"strings"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
	"github.com/hasty/alchemy/zap"
)

func renderEvents(cluster *matter.Cluster, cx *etree.Element) {
	for _, e := range cluster.Events {

		ex := cx.CreateElement("event")
		ex.CreateAttr("code", e.ID.HexString())
		ex.CreateAttr("name", zap.CleanName(e.Name))
		ex.CreateAttr("priority", strings.ToLower(e.Priority))
		if e.FabricSensitive {
			ex.CreateAttr("isFabricSensitive", "true")
		}
		if !conformance.IsMandatory(e.Conformance) {
			ex.CreateAttr("optional", "true")
		}

		if len(e.Description) > 0 {
			ex.CreateElement("description").SetText(e.Description)
		} else {
			ex.CreateElement("description").SetText(e.Name)

		}
		for _, f := range e.Fields {
			if conformance.IsZigbee(e.Fields, f.Conformance) {
				continue
			}
			if !f.ID.Valid() {
				continue
			}
			fx := ex.CreateElement("field")
			fx.CreateAttr("id", f.ID.IntString())
			fx.CreateAttr("name", zap.CleanName(f.Name))
			writeDataType(fx, e.Fields, f)
			renderConstraint(e.Fields, f, fx)
			if f.Quality.Has(matter.QualityNullable) {
				fx.CreateAttr("isNullable", "true")
			}
			if !conformance.IsMandatory(f.Conformance) {
				fx.CreateAttr("optional", "true")
			}
		}
		if e.Access.Read != matter.PrivilegeUnknown {
			ax := ex.CreateElement("access")
			ax.CreateAttr("op", "read")
			ax.CreateAttr("privilege", renderPrivilege(e.Access.Read))
		}
		if e.Access.Write != matter.PrivilegeUnknown {
			ax := ex.CreateElement("access")
			ax.CreateAttr("op", "write")
			ax.CreateAttr("privilege", renderPrivilege(e.Access.Write))
		}
	}
}
