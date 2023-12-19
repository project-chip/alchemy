package parse

import (
	"encoding/xml"
	"fmt"
	"io"
	"strconv"

	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
	"github.com/hasty/alchemy/parse"
	"github.com/hasty/alchemy/zap"
)

func readBitmap(d *xml.Decoder, e xml.StartElement) (bitmap *matter.Bitmap, clusterIDs []*matter.Number, err error) {
	bitmap = &matter.Bitmap{}
	for _, a := range e.Attr {
		switch a.Name.Local {
		case "name":
			bitmap.Name = a.Value
		case "type":
			bitmap.Type = matter.NewDataType(zap.ConvertZapToDataTypeName(a.Value), false)
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
				var bit *matter.Bit
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
		//case xml.CharData:
		default:
			err = fmt.Errorf("unexpected bitmap level type: %T", t)
		}
		if err != nil {
			err = fmt.Errorf("error parsing bitmap: %w", err)
			return
		}
	}
}

func readBitmapField(bitmap *matter.Bitmap, d *xml.Decoder, e xml.StartElement) (bv *matter.Bit, err error) {
	bv = &matter.Bit{}
	for _, a := range e.Attr {
		switch a.Name.Local {
		case "name":
			bv.Name = a.Value
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
			case matter.BaseDataTypeMap8:
				maxBit = 8
			case matter.BaseDataTypeMap16:
				maxBit = 16
			case matter.BaseDataTypeMap32:
				maxBit = 32
			case matter.BaseDataTypeMap64:
				maxBit = 64
			default:
				err = fmt.Errorf("unknown bitmap type: %v", bitmap.Type)
				return
			}
			for offset := 0; offset < maxBit; offset++ {
				if mask&(1<<offset) == 1 {
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
				if startBit != endBit {
					bv.Bit = fmt.Sprintf("%d..%d", startBit, endBit)
				} else {
					bv.Bit = strconv.Itoa(startBit)
				}
			}
		case "optional":
			if a.Value != "true" {
				bv.Conformance = conformance.Set{&conformance.Mandatory{}}
			}
		default:
			return nil, fmt.Errorf("unexpected bitmap field attribute: %s", a.Name.Local)
		}
	}
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
		case xml.CharData:
		default:
			err = fmt.Errorf("unexpected field level type: %T", t)
		}
		if err != nil {
			return
		}
	}
}
