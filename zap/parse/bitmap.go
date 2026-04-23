package parse

import (
	"encoding/xml"
	"fmt"
	"io"
	"strconv"

	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/zap"
)

func (sp *ZapParser) readBitmap(d *xml.Decoder, e xml.StartElement) (bitmap *matter.Bitmap, clusterIDs []*matter.Number, err error) {
	bitmap = &matter.Bitmap{}
	for _, a := range e.Attr {
		switch a.Name.Local {
		case "name":
			bitmap.Name = a.Value
		case "type":
			bitmap.Type = types.NewDataType(zap.ToBaseDataType(a.Value), types.DataTypeRankScalar)
		case "apiMaturity":
		default:
			return nil, nil, fmt.Errorf("unexpected bitmap attribute: %s", a.Name.Local)
		}
	}
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of bitmap")
		}
		if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "cluster":
				var cid *matter.Number
				cid, err = readClusterCode(d, t)
				if err == nil {
					clusterIDs = append(clusterIDs, cid)
				}
			case "description":
				bitmap.Description, err = readSimpleElement(d, t.Name.Local)
			case "field":
				var bit *matter.BitmapBit
				bit, err = readBitmapField(bitmap, d, t)
				if err == nil {
					bitmap.Bits = append(bitmap.Bits, bit)
				}
			default:
				err = fmt.Errorf("unexpected bitmap level element: %s", t.Name.Local)
			}
		case xml.EndElement:
			switch t.Name.Local {
			case "bitmap":
				return
			default:
				err = fmt.Errorf("unexpected bitmap end element: %s", t.Name.Local)
			}
		case xml.CharData, xml.Comment:
		default:
			err = fmt.Errorf("unexpected bitmap level type: %T", t)
		}
		if err != nil {
			err = fmt.Errorf("error parsing bitmap: %w", err)
			return
		}
	}
}

func readBitmapField(bitmap *matter.Bitmap, d *xml.Decoder, e xml.StartElement) (bv *matter.BitmapBit, err error) {
	var name, bit string
	var conf conformance.Set
	for _, a := range e.Attr {
		switch a.Name.Local {
		case "name":
			name = a.Value
		case "mask":
			var mask uint64
			mask, err = parse.HexOrDec(a.Value)
			if err != nil {
				return
			}
			startBit := -1
			endBit := -1

			var maxBit int
			switch bitmap.Type.BaseType {
			case types.BaseDataTypeMap8:
				maxBit = 8
			case types.BaseDataTypeMap16:
				maxBit = 16
			case types.BaseDataTypeMap32:
				maxBit = 32
			case types.BaseDataTypeMap64:
				maxBit = 64
			default:
				err = fmt.Errorf("unknown bitmap type: %v", bitmap.Type)
				return
			}
			for offset := 0; offset < maxBit; offset++ {
				if mask&(1<<offset) == (1 << offset) {
					if startBit == -1 {
						startBit = offset
					} else {
						endBit = offset
					}
				} else if startBit >= 0 {
					if endBit == -1 {
						endBit = startBit
					}
					break
				}
			}
			if startBit >= 0 {
				if startBit != endBit && endBit != -1 {
					bit = fmt.Sprintf("%d..%d", startBit, endBit)
				} else {
					bit = strconv.Itoa(startBit)
				}
			}
		case "optional":
			if a.Value != "true" {
				conf = conformance.Set{&conformance.Mandatory{}}
			}
		default:
			return nil, fmt.Errorf("unexpected bitmap field attribute: %s", a.Name.Local)
		}
	}
	bv = matter.NewBitmapBit(nil, bitmap, bit, name, "", conf)
	for {
		var tok xml.Token
		tok, err = d.Token()
		if tok == nil || err == io.EOF {
			err = fmt.Errorf("EOF before end of field")
		}
		if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.EndElement:
			switch t.Name.Local {
			case "field":
				return
			default:
				err = fmt.Errorf("unexpected field end element: %s", t.Name.Local)
			}
		case xml.CharData, xml.Comment:
		default:
			err = fmt.Errorf("unexpected field level type: %T", t)
		}
		if err != nil {
			return
		}
	}
}
