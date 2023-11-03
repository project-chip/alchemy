package zap

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/hasty/matterfmt/matter"
	"github.com/hasty/matterfmt/parse"
)

type XMLBitmapItem struct {
	XMLName      xml.Name `xml:"field"`
	Name         string   `xml:"name,attr"`
	Mask         string   `xml:"mask,attr"`
	FieldID      int      `xml:"fieldId,attr"`
	IntroducedIn string   `xml:"introducedIn,attr"`
}

type XMLBitmap struct {
	XMLName xml.Name        `xml:"bitmap"`
	Name    string          `xml:"name,attr"`
	Type    string          `xml:"type,attr"`
	Cluster XMLClusterCode  `xml:"cluster"`
	Items   []XMLBitmapItem `xml:"field"`
}

func (b *XMLBitmap) ToModel() (mb *matter.Bitmap, err error) {
	mb = &matter.Bitmap{
		Name: b.Name,
		Type: b.Type,
	}
	for _, bi := range b.Items {
		var mask uint64
		mask, err = parse.HexOrDec(bi.Mask)
		if err != nil {
			return
		}
		bit := 0
		for mask > 0 {
			if (mask & 1) == 1 {
				if mask > 1 {
					err = fmt.Errorf("non-power of 2 mask: %s", bi.Mask)
					return
				}
				break
			}
			bit++
			mask >>= 1
		}

		mb.Bits = append(mb.Bits, &matter.BitmapValue{
			Name: bi.Name,
			Bit:  strconv.Itoa(bit),
		})
	}
	return
}
