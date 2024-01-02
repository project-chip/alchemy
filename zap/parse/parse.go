package parse

import (
	"encoding/xml"
	"fmt"
	"io"

	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/types"
)

func ZAP(r io.Reader) (entities []types.Entity, err error) {
	d := xml.NewDecoder(r)
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = nil
			return
		} else if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.ProcInst:
		case xml.CharData:
		case xml.StartElement:
			switch t.Name.Local {
			case "configurator":
				var cm []types.Entity
				cm, err = readConfigurator(d)
				if err == nil {
					entities = append(entities, cm...)
				}
			default:
				err = fmt.Errorf("unexpected top level element: %s", t.Name.Local)
			}
		case xml.Comment:
		default:
			err = fmt.Errorf("unexpected top level type: %T", t)
		}
		if err != nil {
			return
		}
	}
}

func Privilege(a string) matter.Privilege {
	switch a {
	case "view":
		return matter.PrivilegeView
	case "manage":
		return matter.PrivilegeManage
	case "administer":
		return matter.PrivilegeAdminister
	case "operate":
		return matter.PrivilegeOperate
	default:
		return matter.PrivilegeUnknown
	}
}

func readSimpleElement(d *xml.Decoder, name string) (val string, err error) {
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of %s", name)
		}
		if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.EndElement:
			switch t.Name.Local {
			case name:
				return
			default:
				err = fmt.Errorf("unexpected %s end element: %s", name, t.Name.Local)
			}
		case xml.CharData:
			val = string(t)
		default:
			err = fmt.Errorf("unexpected %s level type: %T", name, t)
		}
		if err != nil {
			return
		}
	}
}

func Ignore(d *xml.Decoder, name string) (err error) {
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			return fmt.Errorf("EOF before end of %s", name)
		} else if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.EndElement:
			switch t.Name.Local {
			case name:
				return nil
			default:
			}
		default:
		}
		if err != nil {
			return
		}
	}
}

func Extract(d *xml.Decoder, el xml.StartElement) (tokens []xml.Token, err error) {
	tokens = append(tokens, el)
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of %s", el.Name.Local)
		}
		if err != nil {
			return
		}
		tokens = append(tokens, tok)
		switch t := tok.(type) {
		case xml.EndElement:
			switch t.Name.Local {
			case el.Name.Local:
				return
			default:
			}
		default:
		}
		if err != nil {
			return
		}
	}
}

func readTag(d *xml.Decoder, e xml.StartElement) (c *matter.Bitmap, err error) {
	c = &matter.Bitmap{Name: "Feature", Type: types.NewDataType("map32", false)}
	for _, a := range e.Attr {
		switch a.Name.Local {
		case "name":
			c.Name = a.Value
		case "description":
			c.Description = a.Value
		default:
			return nil, fmt.Errorf("unexpected tag attribute: %s", a.Name.Local)
		}
	}
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of field")
		}
		if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.EndElement:
			switch t.Name.Local {
			case "tag":
				return
			default:
				err = fmt.Errorf("unexpected tag end element: %s", t.Name.Local)
			}
		case xml.CharData:
		default:
			err = fmt.Errorf("unexpected tag level type: %T", t)
		}
		if err != nil {
			return
		}
	}
}
