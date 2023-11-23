package amend

import (
	"encoding/xml"

	"github.com/hasty/alchemy/conformance"
	"github.com/hasty/alchemy/matter"
)

func (r *renderer) amendCommand(ts *tokenSet, e xmlEncoder, el xml.StartElement, commands map[*matter.Command]struct{}) (err error) {
	code := getAttributeValue(el.Attr, "code")
	source := getAttributeValue(el.Attr, "source")
	commandID := matter.ParseID(code)

	var matchingCommand *matter.Command
	for c := range commands {
		if c.ID.Equals(commandID) {
			if c.Direction == matter.InterfaceServer && source == "client" {
				matchingCommand = c
				delete(commands, c)
				break
			}
			if c.Direction == matter.InterfaceClient && source == "server" {
				matchingCommand = c
				delete(commands, c)
				break
			}
		}
	}

	Ignore(ts, "command")

	if matchingCommand == nil {
		return nil
	}

	return r.writeCommand(e, el, matchingCommand)
}

func (r *renderer) writeCommand(e xmlEncoder, el xml.StartElement, c *matter.Command) (err error) {

	xfb := el.Copy()

	xfb = r.setCommandElementAttributes(c, e, xfb)

	err = e.EncodeToken(xfb)
	if err != nil {
		return
	}

	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "description"}})
	e.EncodeToken(xml.CharData(c.Description))
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "description"}})

	if c.Access.Invoke != matter.PrivilegeUnknown && c.Access.Invoke != matter.PrivilegeOperate {

		err = r.renderAccess(e, "invoke", c.Access.Invoke)
		if err != nil {
			return
		}
	}

	for _, f := range c.Fields {
		if conformance.IsZigbee(f.Conformance) {
			continue
		}

		elName := xml.Name{Local: "arg"}
		xfs := xml.StartElement{Name: elName}
		mandatory := conformance.IsMandatory(f.Conformance)
		xfs.Attr = setAttributeValue(xfs.Attr, "name", f.Name)
		xfs.Attr = writeDataType(c.Fields, f, xfs.Attr)
		if !mandatory {
			xfs.Attr = setAttributeValue(xfs.Attr, "optional", "true")
		} else {
			xfs.Attr = removeAttribute(xfs.Attr, "optional")
		}
		if f.Quality.Has(matter.QualityNullable) {
			xfs.Attr = setAttributeValue(xfs.Attr, "isNullable", "true")
		} else {
			xfs.Attr = removeAttribute(xfs.Attr, "isNullable")
		}
		xfs.Attr = r.renderConstraint(c.Fields, f, xfs.Attr)
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

	return e.EncodeToken(xml.EndElement{Name: xfb.Name})
}

func (*renderer) setCommandElementAttributes(c *matter.Command, e xmlEncoder, xfb xml.StartElement) xml.StartElement {
	mandatory := conformance.IsMandatory(c.Conformance)

	xfb.Name = xml.Name{Local: "command"}

	var serverSource bool
	if c.Direction == matter.InterfaceServer {
		xfb.Attr = setAttributeValue(xfb.Attr, "source", "client")
	} else if c.Direction == matter.InterfaceClient {
		xfb.Attr = setAttributeValue(xfb.Attr, "source", "server")
		serverSource = true
	}
	xfb.Attr = setAttributeValue(xfb.Attr, "code", c.ID.ShortHexString())
	xfb.Attr = setAttributeValue(xfb.Attr, "name", c.Name)
	if c.Access.FabricScoped {
		xfb.Attr = setAttributeValue(xfb.Attr, "isFabricScoped", "true")
	} else {
		xfb.Attr = removeAttribute(xfb.Attr, "isFabricScoped")
	}
	if !mandatory {
		xfb.Attr = setAttributeValue(xfb.Attr, "optional", "true")
	} else {
		xfb.Attr = setAttributeValue(xfb.Attr, "optional", "false")
	}
	if len(c.Response) > 0 && c.Response != "Y" && c.Response != "N" {
		xfb.Attr = setAttributeValue(xfb.Attr, "response", c.Response)
	} else {
		xfb.Attr = removeAttribute(xfb.Attr, "response")
	}
	if c.Response == "N" && !serverSource {
		xfb.Attr = setAttributeValue(xfb.Attr, "disableDefaultResponse", "true")
	} else {
		xfb.Attr = removeAttribute(xfb.Attr, "disableDefaultResponse")
	}
	if c.Access.Timed {
		xfb.Attr = setAttributeValue(xfb.Attr, "mustUseTimedInvoke", "true")
	} else {
		xfb.Attr = removeAttribute(xfb.Attr, "mustUseTimedInvoke")
	}
	return xfb
}

func (r *renderer) renderAccess(e xmlEncoder, op string, p matter.Privilege) (err error) {
	ax := xml.StartElement{Name: xml.Name{Local: "access"}}
	ax.Attr = setAttributeValue(ax.Attr, "op", op)
	var px xml.Attr
	if r.errata.WriteRoleAsPrivilege {
		px, _ = p.MarshalXMLAttr(xml.Name{Local: "privilege"})
	} else {
		px, _ = p.MarshalXMLAttr(xml.Name{Local: "role"})
	}
	ax.Attr = append(ax.Attr, px)
	err = e.EncodeToken(ax)
	if err != nil {
		return
	}
	err = e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "access"}})
	if err != nil {
		return
	}
	return
}
