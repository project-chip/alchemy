package amend

import (
	"encoding/xml"
	"strings"

	"github.com/hasty/alchemy/conformance"
	"github.com/hasty/alchemy/matter"
)

func (r *renderer) amendEvent(ts *tokenSet, e xmlEncoder, el xml.StartElement, events map[*matter.Event]struct{}) (err error) {
	code := getAttributeValue(el.Attr, "code")
	eventID := matter.ParseID(code)

	var matchingEvent *matter.Event
	for e := range events {
		if e.ID.Equals(eventID) {
			matchingEvent = e
			delete(events, e)
			break
		}
	}

	Ignore(ts, "event")

	if matchingEvent == nil {
		return nil
	}

	return r.writeEvent(e, el, matchingEvent, false)
}

func (r *renderer) writeEvent(e xmlEncoder, el xml.StartElement, ev *matter.Event, provisional bool) (err error) {

	xfb := el.Copy()
	xfb.Name = xml.Name{Local: "event"}

	xfb.Attr = setAttributeValue(xfb.Attr, "code", ev.ID.HexString())
	xfb.Attr = setAttributeValue(xfb.Attr, "name", ev.Name)
	xfb.Attr = setAttributeValue(xfb.Attr, "priority", strings.ToLower(ev.Priority))
	xfb.Attr = setAttributeValue(xfb.Attr, "side", "server")

	if ev.FabricSensitive {
		xfb.Attr = setAttributeValue(xfb.Attr, "isFabricSensitive", "true")
	} else {
		xfb.Attr = removeAttribute(xfb.Attr, "isFabricSensitive")
	}
	if !conformance.IsMandatory(ev.Conformance) {
		xfb.Attr = setAttributeValue(xfb.Attr, "optional", "true")
	} else {
		xfb.Attr = removeAttribute(xfb.Attr, "optional")
	}

	err = e.EncodeToken(xfb)
	if err != nil {
		return
	}

	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "description"}})
	desc := strings.TrimSpace(ev.Description)
	if len(desc) > 0 {
		e.EncodeToken(xml.CharData(desc))
	} else {
		e.EncodeToken(xml.CharData(ev.Name))
	}
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "description"}})

	for _, f := range ev.Fields {
		if conformance.IsZigbee(f.Conformance) {
			continue
		}
		if !f.ID.Valid() {
			continue
		}
		mandatory := conformance.IsMandatory(f.Conformance)

		elName := xml.Name{Local: "field"}
		xfs := xml.StartElement{Name: elName}
		xfs.Attr = setAttributeValue(xfs.Attr, "id", f.ID.IntString())
		xfs.Attr = setAttributeValue(xfs.Attr, "name", f.Name)
		xfs.Attr = writeDataType(ev.Fields, f, xfs.Attr)
		xfs.Attr = r.renderConstraint(ev.Fields, f, xfs.Attr)

		if f.Quality.Has(matter.QualityNullable) {
			xfs.Attr = setAttributeValue(xfs.Attr, "isNullable", "true")
		} else {
			xfs.Attr = removeAttribute(xfs.Attr, "isNullable")
		}
		if !mandatory {
			xfs.Attr = setAttributeValue(xfs.Attr, "optional", "true")
		} else {
			xfs.Attr = removeAttribute(xfs.Attr, "optional")
		}

		err = e.EncodeToken(xfs)
		if err != nil {
			return
		}

		if f.Access.Read != matter.PrivilegeUnknown {
			aName := xml.Name{Local: "access"}
			xa := xml.StartElement{Name: aName}
			xa.Attr = setAttributeValue(xa.Attr, "op", "read")
			p, _ := f.Access.Read.MarshalXMLAttr(xml.Name{Local: "privilege"})
			xa.Attr = append(xa.Attr, p)
			err = e.EncodeToken(xa)
			if err != nil {
				return
			}
			xae := xml.EndElement{Name: aName}
			err = e.EncodeToken(xae)
			if err != nil {
				return
			}
		}

		if f.Access.Write != matter.PrivilegeUnknown {
			aName := xml.Name{Local: "access"}
			xa := xml.StartElement{Name: aName}
			xa.Attr = setAttributeValue(xa.Attr, "op", "write")
			p, _ := f.Access.Read.MarshalXMLAttr(xml.Name{Local: "privilege"})
			xa.Attr = append(xa.Attr, p)
			err = e.EncodeToken(xa)
			if err != nil {
				return
			}
			xae := xml.EndElement{Name: aName}
			err = e.EncodeToken(xae)
			if err != nil {
				return
			}
		}

		xfe := xml.EndElement{Name: elName}
		err = e.EncodeToken(xfe)
		if err != nil {
			return
		}

	}
	return e.EncodeToken(xml.EndElement{Name: xfb.Name})
}
