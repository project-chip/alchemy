package render

import (
	"github.com/beevik/etree"
	"github.com/hasty/alchemy/matter"
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
	mandatory := (c.Conformance == "M")

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
	cx.CreateAttr("apiMaturity", "provisional")

	if c.Access.Invoke != matter.PrivilegeUnknown {
		ax := cx.CreateElement("access")
		ax.CreateAttr("op", "invoke")
		ax.CreateAttr("privilege", renderPrivilege(c.Access.Invoke))
	}
	for _, f := range c.Fields {
		if f.Conformance == "Zigbee" {
			continue
		}
		fx := cx.CreateElement("arg")
		mandatory := (f.Conformance == "M")
		fx.CreateAttr("name", f.Name)
		writeDataType(fx, c.Fields, f)
		renderConstraint(c.Fields, f, fx)
		if !mandatory {
			fx.CreateAttr("optional", "true")
		}
		if f.Quality.Has(matter.QualityNullable) {
			fx.CreateAttr("isNullable", "true")
		}
		//fx.CreateAttr("apiMaturity", "provisional")
	}
	cx.CreateElement("description").SetText(c.Description)
}
