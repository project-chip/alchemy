package matter

import (
	"encoding/json"

	"github.com/hasty/alchemy/matter/conformance"
)

type Features struct {
	Bitmap
}

func (f *Features) Clone() *Features {
	return &Features{Bitmap: *f.Bitmap.Clone()}
}

type Feature struct {
	BitmapBit
	Code string
}

func NewFeature(bit string, name string, code string, summary string, conformance conformance.Set) *Feature {
	return &Feature{BitmapBit: BitmapBit{bit: bit, name: name, summary: summary, conformance: conformance}, Code: code}
}

func (c *Feature) MarshalJSON() ([]byte, error) {
	type Alias Feature
	return json.Marshal(
		&struct {
			*Alias
			Code string `json:"code"`
		}{
			Alias: (*Alias)(c),
			Code:  c.Code,
		},
	)

}
