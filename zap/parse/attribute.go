package parse

import (
	"encoding/xml"
	"fmt"
	"io"

	"github.com/hasty/alchemy/matter"
)

func readAttribute(d *xml.Decoder, e xml.StartElement) (attr *matter.Field, err error) {
	attr = &matter.Field{Type: &matter.DataType{}}
	err = readFieldAttributes(e, attr, "attribute")
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
				a := matter.Access{}
				err = readAccess(d, t, &a)
			case "description":
				_, err = readSimpleElement(d, t.Name.Local)
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
			attr.Name = string(t)
		default:
			err = fmt.Errorf("unexpected attribute level type: %T", t)
		}
		if err != nil {
			return
		}
	}
}
