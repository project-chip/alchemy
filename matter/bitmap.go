package matter

import (
	"encoding/json"
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/hasty/alchemy/matter/conformance"
	"github.com/hasty/alchemy/matter/types"
	"github.com/hasty/alchemy/parse"
)

type Bitmap struct {
	Name        string          `json:"name,omitempty"`
	Description string          `json:"description,omitempty"`
	Type        *types.DataType `json:"type,omitempty"`
	Bits        BitSet          `json:"bits,omitempty"`
}

func (c *Bitmap) EntityType() types.EntityType {
	return types.EntityTypeBitmap
}

func (c *Bitmap) BaseDataType() types.BaseDataType {
	return c.Type.BaseType
}

func (c *Bitmap) NullValue() uint64 {
	return c.Type.NullValue()
}

func (c *Bitmap) Size() int {
	if c.Type == nil {
		return 8
	}
	switch c.Type.BaseType {
	case types.BaseDataTypeMap64:
		return 64
	case types.BaseDataTypeMap32:
		return 32
	case types.BaseDataTypeMap16:
		return 16
	default:
		return 8
	}
}

func (bm *Bitmap) Clone() *Bitmap {
	nbm := &Bitmap{Name: bm.Name, Description: bm.Description}
	if bm.Type != nil {
		nbm.Type = bm.Type.Clone()
	}
	for _, b := range bm.Bits {
		nbm.Bits = append(nbm.Bits, b.Clone())
	}
	return nbm
}

func (bm Bitmap) Identifier(id string) (types.Entity, bool) {
	if len(bm.Bits) == 0 {
		return nil, false
	}
	for _, b := range bm.Bits {
		if b.Name() == id {
			return b, true
		}
	}
	return nil, false
}

func (bm *Bitmap) Inherit(parent *Bitmap) error {
	mergedBits := make(BitSet, 0, len(parent.Bits))
	for _, b := range parent.Bits {
		mergedBits = append(mergedBits, b.Clone())
	}
	for _, b := range bm.Bits {
		var matching Bit
		for _, mb := range mergedBits {
			if b.Bit() == mb.Bit() {
				matching = b
				break
			}
		}
		if matching == nil {
			mergedBits = append(mergedBits, b.Clone())
			continue
		}
		err := b.Inherit(matching)
		if err != nil {
			return err
		}
	}
	if bm.Type == nil {
		bm.Type = parent.Type
	}
	if len(bm.Description) == 0 {
		bm.Description = parent.Description
	}
	slices.SortFunc(mergedBits, func(a, b Bit) int {
		return strings.Compare(a.Bit(), b.Bit())
	})
	bm.Bits = mergedBits
	return nil
}

type BitmapSet []*Bitmap

func (bs BitmapSet) Identifier(name string) (types.Entity, bool) {
	for _, e := range bs {
		if e.Name == name {
			return e, true
		}
	}
	return nil, false
}

type Bit interface {
	types.Entity

	Bit() string
	Name() string
	Summary() string
	Conformance() conformance.Set

	Inherit(parent Bit) error
	Clone() Bit

	Bits() (from uint64, to uint64, err error)
	Mask() (uint64, error)
}

type BitmapBit struct {
	bit         string
	name        string
	summary     string
	conformance conformance.Set
}

func NewBitmapBit(bit string, name string, summary string, conformance conformance.Set) *BitmapBit {
	return &BitmapBit{bit: bit, name: name, summary: summary, conformance: conformance}
}

func (c *BitmapBit) EntityType() types.EntityType {
	return types.EntityTypeBitmapValue
}

func (bmb *BitmapBit) Bit() string {
	return bmb.bit
}

func (bmb *BitmapBit) Name() string {
	return bmb.name
}

func (bmb *BitmapBit) Summary() string {
	return bmb.summary
}

func (bmb *BitmapBit) Conformance() conformance.Set {
	return bmb.conformance
}

func (c *BitmapBit) Clone() Bit {
	nb := &BitmapBit{bit: c.bit, name: c.name, summary: c.summary}
	if len(c.conformance) > 0 {
		nb.conformance = c.conformance.CloneSet()
	}
	return nb
}

func (bb *BitmapBit) Inherit(parent Bit) error {
	if len(parent.Summary()) > 0 {
		bb.summary = parent.Summary()
	}
	if len(parent.Conformance()) > 0 {
		bb.conformance = parent.Conformance().CloneSet()
	}
	return nil
}

func (c *BitmapBit) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Bit         string          `json:"bit,omitempty"`
		Name        string          `json:"name,omitempty"`
		Summary     string          `json:"summary,omitempty"`
		Conformance conformance.Set `json:"conformance,omitempty"`
	}{
		Bit:         c.bit,
		Name:        c.name,
		Summary:     c.summary,
		Conformance: c.conformance,
	})
}

var bitRangePattern = regexp.MustCompile(`^(?P<From>[0-9]+)(?:\.{2,}|\s*\-\s*)(?P<To>[0-9]+)$`)

func (bv *BitmapBit) Bits() (from uint64, to uint64, err error) {
	from, err = parse.HexOrDec(bv.bit)
	if err == nil {
		to = from
		return
	}
	matches := bitRangePattern.FindStringSubmatch(bv.bit)
	if len(matches) < 3 {
		err = fmt.Errorf("invalid bit mask range: \"%s\"", bv.bit)
		return
	}
	from, err = parse.HexOrDec(matches[1])
	if err != nil {
		return
	}
	to, err = parse.HexOrDec(matches[2])
	if err != nil {
		return
	}
	return
}

func (bv *BitmapBit) Mask() (uint64, error) {
	from, to, err := bv.Bits()
	if err != nil {
		return 0, err
	}
	if from == to {
		return 1 << (from), nil
	}
	var val uint64
	for i := from; i <= to; i++ {
		val |= (1 << i)
	}
	return val, nil
}

func (bv *BitmapBit) GetConformance() conformance.Set {
	return bv.conformance
}

type BitSet []Bit

func (bs BitSet) Identifier(name string) (types.Entity, bool) {
	for _, b := range bs {
		if b.Name() == name {
			return b, true
		}
	}
	return nil, false
}
