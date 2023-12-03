package render

import (
	"slices"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/conformance"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/zap"
)

func renderCommands(cluster *matter.Cluster, cx *etree.Element, errata *zap.Errata) {
	commands := make([]*matter.Command, 0, len(cluster.Commands))

	for _, c := range cluster.Commands {
		if conformance.IsZigbee(cluster.Commands, c.Conformance) {
			continue
		}
		commands = append(commands, c)
	}

	slices.SortFunc(commands, func(a, b *matter.Command) int {
		if a.Direction == b.Direction {
			return a.ID.Compare(b.ID)
		} else if a.Direction == matter.InterfaceServer {
			return -1
		} else if a.Direction == matter.InterfaceClient {
			return 1
		} else {
			return 0
		}
	})

	for _, c := range commands {
		renderCommand(c, cx, errata)
	}
}

func renderCommand(c *matter.Command, e *etree.Element, errata *zap.Errata) {
	mandatory := conformance.IsMandatory(c.Conformance)

	cx := e.CreateElement("command")
	var serverSource bool
	if c.Direction == matter.InterfaceServer {
		cx.CreateAttr("source", "client")
	} else if c.Direction == matter.InterfaceClient {
		cx.CreateAttr("source", "server")
		serverSource = true
	}
	cx.CreateAttr("code", c.ID.ShortHexString())
	cx.CreateAttr("name", zap.CleanName(c.Name))
	if c.Access.FabricScoped {
		cx.CreateAttr("isFabricScoped", "true")
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
	if !mandatory {
		cx.CreateAttr("optional", "true")
	} else {
		cx.CreateAttr("optional", "false")
	}

	if c.Access.Invoke != matter.PrivilegeUnknown && c.Access.Invoke != matter.PrivilegeOperate {
		ax := cx.CreateElement("access")
		ax.CreateAttr("op", "invoke")
		ax.CreateAttr("privilege", renderPrivilege(c.Access.Invoke))
	}
	cx.CreateElement("description").SetText(c.Description)
	for _, f := range c.Fields {
		if conformance.IsZigbee(c.Fields, f.Conformance) {
			continue
		}
		fx := cx.CreateElement("arg")
		mandatory := conformance.IsMandatory(f.Conformance)
		fx.CreateAttr("name", zap.CleanName(f.Name))
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
}
