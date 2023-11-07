package parse

import (
	"encoding/xml"
	"fmt"
	"io"
	"strings"

	"github.com/hasty/alchemy/matter"
)

func readAccess(d *xml.Decoder, e xml.StartElement, access *matter.Access) (err error) {
	var op string
	var privilege string
	for _, a := range e.Attr {
		switch a.Name.Local {
		case "op":
			op = a.Value
		case "role", "privilege":
			privilege = a.Value
		default:
			return fmt.Errorf("unexpected access attribute: %s", a.Name.Local)
		}
	}
	p := Privilege(strings.ToLower(privilege))
	if p == matter.PrivilegeUnknown {
		return fmt.Errorf("unknown privilege value: %s", privilege)
	}
	switch strings.ToLower(op) {
	case "read":
		access.Read = p
	case "write":
		access.Write = p
	case "invoke":
		access.Invoke = p
	default:
		return fmt.Errorf("unknown privilege value: %s", privilege)
	}
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			return fmt.Errorf("EOF before end of access")

		} else if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.EndElement:
			switch t.Name.Local {
			case "access":
				return nil
			default:
				return fmt.Errorf("unexpected access end element: %s", t.Name.Local)
			}
		case xml.CharData:
		default:
			return fmt.Errorf("unexpected access level type: %T", t)
		}
		if err != nil {
			return
		}
	}
}
