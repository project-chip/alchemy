package render

import (
	"encoding/xml"
	"fmt"

	"github.com/hasty/matterfmt/matter"
	"github.com/hasty/matterfmt/parse"
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

	if matchingBitmap == nil {
		return writeThrough(d, e, el)
	}

	Ignore(d, "bitmap")

	return r.writeBitmap(e, el, matchingBitmap, clusterIDs)
}

func (r *renderer) writeBitmap(e xmlEncoder, xfb xml.StartElement, bitmap *matter.Bitmap, clusterIDs []string) (err error) {
	xfb.Attr = setAttributeValue(xfb.Attr, "name", bitmap.Name)
	xfb.Attr = setAttributeValue(xfb.Attr, "type", bitmap.Type)
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
		xfs.Attr = setAttributeValue(xfs.Attr, "value", fmt.Sprintf("%#02x", 1<<(bit-1)))
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