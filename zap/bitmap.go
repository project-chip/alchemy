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

		startBit := -1
		endBit := -1

		var maxBit int
		switch mb.Type {
		case "map8":
			maxBit = 8
		case "map16":
			maxBit = 16
		case "map32":
			maxBit = 32
		case "map64":
			maxBit = 64
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
			bv := &matter.BitmapValue{
				Name: bi.Name,
			}
			if startBit != endBit {
				bv.Bit = fmt.Sprintf("%d..%d", startBit, endBit)
			} else {
				bv.Bit = strconv.Itoa(startBit)
			}
			mb.Bits = append(mb.Bits, bv)

		}
	}
	return
}
