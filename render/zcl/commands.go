package zcl

import (
	"fmt"
	"strings"

	"github.com/beevik/etree"
	"github.com/hasty/matterfmt/matter"
	"github.com/hasty/matterfmt/parse"
)

func renderCommands(cluster *matter.Cluster, cx *etree.Element) {
	for _, c := range cluster.Commands {
		if !strings.HasSuffix(c.Name, "Response") {
			renderCommand(c, cx)
		}
	}
	for _, c := range cluster.Commands {
		if strings.HasSuffix(c.Name, "Response") {
			renderCommand(c, cx)
		}
	}
}

func renderCommand(c *matter.Command, e *etree.Element) {
	mandatory := (c.Conformance == "M")

	readAccess, _, fabricScoped, _, timed := matter.ParseAccessValues(c.Access)

	cx := e.CreateElement("command")
	if strings.HasPrefix(c.Direction, "client") || strings.HasPrefix(c.Direction, "Client") {
		cx.CreateAttr("source", "client")
	} else if strings.HasPrefix(c.Direction, "server") || strings.HasPrefix(c.Direction, "Server") {
		cx.CreateAttr("source", "server")
	}
	id := c.ID
	cid, err := parse.ID(id)
	if err == nil {
		id = fmt.Sprintf("%#02x", cid)
	}
	cx.CreateAttr("code", id)
	cx.CreateAttr("name", c.Name)
	if fabricScoped == 1 {
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
	if c.Response == "N" {
		cx.CreateAttr("disableDefaultResponse", "true")
	}
	if timed == 1 {
		cx.CreateAttr("mustUseTimedInvoke", "true")
	}

	cx.CreateElement("description").SetText(c.Description)
	if readAccess != "" {
		ax := cx.CreateElement("access")
		ax.CreateAttr("op", "invoke")
		ax.CreateAttr("role", renderAccess(readAccess))
	}
	for _, f := range c.Fields {
		fx := cx.CreateElement("arg")
		mandatory := (f.Conformance == "M")
		/*id, err := parse.ID(f.ID)
		if err != nil {
			continue
		}
		fx.CreateAttr("id", strconv.Itoa(int(id)))*/
		fx.CreateAttr("name", f.Name)
		writeDataType(fx, f.Type)
		if !mandatory {
			fx.CreateAttr("optional", "true")
		}
	}
}
