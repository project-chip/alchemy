package generate

import (
	"log/slog"
	"strings"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/internal/xml"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

func (cr *configuratorRenderer) generateEvents(ce *etree.Element, cluster *matter.Cluster, events map[*matter.Event]struct{}) (err error) {

	for _, eve := range ce.SelectElements("event") {

		code := eve.SelectAttr("code")
		if code == nil {
			slog.Warn("missing code attribute in event", slog.String("path", cr.configurator.OutPath))
			continue
		}
		eventID := matter.ParseNumber(code.Value)
		if !eventID.Valid() {
			slog.Warn("invalid code ID in event", slog.String("path", cr.configurator.OutPath), slog.String("commandId", eventID.Text()))
			continue
		}

		var matchingEvent *matter.Event
		for e := range events {
			if conformance.IsZigbee(cluster.Commands, e.Conformance) || conformance.IsDisallowed(e.Conformance) {
				continue
			}
			if e.ID.Equals(eventID) {
				matchingEvent = e
				delete(events, e)
				break
			}
		}

		if matchingEvent == nil {
			slog.Warn("unknown event ID", slog.String("path", cr.configurator.OutPath), slog.String("eventId", eventID.Text()))
			ce.RemoveChild(eve)
			continue
		}
		cr.populateEvent(eve, matchingEvent, cluster)
	}

	for event := range events {
		ee := etree.NewElement("event")
		cr.populateEvent(ee, event, cluster)
		xml.InsertElementByAttribute(ce, ee, "code", "command", "attribute", "globalAttribute")
	}
	return
}

func (cr *configuratorRenderer) populateEvent(eventElement *etree.Element, event *matter.Event, cluster *matter.Cluster) {
	cr.elementMap[eventElement] = event
	needsAccess := event.Access.Read != matter.PrivilegeUnknown && event.Access.Read != matter.PrivilegeView

	patchNumberAttribute(eventElement, event.ID, "code")
	eventElement.CreateAttr("name", event.Name)
	priority := cr.configurator.Errata.EventPriority(event.Name, strings.ToLower(event.Priority))
	eventElement.CreateAttr("priority", priority)
	eventElement.CreateAttr("side", "server")

	if event.Access.FabricSensitivity == matter.FabricSensitivitySensitive {
		eventElement.CreateAttr("isFabricSensitive", "true")
	} else {
		eventElement.RemoveAttr("isFabricSensitive")
	}
	if !conformance.IsMandatory(event.Conformance) {
		eventElement.CreateAttr("optional", "true")
	} else {
		eventElement.RemoveAttr("optional")
	}

	descriptionElement := eventElement.SelectElement("description")
	if descriptionElement == nil {
		descriptionElement = etree.NewElement("description")
		eventElement.Child = append([]etree.Token{descriptionElement}, eventElement.Child...)
	}
	if len(event.Description) > 0 {
		descriptionElement.SetText(cr.configurator.Errata.TypeDescription(types.EntityTypeEvent, event.Name, event.Description))
	}

	if needsAccess {
		for _, el := range eventElement.SelectElements("access") {
			if needsAccess {
				cr.setAccessAttributes(el, "read", event.Access.Read)
				needsAccess = false
			} else {
				eventElement.RemoveChild(el)
			}
		}
		if needsAccess {
			cr.setAccessAttributes(eventElement.CreateElement("access"), "read", event.Access.Read)
		}

	} else {
		for _, el := range eventElement.SelectElements("access") {
			eventElement.RemoveChild(el)
		}
	}

	fieldIndex := 0
	fieldElements := eventElement.SelectElements("field")
	for _, fe := range fieldElements {
		for {
			if fieldIndex >= len(event.Fields) {
				eventElement.RemoveChild(fe)
				break
			}
			f := event.Fields[fieldIndex]
			fieldIndex++
			if conformance.IsZigbee(event.Fields, f.Conformance) || conformance.IsDisallowed(f.Conformance) {
				continue
			}
			if matter.NonGlobalIDInvalidForEntity(f.ID, types.EntityTypeEventField) {
				continue
			}
			fe.CreateAttr("id", f.ID.IntString())
			cr.setFieldAttributes(fe, types.EntityTypeEvent, event.Name, f, event.Fields)
			break
		}
	}
	for fieldIndex < len(event.Fields) {
		f := event.Fields[fieldIndex]
		fieldIndex++
		if conformance.IsZigbee(event.Fields, f.Conformance) || conformance.IsDisallowed(f.Conformance) {
			continue
		}
		if matter.NonGlobalIDInvalidForEntity(f.ID, types.EntityTypeEventField) {
			continue
		}
		fe := etree.NewElement("field")
		fe.CreateAttr("id", f.ID.IntString())
		cr.setFieldAttributes(fe, types.EntityTypeEvent, event.Name, f, event.Fields)
		xml.AppendElement(eventElement, fe, "access")
	}

	if cluster != nil && cr.configurator != nil {
		if cr.generator.generateConformanceXML {
			renderConformance(cr.generator.spec, event, cluster, event.Conformance, eventElement, "field", "access", "description")
		} else {
			removeConformance(eventElement)
		}
	}
}
