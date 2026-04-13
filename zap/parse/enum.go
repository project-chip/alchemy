package parse

import (
	"encoding/xml"
	"fmt"
	"io"

	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/zap"
)

func (sp *ZapParser) readEnum(d *xml.Decoder, e xml.StartElement) (en *matter.Enum, clusterIDs []*matter.Number, err error) {
	en = &matter.Enum{}
	for _, a := range e.Attr {
		switch a.Name.Local {
		case "name":
			en.Name = a.Value
		case "type":
			en.Type = types.NewDataType(zap.ToBaseDataType(a.Value), types.DataTypeRankScalar)
		case "apiMaturity":
		default:
			return nil, nil, fmt.Errorf("unexpected enum attribute: %s", a.Name.Local)
		}
	}
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of enum")

		}
		if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			//case "access":
			//	err = readAccess(d, t, a)
			case "description":
				en.Description, err = readSimpleElement(d, t.Name.Local)
			case "item":
				var ev *matter.EnumValue
				ev, err = readEnumItem(d, t)
				if err == nil {
					en.Values = append(en.Values, ev)
				}
			case "cluster":
				var cid *matter.Number
				cid, err = readClusterCode(d, t)
				if err == nil {
					clusterIDs = append(clusterIDs, cid)
				}
			default:
				err = fmt.Errorf("unexpected enum level element: %s", t.Name.Local)
			}
		case xml.EndElement:
			switch t.Name.Local {
			case "enum":
				return
			default:
				err = fmt.Errorf("unexpected enum end element: %s", t.Name.Local)
			}
		case xml.CharData, xml.Comment:
		default:
			err = fmt.Errorf("unexpected enum level type: %T", t)
		}
		if err != nil {
			err = fmt.Errorf("error parsing enum: %w", err)
			return
		}
	}
}

func readEnumItem(d *xml.Decoder, e xml.StartElement) (ev *matter.EnumValue, err error) {
	ev = &matter.EnumValue{}
	for _, a := range e.Attr {
		switch a.Name.Local {
		case "name":
			ev.Name = a.Value
		case "value":
			ev.Value = matter.ParseNumber(a.Value)
		case "apiMaturity":
		default:
			return nil, fmt.Errorf("unexpected enum item attribute: %s", a.Name.Local)
		}
	}
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of enum item")
		}
		if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.EndElement:
			switch t.Name.Local {
			case "item":
				return
			default:
				err = fmt.Errorf("unexpected enum item end element: %s", t.Name.Local)
			}
		case xml.CharData, xml.Comment:
		default:
			err = fmt.Errorf("unexpected enum item level type: %T", t)
		}
		if err != nil {
			return
		}
	}
}
