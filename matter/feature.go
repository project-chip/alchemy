package matter

import (
	"encoding/json"

	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

type Features struct {
	Bitmap
}

func (fs *Features) Clone() *Features {
	return &Features{Bitmap: *fs.Bitmap.Clone()}
}

func (fs *Features) Identifier(id string) (types.Entity, bool) {
	if fs == nil {
		return nil, false
	}
	if len(fs.Bits) == 0 {
		return nil, false
	}
	for _, b := range fs.Bits {
		f := b.(*Feature)
		if f.Code == id {
			return b, true
		}
	}
	return fs.Bitmap.Identifier(id)
}

type Feature struct {
	BitmapBit
	Code string
}

func NewFeature(bit string, name string, code string, summary string, conformance conformance.Set) *Feature {
	return &Feature{BitmapBit: BitmapBit{bit: bit, name: name, summary: summary, conformance: conformance}, Code: code}
}

func (f *Feature) Entity() types.EntityType {
	return types.EntityTypeFeature
}

func (f *Feature) Clone() Bit {
	return NewFeature(f.bit, f.name, f.Code, f.summary, f.conformance)
}

func (f *Feature) MarshalJSON() ([]byte, error) {
	type Alias Feature
	return json.Marshal(
		&struct {
			*Alias
			Code string `json:"code"`
		}{
			Alias: (*Alias)(f),
			Code:  f.Code,
		},
	)

}
