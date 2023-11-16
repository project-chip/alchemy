package amend

import (
	"encoding/xml"
	"fmt"

	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/parse"
)

func (r *renderer) amendBitmap(d xmlDecoder, e xmlEncoder, el xml.StartElement, cluster *matter.Cluster, clusterIDs []string, bitmaps map[*matter.Bitmap]struct{}) (err error) {
	name := getAttributeValue(el.Attr, "name")

	var matchingBitmap *matter.Bitmap
	for bm := range bitmaps {
		if bm.Name == name {
			matchingBitmap = bm
			delete(bitmaps, bm)
			break
		}
	}

	Ignore(d, "bitmap")

	if matchingBitmap == nil {
		return nil
	}

	return r.writeBitmap(e, el, matchingBitmap, clusterIDs, false)
}

func (r *renderer) writeBitmap(e xmlEncoder, xfb xml.StartElement, bitmap *matter.Bitmap, clusterIDs []string, provisional bool) (err error) {
	xfb.Attr = setAttributeValue(xfb.Attr, "name", bitmap.Name)
	if bitmap.Type != "" {
		xfb.Attr = setAttributeValue(xfb.Attr, "type", bitmap.Type)
	} else {
		xfb.Attr = setAttributeValue(xfb.Attr, "type", "enum8")
	}
	if provisional {
		xfb.Attr = setAttributeValue(xfb.Attr, "apiMaturity", "provisional")
	}

	err = e.EncodeToken(xfb)
	if err != nil {
		return
	}

	err = r.renderClusterCodes(e, clusterIDs)
	if err != nil {
		return
	}

	for _, b := range bitmap.Bits {
		if b.Conformance == "Zigbee" {
			continue
		}

		bit, er := parse.HexOrDec(b.Bit)
		if er != nil {
			continue
		}

		elName := xml.Name{Local: "field"}
		xfs := xml.StartElement{Name: elName}
		xfs.Attr = setAttributeValue(xfs.Attr, "name", b.Name)
		xfs.Attr = setAttributeValue(xfs.Attr, "mask", fmt.Sprintf("%#02x", 1<<(bit)))
		err = e.EncodeToken(xfs)
		if err != nil {
			return
		}
		xfe := xml.EndElement{Name: elName}
		err = e.EncodeToken(xfe)
		if err != nil {
			return
		}

	}
	err = e.EncodeToken(xml.EndElement{Name: xfb.Name})
	return
}
