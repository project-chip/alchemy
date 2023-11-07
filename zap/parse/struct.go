package parse

import (
	"encoding/xml"
	"fmt"
	"io"

	"github.com/hasty/alchemy/matter"
)

func readStruct(d *xml.Decoder, e xml.StartElement) (s *matter.Struct, clusterIDs []*matter.ID, err error) {
	s = &matter.Struct{}
	for _, a := range e.Attr {
		switch a.Name.Local {
		case "name":
			s.Name = a.Value
		case "isFabricScoped":
		default:
			return nil, nil, fmt.Errorf("unexpected struct attribute: %s", a.Name.Local)
		}
	}
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of struct")
		}
		if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "cluster":
				var cid *matter.ID
				cid, err = readClusterCode(d, t)
				if err == nil {
					clusterIDs = append(clusterIDs, cid)
				}
			case "description":
				s.Description, err = readSimpleElement(d, t.Name.Local)
			case "item":
				var f *matter.Field
				f, err = readField(d, t, "item")
				if err != nil {
					s.Fields = append(s.Fields, f)
				}
			default:
				err = fmt.Errorf("unexpected struct level element: %s", t.Name.Local)
			}
		case xml.EndElement:
			switch t.Name.Local {
			case "struct":
				return
			default:
				err = fmt.Errorf("unexpected struct end element: %s", t.Name.Local)
			}
		case xml.CharData:
		default:
			err = fmt.Errorf("unexpected struct level type: %T", t)
		}
		if err != nil {
			err = fmt.Errorf("error parsing struct: %w", err)
			return
		}
	}
}
