package zcl

import (
	"fmt"
	"strings"

	"github.com/beevik/etree"
	"github.com/hasty/matterfmt/matter"
	"github.com/hasty/matterfmt/parse"
)

func renderCommands(cluster *matter.Cluster, cx *etree.Element, errata *errata) {
	for _, c := range cluster.Commands {
		if !strings.HasSuffix(c.Name, "Response") {
			renderCommand(c, cx, errata)
		}
	}
	for _, c := range cluster.Commands {
		if strings.HasSuffix(c.Name, "Response") {
			renderCommand(c, cx, errata)
		}
	}
}

func renderCommand(c *matter.Command, e *etree.Element, errata *errata) {
	mandatory := (c.Conformance == "M")

	cx := e.CreateElement("command")
	var serverSource bool
	if strings.HasPrefix(c.Direction, "client") || strings.HasPrefix(c.Direction, "Client") {
		cx.CreateAttr("source", "client")
	} else if strings.HasPrefix(c.Direction, "server") || strings.HasPrefix(c.Direction, "Server") {
		cx.CreateAttr("source", "server")
		serverSource = true
	}
	id := c.ID
	cid, err := parse.HexOrDec(id)
	if err == nil {
		id = fmt.Sprintf("%#02x", cid)
	}
	cx.CreateAttr("code", id)
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

	cx.CreateElement("description").SetText(c.Description)
	if c.Access.Invoke != matter.PrivilegeUnknown {
		ax := cx.CreateElement("access")
		ax.CreateAttr("op", "invoke")
		ax.CreateAttr("role", renderPrivilege(c.Access.Invoke))
	}
	for _, f := range c.Fields {
		fx := cx.CreateElement("arg")
		mandatory := (f.Conformance == "M")
		renderConstraint(f.Constraint, errata, fx)
		fx.CreateAttr("name", f.Name)
		writeDataType(fx, f.Type)
		if !mandatory {
			fx.CreateAttr("optional", "true")
		}
	}
}
