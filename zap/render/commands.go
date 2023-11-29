package render

import (
	"github.com/beevik/etree"
	"github.com/hasty/alchemy/conformance"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/zap"
)

func renderCommands(cluster *matter.Cluster, cx *etree.Element, errata *Errata) {
	for _, c := range cluster.Commands {
		if c.Direction == matter.InterfaceServer {
			renderCommand(c, cx, errata)
		}
	}
	for _, c := range cluster.Commands {
		if c.Direction == matter.InterfaceClient {
			renderCommand(c, cx, errata)
		}
	}
}

func renderCommand(c *matter.Command, e *etree.Element, errata *Errata) {
	mandatory := conformance.IsMandatory(c.Conformance)

	cx := e.CreateElement("command")
	var serverSource bool
	if c.Direction == matter.InterfaceServer {
		cx.CreateAttr("source", "client")
	} else if c.Direction == matter.InterfaceClient {
		cx.CreateAttr("source", "server")
		serverSource = true
	}
	cx.CreateAttr("code", c.ID.HexString())
	cx.CreateAttr("name", c.Name)
	if c.Access.FabricScoped {
		cx.CreateAttr("isFabricScoped", "true")
	}
	if !mandatory {
		cx.CreateAttr("optional", "true")
	} else {
		cx.CreateAttr("optional", "false")
	}
	if len(c.Response) > 0 && c.Response != "Y" && c.Response != "N" {
		cx.CreateAttr("response", c.Response)
	}
	if c.Response == "N" && !serverSource {
		cx.CreateAttr("disableDefaultResponse", "true")
	}
	if c.Access.Timed {
		cx.CreateAttr("mustUseTimedInvoke", "true")
	}

	if c.Access.Invoke != matter.PrivilegeUnknown && c.Access.Invoke != matter.PrivilegeOperate {
		ax := cx.CreateElement("access")
		ax.CreateAttr("op", "invoke")
		ax.CreateAttr("privilege", renderPrivilege(c.Access.Invoke))
	}
	for _, f := range c.Fields {
		if conformance.IsZigbee(f.Conformance) {
			continue
		}
		fx := cx.CreateElement("arg")
		mandatory := conformance.IsMandatory(f.Conformance)
		fx.CreateAttr("name", f.Name)
		writeDataType(fx, c.Fields, f)
		renderConstraint(c.Fields, f, fx)
		if !mandatory {
			fx.CreateAttr("optional", "true")
		}
		if f.Quality.Has(matter.QualityNullable) {
			fx.CreateAttr("isNullable", "true")
		}
		defaultValue := zap.GetDefaultValue(&matter.ConstraintContext{Field: f, Fields: c.Fields})
		if defaultValue.Defined() {
			fx.CreateAttr("default", defaultValue.ZapString(f.Type))
		}

	}
	cx.CreateElement("description").SetText(c.Description)
}
