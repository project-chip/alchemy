package parse

import (
	"encoding/xml"
	"fmt"
	"io"
)

func readServer(d *xml.Decoder, e xml.StartElement) (err error) {
	for _, a := range e.Attr {
		switch a.Name.Local {
		case "init":
		case "tick":
		case "tickFrequency":
		default:
			return fmt.Errorf("unexpected server attribute: %s", a.Name.Local)
		}
	}
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			return fmt.Errorf("EOF before end of server")
		} else if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.EndElement:
			switch t.Name.Local {
			case "server":
				return nil
			default:
				return fmt.Errorf("unexpected server end element: %s", t.Name.Local)
			}
		case xml.CharData:
		default:
			return fmt.Errorf("unexpected server level type: %T", t)
		}
		if err != nil {
			return
		}
	}
}
