package render

import (
	"encoding/xml"

	"github.com/hasty/alchemy/matter"
)

func (r *renderer) amendCommand(ts *tokenSet, e xmlEncoder, el xml.StartElement, commands map[*matter.Command]struct{}) (err error) {
	code := getAttributeValue(el.Attr, "code")
	source := getAttributeValue(el.Attr, "source")
	commandID := matter.ParseID(code)

	var matchingCommand *matter.Command
	for c := range commands {
		if c.ID.Equals(commandID) {
			if c.Direction == matter.CommandDirectionClientToServer && source == "client" {
				matchingCommand = c
				delete(commands, c)
				break
			}
			if c.Direction == matter.CommandDirectionServerToClient && source == "server" {
				matchingCommand = c
				delete(commands, c)
				break
			}
		}
	}

	if matchingCommand == nil {
		return writeThrough(ts, e, el)
	}

	Ignore(ts, "command")

	return r.writeCommand(e, el, matchingCommand)
}

func (r *renderer) writeCommand(e xmlEncoder, el xml.StartElement, c *matter.Command) (err error) {
	mandatory := (c.Conformance == "M")

	xfb := el.Copy()
	xfb.Name = xml.Name{Local: "command"}

	var serverSource bool
	if c.Direction == matter.CommandDirectionClientToServer {
		xfb.Attr = setAttributeValue(xfb.Attr, "source", "client")
	} else if c.Direction == matter.CommandDirectionServerToClient {
		xfb.Attr = setAttributeValue(xfb.Attr, "source", "server")
		serverSource = true
	}
	xfb.Attr = setAttributeValue(xfb.Attr, "code", c.ID.ShortHexString())
	xfb.Attr = setAttributeValue(xfb.Attr, "name", c.Name)
	if c.Access.FabricScoped {
		xfb.Attr = setAttributeValue(xfb.Attr, "isFabricScoped", "true")
	}
	if !mandatory {
		xfb.Attr = setAttributeValue(xfb.Attr, "optional", "true")
	} else {
		xfb.Attr = setAttributeValue(xfb.Attr, "optional", "false")
	}
	if len(c.Response) > 0 && c.Response != "Y" && c.Response != "N" {
		xfb.Attr = setAttributeValue(xfb.Attr, "response", c.Response)
	}
	if c.Response == "N" && !serverSource {
		xfb.Attr = setAttributeValue(xfb.Attr, "disableDefaultResponse", "true")
	}
	if c.Access.Timed {
		xfb.Attr = setAttributeValue(xfb.Attr, "mustUseTimedInvoke", "true")
	}

	err = e.EncodeToken(xfb)
	if err != nil {
		return
	}

	if c.Access.Invoke != matter.PrivilegeUnknown {

		err = r.renderAccess(e, "invoke", c.Access.Invoke)
		if err != nil {
			return
		}
	}
	for _, f := range c.Fields {
		if f.Conformance == "Zigbee" {
			continue
		}

		elName := xml.Name{Local: "arg"}
		xfs := xml.StartElement{Name: elName}
		mandatory := (f.Conformance == "M")
		xfs.Attr = setAttributeValue(xfs.Attr, "name", f.Name)
		xfs.Attr = writeCommandDataType(f.Type, xfs.Attr)
		if !mandatory {
			xfs.Attr = setAttributeValue(xfs.Attr, "optional", "true")
		}
		xfs.Attr = r.renderConstraint(xfs.Attr, f.Constraint)
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
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "description"}})
	e.EncodeToken(xml.CharData(c.Description))
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "description"}})

	return e.EncodeToken(xml.EndElement{Name: xfb.Name})
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
