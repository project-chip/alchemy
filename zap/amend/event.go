package amend

import (
	"encoding/xml"
	"io"
	"log/slog"
	"strings"

	"github.com/hasty/alchemy/conformance"
	"github.com/hasty/alchemy/matter"
)

func (r *renderer) amendEvent(cluster *matter.Cluster, ts *tokenSet, e xmlEncoder, el xml.StartElement, events map[*matter.Event]struct{}) (err error) {
	code := getAttributeValue(el.Attr, "code")
	eventID := matter.ParseID(code)

	var matchingEvent *matter.Event
	for e := range events {
		if conformance.IsZigbee(cluster.Events, e.Conformance) {
			continue
		}
		if e.ID.Equals(eventID) {
			matchingEvent = e
			delete(events, e)
			break
		}
	}

	if matchingEvent == nil {
		Ignore(ts, "event")
		return nil
	}

	el.Attr = r.setEventAttributes(el.Attr, matchingEvent)
	err = e.EncodeToken(el)
	if err != nil {
		return
	}

	var fieldIndex int
	needsAccess := matchingEvent.Access.Read != matter.PrivilegeUnknown && matchingEvent.Access.Read != matter.PrivilegeView

	for {
		var tok xml.Token
		tok, err = ts.Token()
		if tok == nil || err == io.EOF {
			err = io.EOF
			return
		} else if err != nil {
			return
		}

		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "description":
				err = writeThrough(ts, e, t)
			case "field":
				for {
					if fieldIndex >= len(matchingEvent.Fields) {
						Ignore(ts, "field")
						break
					} else {
						f := matchingEvent.Fields[fieldIndex]
						fieldIndex++
						if conformance.IsZigbee(matchingEvent.Fields, f.Conformance) {
							continue
						}

						t.Attr = setAttributeValue(t.Attr, "id", f.ID.IntString())
						t.Attr = r.setFieldAttributes(f, t.Attr, matchingEvent.Fields)
						err = writeThrough(ts, e, t)
						if err != nil {
							return
						}
						break
					}
				}
			case "access":
				{
					if !needsAccess {
						Ignore(ts, "access")
					} else {
						r.setAccessAttributes(t.Attr, "read", matchingEvent.Access.Read)
						err = writeThrough(ts, e, t)
						needsAccess = false
					}
				}
			default:
				slog.Warn("unexpected element in event", "name", t.Name.Local)
				err = Ignore(ts, t.Name.Local)
			}
		case xml.EndElement:
			switch t.Name.Local {
			case "event":
				for fieldIndex < len(matchingEvent.Fields) {
					f := matchingEvent.Fields[fieldIndex]
					fieldIndex++
					elName := xml.Name{Local: "field"}
					xfs := xml.StartElement{Name: elName}
					xfs.Attr = setAttributeValue(xfs.Attr, "id", f.ID.IntString())
					xfs.Attr = r.setFieldAttributes(f, xfs.Attr, matchingEvent.Fields)
					err = e.EncodeToken(xfs)
					if err != nil {
						return
					}
					xfe := xml.EndElement{Name: elName}
					err = e.EncodeToken(xfe)
					if err != nil {
						return
					}
				}
				if needsAccess {
					err = r.renderAccess(e, "read", matchingEvent.Access.Read)
					if err != nil {
						return
					}
				}
				err = e.EncodeToken(t)
				return
			default:
				err = e.EncodeToken(tok)

			}
		case xml.CharData:
		default:
			err = e.EncodeToken(t)
		}
		if err != nil {
			return
		}
	}

	//return r.writeEvent(e, el, matchingEvent, false)
}

func (r *renderer) writeEvent(e xmlEncoder, el xml.StartElement, ev *matter.Event, provisional bool) (err error) {

	xfb := el.Copy()
	xfb.Name = xml.Name{Local: "event"}

	xfb.Attr = r.setEventAttributes(xfb.Attr, ev)

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
		if conformance.IsZigbee(ev.Fields, f.Conformance) {
			continue
		}
		if !f.ID.Valid() {
			continue
		}

		elName := xml.Name{Local: "field"}
		xfs := xml.StartElement{Name: elName}
		xfs.Attr = setAttributeValue(xfs.Attr, "id", f.ID.IntString())
		xfs.Attr = r.setFieldAttributes(f, xfs.Attr, ev.Fields)

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

func (*renderer) setEventAttributes(xfb []xml.Attr, ev *matter.Event) []xml.Attr {
	xfb = setAttributeValue(xfb, "code", ev.ID.HexString())
	xfb = setAttributeValue(xfb, "name", ev.Name)
	xfb = setAttributeValue(xfb, "priority", strings.ToLower(ev.Priority))
	xfb = removeAttribute(xfb, "side")

	if ev.FabricSensitive {
		xfb = setAttributeValue(xfb, "isFabricSensitive", "true")
	} else {
		xfb = removeAttribute(xfb, "isFabricSensitive")
	}
	if !conformance.IsMandatory(ev.Conformance) {
		xfb = setAttributeValue(xfb, "optional", "true")
	} else {
		xfb = removeAttribute(xfb, "optional")
	}
	return xfb
}
