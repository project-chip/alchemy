package matter

import (
	"encoding/json"
	"iter"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

type Features struct {
	Bitmap
}

func NewFeatures(source asciidoc.Element, parent types.Entity) *Features {
	f := &Features{
		Bitmap: Bitmap{
			Name: "Features",
			Type: types.NewDataType(types.BaseDataTypeMap32, false),
			entity: entity{
				source: source,
				parent: parent,
			},
		},
	}
	return f
}

func (fs *Features) Clone() *Features {
	return &Features{Bitmap: *fs.Bitmap.Clone()}
}

func (fs *Features) CloneTo(cluster *Cluster) *Features {
	f := &Features{Bitmap: *fs.Bitmap.Clone()}
	f.entity.parent = cluster
	return f
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

func (fs *Features) AddFeatureBit(b *Feature) {
	b.parent = fs
	fs.Bits = append(fs.Bits, b)
}

func (fs *Features) FeatureBits() iter.Seq[*Feature] {
	return func(yield func(*Feature) bool) {
		for _, b := range fs.Bits {
			f, ok := b.(*Feature)
			if ok && !yield(f) {
				return
			}
		}
	}
}

type Feature struct {
	BitmapBit
	Code string
}

func NewFeature(source asciidoc.Element, bit string, name string, code string, summary string, conformance conformance.Set) *Feature {
	return &Feature{BitmapBit: BitmapBit{entity: entity{source: source}, bit: bit, name: name, summary: summary, conformance: conformance}, Code: code}
}

func (f *Feature) Entity() types.EntityType {
	return types.EntityTypeFeature
}

func (f *Feature) Clone() Bit {
	return NewFeature(f.source, f.bit, f.name, f.Code, f.summary, f.conformance)
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
