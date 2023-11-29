package amend

import (
	"encoding/xml"
	"io"

	"github.com/hasty/alchemy/conformance"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/zap"
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

	if matchingCommand == nil {
		Ignore(ts, "command")
		return nil
	}

	el = r.setCommandElementAttributes(matchingCommand, e, el)
	err = e.EncodeToken(el)
	if err != nil {
		return
	}

	var argIndex int

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
			case "access":
				r.setAccessAttributes(t.Attr, "invoke", matchingCommand.Access.Invoke)
				writeThrough(ts, e, t)
			case "description":
				writeThrough(ts, e, t)
			case "arg":
				for {
					if argIndex >= len(matchingCommand.Fields) {
						Ignore(ts, "arg")
						break
					} else {
						f := matchingCommand.Fields[argIndex]
						argIndex++
						if conformance.IsZigbee(f.Conformance) {
							continue
						}
						t.Attr = r.setFieldAttributes(f, t.Attr, matchingCommand.Fields)
						writeThrough(ts, e, t)
						break
					}
				}

			default:

			}
		case xml.EndElement:
			switch t.Name.Local {
			case "command":
				for argIndex < len(matchingCommand.Fields) {
					f := matchingCommand.Fields[argIndex]
					argIndex++
					elName := xml.Name{Local: "arg"}
					xfs := xml.StartElement{Name: elName}
					xfs.Attr = r.setFieldAttributes(f, xfs.Attr, matchingCommand.Fields)
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
		xfs.Attr = r.setFieldAttributes(f, xfs.Attr, c.Fields)
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

func (r *renderer) setFieldAttributes(f *matter.Field, xfs []xml.Attr, fs matter.FieldSet) []xml.Attr {
	mandatory := conformance.IsMandatory(f.Conformance)
	xfs = setAttributeValue(xfs, "name", f.Name)
	xfs = writeDataType(fs, f, xfs)
	if !mandatory {
		xfs = setAttributeValue(xfs, "optional", "true")
	} else {
		xfs = removeAttribute(xfs, "optional")
	}
	if f.Quality.Has(matter.QualityNullable) {
		xfs = setAttributeValue(xfs, "isNullable", "true")
	} else {
		xfs = removeAttribute(xfs, "isNullable")
	}
	if f.Default != "" {
		defaultValue := zap.GetDefaultValue(&matter.ConstraintContext{Field: f, Fields: fs})
		if defaultValue.Defined() {
			xfs = setAttributeValue(xfs, "default", defaultValue.ZapString(f.Type))
		} else {
			xfs = removeAttribute(xfs, "default")
		}
	} else {
		xfs = removeAttribute(xfs, "default")
	}
	xfs = r.renderConstraint(fs, f, xfs)
	return xfs
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
	ax.Attr = r.setAccessAttributes(ax.Attr, op, p)
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

func (r *renderer) setAccessAttributes(ax []xml.Attr, op string, p matter.Privilege) []xml.Attr {
	ax = setAttributeValue(ax, "op", op)
	var px xml.Attr
	if r.errata.WriteRoleAsPrivilege {
		px, _ = p.MarshalXMLAttr(xml.Name{Local: "privilege"})
	} else {
		px, _ = p.MarshalXMLAttr(xml.Name{Local: "role"})
	}
	ax = setAttributeValue(ax, px.Name.Local, px.Value)
	return ax
}
