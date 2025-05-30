package parse

import (
	"encoding/xml"
	"fmt"
	"io"
	"strings"

	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

func readAttribute(d *xml.Decoder, e xml.StartElement, c *matter.Cluster) (attr *matter.Field, err error) {
	attr = matter.NewAttribute(nil, c)
	attr.Access = matter.DefaultAccess(types.EntityTypeAttribute)
	err = readFieldAttributes(e, attr, "attribute")
	if err != nil {
		return
	}
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of attribute")
		}
		if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "access":
				err = readAccess(d, t, &attr.Access)
			case "description":
				attr.Name, err = readSimpleElement(d, t.Name.Local)
			case "quality":
				c.Quality, err = parseQuality(d, t)
			default:
				if isConformanceElement(t) {
					var cs conformance.Conformance
					cs, err = parseConformance(d, t)
					if err == nil {
						attr.Conformance = append(attr.Conformance, cs)
					}
				} else {
					err = fmt.Errorf("unexpected attribute level element: %s", t.Name.Local)
				}
			}
		case xml.EndElement:
			switch t.Name.Local {
			case "attribute":
				return
			default:
				err = fmt.Errorf("unexpected attribute end element: %s", t.Name.Local)
			}
		case xml.CharData:
			if len(attr.Name) == 0 {
				attr.Name = strings.TrimSpace(string(t))
			}
		case xml.Comment:
		default:
			err = fmt.Errorf("unexpected attribute level type: %T", t)
		}
		if err != nil {
			return
		}
	}
}
