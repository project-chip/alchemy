package generate

import (
	"encoding/xml"
	"log/slog"
	"strings"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/log"
	axml "github.com/project-chip/alchemy/internal/xml"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/zap"
)

func generateEvents(configurator *zap.Configurator, ce *etree.Element, cluster *matter.Cluster, events map[*matter.Event]struct{}, errata *errata.ZAP) (err error) {

	for _, eve := range ce.SelectElements("event") {

		code := eve.SelectAttr("code")
		if code == nil {
			slog.Warn("missing code attribute in event", log.Path("path", configurator.Doc.Path))
			continue
		}
		eventID := matter.ParseNumber(code.Value)
		if !eventID.Valid() {
			slog.Warn("invalid code ID in event", log.Path("path", configurator.Doc.Path), slog.String("commandId", eventID.Text()))
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
			slog.Warn("unknown event ID", log.Path("path", configurator.Doc.Path), slog.String("eventId", eventID.Text()))
			ce.RemoveChild(eve)
			continue
		}
		populateEvent(eve, matchingEvent, cluster, errata)
	}

	for event := range events {
		ee := etree.NewElement("event")
		populateEvent(ee, event, cluster, errata)
		axml.InsertElementByAttribute(ce, ee, "code", "command", "attribute")
	}
	return
}

func populateEvent(ee *etree.Element, e *matter.Event, cluster *matter.Cluster, errata *errata.ZAP) {
	needsAccess := e.Access.Read != matter.PrivilegeUnknown && e.Access.Read != matter.PrivilegeView

	patchNumberAttribute(ee, e.ID, "code")
	ee.CreateAttr("name", e.Name)
	ee.CreateAttr("priority", strings.ToLower(e.Priority))
	ee.CreateAttr("side", "server")

	if e.Access.FabricSensitivity == matter.FabricSensitivitySensitive {
		ee.CreateAttr("isFabricSensitive", "true")
	} else {
		ee.RemoveAttr("isFabricSensitive")
	}
	if !conformance.IsMandatory(e.Conformance) {
		ee.CreateAttr("optional", "true")
	} else {
		ee.RemoveAttr("optional")
	}

	de := ee.SelectElement("description")
	if de == nil {
		de = etree.NewElement("description")
		ee.Child = append([]etree.Token{de}, ee.Child...)
	}
	if len(e.Description) > 0 {
		de.SetText(e.Description)
	}

	fieldIndex := 0
	fieldElements := ee.SelectElements("field")
	for _, fe := range fieldElements {
		for {
			if fieldIndex >= len(e.Fields) {
				ee.RemoveChild(fe)
				break
			}
			f := e.Fields[fieldIndex]
			fieldIndex++
			if conformance.IsZigbee(e.Fields, f.Conformance) || conformance.IsDisallowed(f.Conformance) {
				continue
			}
			fe.CreateAttr("id", f.ID.IntString())
			setFieldAttributes(fe, f, e.Fields, errata)
			break
		}
	}
	for fieldIndex < len(e.Fields) {
		f := e.Fields[fieldIndex]
		fieldIndex++
		if conformance.IsZigbee(e.Fields, f.Conformance) || conformance.IsDisallowed(f.Conformance) {
			continue
		}
		fe := etree.NewElement("field")
		fe.CreateAttr("id", f.ID.IntString())
		setFieldAttributes(fe, f, e.Fields, errata)
		axml.AppendElement(ee, fe)
	}
	if needsAccess {
		for _, el := range ee.SelectElements("access") {
			if needsAccess {
				setAccessAttributes(el, "read", e.Access.Read, errata)
				needsAccess = false
			} else {
				ee.RemoveChild(el)
			}
		}
		if needsAccess {
			el := ee.CreateElement("access")
			setAccessAttributes(el, "read", e.Access.Read, errata)
		}

	} else {
		for _, el := range ee.SelectElements("access") {
			ee.RemoveChild(el)
		}
	}
}

func setFieldAttributes(e *etree.Element, f *matter.Field, fs matter.FieldSet, errata *errata.ZAP) {
	mandatory := conformance.IsMandatory(f.Conformance)
	e.CreateAttr("name", f.Name)
	writeDataType(e, fs, f, errata)
	if !mandatory {
		e.CreateAttr("optional", "true")
	} else {
		e.RemoveAttr("optional")
	}
	if f.Quality.Has(matter.QualityNullable) {
		e.CreateAttr("isNullable", "true")
	} else {
		e.RemoveAttr("isNullable")
	}
	if f.Access.IsFabricSensitive() {
		e.CreateAttr("isFabricSensitive", "true")
	} else {
		e.RemoveAttr("isFabricSensitive")
	}
	setFieldDefault(e, f, fs)
	renderConstraint(e, fs, f)
}

func writeDataType(e *etree.Element, fs matter.FieldSet, f *matter.Field, errata *errata.ZAP) {
	if f.Type == nil {
		return
	}
	dts := getDataTypeString(fs, f)
	dts = errata.TypeName(dts)
	if f.Type.IsArray() {
		e.CreateAttr("array", "true")
		e.CreateAttr("type", dts)
	} else {
		e.CreateAttr("type", dts)
		e.RemoveAttr("array")
	}
}

func getDataTypeString(fs matter.FieldSet, f *matter.Field) string {
	switch f.Type.BaseType {
	case types.BaseDataTypeTag:
		if f.Type.Entity != nil {
			if namespace, ok := f.Type.Entity.(*matter.Namespace); ok {
				return matterNamespaceName(namespace)
			}
		} else {
			return "enum8"
		}
	case types.BaseDataTypeNamespaceID:
		return "enum8"
	}
	return zap.FieldToZapDataType(fs, f)
}

func setAccessAttributes(el *etree.Element, op string, p matter.Privilege, errata *errata.ZAP) {
	el.CreateAttr("op", op)
	role := el.SelectAttr("role")
	var name string
	if role != nil {
		name = "role"
	} else if errata.WritePrivilegeAsRole {
		name = "role"
		el.RemoveAttr("privilege")
	} else {
		name = "privilege"
		el.RemoveAttr("role")
	}
	px, _ := p.MarshalXMLAttr(xml.Name{Local: name})
	el.CreateAttr(name, px.Value)
}
