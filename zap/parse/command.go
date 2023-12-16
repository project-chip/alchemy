package parse

import (
	"encoding/xml"
	"fmt"
	"io"

	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
)

func readCommand(d *xml.Decoder, e xml.StartElement) (c *matter.Command, err error) {
	c = &matter.Command{}
	for _, a := range e.Attr {
		switch a.Name.Local {
		case "source":

		case "code":
			c.ID = matter.ParseNumber(a.Value)
		case "name":
			c.Name = a.Value
		case "isFabricScoped":
			c.IsFabricScoped = a.Value == "true"
		case "optional":
			if a.Value == "false" {
				c.Conformance = &conformance.MandatoryConformance{}
			}
		case "response":
			c.Response = a.Value
		case "mustUseTimedInvoke":
			c.Access.Timed = a.Value == "true"
		case "cli":
		case "disableDefaultResponse":
			c.Response = "N"
		case "apiMaturity":
		case "cliFunctionName":
		case "noDefaultImplementation":

		default:
			return nil, fmt.Errorf("unexpected command attribute: %s", a.Name.Local)
		}
	}
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of command")
		}
		if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "access":
				err = readAccess(d, t, &c.Access)
			case "description":
				_, err = readSimpleElement(d, t.Name.Local)
			case "arg":
				var f *matter.Field
				f, err = readField(d, t, "arg")
				if err != nil {
					c.Fields = append(c.Fields, f)
				}
			default:
				err = fmt.Errorf("unexpected command level element: %s", t.Name.Local)
			}
		case xml.EndElement:
			switch t.Name.Local {
			case "command":
				return
			default:
				err = fmt.Errorf("unexpected command end element: %s", t.Name.Local)
			}
		case xml.CharData:
		default:
			err = fmt.Errorf("unexpected command level type: %T", t)
		}
		if err != nil {
			err = fmt.Errorf("error parsing command: %w", err)
			return
		}
	}
}
