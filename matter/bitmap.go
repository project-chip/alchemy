package matter

import (
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

func (bm Bitmap) Reference(id string) conformance.HasConformance {
	if len(bm.Bits) == 0 {
		return nil
	}
	for _, b := range bm.Bits {
		if b.Code == id {
			return b
		}
	}
	return nil
}

func (bm *Bitmap) Inherit(parent *Bitmap) error {
	mergedBits := make(BitSet, 0, len(parent.Bits))
	for _, b := range parent.Bits {
		mergedBits = append(mergedBits, b.Clone())
	}
	for _, b := range bm.Bits {
		var matching *Bit
		for _, mb := range mergedBits {
			if b.Bit == mb.Bit {
				matching = b
				break
			}
		}
		if matching == nil {
			mergedBits = append(mergedBits, b.Clone())
			continue
		}
		if len(b.Summary) > 0 {
			matching.Summary = b.Summary
		}
		if len(b.Conformance) > 0 {
			matching.Conformance = b.Conformance.CloneSet()
		}
	}
	if bm.Type == nil {
		bm.Type = parent.Type
	}
	if len(bm.Description) == 0 {
		bm.Description = parent.Description
	}
	slices.SortFunc(mergedBits, func(a, b *Bit) int {
		return strings.Compare(a.Bit, b.Bit)
	})
	bm.Bits = mergedBits
	return nil
}

type Bit struct {
	Bit         string          `json:"bit,omitempty"`
	Code        string          `json:"code,omitempty"`
	Name        string          `json:"name,omitempty"`
	Summary     string          `json:"summary,omitempty"`
	Conformance conformance.Set `json:"conformance,omitempty"`
}

func (c *Bit) EntityType() types.EntityType {
	return types.EntityTypeBitmapValue
}

func (c *Bit) Clone() *Bit {
	nb := &Bit{Bit: c.Bit, Code: c.Code, Name: c.Name, Summary: c.Summary}
	if len(c.Conformance) > 0 {
		nb.Conformance = c.Conformance.CloneSet()
	}
	return nb
}

var bitRangePattern = regexp.MustCompile(`^(?P<From>[0-9]+)(?:\.{2,}|\s*\-\s*)(?P<To>[0-9]+)$`)

func (bv *Bit) Mask() (uint64, error) {
	val, err := parse.HexOrDec(bv.Bit)
	if err == nil {
		return 1 << (val), nil
	}
	matches := bitRangePattern.FindStringSubmatch(bv.Bit)
	if len(matches) > 2 {
		from, err := parse.HexOrDec(matches[1])
		if err != nil {
			return 0, err
		}
		to, err := parse.HexOrDec(matches[2])
		if err != nil {
			return 0, err
		}
		if from > to {
			return 0, fmt.Errorf("incorrect order of bit mask range: %d..%d", from, to)
		}
		var val uint64
		for i := from; i <= to; i++ {
			val |= (1 << i)
		}
		return val, err
	}
	return 0, fmt.Errorf("invalid bit mask range: \"%s\"", bv.Bit)
}

func (bv *Bit) GetConformance() conformance.Set {
	return bv.Conformance
}

type BitSet []*Bit

func (bs BitSet) Reference(name string) conformance.HasConformance {
	for _, b := range bs {
		if b.Name == name {
			return b
		}
	}
	return nil
}
