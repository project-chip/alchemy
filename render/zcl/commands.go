package zcl

import (
	"github.com/beevik/etree"
	"github.com/hasty/alchemy/matter"
)

func renderCommands(cluster *matter.Cluster, cx *etree.Element, errata *Errata) {
	for _, c := range cluster.Commands {
		if c.Direction == matter.CommandDirectionClientToServer {
			renderCommand(c, cx, errata)
		}
	}
	for _, c := range cluster.Commands {
		if c.Direction == matter.CommandDirectionServerToClient {
			renderCommand(c, cx, errata)
		}
	}
}

func renderCommand(c *matter.Command, e *etree.Element, errata *Errata) {
	mandatory := (c.Conformance == "M")

	cx := e.CreateElement("command")
	var serverSource bool
	if c.Direction == matter.CommandDirectionClientToServer {
		cx.CreateAttr("source", "client")
	} else if c.Direction == matter.CommandDirectionServerToClient {
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
		ax.CreateAttr("role", renderPrivilege(c.Access.Invoke))
	}
	for _, f := range c.Fields {
		if f.Conformance == "Zigbee" {
			continue
		}
		fx := cx.CreateElement("arg")
		mandatory := (f.Conformance == "M")
		renderConstraint(f.Constraint, fx)
		fx.CreateAttr("name", f.Name)
		writeCommandDataType(fx, f.Type)
		if !mandatory {
			fx.CreateAttr("optional", "true")
		}
		renderConstraint(f.Constraint, fx)
		fx.CreateAttr("apiMaturity", "provisional")
	}
	cx.CreateElement("description").SetText(c.Description)
}
