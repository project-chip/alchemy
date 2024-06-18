package parse

import (
	"encoding/xml"
	"fmt"
	"io"
)

func readClient(d *xml.Decoder, e xml.StartElement) (err error) {
	for _, a := range e.Attr {
		switch a.Name.Local {
		case "init":
		case "tick":
		default:
			return fmt.Errorf("unexpected client attribute: %s", a.Name.Local)
		}
	}
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			return fmt.Errorf("EOF before end of client")

		} else if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.EndElement:
			switch t.Name.Local {
			case "client":
				return nil
			default:
				return fmt.Errorf("unexpected client end element: %s", t.Name.Local)
			}
		case xml.CharData, xml.Comment:
		default:
			return fmt.Errorf("unexpected client level type: %T", t)
		}
	}
}
