package generate

import (
	"log/slog"
	"slices"
	"strings"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/internal/xml"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
	"github.com/hasty/alchemy/zap"
)

func generateCommands(configurator *zap.Configurator, ce *etree.Element, cluster *matter.Cluster, commands map[*matter.Command]struct{}, errata *zap.Errata) (err error) {

	for _, cmde := range ce.SelectElements("command") {

		code := cmde.SelectAttr("code")
		if code == nil {
			slog.Warn("missing code attribute in command", slog.String("path", configurator.Doc.Path))
			continue
		}
		source := cmde.SelectAttr("source")
		if code == nil {
			slog.Warn("missing source attribute in command", slog.String("path", configurator.Doc.Path))
			continue
		}
		commandID := matter.ParseNumber(code.Value)
		if !commandID.Valid() {
			slog.Warn("invalid code ID in command", slog.String("path", configurator.Doc.Path), slog.String("commandId", commandID.Text()))
			continue
		}

		var matchingCommand *matter.Command
		for c := range commands {
			if conformance.IsZigbee(cluster.Commands, c.Conformance) || conformance.IsDisallowed(c.Conformance) {
				continue
			}
			if c.ID.Equals(commandID) {
				if c.Direction == matter.InterfaceServer && source.Value == "client" {
					matchingCommand = c
					delete(commands, c)
					break
				}
				if c.Direction == matter.InterfaceClient && source.Value == "server" {
					matchingCommand = c
					delete(commands, c)
					break
				}
			}
		}

		if matchingCommand == nil {
			slog.Warn("unknown command ID", slog.String("path", configurator.Doc.Path), slog.String("commandId", commandID.Text()))
			ce.RemoveChild(cmde)
			continue
		}
		populateCommand(cmde, matchingCommand, cluster, errata)
	}

	var remainingCommands []*matter.Command
	for command := range commands {
		remainingCommands = append(remainingCommands, command)
	}
	slices.SortFunc(remainingCommands, func(a, b *matter.Command) int { return strings.Compare(a.Name, b.Name) })

	for _, command := range remainingCommands {
		if conformance.IsZigbee(cluster.Commands, command.Conformance) || conformance.IsDisallowed(command.Conformance) {
			continue
		}
		cme := etree.NewElement("command")
		cme.CreateAttr("code", command.ID.HexString())
		populateCommand(cme, command, cluster, errata)
		xml.InsertElementByAttribute(ce, cme, "code", "attribute")
	}
	return
}

func populateCommand(ce *etree.Element, c *matter.Command, cluster *matter.Cluster, errata *zap.Errata) {
	mandatory := conformance.IsMandatory(c.Conformance)

	var serverSource bool
	if c.Direction == matter.InterfaceServer {
		ce.CreateAttr("source", "client")
	} else if c.Direction == matter.InterfaceClient {
		ce.CreateAttr("source", "server")
		serverSource = true
	}
	ce.CreateAttr("code", c.ID.ShortHexString())
	ce.CreateAttr("name", zap.CleanName(c.Name))
	if c.Access.IsFabricScoped() {
		ce.CreateAttr("isFabricScoped", "true")
	} else {
		ce.RemoveAttr("isFabricScoped")
	}
	if !mandatory {
		ce.CreateAttr("optional", "true")
	} else {
		ce.CreateAttr("optional", "false")
	}
	if len(c.Response) > 0 && c.Response != "Y" && c.Response != "N" {
		ce.CreateAttr("response", c.Response)
	} else {
		ce.RemoveAttr("response")
	}
	if c.Response == "N" && serverSource {
		ce.CreateAttr("disableDefaultResponse", "true")
	} else {
		ce.RemoveAttr("disableDefaultResponse")
	}
	if c.Access.IsTimed() {
		ce.CreateAttr("mustUseTimedInvoke", "true")
	} else {
		ce.RemoveAttr("mustUseTimedInvoke")
	}

	argIndex := 0
	argElements := ce.SelectElements("arg")
	for _, fe := range argElements {
		for {
			if argIndex >= len(c.Fields) {
				ce.RemoveChild(fe)
				break
			}
			f := c.Fields[argIndex]
			argIndex++
			if conformance.IsZigbee(c.Fields, f.Conformance) || conformance.IsDisallowed(f.Conformance) {
				continue
			}
			fe.CreateAttr("id", f.ID.IntString())
			setFieldAttributes(fe, f, c.Fields)
			break
		}
	}
	for argIndex < len(c.Fields) {
		f := c.Fields[argIndex]
		argIndex++
		fe := etree.NewElement("arg")
		fe.CreateAttr("id", f.ID.IntString())
		setFieldAttributes(fe, f, c.Fields)
		xml.AppendElement(ce, fe)
	}
	needsAccess := c.Access.Invoke != matter.PrivilegeUnknown && c.Access.Invoke != matter.PrivilegeOperate
	if needsAccess {
		for _, el := range ce.SelectElements("access") {
			if needsAccess {
				setAccessAttributes(el, "invoke", c.Access.Invoke, errata)
				needsAccess = false
			} else {
				ce.RemoveChild(el)
			}
		}
	} else {
		for _, el := range ce.SelectElements("access") {
			ce.RemoveChild(el)
		}
	}
}
