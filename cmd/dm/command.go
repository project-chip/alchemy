package dm

import (
	"strings"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/conformance"
	"github.com/hasty/alchemy/matter"
)

func renderCommands(cluster *matter.Cluster, c *etree.Element) (err error) {
	if len(cluster.Commands) == 0 {
		return
	}
	commands := c.CreateElement("commands")
	for _, cmd := range cluster.Commands {
		if conformance.IsZigbee(cmd.Conformance) {
			continue
		}
		cx := commands.CreateElement("command")
		cx.CreateAttr("id", cmd.ID.ShortHexString())
		cx.CreateAttr("name", cmd.Name)
		if cmd.Access.Invoke != matter.PrivilegeUnknown {
			a := cx.CreateElement("access")
			a.CreateAttr("invokePrivilege", strings.ToLower(matter.PrivilegeNamesShort[cmd.Access.Invoke]))
			if cmd.Access.FabricScoped {
				a.CreateAttr("fabricScoped", "true")
			}
			if cmd.Access.Timed {
				a.CreateAttr("timed", "true")
			}
		}

		if cmd.Direction == matter.InterfaceClient {
			cx.CreateAttr("direction", "responseFromServer")
		} else if cmd.Response != "" {
			cx.CreateAttr("response", cmd.Response)
		}
		err = renderConformanceString(cluster, cmd.Conformance, cx)
		if err != nil {
			return
		}

		for _, f := range cmd.Fields {
			if !f.ID.Valid() {
				continue
			}
			i := cx.CreateElement("field")
			i.CreateAttr("id", f.ID.IntString())
			i.CreateAttr("name", f.Name)
			renderDataType(f, i)
			if len(f.Default) > 0 {
				i.CreateAttr("default", f.Default)
			}
			err = renderConformanceString(cluster, f.Conformance, i)
			if err != nil {
				return
			}

			err = renderConstraint(f.Constraint, f.Type, i)
			if err != nil {
				return
			}

		}
	}

	return
}
