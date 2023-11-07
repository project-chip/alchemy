package render

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/hasty/matterfmt/matter"
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

	if matchingEvent == nil {
		fmt.Errorf("no matching event for %s\n", eventID.HexString())
		return writeThrough(ts, e, el)
	}

	Ignore(ts, "event")

	return r.writeEvent(e, el, matchingEvent)
}

func (r *renderer) writeEvent(e xmlEncoder, el xml.StartElement, ev *matter.Event) (err error) {

	xfb := el.Copy()
	xfb.Name = xml.Name{Local: "event"}

	xfb.Attr = setAttributeValue(xfb.Attr, "code", ev.ID.HexString())
	xfb.Attr = setAttributeValue(xfb.Attr, "name", ev.Name)
	xfb.Attr = setAttributeValue(xfb.Attr, "priority", strings.ToLower(ev.Priority))
	xfb.Attr = setAttributeValue(xfb.Attr, "side", "server")

	if ev.Access.FabricSensitive {
		xfb.Attr = setAttributeValue(xfb.Attr, "isFabricSensitive", "true")
	}
	if ev.Conformance != "M" {
		xfb.Attr = setAttributeValue(xfb.Attr, "optional", "true")
	}

	err = e.EncodeToken(xfb)
	if err != nil {
		return
	}

	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "description"}})
	if len(ev.Description) > 0 {
		e.EncodeToken(xml.CharData(ev.Description))
	} else {
		e.EncodeToken(xml.CharData(ev.Name))
	}
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "description"}})

	for _, f := range ev.Fields {
		if f.Conformance == "Zigbee" {
			continue
		}
		if !f.ID.Valid() {
			continue
		}

		elName := xml.Name{Local: "field"}
		xfs := xml.StartElement{Name: elName}
		xfs.Attr = setAttributeValue(xfs.Attr, "name", f.Name)
		xfs.Attr = writeDataType(f.Type, xfs.Attr)

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
