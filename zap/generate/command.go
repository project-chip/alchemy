package generate

import (
	"log/slog"
	"slices"
	"strings"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/internal/xml"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/zap"
)

func (cr *configuratorRenderer) generateCommands(commands map[*matter.Command]struct{}, parent *etree.Element, cluster *matter.Cluster) (err error) {

	for _, cmde := range parent.SelectElements("command") {

		code := cmde.SelectAttr("code")
		if code == nil {
			slog.Warn("missing code attribute in command", slog.String("path", cr.configurator.OutPath))
			continue
		}
		source := cmde.SelectAttr("source")
		if code == nil {
			slog.Warn("missing source attribute in command", slog.String("path", cr.configurator.OutPath))
			continue
		}
		commandID := matter.ParseNumber(code.Value)
		if !commandID.Valid() {
			slog.Warn("invalid code ID in command", slog.String("path", cr.configurator.OutPath), slog.String("commandId", commandID.Text()))
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
			slog.Warn("unknown command ID", slog.String("path", cr.configurator.OutPath), slog.String("commandId", commandID.Text()))
			parent.RemoveChild(cmde)
			continue
		}
		if matter.NonGlobalIDInvalidForEntity(matchingCommand.ID, types.EntityTypeCommand) {
			continue
		}
		cr.populateCommand(cmde, cluster, matchingCommand)
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
		cr.populateCommand(cme, cluster, command)
		xml.InsertElementByAttribute(parent, cme, "code", "attribute", "globalAttribute")
	}
	return
}

func (cr *configuratorRenderer) populateCommand(ce *etree.Element, cluster *matter.Cluster, c *matter.Command) {
	cr.elementMap[ce] = c
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
				cr.setAccessAttributes(el, "invoke", c.Access.Invoke)
				needsAccess = false
			} else {
				ce.RemoveChild(el)
			}
		}
		if needsAccess {
			el := etree.NewElement("access")
			xml.AppendElement(ce, el, "description")
			cr.setAccessAttributes(el, "invoke", c.Access.Invoke)
		}
	} else {
		for _, el := range ce.SelectElements("access") {
			ce.RemoveChild(el)
		}
	}

	needsQuality := c.Quality.Has(matter.QualityLargeMessage)
	if needsQuality {
		for _, el := range ce.SelectElements("quality") {
			if needsQuality {
				cr.setQualityAttributes(el, c.Quality)
				needsQuality = false
			} else {
				ce.RemoveChild(el)
			}
		}
		if needsQuality {
			el := etree.NewElement("quality")
			xml.AppendElement(ce, el, "description", "access")
			cr.setQualityAttributes(el, c.Quality)
		}
	} else {
		for _, el := range ce.SelectElements("quality") {
			ce.RemoveChild(el)
		}
	}

	if cluster != nil && cr.generator != nil {
		if cr.generator.generateConformanceXML {
			renderConformance(cr.generator.spec, c, cluster, c.Conformance, ce)
		} else {
			removeConformance(ce)
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
			cr.setFieldAttributes(fe, f, c.Fields)
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
		cr.setFieldAttributes(fe, f, c.Fields)
		xml.AppendElement(ce, fe)
	}
}
