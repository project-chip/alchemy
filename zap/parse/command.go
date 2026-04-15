package parse

import (
	"encoding/xml"
	"fmt"
	"io"
	"log/slog"

	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

func readCommand(path string, d *xml.Decoder, e xml.StartElement) (c *matter.Command, err error) {
	c = &matter.Command{Access: matter.DefaultAccess(types.EntityTypeCommand)}
	var optional, isFabricScoped, disableDefaultResponse string
	for _, a := range e.Attr {
		switch a.Name.Local {
		case "source":
			switch a.Value {
			case "client":
				c.Direction = matter.InterfaceServer
			case "server":
				c.Direction = matter.InterfaceClient
			}
		case "code":
			c.ID = matter.ParseNumber(a.Value)
		case "name":
			c.Name = a.Value
		case "isFabricScoped":
			isFabricScoped = a.Value
		case "optional":
			optional = a.Value
		case "response":
			c.Response = types.NewCustomDataType(a.Value, types.DataTypeRankScalar)
		case "mustUseTimedInvoke":
			if a.Value == "true" {
				c.Access.Timing = matter.TimingTimed
			} else {
				c.Access.Timing = matter.TimingUntimed
			}
		case "cli":
		case "disableDefaultResponse":
			disableDefaultResponse = a.Value
		case "apiMaturity":
		case "cliFunctionName":
		case "noDefaultImplementation":

		default:
			return nil, fmt.Errorf("unexpected command attribute: %s", a.Name.Local)
		}
	}
	if optional == "true" {
		c.Conformance = conformance.Set{&conformance.Optional{}}
	} else {
		c.Conformance = conformance.Set{&conformance.Mandatory{}}
	}

	if isFabricScoped == "true" {
		c.Access.FabricScoping = matter.FabricScopingScoped
	} else {
		c.Access.FabricScoping = matter.FabricScopingUnscoped
	}
	if disableDefaultResponse == "true" || c.Direction == matter.InterfaceClient {
		c.Response = types.NewCustomDataType("N", types.DataTypeRankScalar)
	} else if c.Response == nil {
		c.Response = types.NewCustomDataType("Y", types.DataTypeRankScalar)
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
			case "quality":
				c.Quality, err = parseQuality(d, t)
			case "arg":
				var f *matter.Field
				f, err = readField(path, d, t, types.EntityTypeCommand, "arg", c)
				if err != nil {
					slog.Warn("error reading command field", slog.Any("error", err))
				} else {
					c.Fields = append(c.Fields, f)
				}
			default:
				if isConformanceElement(t) {
					var cs conformance.Conformance
					cs, err = parseConformance(d, t)
					if err == nil {
						c.Conformance = append(c.Conformance, cs)
					}
				} else {
					err = fmt.Errorf("unexpected command level element: %s", t.Name.Local)
				}
			}
		case xml.EndElement:
			switch t.Name.Local {
			case "command":
				return
			default:
				err = fmt.Errorf("unexpected command end element: %s", t.Name.Local)
			}
		case xml.CharData, xml.Comment:
		default:
			err = fmt.Errorf("unexpected command level type: %T", t)
		}
		if err != nil {
			err = fmt.Errorf("error parsing command: %w", err)
			return
		}
	}
}
