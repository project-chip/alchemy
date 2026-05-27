package render

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/internal/xml"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/zap"
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
			if conformance.IsZigbee(e.Conformance) || zap.IsDisallowed(e, e.Conformance) {
				continue
			}
			if e.ID.Equals(eventID) {
				matchingEvent = e
				delete(events, e)
				break
			}
		}

		if matchingEvent == nil {
			slog.Warn("Removing unrecognized event from ZAP XML", slog.String("path", cr.configurator.OutPath), slog.String("eventId", eventID.Text()))
			ce.RemoveChild(eve)
			continue
		}
		cr.populateEvent(eve, matchingEvent, cluster)
	}

	for event := range events {
		if cr.isProvisionalViolation(event) {
			err = fmt.Errorf("new event added without provisional conformance: %s", event.Name)
			return
		}
		ee := etree.NewElement("event")
		err = cr.populateEvent(ee, event, cluster)
		if err != nil {
			return
		}
		xml.InsertElementByAttribute(ce, ee, "code", "command", "attribute", "globalAttribute", "features")
	}
	return
}

func (cr *configuratorRenderer) populateEvent(eventElement *etree.Element, event *matter.Event, cluster *matter.Cluster) (err error) {
	cr.elementMap[eventElement] = event
	needsAccess := event.Access.Read != matter.PrivilegeUnknown && event.Access.Read != matter.PrivilegeView

	patchNumberAttribute(eventElement, event.ID, "code")
	eventElement.CreateAttr("name", event.Name)
	priority := strings.ToLower(event.Priority)
	eventElement.CreateAttr("priority", priority)
	eventElement.CreateAttr("side", "server")

	cr.setProvisional(eventElement, event)

	if event.Access.IsFabricSensitive() {
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
		descriptionElement.SetText(event.Description)
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
			if conformance.IsZigbee(f.Conformance) || zap.IsDisallowed(f, f.Conformance) {
				continue
			}
			if matter.NonGlobalIDInvalidForEntity(f.ID, types.EntityTypeEventField) {
				continue
			}
			fe.CreateAttr("fieldId", f.ID.IntString())
			cr.setFieldAttributes(fe, types.EntityTypeEvent, f, event.Fields, false)
			break
		}
	}
	for fieldIndex < len(event.Fields) {
		f := event.Fields[fieldIndex]
		fieldIndex++
		if conformance.IsZigbee(f.Conformance) || zap.IsDisallowed(f, f.Conformance) {
			continue
		}
		if matter.NonGlobalIDInvalidForEntity(f.ID, types.EntityTypeEventField) {
			continue
		}
		if cr.isProvisionalViolation(f) {
			err = fmt.Errorf("new event field added without provisional conformance: %s.%s.%s", cluster.Name, event.Name, f.Name)
			return
		}
		fe := etree.NewElement("field")
		fe.CreateAttr("fieldId", f.ID.IntString())
		cr.setFieldAttributes(fe, types.EntityTypeEvent, f, event.Fields, true)
		xml.AppendElement(eventElement, fe, "access")
	}

	if cluster != nil && cr.configurator != nil {
		if cr.generator.options.ConformanceXML {
			renderConformance(cr.generator.spec, event, event.Conformance, eventElement, "field", "access", "description")
		} else {
			removeConformance(eventElement)
		}
	}
	return
}
