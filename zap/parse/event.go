package parse

import (
	"encoding/xml"
	"fmt"
	"io"

	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
)

func readEvent(d *xml.Decoder, e xml.StartElement) (event *matter.Event, err error) {
	event = &matter.Event{}
	for _, a := range e.Attr {
		switch a.Name.Local {
		case "side":
		case "code":
			event.ID = matter.ParseNumber(a.Value)
		case "priority":
			event.Priority = a.Value
		case "name":
			event.Name = a.Value
		case "isFabricSensitive":
			event.FabricSensitive = (a.Value == "true")
		case "optional":
			if a.Value == "true" {
				event.Conformance = &conformance.MandatoryConformance{}
			}
		default:
			return nil, fmt.Errorf("unexpected event attribute: %s", a.Name.Local)
		}
	}
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of event")

		}
		if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "access":
				err = readAccess(d, t, &event.Access)
			case "description":
				event.Description, err = readSimpleElement(d, t.Name.Local)
			case "field":
				var field *matter.Field
				field, err = readField(d, t, "field")
				if err == nil {
					event.Fields = append(event.Fields, field)
				}
			default:
				err = fmt.Errorf("unexpected event level element: %s", t.Name.Local)
			}
		case xml.EndElement:
			switch t.Name.Local {
			case "event":
				return
			default:
				err = fmt.Errorf("unexpected event end element: %s", t.Name.Local)
			}
		case xml.CharData:
		default:
			err = fmt.Errorf("unexpected event level type: %T", t)
		}
		if err != nil {
			return
		}
	}
}
