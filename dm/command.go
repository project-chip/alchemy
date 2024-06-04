package dm

import (
	"slices"
	"strings"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
	"github.com/hasty/alchemy/matter/spec"
)

func renderCommands(doc *spec.Doc, cluster *matter.Cluster, c *etree.Element) (err error) {
	if len(cluster.Commands) == 0 {
		return
	}

	cmds := make([]*matter.Command, 0, len(cluster.Commands))
	for _, c := range cluster.Commands {
		if conformance.IsZigbee(cluster.Commands, c.Conformance) {
			continue
		}
		cmds = append(cmds, c)
	}

	slices.SortFunc(cmds, func(a, b *matter.Command) int {
		if a.ID.Equals(b.ID) {
			if a.Direction == b.Direction {
				return 0
			}
			if a.Direction == matter.InterfaceServer {
				return -1
			}
			return 1
		}
		return a.ID.Compare(b.ID)
	})

	commands := c.CreateElement("commands")
	for _, cmd := range cmds {
		cx := commands.CreateElement("command")
		cx.CreateAttr("id", cmd.ID.ShortHexString())
		cx.CreateAttr("name", cmd.Name)
		switch cmd.Direction {
		case matter.InterfaceClient:
			cx.CreateAttr("direction", "responseFromServer")
			if cmd.Access.IsFabricScoped() {
				a := cx.CreateElement("access")
				if cmd.Access.IsFabricScoped() {
					a.CreateAttr("fabricScoped", "true")
				}
				if cmd.Access.IsTimed() {
					a.CreateAttr("timed", "true")
				}
			}
		case matter.InterfaceServer:
			cx.CreateAttr("direction", "commandToServer")
			if cmd.Response != "" {
				cx.CreateAttr("response", cmd.Response)
			}
			if cmd.Access.Invoke != matter.PrivilegeUnknown || cmd.Access.IsFabricScoped() || cmd.Access.IsTimed() {
				a := cx.CreateElement("access")
				if cmd.Access.Invoke != matter.PrivilegeUnknown {
					a.CreateAttr("invokePrivilege", strings.ToLower(matter.PrivilegeNamesShort[cmd.Access.Invoke]))
				}
				if cmd.Access.IsFabricScoped() {
					a.CreateAttr("fabricScoped", "true")
				}
				if cmd.Access.IsTimed() {
					a.CreateAttr("timed", "true")
				}
			}
		}

		err = renderConformanceString(doc, cluster, cmd.Conformance, cx)
		if err != nil {
			return
		}

		for _, f := range cmd.Fields {
			i := cx.CreateElement("field")
			if f.ID.Valid() {
				i.CreateAttr("id", f.ID.IntString())
			}
			i.CreateAttr("name", f.Name)
			renderDataType(f, i)
			if len(f.Default) > 0 {
				i.CreateAttr("default", f.Default)
			}
			renderQuality(i, f.Quality, matter.QualityNullable)
			err = renderConformanceString(doc, cmd.Fields, f.Conformance, i)
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
