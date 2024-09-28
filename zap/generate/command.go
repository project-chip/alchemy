package generate

import (
	"log/slog"
	"slices"
	"strings"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/xml"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/zap"
)

func generateCommands(commands map[*matter.Command][]*matter.Number, docPath string, parent *etree.Element, errata *errata.ZAP) (err error) {

	for _, cmde := range parent.SelectElements("command") {

		code := cmde.SelectAttr("code")
		if code == nil {
			slog.Warn("missing code attribute in command", slog.String("path", docPath))
			continue
		}
		source := cmde.SelectAttr("source")
		if code == nil {
			slog.Warn("missing source attribute in command", slog.String("path", docPath))
			continue
		}
		commandID := matter.ParseNumber(code.Value)
		if !commandID.Valid() {
			slog.Warn("invalid code ID in command", slog.String("path", docPath), slog.String("commandId", commandID.Text()))
			continue
		}

		var matchingCommand *matter.Command
		for c := range commands {
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
			slog.Warn("unknown command ID", slog.String("path", docPath), slog.String("commandId", commandID.Text()))
			parent.RemoveChild(cmde)
			continue
		}
		if matter.NonGlobalIDInvalidForEntity(matchingCommand.ID, types.EntityTypeCommand) {
			continue
		}
		populateCommand(cmde, matchingCommand, errata)
	}

	var remainingCommands []*matter.Command
	for command := range commands {
		remainingCommands = append(remainingCommands, command)
	}
	slices.SortStableFunc(remainingCommands, func(a, b *matter.Command) int { return strings.Compare(a.Name, b.Name) })

	for _, command := range remainingCommands {
		if matter.NonGlobalIDInvalidForEntity(command.ID, types.EntityTypeCommand) {
			continue
		}
		cme := etree.NewElement("command")
		cme.CreateAttr("code", command.ID.HexString())
		populateCommand(cme, command, errata)
		xml.InsertElementByAttribute(parent, cme, "code", "attribute")
	}
	return
}

func populateCommand(ce *etree.Element, c *matter.Command, errata *errata.ZAP) {
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
	if c.Response != nil && c.Response.Name != "Y" && c.Response.Name != "N" {
		ce.CreateAttr("response", c.Response.Name)
	} else {
		ce.RemoveAttr("response")
	}
	if c.Response != nil && c.Response.Name == "N" && serverSource {
		ce.CreateAttr("disableDefaultResponse", "true")
	} else {
		ce.RemoveAttr("disableDefaultResponse")
	}
	if c.Access.IsTimed() {
		ce.CreateAttr("mustUseTimedInvoke", "true")
	} else {
		ce.RemoveAttr("mustUseTimedInvoke")
	}

	de := ce.SelectElement("description")
	if de == nil {
		de = etree.NewElement("description")
		ce.Child = append([]etree.Token{de}, ce.Child...)
	}
	if len(c.Description) > 0 {
		de.SetText(c.Description)
	}

	needsAccess := c.Access.Invoke != matter.PrivilegeUnknown && c.Access.Invoke != matter.PrivilegeOperate && c.Direction != matter.InterfaceClient
	if needsAccess {
		for _, el := range ce.SelectElements("access") {
			if needsAccess {
				setAccessAttributes(el, "invoke", c.Access.Invoke, errata)
				needsAccess = false
			} else {
				ce.RemoveChild(el)
			}
		}
		if needsAccess {
			el := etree.NewElement("access")
			xml.AppendElement(ce, el, "description")
			setAccessAttributes(el, "invoke", c.Access.Invoke, errata)
		}
	} else {
		for _, el := range ce.SelectElements("access") {
			ce.RemoveChild(el)
		}
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
			if matter.NonGlobalIDInvalidForEntity(f.ID, types.EntityTypeCommandField) {
				continue
			}
			xml.PrependAttribute(fe, "id", f.ID.IntString())
			setFieldAttributes(fe, f, c.Fields, errata)
			break
		}
	}
	for argIndex < len(c.Fields) {
		f := c.Fields[argIndex]
		argIndex++
		if conformance.IsZigbee(c.Fields, f.Conformance) || conformance.IsDisallowed(f.Conformance) {
			continue
		}
		if matter.NonGlobalIDInvalidForEntity(f.ID, types.EntityTypeCommandField) {
			continue
		}
		fe := ce.CreateElement("arg")
		fe.CreateAttr("id", f.ID.IntString())
		setFieldAttributes(fe, f, c.Fields, errata)
		xml.AppendElement(ce, fe)
	}
}
