package parse

import (
	"encoding/xml"
	"fmt"
	"io"
	"strings"

	"github.com/hasty/alchemy/matter"
)

func readAttribute(d *xml.Decoder, e xml.StartElement) (attr *matter.Field, err error) {
	attr = matter.NewAttribute()
	attr.Access = matter.DefaultAccess(false)
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
			default:
				err = fmt.Errorf("unexpected attribute level element: %s", t.Name.Local)
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
