package parse

import (
	"encoding/xml"
	"fmt"
	"io"
	"log/slog"

	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
	"github.com/hasty/alchemy/matter/types"
)

func readEvent(path string, d *xml.Decoder, e xml.StartElement) (event *matter.Event, err error) {
	event = &matter.Event{Access: matter.DefaultAccess(types.EntityTypeEvent)}
	var optional, isFabricSensitive string
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
			isFabricSensitive = a.Value
		case "optional":
			optional = a.Value
		case "apiMaturity":
		default:
			return nil, fmt.Errorf("unexpected event attribute: %s", a.Name.Local)
		}
	}

	if optional == "true" {
		event.Conformance = conformance.Set{&conformance.Optional{}}
	} else {
		event.Conformance = conformance.Set{&conformance.Mandatory{}}
	}
	if isFabricSensitive == "true" {
		event.Access.FabricSensitivity = matter.FabricSensitivitySensitive
	} else {
		event.Access.FabricSensitivity = matter.FabricSensitivityInsensitive
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
				field, err = readField(path, d, t, types.EntityTypeEvent, "field")
				if err != nil {
					slog.Warn("error reading event field", slog.Any("error", err))
				} else {
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
		case xml.CharData, xml.Comment:
		default:
			err = fmt.Errorf("unexpected event level type: %T", t)
		}
		if err != nil {
			return
		}
	}
}
